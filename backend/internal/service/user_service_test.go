package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/domain"
	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/external"
	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/repository"
	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/service"
)

// mock for UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	args := m.Called(ctx, id)
	var user *domain.User
	if u := args.Get(0); u != nil {
		user = u.(*domain.User)
	}
	return user, args.Error(1)
}

type MockNotifier struct {
	mock.Mock
}

func (m *MockNotifier) SendWelcomeEmail(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

var _ repository.UserRepository = (*MockUserRepository)(nil)
var _ external.Notifier = (*MockNotifier)(nil)

func TestWelcomeUser_Success(t *testing.T) {
	ctx := context.Background()

	userRepo := new(MockUserRepository)
	notifier := new(MockNotifier)

	user := &domain.User{ID: 1, Name: "Tai", Email: "tai@example.com"}

	userRepo.On("FindByID", mock.Anything, int64(1)).
		Return(user, nil).
		Once()

	notifier.On("SendWelcomeEmail", mock.Anything, user).
		Return(nil).
		Once()

	svc := service.NewUserService(userRepo, notifier)

	err := svc.WelcomeUser(ctx, 1)

	assert.NoError(t, err)
	userRepo.AssertExpectations(t)
	notifier.AssertExpectations(t)
}

func TestWelcomeUser_NotFound(t *testing.T) {
	ctx := context.Background()

	userRepo := new(MockUserRepository)
	notifier := new(MockNotifier)

	userRepo.On("FindByID", mock.Anything, int64(42)).
		Return((*domain.User)(nil), nil).
		Once()

	svc := service.NewUserService(userRepo, notifier)

	err := svc.WelcomeUser(ctx, 42)

	assert.Error(t, err)
	notifier.AssertNotCalled(t, "SendWelcomeEmail", mock.Anything, mock.Anything)
}
