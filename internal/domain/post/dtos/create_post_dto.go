package postdtos

type CreatePostDto struct {
	UserId      int64
	Description string
	Tags        []string
}
