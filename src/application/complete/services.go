package complete

import (
	"context"

	"onlineApplicationAPI/src/application/complete/entity"
)

type EmailService interface {
	SendMessage(ctx context.Context, email string, message string) error
}

type BullySystemService interface {
	GetCriminalRecords(ctx context.Context, personalID string) (entity.CriminalRecord, error)
}
