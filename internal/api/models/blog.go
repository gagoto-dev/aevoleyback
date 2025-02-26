package models

import (
	"time"
	// "github.com/google/uuid"
)

type BlogEntry struct {
	Id         int64
	Title      string
	AuthorId   int64
	AuthorName string
	CreatedAt  time.Time
}

func NewBlogEntry(title string, authorId int64) *BlogEntry {
	return &BlogEntry{
		Title:     title,
		AuthorId:  authorId,
		CreatedAt: time.Now(),
	}
}
