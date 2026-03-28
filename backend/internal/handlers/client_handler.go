package handlers

import (
	"encoding/json"
	"errors"
	"hotel-booking-backend/internal/models"
	"hotel-booking-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ClientHandler struct {
	service *service.ClientService
}

func NewClientHandler(service *service.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		writeError(w, errors.New("invalid request body"))
		return
	}

	if err := h.service.Create(r.Context(), &client); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func (h *ClientHandler) GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.service.GetAll(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if clients == nil {
		clients = []models.Client{}
	}

	json.NewEncoder(w).Encode(clients)
}

func (h *ClientHandler) GetClient(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, errors.New("invalid id"))
		return
	}

	client, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, err)
		return
	}

	json.NewEncoder(w).Encode(client)
}

func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, errors.New("invalid id"))
		return
	}

	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		writeError(w, errors.New("invalid request body"))
		return
	}

	updatedClient, err := h.service.Update(r.Context(), id, &client)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedClient)
}

func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, errors.New("invalid id"))
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
