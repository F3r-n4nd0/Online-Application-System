package complete

import (
	"context"

	"onlineApplicationAPI/src/application/common/entity"
)

type UseCase interface {
	CompleteApplication(ctx context.Context, application entity.Application) error
}
