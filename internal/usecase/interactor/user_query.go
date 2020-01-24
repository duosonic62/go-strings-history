package interactor

import (
	"fmt"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
)

type UserQueryUseCaseInteracter struct {
	repository     repository.UserQueryRepository
	presenter      outputboundary.UserQueryPresenter
	errorPresenter outputboundary.ErrorPresenter
}

func NewUserQueyUseCase(
	repository repository.UserQueryRepository,
	presenter outputboundary.UserQueryPresenter,
	errorPresenter outputboundary.ErrorPresenter,
) inputboundary.UserQueryUseCase {
	return UserQueryUseCaseInteracter{
		repository:     repository,
		presenter:      presenter,
		errorPresenter: errorPresenter,
	}
}

func (useCase UserQueryUseCaseInteracter) Show(token *valueobject.AuthorizationToken, ctx input.Context) {
	user, err := useCase.repository.Find(token)
	if err != nil {
		fmt.Println(err)
		useCase.errorPresenter.OutputError(ctx, err)
		return
	}

	useCase.presenter.OutputUser(user, ctx)
}
