package main

import (
	"fmt"
	"log"
	"net/http"

	handler "bookingtogo/internal/delivery/http"
	route "bookingtogo/internal/delivery/http/routes"
	usecase "bookingtogo/internal/usecases"
	"bookingtogo/pkg/postgres"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Config file 'config.json' is required")
	}

	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")
	}

	postgres.InitConnection()
}

func main() {
	// Load host and port from config
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "8334"
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	// Initialize usecases
	nationalityUC := usecase.NewNationalityUsecase()
	customerUC := usecase.NewCustomerUsecase()
	familyUC := usecase.NewFamilyUsecase()

	// Initialize handlers with dependencies
	nationalityHandler := handler.NewNationalityHandler(nationalityUC)
	customerHandler := handler.NewCustomerHandler(customerUC)
	familyHandler := handler.NewFamilyHandler(familyUC)

	// Initialize router
	router := route.NewRouter(customerHandler, familyHandler, nationalityHandler)

	log.Printf("Server started at http://%s\n", addr)

	// Start HTTP server with mux
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
