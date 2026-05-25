package domain

import (
	"time"

	"github.com/google/uuid"
)

type BookingStatus string

const (
	BookingStatusActive    BookingStatus = "active"
	BookingStatusCancelled BookingStatus = "cancelled"
)

type Booking struct {
	ID        uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	RoomID    uuid.UUID     `gorm:"type:uuid;not null;index" json:"room_id"`
	UserID    string        `gorm:"type:varchar(128);not null;index" json:"user_id"`
	Status    BookingStatus `gorm:"type:varchar(32);not null;default:active" json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`

	Room *Room `gorm:"foreignKey:RoomID" json:"room,omitempty"`
}

type CreateBookingInput struct {
	RoomID uuid.UUID `json:"room_id" binding:"required"`
}
