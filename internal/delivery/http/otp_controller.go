package http

import (
	"sqe/internal/domain"
	"sqe/internal/presenters"
	usecases "sqe/internal/usecases/otp"

	"github.com/gofiber/fiber/v2"
)

type OTPController struct {
	useCase *usecases.OTPUsecase
}

func NewOTPController() *OTPController {
	return &OTPController{
		useCase: usecases.NewOTPUsecase(nil),
	}
}

func (h *OTPController) RequestOTP(c *fiber.Ctx) error {
	var req domain.OTPRequest
	if err := c.BodyParser(&req); err != nil || req.UserID == "" {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid request payload",
			Success:    false,
		})
	}

	otpData, err := h.useCase.SendOTP(c.Context(), req.UserID)
	if err != nil {
		return c.Status(500).JSON(presenters.Response{
			StatusCode: 500,
			Message:    err.Error(),
			Success:    false,
		})
	}

	return c.JSON(presenters.Response{
		StatusCode: 200,
		Message:    "otp saved to db.",
		Success:    true,
		Data:       otpData,
	})
}

func (h *OTPController) VerifyOTP(c *fiber.Ctx) error {
	var req domain.OTPRequest
	if err := c.BodyParser(&req); err != nil || req.UserID == "" || req.OTP == "" {
		return c.Status(400).JSON(presenters.Response{
			StatusCode: 400,
			Message:    "invalid request payload",
			Success:    false,
		})
	}

	ok, err := h.useCase.VerifyOTP(c.Context(), req.UserID, req.OTP)
	if err != nil || !ok {
		return c.Status(401).JSON(presenters.Response{
			StatusCode: 401,
			Message:    "invalid or expired otp",
			Success:    false,
		})
	}

	return c.JSON(presenters.Response{
		StatusCode: 200,
		Message:    "OTP validated successfully.",
		Success:    true,
	})
}
