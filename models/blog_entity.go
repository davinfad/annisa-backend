package models

import "time"

type Blog struct {
	ID          int
	Slug        string
	Title       string
	Description string
	FileName    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}