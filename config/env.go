package config

import "os"

func CoudinaryUrl() string {
	env := os.Getenv("CLOUDINARY_URL")
	if env == "" {
		panic("can't load coudinary url")
	}
	return env
}
