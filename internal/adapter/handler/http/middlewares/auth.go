package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/g-villarinho/hexagonal-demo/internal/core/port"
)

type contextKey string

const (
	userContextKey = contextKey("user_id")
)

func AuthMiddleware(tokenMaker port.TokenMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "cabeçalho de autorização não fornecido", http.StatusUnauthorized)
				return
			}

			fields := strings.Fields(authHeader)
			if len(fields) < 2 {
				http.Error(w, "formato do cabeçalho de autorização inválido", http.StatusUnauthorized)
				return
			}

			if strings.ToLower(fields[0]) != "bearer" {
				http.Error(w, "esquema de autorização não suportado", http.StatusUnauthorized)
				return
			}

			accessToken := fields[1]
			userID, err := tokenMaker.Verify(r.Context(), accessToken)
			if err != nil {
				http.Error(w, "token inválido", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userContextKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
