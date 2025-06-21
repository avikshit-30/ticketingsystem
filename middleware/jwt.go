package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: No token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["userId"].(float64))

		// Attach user ID to context
		ctx := context.WithValue(r.Context(), "userId", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
