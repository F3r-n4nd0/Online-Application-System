package entity

import (
	"net/url"
	"time"
)

type Application struct {
	UUID string
	Email      string
	Name      string
	PersonalID   string
	Birth  Birth
	CurriculumVitae	 url.URL
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}