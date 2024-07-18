package http

import (
	"net/http"

	"github.com/fandyputram/go-project-template/internal/usecase"
)

type Handler struct {
	usecase usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{usecase: uc}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handle HTTP requests
}
