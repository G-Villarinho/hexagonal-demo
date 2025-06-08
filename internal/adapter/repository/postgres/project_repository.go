package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/g-villarinho/hexagonal-demo/internal/core/domain"
	"github.com/g-villarinho/hexagonal-demo/internal/core/port"
	"github.com/google/uuid"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) port.ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Save(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	project.ID = uuid.New().String()
	project.CreatedAt = time.Now()

	query := `INSERT INTO projects (id, user_id, title, description, created_at) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING id, user_id, title, description, created_at`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx,
		project.ID,
		project.UserID,
		project.Title,
		project.Description,
		project.CreatedAt,
		project.UpdatedAt,
	).Scan(&project.ID, &project.UserID, &project.Title, &project.Description, &project.CreatedAt, &project.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (r *ProjectRepository) FindByID(ctx context.Context, id string) (*domain.Project, error) {
	project := &domain.Project{}
	query := `SELECT id, user_id, title, description, created_at, updated_at 
			  FROM projects WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, id).Scan(&project.ID, &project.UserID, &project.Title, &project.Description, &project.CreatedAt, &project.UpdatedAt)

	if err != nil {
		return nil, err
	}

	if project.ID == "" {
		return nil, nil
	}

	return project, nil
}
