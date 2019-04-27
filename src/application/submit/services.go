package submit

import (
	"context"
	"net/http"

	commonEntity "onlineApplicationAPI/src/application/common/entity"
)

type EmailService interface {
	SendMessage(ctx context.Context, email string, message string) error
}

type QueueToCompleteApplicationsService interface {
	CompleteApplication(ctx context.Context, application commonEntity.Application) error
}

type AuthenticationService interface {
	VerifyToken(ctx context.Context, request *http.Request) (string, error)
}
