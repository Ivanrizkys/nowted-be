package model

import "time"

type Notes struct {
	NoteId    string
	UserId    string
	FolderId  string
	Title     string
	Content   []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}
