package main

import (
	"log"
	"notes/internal/config"
	"notes/internal/controller"
	"notes/internal/helper"
	"notes/internal/middleware"
	"notes/internal/repository"
	"notes/internal/service"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	// * initialize deppendency configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cld, err := cloudinary.NewFromURL(config.CoudinaryUrl())
	if err != nil {
		log.Fatal("Can't load cloudinary instance")
	}
	validate := validator.New()
	pg := config.PgConnection()
	googleOAuthConf := &oauth2.Config{
		ClientID:     config.GoogleOAuthClientId(),
		ClientSecret: config.GoogleOAuthClientSecreet(),
		RedirectURL:  config.GoogleOAuthRedirectUrl(),
		Endpoint:     google.Endpoint,
	}

	userRepository := repository.NewUserRepository(pg)
	folderRepository := repository.NewFolderRepository(pg)
	noteRepository := repository.NewNoteRepository(pg)
	authService := service.NewAuthService(pg, validate, googleOAuthConf, userRepository, folderRepository, noteRepository)
	noteService := service.NewNoteService(validate, cld)
	noteController := controller.NewNoteController(noteService)
	authController := controller.NewAuthController(authService)

	r := gin.Default()
	r.Use(middleware.CORS())
	r.Use(middleware.Error())

	r.GET("/health-check", func(c *gin.Context) {
		helper.SendResponse(c, nil, helper.Meta{
			Message:    "Activated",
			StatusCode: 200,
		})
	})
	r.POST("/note/image-upload", noteController.ImageUpload)
	r.POST("/auth/signin/google", authController.LoginWithGoogle)

	r.Run()
}
