package otp

import (
	"context"
	"fmt"
	"time"

	"sqe/internal/domain"
	"sqe/internal/presenters"
	otp "sqe/internal/repositories/otp"
	"sqe/pkg/postgres"
)

type OTPRepository interface {
	StoreOTP(userID string, otp int) error
	VerifyOTP(userID string, otp int) (bool, error)
}

type OTPUsecase struct {
	Repo OTPRepository
}

func NewOTPUsecase(dbHandler postgres.DatabaseHandlerFunc) *OTPUsecase {
	if dbHandler == nil {
		dbHandler = postgres.NewDatabase(postgres.DbDetails).CreateDatabaseHandler()
	}
	repo := otp.NewOTPRepo(dbHandler)
	return &OTPUsecase{
		Repo: repo,
	}
}

func (uc *OTPUsecase) SendOTP(ctx context.Context, userID string) (domain.OTPResponse, error) {
	code := presenters.GenerateOTP()
	message := fmt.Sprintf("Your authentication code is: %s", code)

	otpData := domain.OTPRequest{
		UserID:    userID,
		OTP:       code,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(2 * time.Minute),
	}

	if err := uc.Repo.StoreOTP(otpData.UserID, presenters.StringToInt(otpData.OTP)); err != nil {
		return domain.OTPResponse{}, err
	}

	response := domain.OTPResponse{
		UserID: otpData.UserID,
		OTP:    otpData.OTP,
	}

	_ = message

	return response, nil
}

func (uc *OTPUsecase) VerifyOTP(ctx context.Context, userID string, otp string) (bool, error) {
	return uc.Repo.VerifyOTP(userID, presenters.StringToInt(otp))
}
