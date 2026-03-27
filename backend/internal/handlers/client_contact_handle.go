package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"clientmanager/internal/models"
	"clientmanager/internal/service"

	"github.com/go-chi/chi/v5"
)

type ClientContactHandler struct {
	service *service.ClientContactService
}

func NewClientContactHandler(service *service.ClientContactService) *ClientContactHandler {
	return &ClientContactHandler{service: service}
}

// GetContactsByClientID handles GET /clients/{clientId}/contacts
func (h *ClientContactHandler) GetContactsByClientID(w http.ResponseWriter, r *http.Request) {
	clientID, err := strconv.Atoi(chi.URLParam(r, "clientId"))
	if err != nil {
		writeError(w, errors.New("invalid client id"))
		return
	}

	contacts, err := h.service.GetByClientID(r.Context(), clientID)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if contacts == nil {
		contacts = []models.ClientContact{}
	}

	json.NewEncoder(w).Encode(contacts)
}

// CreateContact handles POST /clients/{clientId}/contacts
func (h *ClientContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	clientID, err := strconv.Atoi(chi.URLParam(r, "clientId"))
	if err != nil {
		writeError(w, errors.New("invalid client id"))
		return
	}

	var contact models.ClientContact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		writeError(w, errors.New("invalid request body"))
		return
	}

	// Set the client ID from the URL into the contact struct
	contact.IDClient = clientID

	if err := h.service.Create(r.Context(), &contact); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}

// GetContact handles GET /clients/{clientId}/contacts/{id}
func (h *ClientContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	contactID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, errors.New("invalid contact id"))
		return
	}

	contact, err := h.service.GetByID(r.Context(), contactID)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contact)
}

func (h *ClientContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	contactID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, errors.New("invalid contact id"))
		return
	}

	clientID, _ := strconv.Atoi(chi.URLParam(r, "clientId"))

	var contact models.ClientContact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		writeError(w, errors.New("invalid request body"))
		return
	}

	if contact.IDClient != 0 && contact.IDClient != clientID {
		writeError(w, errors.New("client id mismatch"))
		return
	}

	if err := h.service.Update(r.Context(), contactID, &contact); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contact)
}

func (h *ClientContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	contactID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, errors.New("invalid contact id"))
		return
	}

	if err := h.service.Delete(r.Context(), contactID); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
