package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"notes/internal/constant"
	"notes/internal/dtos"
	"notes/internal/helper"
	"notes/internal/model"
	"notes/internal/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/ksuid"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
)

type AuthService interface {
	GoogleOAuth(ctx context.Context, data dtos.GoogleLoginReq) (dtos.GoogleLoginRes, error)
}

type AuthServiceImpl struct {
	Pg                *sql.DB
	Validate          *validator.Validate
	GoogleOAuthConfig *oauth2.Config
	UserRepository    repository.UserRepository
	FolderRepository  repository.FolderRepository
	NoteRepository    repository.NoteRepository
}

func NewAuthService(pg *sql.DB, validate *validator.Validate, googleOAuthConfig *oauth2.Config,
	userRepository repository.UserRepository, folderRepository repository.FolderRepository, noteRepository repository.NoteRepository) AuthService {
	return &AuthServiceImpl{
		Pg:                pg,
		Validate:          validate,
		GoogleOAuthConfig: googleOAuthConfig,
		UserRepository:    userRepository,
		FolderRepository:  folderRepository,
		NoteRepository:    noteRepository,
	}
}

func (s *AuthServiceImpl) GoogleOAuth(ctx context.Context, data dtos.GoogleLoginReq) (dtos.GoogleLoginRes, error) {
	res := dtos.GoogleLoginRes{}

	err := s.Validate.Struct(&data)
	if err != nil {
		return res, err
	}
	tx, err := s.Pg.Begin()
	if err != nil {
		return res, err
	}
	defer helper.CommitOrRollback(tx, &err)

	token, err := s.GoogleOAuthConfig.Exchange(ctx, data.Code)
	if err != nil {
		return res, errors.Join(constant.ErrBadRequest, err)
	}
	peopleSrv, err := people.NewService(ctx, option.WithTokenSource(s.GoogleOAuthConfig.TokenSource(ctx, token)))
	if err != nil {
		return res, err
	}
	user, err := peopleSrv.People.Get("people/me").PersonFields("names,emailAddresses").Do()
	if err != nil {
		return res, err
	}

	// * check user in database (already exist or not)
	userCount, err := s.UserRepository.GetUserCountWhereEmail(ctx, s.Pg, user.EmailAddresses[0].Value)
	if err != nil {
		return res, err
	}
	// * save user to databases if user not exist and generate default notes and folders
	if userCount == 0 {
		userId := ksuid.New().String()
		workFolderId := ksuid.New().String()
		personalFolderId := ksuid.New().String()
		folders := []model.Folder{
			{
				FolderId:  personalFolderId,
				UserId:    userId,
				Name:      "Personal",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				FolderId:  workFolderId,
				UserId:    userId,
				Name:      "Work",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		notes := []model.Notes{
			{
				NoteId:    ksuid.New().String(),
				UserId:    userId,
				FolderId:  personalFolderId,
				Title:     "My Goals for the Next Year",
				Content:   json.RawMessage(`{}`),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				NoteId:    ksuid.New().String(),
				UserId:    userId,
				FolderId:  personalFolderId,
				Title:     "Thoughts on the Pandemic",
				Content:   json.RawMessage(`{}`),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				NoteId:    ksuid.New().String(),
				UserId:    userId,
				FolderId:  workFolderId,
				Title:     "Reflection on the Month of June",
				Content:   json.RawMessage(`{}`),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				NoteId:    ksuid.New().String(),
				UserId:    userId,
				FolderId:  workFolderId,
				Title:     "My Planning in Q2 2024",
				Content:   json.RawMessage(`{}`),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		err = s.UserRepository.AddUserTx(ctx, tx, model.User{
			UserId:    userId,
			Name:      user.Names[0].DisplayName,
			Email:     user.EmailAddresses[0].Value,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		if err != nil {
			return res, err
		}
		err = s.FolderRepository.AddMultipleFolderTx(ctx, tx, folders)
		if err != nil {
			return res, err
		}
		err = s.NoteRepository.AddMultipleNoteTx(ctx, tx, notes)
		if err != nil {
			return res, err
		}
	}

	res = dtos.GoogleLoginRes{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		TokenType:    token.TokenType,
		User: dtos.User{
			Name:  user.Names[0].DisplayName,
			Email: user.EmailAddresses[0].Value,
		},
	}
	return res, nil
}
