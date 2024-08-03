package service

import (
	"context"
	"errors"
	"fmt"
	"notes/internal/constant"
	"notes/internal/dtos"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type NoteService interface {
	ImageUpload(ctx context.Context, data dtos.ImageUploadReq) (dtos.ImageUploadRes, error)
}

type NoteServiceImpl struct {
	Validate   *validator.Validate
	Cloudinary *cloudinary.Cloudinary
}

func NewNoteService(validate *validator.Validate, cloudinary *cloudinary.Cloudinary) NoteService {
	return &NoteServiceImpl{
		Validate:   validate,
		Cloudinary: cloudinary,
	}
}

func (s *NoteServiceImpl) ImageUpload(ctx context.Context, data dtos.ImageUploadReq) (dtos.ImageUploadRes, error) {
	var res dtos.ImageUploadRes

	if data.Image.Size == 0 {
		return res, errors.Join(constant.ErrBadRequest, errors.New("image url can't be empty"))
	}

	uuid := uuid.Must(uuid.NewV6())
	cldRes, err := s.Cloudinary.Upload.Upload(ctx, data.Image, uploader.UploadParams{
		Folder:   "note",
		PublicID: fmt.Sprintf("%v", uuid),
	})

	if err != nil {
		return res, errors.Join(constant.ErrService, err)
	}

	res = dtos.ImageUploadRes{
		ImageUrl: cldRes.SecureURL,
	}
	return res, nil
}
