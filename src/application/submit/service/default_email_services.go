package service

import (
	"context"

	"onlineApplicationAPI/src/application/submit"
)

type defaultEmailServices struct {
}

func NewDefaultEmailServices() submit.EmailService {
	return &defaultEmailServices{}
}

func (service *defaultEmailServices) SendMessage(ctx context.Context, email string, message string) error {
	// SEND EMAIL
	return nil
}
