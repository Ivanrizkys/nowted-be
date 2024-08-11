package model

import "time"

type Folder struct {
	FolderId  string
	UserId    string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
