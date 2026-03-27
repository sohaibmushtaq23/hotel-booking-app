package main

import (
	"clientmanager/internal/config"
	"clientmanager/internal/database"
	"clientmanager/internal/handlers"
	"clientmanager/internal/repository"
	"clientmanager/internal/router"
	"clientmanager/internal/service"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	connString := config.GetConnectionString()
	database.Connect(connString)

	clientRepo := repository.NewClientRepository(database.DB)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handlers.NewClientHandler(clientService)

	contactRepo := repository.NewClientContactRepository(database.DB)
	contactService := service.NewClientContactService(contactRepo)
	contactHandler := handlers.NewClientContactHandler(contactService)

	r := router.NewRouter(clientHandler, contactHandler)

	log.Println("Server running on :8080")
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
