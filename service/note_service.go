package service

import (
	"context"
	"notes/model/client"
)

type NoteService interface {
	ImageUpload(ctx context.Context, data client.ImageUploadReq) (client.ImageUploadRes, error)
}
