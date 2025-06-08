package domain

import (
	"database/sql"
	"time"
)

type Project struct {
	ID          string
	UserID      string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}
