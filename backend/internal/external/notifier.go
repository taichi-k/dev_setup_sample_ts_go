package external

import (
	"context"

	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/domain"
)

type Notifier interface {
	SendWelcomeEmail(ctx context.Context, user *domain.User) error
}
