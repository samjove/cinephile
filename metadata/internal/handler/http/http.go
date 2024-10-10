package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/samjove/cinephile/metadata/internal/controller/metadata"
	"github.com/samjove/cinephile/metadata/internal/repository"
)

// Handler defines a film metadata HTTP handler.
type Handler struct {
	c *metadata.Controller
}

// New creates a new film metadata HTTP handler.
func New(c *metadata.Controller) *Handler {
	return &Handler{c}
}

// GetMetadata handles GET /metadata requests.
func (h *Handler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	m, err := h.c.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}