package userdto

import "social/internal/domain/user/views"

type userPostResponse struct {
	Id          int64           `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	UserId      int64           `json:"user_id"`
	Tags        []string        `json:"tags"`
	MediaURLs   []mediaResponse `json:"media_urls"`
	CreatedAt   int64           `json:"created_at"`
}

type mediaResponse struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

func ToUserPostsResponse(views []*views.PostView) []*userPostResponse {
	res := make([]*userPostResponse, len(views))
	for i, v := range views {
		media := make([]mediaResponse, len(v.MediaURLs))
		for j, m := range v.MediaURLs {
			media[j] = mediaResponse{Type: m.Type, Path: m.Path}
		}
		res[i] = &userPostResponse{
			Id:          v.Id,
			UserId:      v.UserId,
			Title:       v.Title,
			Description: v.Description,
			Tags:        v.Tags,
			MediaURLs:   media,
			CreatedAt:   v.CreatedAt.Unix(),
		}
	}
	return res
}
