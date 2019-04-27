package entity

import (
	"time"

	"onlineApplicationAPI/src/application/common/entity"
)

type FullApplication struct {
	UUID           string
	Application    entity.Application
	CriminalRecord CriminalRecord
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
