package middleware

import (
	"context"
	"net/http"
	"rent-server/config"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string
const userIDKey ContextKey = "user_id"

func AuthMiddleware(cfg *config.Config) func (http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenString := cookie.Value
			claims := &jwt.MapClaims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error){
				return []byte(cfg.JwtKey), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), userIDKey, (*claims)["user_id"]))

            next.ServeHTTP(w, r)
		})
	}
}

func GetKey() ContextKey {
	return userIDKey
}