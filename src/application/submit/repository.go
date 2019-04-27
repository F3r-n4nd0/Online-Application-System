package submit

import (
	"context"
	"mime/multipart"
	"net/url"
)

type FileRepository interface {
	StoreCV(ctx context.Context, file multipart.File) (*url.URL, error)
}
