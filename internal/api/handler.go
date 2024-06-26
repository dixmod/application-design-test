package api

import (
	"applicationDesignTest/internal/booking"
	"applicationDesignTest/pkg/log"
	"applicationDesignTest/pkg/order"
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller struct {
	service booking.Service
	logger  log.LocalLogger
}

func NewController(serviceBooking booking.Service, logger log.LocalLogger) *Controller {
	return &Controller{
		service: serviceBooking,
		logger:  logger,
	}
}

func (c Controller) CreateOrder(w http.ResponseWriter, request *http.Request) {
	newOrder, err := order.Fabric{}.CreateFromRequest(request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Validation error order: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = c.service.Booking(newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		c.logger.LogErrorf(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newOrder)

	if err != nil {
		c.logger.LogErrorf("Order error encoder: %v", newOrder)
		return
	}

	c.logger.LogInfo("Order successfully created: %v", newOrder)
}
