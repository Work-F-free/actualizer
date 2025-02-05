package dto

import (
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	SeatId       uuid.UUID `json:"seat_id" db:"seat_id"`
	Expired      bool      `json:"expired" db:"expired"`
	FromDateTime time.Time `json:"from_datetime" db:"from_datetime"`
	ToDateTime   time.Time `json:"to_datetime" db:"to_datetime"`
}
