package service

import (
	"context"
	"fmt"

	"github.com/g-villarinho/hexagonal-demo/internal/core/domain"
	"github.com/g-villarinho/hexagonal-demo/internal/core/port"
)

type userService struct {
	userRepo   port.UserRepository
	tokenMaker port.TokenMaker
}

func NewUserService(repo port.UserRepository, tokenMaker port.TokenMaker) port.UserService {
	return &userService{
		userRepo:   repo,
		tokenMaker: tokenMaker,
	}
}

func (u *userService) Create(ctx context.Context, name string, email string, password string) (*domain.User, error) {
	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	userFromEmail, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("checking existing user by email: %w", err)
	}

	if userFromEmail != nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	savedUser, err := u.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}

func (u *userService) Get(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) Login(ctx context.Context, email string, password string) (token string, err error) {
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("finding user by email: %w", err)
	}

	if user == nil {
		return "", domain.ErrInvalidCredentials
	}

	if !user.CheckPassword(password) {
		return "", domain.ErrInvalidCredentials
	}

	token, err = u.tokenMaker.Generate(ctx, user.ID)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return token, nil
}
