package service

import (
	"context"

	"onlineApplicationAPI/src/application/submit"
	"onlineApplicationAPI/src/application/submit/entity"
)

type defaultQueueCompleApplicationrvices struct {
}

func NewDefaultQueueCompleteApplicationServices() submit.QueueToCompleteApplicationsService {
	return &defaultQueueCompleApplicationrvices{}
}

func (defaultQueueCompleApplicationrvices) ProcessNewApplication(ctx context.Context, newApplication entity.NewApplication) error {
	// Add application to the queue
	return nil
}
