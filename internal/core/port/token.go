package port

import "context"

type TokenMaker interface {
	Generate(ctx context.Context, userID string) (string, error)
	Verify(ctx context.Context, token string) (string, error)
}
