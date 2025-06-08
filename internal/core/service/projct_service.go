package service

import (
	"context"
	"fmt"

	"github.com/g-villarinho/hexagonal-demo/internal/core/domain"
	"github.com/g-villarinho/hexagonal-demo/internal/core/port"
)

type ProjectService struct {
	ProjectRepo port.ProjectRepository
	aiGenerator port.AIGenerator
}

func NewProjectService(
	repo port.ProjectRepository,
	aiGenerator port.AIGenerator) port.ProjectService {
	return &ProjectService{
		ProjectRepo: repo,
		aiGenerator: aiGenerator,
	}
}

func (p *ProjectService) Create(ctx context.Context, title string) (*domain.Project, error) {
	project := &domain.Project{
		Title: title,
	}

	savedProject, err := p.ProjectRepo.Save(ctx, project)
	if err != nil {
		return nil, err
	}

	return savedProject, nil
}

func (p *ProjectService) Get(ctx context.Context, projectID, userID string) (*domain.Project, error) {
	project, err := p.ProjectRepo.FindByID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	if project == nil {
		return nil, domain.ErrNotFound
	}

	if project.UserID != userID {
		return nil, domain.ErrNotFound
	}

	return project, nil
}

func (p *ProjectService) GenerateProjectDescription(ctx context.Context, projectID string, userID string) (*domain.Project, error) {
	project, err := p.ProjectRepo.FindByID(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("find project by id %s: %w", projectID, err)
	}

	if project == nil {
		return nil, domain.ErrNotFound
	}

	if project.UserID != userID {
		return nil, domain.ErrNotFound
	}

	prompt := fmt.Sprintf("Generate a detailed description for the project titled in pt-BR'%s'.", project.Title)

	description, err := p.aiGenerator.GenerateText(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("generate project description: %w", err)
	}

	project.Description = description

	return project, nil
}
