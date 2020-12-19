package models

import (
	"errors"
	"strings"
	"time"
)

// Publication represents a user publication
type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

// Prepare validates and formats a publication
func (publication *Publication) Prepare() error {

	if error := publication.validate(); error != nil {
		return error
	}
	publication.format()
	return nil
}

func (publication *Publication) validate() error {

	if publication.Title == "" {
		return errors.New("Invalid title")
	}

	if publication.Content == "" {
		return errors.New("Invalid content")
	}

	if publication.AuthorId == 0 {
		return errors.New("Invalid author")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}