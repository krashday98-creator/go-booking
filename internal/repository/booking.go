package repository

import (
	"context"
	"errors"

	"booking/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrBookingNotFound = errors.New("booking not found")
	ErrRoomAlreadyBooked = errors.New("room is already booked")
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

// Create вставляет бронь; в транзакции проверяем, что нет другой active на эту комнату.
func (r *BookingRepository) Create(ctx context.Context, booking *domain.Booking) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.Model(&domain.Booking{}).
			Where("room_id = ? AND status = ?", booking.RoomID, domain.BookingStatusActive).
			Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return ErrRoomAlreadyBooked
		}
		return tx.Create(booking).Error
	})
}

func (r *BookingRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Booking, error) {
	var booking domain.Booking
	err := r.db.WithContext(ctx).
		Preload("Room").
		First(&booking, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrBookingNotFound
	}
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *BookingRepository) HasActiveBooking(ctx context.Context, roomID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&domain.Booking{}).
		Where("room_id = ? AND status = ?", roomID, domain.BookingStatusActive).
		Count(&count).Error
	return count > 0, err
}
