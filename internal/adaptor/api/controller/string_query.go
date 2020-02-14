package controller

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/query"
	"github.com/volatiletech/null"
	"strconv"
)

type StringQueryController interface {
	GetGuitarString(ctx input.Context)
	SearchGuitarString(ctx input.Context)
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
		queries.Name = null.StringFrom(name)
	} else {
		null.NewString(name, false)
	}

	if maker, ok := ctx.GetQuery("maker"); ok {
		queries.Maker = null.StringFrom(maker)
	} else {
		null.NewString(maker, false)
	}

	if thinGauge, ok := ctx.GetQuery("thinGauge"); ok {
		g, err := strconv.Atoi(thinGauge)
		if err != nil {
			controller.errorUseCase.BadRequestError(ctx, errors.New("thinGauge must be number"))
			return
		}
		queries.ThinGauge = null.IntFrom(g)
	} else {
		queries.ThinGauge = null.NewInt(0, false)
	}

	if thickGauge, ok := ctx.GetQuery("thickGauge"); ok {
		g, err := strconv.Atoi(thickGauge)
		if err != nil {
			controller.errorUseCase.BadRequestError(ctx, errors.New("thickGauge must be number"))
			return
		}
		queries.ThickGauge = null.IntFrom(g)
	} else {
		queries.ThickGauge = null.NewInt(0, false)
	}

	controller.useCase.SearchGuitarString(queries, authToken, ctx)
}
