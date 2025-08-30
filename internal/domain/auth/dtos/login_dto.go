package authdto

type LoginDto struct {
	UserId   int64  `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
