package user_api

import (
	"encoding/json"
	"errors"
	"net/http"
	"social/internal/api/middleware"
	"social/internal/api/user_api/userdto"
	"social/internal/domain/user"
	"social/internal/pkg/utils"
	"social/internal/pkg/validation"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userService *user.UserService
}

func NewUserHandler(s *user.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.userService.GetUserByUserId(r.Context(), id)

	if err != nil {
		if err.Error() == "user not found" {
			utils.ResponseError(w, http.StatusNotFound, err.Error())
		} else {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	userResponse := userdto.ToUserResponse(user)
	utils.ResponseJSON(w, http.StatusOK, userResponse)
}

func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUser(r.Context())
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	responseUsers := make([]*userdto.UserResponse, 0, len(users))

	for _, item := range users {
		responseUsers = append(responseUsers, userdto.ToUserResponse(item))
	}
	utils.ResponseJSON(w, http.StatusOK, responseUsers)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dto := &userdto.CreateUserDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if errs := validation.ValidateDTO(dto); errs != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{"errors": errs})
		return
	}

	newUser := dto.ToDomain()

	createdUser, err := h.userService.CreateNewUser(r.Context(), newUser)

	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			utils.ResponseError(w, http.StatusConflict, err.Error())
			return
		}
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userResponse := userdto.ToUserResponse(createdUser)
	utils.ResponseJSON(w, http.StatusCreated, userResponse)
}

/////////////////////////////////posts////////////////////////////////////////

func (h *UserHandler) getMyPosts(w http.ResponseWriter, r *http.Request) {
	dto := &userdto.GetPostsByUserIdQueryDto{}

	errs := validation.ValidateQueryDTO(r, dto)

	if errs != nil {
		utils.ResponseError(w, http.StatusBadRequest, map[string]any{"errors": errs})
		return
	}

	claims := middleware.MustGetClaims(w, r)

	filter := user.PostFilter{
		Page:   int64(dto.Page),
		Limit:  int64(dto.Limit),
		Search: dto.Search,
		Sort:   dto.Sort,
		Filters: struct {
			Title string
			Desc  string
			Tags  string
		}{
			Title: dto.Filters.Title,
			Desc:  dto.Filters.Desc,
			Tags:  dto.Filters.Tags,
		},
	}

	posts, err := h.userService.GetPostsByUserId(r.Context(), claims.UserID, filter)

	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses := userdto.ToUserPostsResponse(posts)

	utils.ResponseJSON(w, http.StatusOK, responses)
}

func (h *UserHandler) getPostsByUserId(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid user Id")
		return
	}

	dto := &userdto.GetPostsByUserIdQueryDto{}

	errs := validation.ValidateQueryDTO(r, dto)

	if errs != nil {
		utils.ResponseError(w, http.StatusBadRequest, map[string]any{"errors": errs})
		return
	}

	filter := user.PostFilter{
		Page:   int64(dto.Page),
		Limit:  int64(dto.Limit),
		Search: dto.Search,
		Sort:   dto.Sort,
		Filters: struct {
			Title string
			Desc  string
			Tags  string
		}{
			Title: dto.Filters.Title,
			Desc:  dto.Filters.Desc,
			Tags:  dto.Filters.Tags,
		},
	}

	posts, err := h.userService.GetPostsByUserId(r.Context(), userId, filter)

	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			utils.ResponseError(w, http.StatusNotFound, err.Error())
			return
		}

		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses := userdto.ToUserPostsResponse(posts)

	utils.ResponseJSON(w, http.StatusOK, responses)
}
