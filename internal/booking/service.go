package booking

import (
	"applicationDesignTest/pkg/order"
	"applicationDesignTest/pkg/room"
	"applicationDesignTest/pkg/tools"
	"fmt"
	"time"
)

const limitQuota = 1

type Service interface {
	Booking(newOrder order.Order) error
}

type ServiceBooking struct {
	room  room.Repository
	order order.Repository
}

func NewService(roomRepository room.Repository, orderRepository order.Repository) Service {
	return &ServiceBooking{
		roomRepository,
		orderRepository,
	}
}

func (s ServiceBooking) Booking(newOrder order.Order) error {
	daysToBook := tools.DaysBetween(newOrder.From, newOrder.To)
	unavailableDays := s.unavailableDays(daysToBook)

	for _, dayToBook := range daysToBook {
		for id, availableRoom := range *s.room.GetAvailability() {
			if !s.isAvailableRoom(availableRoom, newOrder) {
				continue
			}

			if !s.isAvailableDate(availableRoom, dayToBook) {
				continue
			}

			availableRoom.Quota -= limitQuota
			s.room.Update(id, availableRoom)

			delete(unavailableDays, dayToBook)
		}
	}

	s.order.Save(newOrder)

	if len(unavailableDays) != 0 {
		return fmt.Errorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
	}

	return nil
}

func (s ServiceBooking) unavailableDays(daysToBook []time.Time) map[time.Time]struct{} {
	unavailableDays := make(map[time.Time]struct{})

	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	return unavailableDays
}

func (s ServiceBooking) isAvailableRoom(availableRoom room.RoomAvailability, order order.Order) bool {
	if availableRoom.HotelID != order.HotelID {
		return false
	}

	if availableRoom.RoomID != order.RoomID {
		return false
	}

	return true
}

func (s ServiceBooking) isAvailableDate(availability room.RoomAvailability, dayToBook time.Time) bool {
	return !availability.Date.Equal(dayToBook) || availability.Quota < limitQuota
}
