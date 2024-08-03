package client

import "mime/multipart"

type ImageUploadReq struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}
type ImageUploadRes struct {
	ImageUrl string `json:"image_url"`
}
