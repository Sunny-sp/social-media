package post_api

import (
	"encoding/json"
	"log"
	"net/http"
	"social/internal/api/middleware"
	"social/internal/api/post_api/postdto"
	"social/internal/domain/post"
	"social/internal/pkg/utils"
	"social/internal/pkg/validation"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	postService *post.PostService
}

func NewPosthandler(postService *post.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (p *PostHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//extract from body
	postDto := &postdto.CreatePostDto{}

	err := json.NewDecoder(r.Body).Decode(postDto)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest)
		return
	}

	claims := middleware.MustGetClaims(w, r)
	log.Println("\nclaims.UserID", claims.UserID)
	postDto.UserId = claims.UserID

	// validate body with dto
	errs := validation.ValidateDTO(postDto)

	if errs != nil {
		utils.ResponseError(w, http.StatusBadRequest, map[string]any{"erros": errs})
		return
	}

	// convert into PostDomain data with userId
	newPost := postDto.ToDomain()

	log.Println("new posts userId:", newPost.UserId)

	log.Printf("new posts: %+v", newPost)
	// service call to addpost
	err = p.postService.AddPost(r.Context(), newPost)

	// response validate
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//http res
	utils.ResponseJSON(w, http.StatusCreated, "Post created Successfully")
}

func (p *PostHandler) getPostById(w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idstr, 10, 64)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid Post Id")
		return
	}

	post, err := p.postService.GetPostById(r.Context(), id)

	if err != nil {
		if err.Error() == "post not found" {
			utils.ResponseError(w, http.StatusNotFound, err.Error())
		} else {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	payload := postdto.ToPostResponse(post)

	utils.ResponseJSON(w, http.StatusOK, payload)
}
