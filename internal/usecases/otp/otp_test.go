package otp_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"sqe/internal/repositories/otp"
	usecase "sqe/internal/usecases/otp"
)

type MockOTPRepo struct {
	mock.Mock
}

func (m *MockOTPRepo) StoreOTP(userID string, otp int) error {
	args := m.Called(userID, otp)
	return args.Error(0)
}

func (m *MockOTPRepo) VerifyOTP(userID string, otp int) (bool, error) {
	args := m.Called(userID, otp)
	return args.Bool(0), args.Error(1)
}

func TestNewOTPUsecase(t *testing.T) {
	t.Run("with custom database handler", func(t *testing.T) {
		mockDBHandler := func(dest interface{}, isExec bool, query string, values ...interface{}) error {
			return nil
		}

		uc := usecase.NewOTPUsecase(mockDBHandler)

		assert.NotNil(t, uc)
		assert.NotNil(t, uc.Repo)
	})

	t.Run("with default constructor", func(t *testing.T) {
		mockHandler := func(dest interface{}, isExec bool, query string, values ...interface{}) error {
			return nil
		}

		uc := usecase.NewOTPUsecase(mockHandler)

		assert.NotNil(t, uc)
		assert.NotNil(t, uc.Repo)

		_, ok := uc.Repo.(*otp.OTPRepo)
		assert.True(t, ok, "Expected repository to be of type *otp.OTPRepo")
	})
}

func TestSendOTP(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := &MockOTPRepo{}
		uc := &usecase.OTPUsecase{
			Repo: mockRepo,
		}

		mockRepo.On("StoreOTP", "user123", mock.AnythingOfType("int")).Return(nil)

		resp, err := uc.SendOTP(context.Background(), "user123")

		assert.NoError(t, err)
		assert.NotEmpty(t, resp.OTP)
		assert.Equal(t, "user123", resp.UserID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error storing OTP", func(t *testing.T) {
		mockRepo := &MockOTPRepo{}
		uc := &usecase.OTPUsecase{
			Repo: mockRepo,
		}

		expectedErr := assert.AnError
		mockRepo.On("StoreOTP", "user123", mock.AnythingOfType("int")).Return(expectedErr)

		resp, err := uc.SendOTP(context.Background(), "user123")

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Empty(t, resp)
		mockRepo.AssertExpectations(t)
	})
}

func TestVerifyOTP(t *testing.T) {
	mockRepo := &MockOTPRepo{}
	uc := &usecase.OTPUsecase{
		Repo: mockRepo,
	}

	mockRepo.On("VerifyOTP", "user123", 123456).Return(true, nil)

	ok, err := uc.VerifyOTP(context.Background(), "user123", "123456")

	assert.True(t, ok)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
