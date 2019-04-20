package repository

import (
	"context"
	"mime/multipart"
	"net/url"

	"onlineApplicationAPI/src/application/submit"
)

type defaultFileRepository struct {
}

func NewDefaultFileRepository() submit.FileRepository {
	return &defaultFileRepository{}
}

func (defaultFileRepository) SaveCV(ctx context.Context, file multipart.File) (*url.URL, error) {

	url, err := url.Parse("https://www.sample.com/image") //Store file and get the url to aces to the file
	if err != nil {
		return nil, err
	}
	return url, nil
}
