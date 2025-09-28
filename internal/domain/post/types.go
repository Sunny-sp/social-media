package post

type PostPresignedResult struct {
	Key              string
	OriginalFilename string
	SignedURL        string
}
