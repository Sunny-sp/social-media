package auth

type AuthToken struct {
	TokenId   string `json:"token_id"`
	UserId    int64  `json:"user_id"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type LoginCredentials struct {
	UserId   int64
	Password string
}
