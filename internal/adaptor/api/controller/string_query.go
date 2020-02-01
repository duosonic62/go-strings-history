package controller

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
)

type StringQueryController interface {
	GetGuitarString(ctx input.Context)
}

type stringQueryController struct {
	useCase      inputboundary.StringQueryUseCase
	errorUseCase inputboundary.ErrorUseCase
}

func NewStringQueryController(useCase inputboundary.StringQueryUseCase, errorUseCase inputboundary.ErrorUseCase) StringQueryController {
	return stringQueryController{
		useCase:      useCase,
		errorUseCase: errorUseCase,
	}
}

func (controller stringQueryController) GetGuitarString(ctx input.Context) {
	id := ctx.Param("id")

	authToken, err := getAuthorizationToken(ctx)
	if err != nil {
		controller.errorUseCase.UnauthorizedError(ctx, err)
		return
	}

	controller.useCase.GetGuitarString(id, authToken, ctx)
}