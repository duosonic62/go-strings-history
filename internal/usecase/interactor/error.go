package interactor

import (
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

func (useCase ErrorUseCaseInteractor) Error(ctx input.Context, err error) {
	useCase.errorPresenter.OutputError(ctx, err)
}