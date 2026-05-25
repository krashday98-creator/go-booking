// @title           Booking API
// @version         1.0
// @description     REST API для управления комнатами и бронированиями.
// @host            localhost:8080
// @BasePath        /
//
// @securityDefinitions.apikey AdminKey
// @in              header
// @name            X-Admin-Key
//
// @securityDefinitions.apikey UserID
// @in              header
// @name            X-User-ID
package main
import (
	"fmt"
	"log"

	_ "booking/docs"

	"booking/internal/config"
	"booking/internal/handler"
	"booking/internal/repository"
	"booking/internal/service"
	"booking/pkg/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	db, err := database.Connect(cfg.DatabaseURL, cfg.GinMode != "release")
	if err != nil {
		log.Fatalf("database: %v", err)
	}

	roomRepo := repository.NewRoomRepository(db)
	bookingRepo := repository.NewBookingRepository(db)

	roomService := service.NewRoomService(roomRepo, bookingRepo)
	bookingService := service.NewBookingService(roomRepo, bookingRepo)

	router := handler.NewRouter(handler.RouterDeps{
		RoomHandler:    handler.NewRoomHandler(roomService),
		BookingHandler: handler.NewBookingHandler(bookingService),
		AdminAPIKey:    cfg.AdminAPIKey,
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("server listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("server: %v", err)
	}
}
