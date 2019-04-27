package usecase

import (
	"context"
	"errors"
	"time"

	entityCommon "onlineApplicationAPI/src/application/common/entity"
	"onlineApplicationAPI/src/application/complete"
	"onlineApplicationAPI/src/application/complete/entity"

	uuid "github.com/satori/go.uuid"
)

type completeUseCase struct {
	emailService          complete.EmailService
	applicationRepository complete.LegacyDataBaseRepository
	bullySystemService    complete.BullySystemService
	timeoutSeconds        time.Duration
}

func NewCompleteUseCase(
	emailService complete.EmailService,
	bullySystemService complete.BullySystemService,
	applicationRepository complete.LegacyDataBaseRepository) complete.UseCase {
	return &completeUseCase{
		emailService:          emailService,
		applicationRepository: applicationRepository,
		bullySystemService:    bullySystemService,
		timeoutSeconds:        80,
	}
}

func (uc *completeUseCase) CompleteApplication(ctx context.Context, application entityCommon.Application) error {
	ctx, cancel := context.WithTimeout(ctx, uc.timeoutSeconds)
	defer cancel()

	err := validateDataApplication(application)
	if err != nil {
		return err
	}

	go func() {
		criminalRecord, err := uc.bullySystemService.GetCriminalRecords(ctx, application.PersonalID)
		if err != nil {
			uc.emailService.SendMessage(ctx, application.Email, err.Error())
			return
		}
		uuid := uuid.NewV4()
		fullApplication := entity.FullApplication{
			UUID:           uuid.String(),
			CriminalRecord: criminalRecord,
			Application:    application,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		err = uc.applicationRepository.Store(ctx, fullApplication)
		if err != nil {
			uc.emailService.SendMessage(ctx, application.Email, err.Error())
			return
		}
		uc.emailService.SendMessage(ctx, application.Email, "Submitted")
	}()

	return nil
}

func validateDataApplication(application entityCommon.Application) error {
	if application.UUID == "" {
		return errors.New("the application needs uuid")
	}
	if application.Email == "" {
		return errors.New("the application needs email")
	}
	if application.PersonalID == "" {
		return errors.New("the application needs personal id")
	}
	return nil
}
