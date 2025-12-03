package service

import (
	"context"
	"fmt"

	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/external"
	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/repository"
)

type UserService struct {
	repo     repository.UserRepository
	notifier external.Notifier
}

func NewUserService(
	repo repository.UserRepository,
	notifier external.Notifier,
) *UserService {
	return &UserService{repo: repo, notifier: notifier}
}

func (s *UserService) WelcomeUser(ctx context.Context, id int64) error {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("find user: %w", err)
	}
	if user == nil {
		return fmt.Errorf("user not found: %d", id)
	}

	if err := s.notifier.SendWelcomeEmail(ctx, user); err != nil {
		return fmt.Errorf("send welcome email: %w", err)
	}

	return nil
}
