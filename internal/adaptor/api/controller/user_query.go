package controller

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
)

type UserQueryController interface {
	Show(ctx input.Context)
}

type UserQueryControllerImpl struct {
	useCase      inputboundary.UserQueryUseCase
	errorUseCase inputboundary.ErrorUseCase
}

func NewUserQueryController(useCase inputboundary.UserQueryUseCase, errorUseCase inputboundary.ErrorUseCase) UserQueryController {
	return UserQueryControllerImpl{
		useCase:      useCase,
		errorUseCase: errorUseCase,
	}
}

func (controller UserQueryControllerImpl) Show(ctx input.Context) {
	authToken, err := getAuthorizationToken(ctx)
	if err != nil {
		controller.errorUseCase.UnauthorizedError(ctx, err)
		return
	}

	controller.useCase.Show(&authToken, ctx)
}
