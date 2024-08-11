package model

import "time"

type User struct {
	UserId    string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
