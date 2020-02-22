package interactor

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/service"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/query"
)

type stringQueryUseCase struct {
	authorizedService service.AuthorizationService
	errorPresenter    outputboundary.ErrorPresenter
	stringRepository  repository.StringQueryRepository
	presenter         outputboundary.StringQueryPresenter
}

func NewStringQueryUseCase(
	authorizedService service.AuthorizationService,
	presenter outputboundary.StringQueryPresenter,
	errorPresenter outputboundary.ErrorPresenter,
	stringRepository repository.StringQueryRepository,
) inputboundary.StringQueryUseCase {
	return stringQueryUseCase{
		authorizedService: authorizedService,
		errorPresenter:    errorPresenter,
		stringRepository:  stringRepository,
		presenter:         presenter,
	}
}

func (useCase stringQueryUseCase) GetGuitarString(
	stringID string,
	token *valueobject.AuthorizationToken,
	ctx input.Context,
) {
	if _, err := useCase.authorizedService.Authorized(token); err != nil {
		useCase.errorPresenter.OutputError(ctx, err)
		return
	}

	guitarString, err := useCase.stringRepository.Find(stringID)
	if err != nil {
		useCase.errorPresenter.OutputError(ctx, entity.NewBadRequestError(err))
		return
	}

	useCase.presenter.OutputGuitarString(guitarString, ctx)
}

func (useCase stringQueryUseCase) SearchGuitarString(
	queries query.SearchGuitarString,
	token *valueobject.AuthorizationToken,
	ctx input.Context,
) {
	if _, err := useCase.authorizedService.Authorized(token); err != nil {
		useCase.errorPresenter.OutputError(ctx, err)
		return
	}

	guitarStrings, err := useCase.stringRepository.Search(
		queries.Name,
		queries.Maker,
		queries.ThinGauge,
		queries.ThickGauge,
	)

	if err != nil {
		useCase.errorPresenter.OutputError(ctx, err)
		return
	}

	useCase.presenter.OutputGuitarStrings(guitarStrings, ctx)
}
