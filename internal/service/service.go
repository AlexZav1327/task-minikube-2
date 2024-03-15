package service

import (
	"context"
	"fmt"
	"time"

	"github.com/AlexZav1327/task-minikube-2/models"
	"github.com/sirupsen/logrus"
)

type Service struct {
	pg  store
	log *logrus.Entry
}

func New(pg store, log *logrus.Logger) *Service {
	return &Service{
		pg:  pg,
		log: log.WithField("module", "service"),
	}
}

type store interface {
	SaveAccessData(ctx context.Context, data models.AccessData) error
}

func (s *Service) SaveAccessData(ctx context.Context, ipAddr string) error {
	data := models.AccessData{
		Time: getCurrentTime(),
		IP:   ipAddr,
	}

	err := s.pg.SaveAccessData(ctx, data)
	if err != nil {
		return fmt.Errorf("s.pg.SaveAccessData(ctx, data): %w", err)
	}

	return nil
}

func getCurrentTime() string {
	now := time.Now()

	return now.Format("2006-01-02 15:04:05")
}
