package controller

import (
	"encoding/json"
	"net/http"
	dto "social/internal/api/dto/auth"
	"social/internal/api/service"
	"social/internal/api/utils"
	"social/internal/validation"
)

type AuthController struct {
	// AuthService
	authService *service.AuthService
}

func NewAuthController(s *service.AuthService) *AuthController {
	return &AuthController{
		authService: s,
	}
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// validate r.body
	defer r.Body.Close()

	loginDto := &dto.LoginDto{}

	err := json.NewDecoder(r.Body).Decode(loginDto)

	if err != nil {
		// err internal server err
		utils.ResponseError(w, http.StatusBadRequest, "Invalid Json")
		return
	}

	// dto validation
	errs := validation.ValidateDTO(loginDto)

	if errs != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"errors": errs,
		})
		return
	}
	// validate user credentials
	user, token, err := a.authService.ValidateUser(r.Context(), loginDto)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// return success/failure
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",                  // cookie valid for entire domain
		HttpOnly: true,                 // JS cannot access cookie
		Secure:   false,                // only sent over HTTPS (set false for local dev)
		SameSite: http.SameSiteLaxMode, // adjust depending on frontend
		MaxAge:   3600,                 // 1 hour expiry
	})

	utils.ResponseJSON(w, http.StatusCreated, map[string]any{"message": "Logged In Succesfully", "user": user})
}
