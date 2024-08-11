package dtos

import "time"

type GoogleLoginReq struct {
	Code string `json:"code" binding:"required"`
}

type GoogleLoginRes struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
	TokenType    string    `json:"token_type"`
	User         User      `json:"user"`
}
