package order

import (
	"time"
)

type Repository interface {
	Save(order Order)
}

type OrderRepository struct {
	data *[]Order
}

func NewRepository() Repository {
	return &OrderRepository{
		data: &[]Order{},
	}
}

func (r OrderRepository) Save(order Order) {
	*r.data = append(*r.data, order)
}

type Order struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}
