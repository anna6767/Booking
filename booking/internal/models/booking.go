package models

import "time"

type Booking struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	RoomID    uint      `json:"room_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
}

type BookingRequest struct {
	UserID    uint      `json:"user_id"`
	RoomID    uint      `json:"room_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
