package service

import (
	"context"

	"booking/internal/domain"
	"booking/internal/repository"

	"github.com/google/uuid"
)

type RoomService struct {
	rooms *repository.RoomRepository
	bookings *repository.BookingRepository
}

func NewRoomService(rooms *repository.RoomRepository, bookings *repository.BookingRepository) *RoomService {
	return &RoomService{rooms: rooms, bookings: bookings}
}

func (s *RoomService) Create(ctx context.Context, input domain.CreateRoomInput) (*domain.Room, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	room := &domain.Room{
		ID:          uuid.New(),
		Class:       input.Class,
		Cost:        input.Cost,
		Description: input.Description,
	}
	if err := s.rooms.Create(ctx, room); err != nil {
		return nil, err
	}
	return room, nil
}

func (s *RoomService) GetByID(ctx context.Context, id uuid.UUID) (*domain.Room, error) {
	return s.rooms.GetByID(ctx, id)
}

func (s *RoomService) List(ctx context.Context) ([]domain.Room, error) {
	return s.rooms.List(ctx)
}

func (s *RoomService) Delete(ctx context.Context, id uuid.UUID) error {
	hasBooking, err := s.bookings.HasActiveBooking(ctx, id)
	if err != nil {
		return err
	}
	if hasBooking {
		return ErrRoomHasActiveBooking
	}
	return s.rooms.Delete(ctx, id)
}
