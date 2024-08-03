package controller

import "github.com/gin-gonic/gin"

type NoteController interface {
	ImageUpload(c *gin.Context)
}
