package service

import (
	"context"

	"onlineApplicationAPI/src/application/complete"
	"onlineApplicationAPI/src/application/complete/entity"
)

type defaultBullyServices struct {
}

func NewDefaultBullyServices() complete.BullySystemService {
	return &defaultBullyServices{}
}

func (defaultBullyServices) GetCriminalRecords(ctx context.Context, personalID string) (entity.CriminalRecord, error) {

	// get criminal records

	return entity.CriminalRecord{}, nil
}
