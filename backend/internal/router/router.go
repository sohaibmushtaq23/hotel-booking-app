package router

import (
	"hotel-booking-backend/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(
	roomHandler *handlers.RoomHandler,
	clientHandler *handlers.ClientHandler,
	userHandler *handlers.UserHandler,
	reservationHandler *handlers.ReservationHandler,
) *chi.Mux {
	r := chi.NewRouter()

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // allow all origins (dev only)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))

	fileServer := http.FileServer(http.Dir("./uploads"))
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", fileServer))

	r.Post("/upload", handlers.UploadRoomImage)

	// Room routes
	r.Route("/rooms", func(r chi.Router) {
		r.Get("/", roomHandler.GetRooms)
		r.Post("/", roomHandler.CreateRoom)
		r.Get("/{id}", roomHandler.GetRoom)
		r.Put("/{id}", roomHandler.UpdateRoom)
		r.Delete("/{id}", roomHandler.DeleteRoom)

		// Room reservation detail routes
		r.Route("/{idRoom}/details", func(r chi.Router) {
			r.Get("/", reservationHandler.GetRoomReservations)
		})
	})

	// Client routes
	r.Route("/clients", func(r chi.Router) {
		r.Get("/", clientHandler.GetClients)
		r.Post("/", clientHandler.CreateClient)
		r.Get("/{id}", clientHandler.GetClient)
		r.Put("/{id}", clientHandler.UpdateClient)
		r.Delete("/{id}", clientHandler.DeleteClient)

		// Client reservation detail routes
		r.Route("/{idClient}/details", func(r chi.Router) {
			r.Get("/", reservationHandler.GetCustomerReservations)
		})
	})

	// User routes
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.GetUser)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})

	// Reservation routes
	r.Route("/reservations", func(r chi.Router) {
		r.Get("/", reservationHandler.GetBookingsWithDetails)
		r.Post("/", reservationHandler.CreateReservation)
		r.Get("/{id}", reservationHandler.GetReservation)
		r.Put("/{id}", reservationHandler.UpdateReservation)
		r.Delete("/{id}", reservationHandler.DeleteReservation)

	})

	return r
}
