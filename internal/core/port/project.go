package port

import (
	"context"

	"github.com/g-villarinho/hexagonal-demo/internal/core/domain"
)

type ProjectRepository interface {
	Save(ctx context.Context, project *domain.Project) (*domain.Project, error)
	FindByID(ctx context.Context, id string) (*domain.Project, error)
}

type ProjectService interface {
	Create(ctx context.Context, name string) (*domain.Project, error)
	Get(ctx context.Context, projectID, userID string) (*domain.Project, error)
	GenerateProjectDescription(ctx context.Context, projectID, userID string) (*domain.Project, error)
}
