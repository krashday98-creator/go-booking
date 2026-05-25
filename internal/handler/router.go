package handler

import (
	"booking/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterDeps struct {
	RoomHandler    *RoomHandler
	BookingHandler *BookingHandler
	AdminAPIKey    string
}

func NewRouter(deps RouterDeps) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	r.GET("/health", Health)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	admin := r.Group("/api/admin", middleware.AdminAuth(deps.AdminAPIKey))
	{
		admin.POST("/rooms", deps.RoomHandler.Create)
		admin.GET("/rooms", deps.RoomHandler.List)
		admin.GET("/rooms/:id", deps.RoomHandler.Get)
		admin.DELETE("/rooms/:id", deps.RoomHandler.Delete)
	}

	user := r.Group("/api", middleware.UserAuth())
	{
		user.POST("/bookings", deps.BookingHandler.Create)
		user.GET("/bookings/:id", deps.BookingHandler.Get)
	}

	return r
}
