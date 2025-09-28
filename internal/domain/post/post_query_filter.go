package post

type PostFilters struct {
	Title string
	Desc  string
	Tags  []string
}

type PostByUserIdFilter struct {
	Page    int64
	Limit   int64
	Search  string
	Sort    string
	Filters PostFilters
}
