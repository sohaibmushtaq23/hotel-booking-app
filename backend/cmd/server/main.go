package main

import (
	"hotel-booking-backend/internal/config"
	"hotel-booking-backend/internal/database"
	"hotel-booking-backend/internal/handlers"
	"hotel-booking-backend/internal/repository"
	"hotel-booking-backend/internal/router"
	"hotel-booking-backend/internal/service"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	connString := config.GetConnectionString()
	database.Connect(connString)

	roomRepo := repository.NewRoomRepository(database.DB)
	roomService := service.NewRoomService(roomRepo)
	roomHandler := handlers.NewRoomHandler(roomService)

	r := router.NewRouter(roomHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
