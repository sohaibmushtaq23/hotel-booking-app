package router

import (
	"clientmanager/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(
	clientHandler *handlers.ClientHandler,
	contactHandler *handlers.ClientContactHandler,
) *chi.Mux {
	r := chi.NewRouter()

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // allow all origins (dev only)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))

	r.Route("/clients", func(r chi.Router) {
		// Client routes
		r.Get("/", clientHandler.GetClients)
		r.Post("/", clientHandler.CreateClient)
		r.Get("/{id}", clientHandler.GetClient)
		r.Put("/{id}", clientHandler.UpdateClient)
		r.Delete("/{id}", clientHandler.DeleteClient)

		// Nested contact routes – /clients/{clientId}/contacts
		r.Route("/{clientId}/contacts", func(r chi.Router) {
			r.Get("/", contactHandler.GetContactsByClientID)
			r.Post("/", contactHandler.CreateContact)
			r.Get("/{id}", contactHandler.GetContact)
			r.Put("/{id}", contactHandler.UpdateContact)
			r.Delete("/{id}", contactHandler.DeleteContact)
		})
	})

	return r
}
