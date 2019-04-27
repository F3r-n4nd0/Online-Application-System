package repository

import (
	"context"
	"time"

	"onlineApplicationAPI/src/application/complete"
	"onlineApplicationAPI/src/application/complete/entity"
)

type mysqlApplicationRepository struct {
}

func NewMysqlApplicationRepository() complete.LegacyDataBaseRepository {
	return &mysqlApplicationRepository{}
}

func (mysqlApplicationRepository) Store(ctx context.Context, fullApplication entity.FullApplication) error {
	// store full application
	time.Sleep(1 * time.Minute)
	return nil
}
