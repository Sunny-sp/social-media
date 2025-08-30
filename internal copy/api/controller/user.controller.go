package controller

import (
	"encoding/json"
	"net/http"
	dto "social/internal/api/dto/user"
	"social/internal/api/service"
	"social/internal/api/utils"
	"social/internal/validation"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(s *service.UserService) *UserController {
	return &UserController{
		userService: s,
	}
}

func (u *UserController) GetByUserId(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := u.userService.GetUserByUserId(r.Context(), id)

	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if user == nil {
		utils.ResponseError(w, http.StatusNotFound, "User not found")
		return
	}

	utils.ResponseJSON(w, http.StatusOK, user)
}

func (u *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.GetAllUser(r.Context())
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	utils.ResponseJSON(w, http.StatusOK, users)
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dto := &dto.CreateUserDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if errs := validation.ValidateDTO(dto); errs != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"errors": errs,
		})
		return
	}

	newUser := dto.ToModel()

	createdUser, err := u.userService.CreateNewUser(r.Context(), newUser)

	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			utils.ResponseError(w, http.StatusConflict, err.Error())
			return
		}

		utils.ResponseError(w, http.StatusInternalServerError, err.Error())

		return
	}

	utils.ResponseJSON(w, http.StatusCreated, createdUser)
}
