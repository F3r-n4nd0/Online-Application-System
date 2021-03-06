package entity

import (
	"net/url"

	"onlineApplicationAPI/src/application/common/entity"
)

type NewApplication struct {
	Email           string
	Name            string
	PersonalID      string
	Birth           entity.Birth
	CurriculumVitae url.URL
	CreatedBy       string
}
