// internal/service/add_test.go
package service_test

import (
	"testing"

	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := service.Add(1, 2)
	assert.Equal(t, 3, result)
}

func TestAddNegative(t *testing.T) {
	result := service.Add(-5, 2)
	assert.Equal(t, -3, result)
}
