package port

import (
	"context"

	"github.com/g-villarinho/hexagonal-demo/internal/core/domain"
)

type UserService interface {
	Create(ctx context.Context, name, email, password string) (*domain.User, error)
	Get(ctx context.Context, id string) (*domain.User, error)
	Login(ctx context.Context, email, password string) (token string, err error)
}

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
