package submit

import (
	"context"
	"net/http"

	"onlineApplicationAPI/src/application/submit/entity"
)

type EmailService interface {
	SendMessage(ctx context.Context, email string, message string) error
}

type QueueToCompleteApplicationsService interface {
	ProcessNewApplication(ctx context.Context, newApplication entity.NewApplication) error
}

type AuthenticationService interface {
	VerifyToken(ctx context.Context, request *http.Request) (string, error)
}
