package controller

import (
	"net/http"
	"notes/internal/dtos"
	"notes/internal/helper"
	"notes/internal/service"

	"github.com/gin-gonic/gin"
)

type NoteController interface {
	ImageUpload(c *gin.Context)
}

type NoteControllerImpl struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
	return &NoteControllerImpl{
		NoteService: noteService,
	}
}

func (cl *NoteControllerImpl) ImageUpload(c *gin.Context) {
	req := dtos.ImageUploadReq{}
	err := c.Bind(&req)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := cl.NoteService.ImageUpload(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}
	helper.SendResponse(c, dtos.ImageUploadRes{
		ImageUrl: res.ImageUrl,
	}, helper.Meta{
		StatusCode: http.StatusCreated,
		Message:    "Upload image succesfully",
	})
}
