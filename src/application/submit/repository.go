package submit

import (
	"context"
	"mime/multipart"
	"net/url"
)

type FileRepository interface {
	SaveCV(ctx context.Context, file multipart.File) (*url.URL, error)
}
