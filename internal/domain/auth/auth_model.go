package auth

type AuthToken struct {
    TokenID   string `json:"token_id"`
    UserID    int64  `json:"user_id"`
    Token     string `json:"token"`
    ExpiresAt int64  `json:"expires_at"`
}