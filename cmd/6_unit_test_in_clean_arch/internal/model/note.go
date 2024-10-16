package model

import (
	"database/sql"
	"time"
)

type Note struct {
	UUID      string
	Info      NoteInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type NoteInfo struct {
	Title   string
	Content string
}
