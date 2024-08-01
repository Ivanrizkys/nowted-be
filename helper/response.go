package helper

import "github.com/gin-gonic/gin"

type Meta struct {
	Message    any `json:"message"`
	StatusCode int `json:"-"`
}

func SendResponse(c *gin.Context, data any, meta Meta) {
	c.JSON(meta.StatusCode, gin.H{
		"data": data,
		"meta": meta,
	})
	if meta.StatusCode >= 400 {
		c.Abort()
	}
}
