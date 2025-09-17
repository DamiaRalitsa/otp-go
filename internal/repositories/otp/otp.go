package otp

import (
	"fmt"
	"log"
	"time"

	"sqe/pkg/postgres"
)

type OTPRepo struct {
	db postgres.DatabaseHandlerFunc
}

func NewOTPRepo(db postgres.DatabaseHandlerFunc) *OTPRepo {
	return &OTPRepo{db: db}
}

func (r *OTPRepo) StoreOTP(userID string, otp int) error {
	query := `
		INSERT INTO otp_requests (user_id, otp, created_at, expires_at)
		VALUES ($1, $2, NOW(), NOW() + interval '2 minutes')
	`

	err := r.db(nil, true, query, userID, otp)
	if err != nil {
		log.Printf("Error storing OTP: %v\n", err)
		return fmt.Errorf("failed to store otp: %w", err)
	}

	log.Println("OTP stored successfully: ", otp)

	return nil
}

func (r *OTPRepo) VerifyOTP(userID string, otp int) (bool, error) {
	query := `
		SELECT otp, expires_at 
		FROM otp_requests
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT 1
	`

	type result struct {
		OTP       int       `db:"otp"`
		ExpiresAt time.Time `db:"expires_at"`
	}
	var results []result

	err := r.db(&results, false, query, userID)
	if err != nil {
		log.Printf("Error retrieving OTP: %v\n", err)
		return false, fmt.Errorf("failed to get otp: %w", err)
	}
	if len(results) == 0 {
		return false, fmt.Errorf("no OTP found for user_id %s", userID)
	}

	stored := results[0]
	if time.Now().After(stored.ExpiresAt) {
		return false, fmt.Errorf("otp expired for user_id %s", userID)
	}

	if stored.OTP != otp {
		return false, fmt.Errorf("otp not found for user_id %s", userID)
	}

	delQuery := `DELETE FROM otp_requests WHERE user_id = $1`
	if err := r.db(nil, true, delQuery, userID); err != nil {
		log.Printf("Error deleting OTP after verify: %v\n", err)
	}

	return true, nil
}
