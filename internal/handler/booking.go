package handler

import (
	"net/http"

	"booking/internal/domain"
	"booking/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookingHandler struct {
	bookings *service.BookingService
}

func NewBookingHandler(bookings *service.BookingService) *BookingHandler {
	return &BookingHandler{bookings: bookings}
}

// Create создаёт бронирование.
//
// @Summary      Забронировать комнату
// @Tags         bookings
// @Accept       json
// @Produce      json
// @Param        body body domain.CreateBookingInput true "ID комнаты"
// @Success      201 {object} domain.Booking
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Failure      409 {object} ErrorResponse
// @Security     UserID
// @Router       /api/bookings [post]
func (h *BookingHandler) Create(c *gin.Context) {
	var input domain.CreateBookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("user_id")
	booking, err := h.bookings.Create(c.Request.Context(), userID, input)
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusCreated, booking)
}

// Get возвращает бронирование по ID (только своё).
//
// @Summary      Получить бронирование
// @Tags         bookings
// @Produce      json
// @Param        id path string true "ID бронирования (UUID)"
// @Success      200 {object} domain.Booking
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      403 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Security     UserID
// @Router       /api/bookings/{id} [get]
func (h *BookingHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking id"})
		return
	}

	userID := c.GetString("user_id")
	booking, err := h.bookings.GetByID(c.Request.Context(), userID, id)
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, booking)
}
