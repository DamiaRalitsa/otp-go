package domain

import "time"

type OTPRequest struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id" validate:"required"`
	OTP       string    `json:"otp"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type OTPResponse struct {
	UserID string `json:"user_id"`
	OTP    string `json:"otp"`
}
