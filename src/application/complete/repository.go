package complete

import (
	"context"

	"onlineApplicationAPI/src/application/complete/entity"
)

type LegacyDataBaseRepository interface {
	Store(ctx context.Context, fullApplication entity.FullApplication) error
}
