package entity

import (
	"net/url"
)

type NewApplication struct {
	Email           string
	Name            string
	PersonalID      string
	Birth           Birth
	CurriculumVitae url.URL
	CreatedBy       string
}
