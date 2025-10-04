package userdto

type GetPostsByUserIdQueryDto struct {
	Page   int    `form:"page" validate:"gte=1" default:"1"`
	Limit  int    `form:"limit" validate:"gte=1,lte=20" default:"10"`
	Search string `form:"search" validate:"omitempty,max=100"`
	Sort   string `form:"sort" validate:"omitempty,oneof=ASC DESC" default:"DESC"`

	Filters PostByUserIdFilterDto
}

type PostByUserIdFilterDto struct {
	Title string `form:"title" validate:"omitempty,max=200"`
	Desc  string `form:"description" validate:"omitempty,max=500"`
	Tags  string `form:"tags" validate:"omitempty,max=50"`
}
