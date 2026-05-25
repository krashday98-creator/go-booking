package handler

import (
	"errors"
	"net/http"

	"booking/internal/domain"
	"booking/internal/repository"
	"booking/internal/service"

	"github.com/gin-gonic/gin"
)

// ErrorResponse описывает тело ошибки API.
type ErrorResponse struct {
	Error string `json:"error" example:"room not found"`
}

// RoomListResponse список комнат.
type RoomListResponse struct {
	Rooms []domain.Room `json:"rooms"`
}

func writeError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidRoomClass):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, repository.ErrRoomNotFound),
		errors.Is(err, service.ErrRoomNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, repository.ErrBookingNotFound),
		errors.Is(err, service.ErrBookingNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, repository.ErrRoomAlreadyBooked),
		errors.Is(err, service.ErrRoomAlreadyBooked):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, service.ErrRoomHasActiveBooking):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, service.ErrForbidden):
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
