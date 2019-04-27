package usecase

import (
	"context"
	"mime/multipart"
	"time"

	commonEntity "onlineApplicationAPI/src/application/common/entity"
	"onlineApplicationAPI/src/application/submit"
	"onlineApplicationAPI/src/application/submit/entity"

	uuid2 "github.com/satori/go.uuid"
)

type submitUseCase struct {
	emailService        submit.EmailService
	fileStoreRepository submit.FileRepository
	timeoutSeconds      time.Duration
	queueApplication    submit.QueueToCompleteApplicationsService
}

func NewSubmitUseCase(
	emailService submit.EmailService,
	queueApplication submit.QueueToCompleteApplicationsService,
	fileStoreRepository submit.FileRepository) submit.UseCase {
	return &submitUseCase{
		emailService:        emailService,
		fileStoreRepository: fileStoreRepository,
		queueApplication:    queueApplication,
		timeoutSeconds:      2,
	}
}

func (uc *submitUseCase) SubmitApplication(ctx context.Context, newApplication entity.NewApplication, cvFile multipart.File) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeoutSeconds)
	defer cancel()
	localURLCV, err := uc.fileStoreRepository.StoreCV(ctx, cvFile)
	if err != nil {
		return false, err
	}
	uuid := uuid2.NewV4()
	application := commonEntity.Application{
		UUID:            uuid.String(),
		PersonalID:      newApplication.PersonalID,
		Email:           newApplication.Email,
		CurriculumVitae: newApplication.CurriculumVitae,
		Birth:           newApplication.Birth,
		Name:            newApplication.Name,
		CreatedBy:       newApplication.CreatedBy,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	newApplication.CurriculumVitae = *localURLCV
	err = uc.queueApplication.CompleteApplication(ctx, application)
	if err != nil {
		return false, err
	}
	return true, err
}
