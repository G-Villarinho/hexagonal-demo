package port

import "context"

type AIGenerator interface {
	GenerateText(ctx context.Context, prompt string) (string, error)
}
