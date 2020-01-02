package interactor

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
)

type ErrorUseCaseInteractor struct {
	errorPresenter outputboundary.ErrorPresenter
}

func NewErrorUseCase(errorPresenter outputboundary.ErrorPresenter) inputboundary.ErrorUseCase {
	return ErrorUseCaseInteractor{
		errorPresenter: errorPresenter,
	}
}

func (useCase ErrorUseCaseInteractor) BadRequestError(ctx input.Context, err error) {
	badRequestErr := entity.NewApplicationError(400, "bad request. cause: " + err.Error(), "Bad Request", err)
	useCase.errorPresenter.OutputError(ctx, badRequestErr)
}

func (useCase ErrorUseCaseInteractor) UnauthorizedError(ctx input.Context, err error) {
	badRequestErr := entity.NewApplicationError(401, "unauthorized. cause: " + err.Error(), "Unauthorized", err)
	useCase.errorPresenter.OutputError(ctx, badRequestErr)
}