package service

import (
	"context"
	"net/http"

	"onlineApplicationAPI/src/application/submit"
)

type defaultAuthenticationServices struct {
}

func NewDefaultAuhtenticationServices() submit.AuthenticationService {
	return &defaultAuthenticationServices{}
}

func (defaultAuthenticationServices) VerifyToken(ctx context.Context, request *http.Request) (string, error) {
	//verify header token and get the user
	return "user-1", nil
}
