package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/handler"
)

// service.UserService の interface 抜き出し用インタフェースを用意してもいいが、
// 簡略化のため service にテスト用のインタフェースを切っても良い。

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) WelcomeUser(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestUserHandler_Welcome_Success(t *testing.T) {
	t.Helper()

	svc := new(MockUserService)
	svc.On("WelcomeUser", mock.Anything, int64(1)).
		Return(nil).
		Once()

	h := handler.NewUserHandler(svc)

	req := httptest.NewRequest(http.MethodPost, "/welcome?id=1", nil)
	w := httptest.NewRecorder()

	h.Welcome(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	svc.AssertExpectations(t)
}

func TestUserHandler_Welcome_NotFound(t *testing.T) {
	svc := new(MockUserService)
	svc.On("WelcomeUser", mock.Anything, int64(42)).
		Return(errors.New("user not found")).
		Once()

	h := handler.NewUserHandler(svc)

	req := httptest.NewRequest(http.MethodPost, "/welcome?id=42", nil)
	w := httptest.NewRecorder()

	h.Welcome(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
