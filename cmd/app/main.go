package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	route "sqe/internal/delivery/http/routes"
	postgres "sqe/pkg/postgres"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic().Msg("Config file 'config.json' is required")
	}

	if viper.GetBool("debug") {
		log.Info().Msg("Service RUN on DEBUG mode")
	}

	postgres.InitConnection()
}

func main() {
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "8335"
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	cfg := route.NewRouteConfig()

	log.Info().Msgf("Server started at http://%s", addr)

	if err := cfg.App.Listen(addr); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
