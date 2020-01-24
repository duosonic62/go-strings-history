package interactor

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/service"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
)

type stringQueryUseCase struct {
	authorizedService service.AuthorizationService
	errorPresenter    outputboundary.ErrorPresenter
	stringRepository  repository.StringQueryRepository
	presenter         outputboundary.StringQueryPresenter
}

func NewStringQueryUseCase(
	authorizedService service.AuthorizationService,
	errorPresenter outputboundary.ErrorPresenter,
) inputboundary.StringQueryUseCase {
	return stringQueryUseCase{
		authorizedService: authorizedService,
		errorPresenter:    errorPresenter,
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
