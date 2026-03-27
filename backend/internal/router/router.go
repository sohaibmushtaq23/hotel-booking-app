package router

import (
	"hotel-booking-backend/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(
	roomHandler *handlers.RoomHandler,
) *chi.Mux {
	r := chi.NewRouter()

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // allow all origins (dev only)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))

	r.Route("/Rooms", func(r chi.Router) {
		// Room routes
		r.Get("/", roomHandler.GetRooms)
		r.Post("/", roomHandler.CreateRoom)
		r.Get("/{id}", roomHandler.GetRoom)
		r.Put("/{id}", roomHandler.UpdateRoom)
		r.Delete("/{id}", roomHandler.DeleteRoom)
	})

	return r
}
