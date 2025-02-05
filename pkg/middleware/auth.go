package middleware

import (
	"context"
	"net/http"
	"strings"

	"go_web_server/pkg/jwt"
	"go_web_server/pkg/response"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHandler := r.Header.Get("Authorization")
		if authHandler == "" {
			response.Error(w, http.StatusUnauthorized, "请登录")
			return
		}

		parts := strings.Split(authHandler, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(w, http.StatusUnauthorized, "无效token")
			return
		}

		tokenString := parts[1]

		claims, err := jwt.ParseToken(tokenString)

		if err != nil {
			response.Error(w, http.StatusUnauthorized, "无效token")
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), userIDKey, claims.UserID))

		next.ServeHTTP(w, r)
	})
}
