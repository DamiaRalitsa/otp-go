package route

import (
	"encoding/json"
	"os"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"

	"sqe/internal/delivery/http"
)

type RouteConfig struct {
	App           *fiber.App
	otpController *http.OTPController
}

func NewRouteConfig() *RouteConfig {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	app := fiber.New(fiber.Config{
		Prefork:     false,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   100 * 1024 * 1024,
	})

	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/metrics" {
			return c.Next()
		}
		return fiberzerolog.New(fiberzerolog.Config{
			Logger: &logger,
		})(c)
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	OTPController := http.NewOTPController()

	routeConfig := RouteConfig{
		App:           app,
		otpController: OTPController,
	}

	routeConfig.SetupRoute()
	return &routeConfig
}

func (rc *RouteConfig) SetupRoute() {
	testGroup := rc.App.Group("/api/test")
	testGroup.Post("/request-otp", rc.otpController.RequestOTP)
	testGroup.Post("/verify-otp", rc.otpController.VerifyOTP)
}

func (rc *RouteConfig) Listen(address string) {
	rc.App.Listen(address)
}
