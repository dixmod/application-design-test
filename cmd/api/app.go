package main

import (
	"applicationDesignTest/internal/api"
	"applicationDesignTest/internal/booking"
	"applicationDesignTest/pkg/log"
	"applicationDesignTest/pkg/order"
	"applicationDesignTest/pkg/room"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type App struct {
	cfg    *api.Config
	router *http.ServeMux
	logger log.LocalLogger
}

func NewApp(cfg *api.Config) *App {
	logger := log.NewLogger()

	bookingService := booking.NewService(room.NewRepository(), order.NewRepository())
	apiHandler := api.NewController(bookingService, logger)

	router := http.NewServeMux()
	router.HandleFunc("/orders", apiHandler.CreateOrder)

	return &App{
		cfg:    cfg,
		router: router,
		logger: logger,
	}
}

func (app App) Run() {
	err := http.ListenAndServe(fmt.Sprintf(":%s", app.cfg.Port), app.router)
	app.logger.LogInfo(fmt.Sprintf("Server listening on localhost:%s", app.cfg.Port))

	if errors.Is(err, http.ErrServerClosed) {
		app.logger.LogInfo("Server closed")
	} else if err != nil {
		app.logger.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
