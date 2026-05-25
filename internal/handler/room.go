package handler

import (
	"net/http"

	"booking/internal/domain"
	"booking/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoomHandler struct {
	rooms *service.RoomService
}

func NewRoomHandler(rooms *service.RoomService) *RoomHandler {
	return &RoomHandler{rooms: rooms}
}

// Create создаёт комнату.
//
// @Summary      Создать комнату
// @Tags         admin-rooms
// @Accept       json
// @Produce      json
// @Param        body body domain.CreateRoomInput true "Данные комнаты"
// @Success      201 {object} domain.Room
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Security     AdminKey
// @Router       /api/admin/rooms [post]
func (h *RoomHandler) Create(c *gin.Context) {
	var input domain.CreateRoomInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := h.rooms.Create(c.Request.Context(), input)
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusCreated, room)
}

// Get возвращает комнату по ID.
//
// @Summary      Получить комнату
// @Tags         admin-rooms
// @Produce      json
// @Param        id path string true "ID комнаты (UUID)"
// @Success      200 {object} domain.Room
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Security     AdminKey
// @Router       /api/admin/rooms/{id} [get]
func (h *RoomHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room id"})
		return
	}

	room, err := h.rooms.GetByID(c.Request.Context(), id)
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, room)
}

// List возвращает список комнат.
//
// @Summary      Список комнат
// @Tags         admin-rooms
// @Produce      json
// @Success      200 {object} RoomListResponse
// @Failure      401 {object} ErrorResponse
// @Security     AdminKey
// @Router       /api/admin/rooms [get]
func (h *RoomHandler) List(c *gin.Context) {
	rooms, err := h.rooms.List(c.Request.Context())
	if err != nil {
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, RoomListResponse{Rooms: rooms})
}

// Delete удаляет комнату.
//
// @Summary      Удалить комнату
// @Tags         admin-rooms
// @Param        id path string true "ID комнаты (UUID)"
// @Success      204
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Failure      409 {object} ErrorResponse
// @Security     AdminKey
// @Router       /api/admin/rooms/{id} [delete]
func (h *RoomHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room id"})
		return
	}

	if err := h.rooms.Delete(c.Request.Context(), id); err != nil {
		writeError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
