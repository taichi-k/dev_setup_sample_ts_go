package repository

import (
	"context"

	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/domain"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*domain.User, error)
}
