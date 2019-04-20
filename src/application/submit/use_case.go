package submit

import (
	"context"
	"mime/multipart"

	"onlineApplicationAPI/src/application/submit/entity"
)

type UseCase interface {
	SubmitApplication(ctx context.Context, newApplication entity.NewApplication, cvFile multipart.File) (bool, error)
}
