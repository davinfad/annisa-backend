package models

import "time"

type Treatments struct {
	ID            int
	Slug          string
	TreatmentName string
	Description   string
	Price         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}