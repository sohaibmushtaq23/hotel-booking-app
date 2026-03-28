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

type RoomReservationHandler struct {
	service *service.RoomReservationService
}

func NewRoomReservationHandler(service *service.RoomReservationService) *RoomReservationHandler {
	return &RoomReservationHandler{service: service}
}

func (h *RoomReservationHandler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var reservation models.RoomReservation
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		writeError(w, errors.New("invalid request body"))
		return
	}

	if err := h.service.Create(r.Context(), &reservation); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reservation)
}

func (h *RoomReservationHandler) GetReservationDetail(w http.ResponseWriter, r *http.Request) {
	idReservation, err := strconv.Atoi(chi.URLParam(r, "idReservation"))
	if err != nil {
		writeError(w, errors.New("invalid reservation id"))
		return
	}

	reservations, err := h.service.GetByIDReservation(r.Context(), idReservation)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if reservations == nil {
		reservations = []models.RoomReservation{}
	}

	json.NewEncoder(w).Encode(reservations)
}

func (h *RoomReservationHandler) GetRoomReservations(w http.ResponseWriter, r *http.Request) {
	idRoom, err := strconv.Atoi(chi.URLParam(r, "idRoom"))
	if err != nil {
		writeError(w, errors.New("invalid reservation id"))
		return
	}

	reservations, err := h.service.GetByIDRoom(r.Context(), idRoom)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if reservations == nil {
		reservations = []models.RoomReservation{}
	}

	json.NewEncoder(w).Encode(reservations)
}

func (h *RoomReservationHandler) GetReservation(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, errors.New("invalid id"))
		return
	}

	reservation, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, err)
		return
	}

	json.NewEncoder(w).Encode(reservation)
}

func (h *RoomReservationHandler) UpdateReservation(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, errors.New("invalid id"))
		return
	}

	var reservation models.RoomReservation
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		writeError(w, errors.New("invalid request body"))
		return
	}

	updatedReservation, err := h.service.Update(r.Context(), id, &reservation)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedReservation)
}

func (h *RoomReservationHandler) DeleteReservation(w http.ResponseWriter, r *http.Request) {
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
