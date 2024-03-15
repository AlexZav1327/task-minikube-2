package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	host    string
	port    int
	server  *http.Server
	service AccessDataService
	log     *logrus.Entry
}

func New(host string, port int, service AccessDataService, log *logrus.Logger) *Server {
	h := NewHandler(service, log)

	s := Server{
		host:    host,
		port:    port,
		service: service,
		log:     log.WithField("module", "http"),
	}

	r := http.NewServeMux()
	r.HandleFunc("/api/v1/log", h.saveAccessData)

	s.server = &http.Server{
		Addr:              fmt.Sprintf("%s:%d", host, port),
		Handler:           r,
		ReadHeaderTimeout: 30 * time.Second,
	}

	return &s
}

func (s *Server) Run(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	defer s.log.Info("Server is stopped")

	go func() {
		<-ctx.Done()

		if err := s.server.Shutdown(shutdownCtx); err != nil {
			s.log.Warningf("s.server.Shutdown(shutdownCtx): %s", err)
		}
	}()

	s.log.Infof("Server is running at port %d", s.port)

	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("s.server.ListenAndServe(): %w", err)
	}

	return nil
}
