package controller

import (
	"net/http"
	"notes/helper"
	"notes/model/client"
	"notes/service"

	"github.com/gin-gonic/gin"
)

type NoteControllerImpl struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
	return &NoteControllerImpl{
		NoteService: noteService,
	}
}

func (cl *NoteControllerImpl) ImageUpload(c *gin.Context) {
	req := client.ImageUploadReq{}
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
	helper.SendResponse(c, client.ImageUploadRes{
		ImageUrl: res.ImageUrl,
	}, helper.Meta{
		StatusCode: http.StatusCreated,
		Message:    "Upload image succesfully",
	})
}
