package service

import (
	"context"
	"errors"

	"booking/internal/domain"
	"booking/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrRoomHasActiveBooking = errors.New("room has an active booking and cannot be deleted")
	ErrRoomNotFound         = repository.ErrRoomNotFound
	ErrBookingNotFound      = repository.ErrBookingNotFound
	ErrRoomAlreadyBooked    = repository.ErrRoomAlreadyBooked
	ErrForbidden            = errors.New("forbidden")
)

type BookingService struct {
	rooms    *repository.RoomRepository
	bookings *repository.BookingRepository
}

func NewBookingService(rooms *repository.RoomRepository, bookings *repository.BookingRepository) *BookingService {
	return &BookingService{rooms: rooms, bookings: bookings}
}

func (s *BookingService) Create(ctx context.Context, userID string, input domain.CreateBookingInput) (*domain.Booking, error) {
	if _, err := s.rooms.GetByID(ctx, input.RoomID); err != nil {
		return nil, err
	}

	booking := &domain.Booking{
		ID:     uuid.New(),
		RoomID: input.RoomID,
		UserID: userID,
		Status: domain.BookingStatusActive,
	}

	if err := s.bookings.Create(ctx, booking); err != nil {
		return nil, err
	}

	return s.bookings.GetByID(ctx, booking.ID)
}

func (s *BookingService) GetByID(ctx context.Context, userID string, id uuid.UUID) (*domain.Booking, error) {
	booking, err := s.bookings.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if booking.UserID != userID {
		return nil, ErrForbidden
	}
	return booking, nil
}
