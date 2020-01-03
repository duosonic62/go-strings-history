package valueobject

import (
	"errors"
)

type AuthorizationToken struct {
	token string
}

func NewAuthorizationToken(token string) (AuthorizationToken, error) {
	if !valid(token) {
		return AuthorizationToken{}, errors.New("this token is invalid: " + token)
	}
	return AuthorizationToken{ token: token }, nil
}

func (token AuthorizationToken) GetToken() string {
	return token.token
}

// トークンのバリデーション
func valid(token string) bool {
	return len(token) > 0
}