package auth_api

import (
	"encoding/json"
	"net/http"
	"social/internal/api/auth_api/authdto"
	"social/internal/api/user_api/userdto"
	"social/internal/domain/auth"
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

	creds := &auth.LoginCredentials{
		UserId:   loginDto.UserId,
		Password: loginDto.Password,
	}

	user, token, err := h.authService.ValidateUser(r.Context(), creds)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	payload := userdto.ToUserResponse(user)

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600,
	})

	utils.ResponseJSON(w, http.StatusCreated, map[string]any{"message": "Logged In Successfully", "user": payload})
}
