package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/g-villarinho/hexagonal-demo/internal/core/domain"
	"github.com/g-villarinho/hexagonal-demo/internal/core/port"
	"github.com/google/uuid"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) port.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()

	stmt, err := u.db.Prepare(`
		INSERT INTO users (id, name, email, password, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.QueryRowContext(ctx, `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE id = $1
	`, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.QueryRowContext(ctx, `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE email = $1
	`, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
