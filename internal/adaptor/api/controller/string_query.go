package controller

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/query"
	"strconv"
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
	authToken, err := getAuthorizationToken(ctx)
	if err != nil {
		controller.errorUseCase.UnauthorizedError(ctx, err)
		return
	}

	id := ctx.Param("id")

	controller.useCase.GetGuitarString(id, authToken, ctx)
}

func (controller stringQueryController) SearchGuitarString(ctx input.Context) {
	authToken, err := getAuthorizationToken(ctx)
	if err != nil {
		controller.errorUseCase.UnauthorizedError(ctx, err)
		return
	}

	queries := query.SearchGuitarString{}
	if name, ok := ctx.GetQuery("name"); ok {
		queries.Name = name
	}
	if maker, ok := ctx.GetQuery("maker"); ok {
		queries.Maker = maker
	}
	if thinGauge, ok := ctx.GetQuery("thinGauge"); ok {
		g, err := strconv.Atoi(thinGauge)
		if err != nil {
			controller.errorUseCase.BadRequestError(ctx, errors.New("thinGauge must be number"))
		}
		queries.ThickGauge = g
	}
	if thickGauge, ok := ctx.GetQuery("thickGauge"); ok {
		g, err := strconv.Atoi(thickGauge)
		if err != nil {
			controller.errorUseCase.BadRequestError(ctx, errors.New("thickGauge must be number"))
		}
		queries.ThickGauge = g
	}

	controller.useCase.SearchGuitarString(queries, authToken, ctx)
}
