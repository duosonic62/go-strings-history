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

func NewBadRequestErrorUseCase(errorPresenter outputboundary.ErrorPresenter) inputboundary.BadRequestUseCase {
	return ErrorUseCaseInteractor{
		errorPresenter: errorPresenter,
	}
}

func (useCase ErrorUseCaseInteractor) BadRequestError(ctx input.Context, err error) {
	badRequesterr := entity.NewApplicationError(400, "bad request. cause: " + err.Error(), "Bad Request", err)
	useCase.errorPresenter.OutputError(ctx, badRequesterr)
}