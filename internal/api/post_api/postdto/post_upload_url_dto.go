package postdto

type PresignedUploadRequest struct {
	OriginalFilename string `json:"original_filename" validate:"required,allowed_ext=jpeg jpg png mp4"`
}
