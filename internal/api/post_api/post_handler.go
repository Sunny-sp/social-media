package post_api

import (
	"encoding/json"
	"net/http"
	"social/internal/domain/post"
	postdtos "social/internal/domain/post/dtos"
	"social/internal/pkg/utils"
	"social/internal/pkg/validation"
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
	postDto := &postdtos.CreatePostDto{}

	err := json.NewDecoder(r.Body).Decode(postDto)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest)
		return
	}

	// validate body with dto
	errs := validation.ValidateDTO(postDto)

	if errs != nil {
		utils.ResponseError(w, http.StatusBadRequest, map[string]any{"erros": errs})
		return
	}
	// service call to addpost
	// newPost := p.postService.AddPost()
	// response validate

	// toresponsemOdel

	//http res
}

func (p *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {

}
