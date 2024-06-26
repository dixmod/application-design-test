package room

import (
	"applicationDesignTest/pkg/tools"
	"time"
)

type Repository interface {
	GetAvailability() *[]RoomAvailability
	Update(id int, room RoomAvailability)
}

type RoomRepository struct {
	data *[]RoomAvailability
}

func NewRepository() Repository {
	return &RoomRepository{
		// TODO: fake data
		&[]RoomAvailability{
			{"reddison", "lux", tools.Date(2024, 1, 1), 1},
			{"reddison", "lux", tools.Date(2024, 1, 2), 1},
			{"reddison", "lux", tools.Date(2024, 1, 3), 1},
			{"reddison", "lux", tools.Date(2024, 1, 4), 1},
			{"reddison", "lux", tools.Date(2024, 1, 5), 0},
		},
	}
}

func (r RoomRepository) GetAvailability() *[]RoomAvailability {
	return r.data
}

func (r RoomRepository) Update(id int, room RoomAvailability) {
	(*r.data)[id] = room
}

type RoomAvailability struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}
