package main

import (
	"log"
	"notes/config"
	"notes/controller"
	"notes/helper"
	"notes/middleware"
	"notes/service"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cld, err := cloudinary.NewFromURL(config.CoudinaryUrl())
	if err != nil {
		log.Fatal("Can't load cloudinary instance")
	}
	validate := validator.New()

	noteService := service.NewNoteService(validate, cld)
	noteController := controller.NewNoteController(noteService)

	r := gin.Default()

	r.Use(middleware.CORS())
	r.Use(middleware.Error())

	r.GET("/test", func(c *gin.Context) {
		helper.SendResponse(c, nil, helper.Meta{
			Message:    "Activated",
			StatusCode: 200,
		})
	})
	r.POST("/note/image-upload", noteController.ImageUpload)

	r.Run()
}
