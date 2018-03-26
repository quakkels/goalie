package models

// TokenAuthentication is a model used for authentication of a token ಠ益ಠ
type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}
