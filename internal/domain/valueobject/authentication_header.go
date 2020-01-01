package valueobject

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"net/http"
)

type AuthorizationToken struct {
	token string
}

func NewAuthorizationToken(token string) (AuthorizationToken, error) {
	if !valid(token) {
		return AuthorizationToken{}, entity.NewApplicationError(http.StatusUnauthorized, "token invalid", "This token is invalid.", nil)
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