package repository

import (
	"context"
	"errors"

	"booking/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrRoomNotFound = errors.New("room not found")

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) Create(ctx context.Context, room *domain.Room) error {
	return r.db.WithContext(ctx).Create(room).Error
}

func (r *RoomRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Room, error) {
	var room domain.Room
	err := r.db.WithContext(ctx).First(&room, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRoomNotFound
	}
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) List(ctx context.Context) ([]domain.Room, error) {
	var rooms []domain.Room
	err := r.db.WithContext(ctx).Order("created_at desc").Find(&rooms).Error
	return rooms, err
}

func (r *RoomRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Room{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrRoomNotFound
	}
	return nil
}
