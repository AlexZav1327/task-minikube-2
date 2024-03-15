package server

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	service AccessDataService
	log     *logrus.Entry
}

type AccessDataService interface {
	SaveAccessData(ctx context.Context, ipAddr string) error
}

func NewHandler(service AccessDataService, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log.WithField("module", "handler"),
	}
}

func (h *Handler) saveAccessData(w http.ResponseWriter, r *http.Request) {
	ipAddr := r.RemoteAddr

	if err := h.service.SaveAccessData(r.Context(), ipAddr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
