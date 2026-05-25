package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	RoomClassStandard = "standard"
	RoomClassDeluxe   = "deluxe"
)

var ErrInvalidRoomClass = errors.New("invalid room class")
type Room struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Class       string    `gorm:"type:varchar(64);not null" json:"class"`
	Cost        float64   `gorm:"type:decimal(12,2);not null" json:"cost"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateRoomInput struct {
	Class       string  `json:"class" binding:"required"`
	Cost        float64 `json:"cost" binding:"required,gt=0"`
	Description string  `json:"description"`
}

func (in CreateRoomInput) Validate() error {
	switch in.Class {
	case RoomClassStandard, RoomClassDeluxe:
		return nil
	default:
		return fmt.Errorf("%w: %q (allowed: %s, %s)",
			ErrInvalidRoomClass, in.Class, RoomClassStandard, RoomClassDeluxe)
	}
}