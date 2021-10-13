package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents a post that is created by users
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

// Prepare validates and formats the post
func (p *Post) Prepare() error {
	if error := p.validate(); error != nil {
		return error
	}

	p.format()
	return nil
}

func (p *Post) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}

func (p *Post) validate() error {
	if p.Title == "" {
		return errors.New("title is mandatory and cannot be blank")
	}
	if p.Content == "" {
		return errors.New("content is mandatory and cannot be blank")
	}
	return nil
}
