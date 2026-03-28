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

type ReservationHandler struct {
	service *service.ReservationService
}

func NewReservationHandler(service *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{service: service}
}

func (h *ReservationHandler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var reservation models.Reservation
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

func (h *ReservationHandler) GetReservations(w http.ResponseWriter, r *http.Request) {
	reservations, err := h.service.GetAll(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if reservations == nil {
		reservations = []models.Reservation{}
	}

	json.NewEncoder(w).Encode(reservations)
}

func (h *ReservationHandler) GetReservation(w http.ResponseWriter, r *http.Request) {
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

func (h *ReservationHandler) UpdateReservation(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, errors.New("invalid id"))
		return
	}

	var reservation models.Reservation
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

func (h *ReservationHandler) DeleteReservation(w http.ResponseWriter, r *http.Request) {
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
