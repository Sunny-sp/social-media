package middleware

import (
	"context"
	"errors"
	"net/http"
	"social/internal/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const claimsKey contextKey = "claims"

type AuthMiddleware struct {
	jwtSecret []byte
}

func NewAuthMiddleware(jwtSecret []byte) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: jwtSecret,
	}
}

func (a *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// extract cookie
		cookie, err := r.Cookie("auth_token")
		//validate cookie,err
		if err != nil {
			utils.ResponseError(w, http.StatusUnauthorized)
			return
		}

		// validate jwt
		claims, err := utils.ValidateJWT(cookie.Value, a.jwtSecret)

		// validate jwt err
		if err != nil {

			if errors.Is(err, jwt.ErrTokenExpired) {
				utils.ResponseError(w, http.StatusUnauthorized, "Token expired")
				return
			}
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				utils.ResponseError(w, http.StatusUnauthorized, "Invalid token signature")
				return
			}

			// Generic validation error
			utils.ResponseError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		ctx := context.WithValue(r.Context(), claimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
