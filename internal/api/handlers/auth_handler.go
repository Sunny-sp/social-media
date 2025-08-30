package handlers

import (
	"encoding/json"
	"net/http"
	"social/internal/domain/auth"
	authdto "social/internal/domain/auth/dtos"
	"social/internal/pkg/utils"
	"social/internal/pkg/validation"
)

type AuthHandler struct {
	authService *auth.AuthService
}

func NewAuthHandler(s *auth.AuthService) *AuthHandler {
	return &AuthHandler{authService: s}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	loginDto := &authdto.LoginDto{}
	if err := json.NewDecoder(r.Body).Decode(loginDto); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if errs := validation.ValidateDTO(loginDto); errs != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"errors": errs})
		return
	}

	user, token, err := h.authService.ValidateUser(r.Context(), loginDto)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600,
	})

	utils.ResponseJSON(w, http.StatusCreated, map[string]any{"message": "Logged In Successfully", "user": user})
}
