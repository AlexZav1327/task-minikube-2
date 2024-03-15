package storage

import (
	"context"
	"fmt"

	"github.com/AlexZav1327/task-minikube-2/models"
)

const (
	saveAccessData = `
	INSERT INTO access_data (time, ip)
	VALUES ($1, $2);
	`
)

func (p *Postgres) SaveAccessData(ctx context.Context, data models.AccessData) error {
	_, err := p.db.Exec(ctx, saveAccessData, data.Time, data.IP)
	if err != nil {
		return fmt.Errorf("p.db.Exec(ctx, saveAccessData, data.Time, data.IP): %w", err)
	}

	return nil
}
