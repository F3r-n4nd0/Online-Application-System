package usecase

import (
	"context"
	"mime/multipart"
	"time"

	"onlineApplicationAPI/src/application/submit"
	"onlineApplicationAPI/src/application/submit/entity"
)

type submitUseCase struct {
	emailService        submit.EmailService
	fileStoreRepository submit.FileRepository
	timeoutSeconds      time.Duration
	queueApplication    submit.QueueToCompleteApplicationsService
}

func NewSubmitUseCase(
	emailService submit.EmailService,
	fileStoreRepository submit.FileRepository) submit.UseCase {
	return &submitUseCase{
		emailService:        emailService,
		fileStoreRepository: fileStoreRepository,
		timeoutSeconds:      2,
	}
}

func (uc *submitUseCase) SubmitApplication(ctx context.Context, newApplication entity.NewApplication, cvFile multipart.File) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeoutSeconds)
	defer cancel()
	localURLCV, err := uc.fileStoreRepository.SaveCV(ctx, cvFile)
	if err != nil {
		return false, err
	}
	newApplication.CurriculumVitae = *localURLCV
	err = uc.queueApplication.ProcessNewApplication(ctx, newApplication)
	if err != nil {
		return false, err
	}
	return true, err
}
