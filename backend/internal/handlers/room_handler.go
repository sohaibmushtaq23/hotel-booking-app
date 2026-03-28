package handlers

import (
	"encoding/json"
	"errors"
	"hotel-booking-backend/internal/models"
	"hotel-booking-backend/internal/repository"
	"hotel-booking-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoomHandler struct {
	service *service.RoomService
}

func NewRoomHandler(service *service.RoomService) *RoomHandler {
	return &RoomHandler{service: service}
}

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case errors.Is(err, repository.ErrRoomNotFound), errors.Is(err, repository.ErrClientNotFound),
		errors.Is(err, repository.ErrUserNotFound):
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		writeError(w, errors.New("invalid request body"))
		return
	}

	if err := h.service.Create(r.Context(), &room); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}

func (h *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.service.GetAll(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if rooms == nil {
		rooms = []models.Room{}
	}

	json.NewEncoder(w).Encode(rooms)
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, errors.New("invalid id"))
		return
	}

	room, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, err)
		return
	}

	json.NewEncoder(w).Encode(room)
}

func (h *RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, errors.New("invalid id"))
		return
	}

	var room models.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		writeError(w, errors.New("invalid request body"))
		return
	}

	updatedRoom, err := h.service.Update(r.Context(), id, &room)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedRoom)
}

func (h *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
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
