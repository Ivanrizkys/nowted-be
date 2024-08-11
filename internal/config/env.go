package config

import "os"

func CoudinaryUrl() string {
	env := os.Getenv("CLOUDINARY_URL")
	if env == "" {
		panic("can't load coudinary url")
	}
	return env
}

func GoogleOAuthClientId() string {
	env := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	if env == "" {
		panic("can't load google oauth client id")
	}
	return env
}

func GoogleOAuthClientSecreet() string {
	env := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	if env == "" {
		panic("can't load google oauth client secreets")
	}
	return env
}

func GoogleOAuthRedirectUrl() string {
	env := os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")
	if env == "" {
		panic("can't load google oauth redirect url")
	}
	return env
}

func PostgresUrl() string {
	env := os.Getenv("POSTGRESQL_URL")
	if env == "" {
		panic("can't load postgresql url")
	}
	return env
}
