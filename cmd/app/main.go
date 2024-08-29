package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/labstack/gommon/log"
	"github.com/qara-qurt/booking_service/config"
	"github.com/qara-qurt/booking_service/internal/handler"
	"github.com/qara-qurt/booking_service/internal/repository"
	"github.com/qara-qurt/booking_service/internal/service"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	repo, err := repository.New(conf)
	if err != nil {
		log.Fatalf("failed to init repository: %v", err)
	}
	service := service.New(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	handler := handler.New(service, r)
	handler.RegisterRoutes()

	log.Printf("starting server on port %s", conf.Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Server.Port), r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
