package controller

import (
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
)

func getAuthorizationToken(ctx input.Context) (*valueobject.AuthorizationToken, error) {
	token := ctx.GetHeader("Authorization")
	authToken, err := valueobject.NewAuthorizationToken(token)
	return &authToken, err
}
