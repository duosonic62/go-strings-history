package interactor

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/factory"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/service"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
)

type StringCommandUseCaseImpl struct {
	authorizedService service.AuthorizationService
	presenter         outputboundary.StringCommandPresenter
	errorPresenter    outputboundary.ErrorPresenter
	userFactory       factory.UserFactory
	stringFactory     factory.StringFactory
	stringRepository  repository.StringCommandRepository
}

func NewStringCommandUseCase(
	authorizedService service.AuthorizationService,
	errorPresenter outputboundary.ErrorPresenter,
) inputboundary.StringCommandUseCase {
	return StringCommandUseCaseImpl{
		authorizedService: authorizedService,
		errorPresenter:    errorPresenter,
	}
}

func (useCase StringCommandUseCaseImpl) Add(data command.StringAddInputData, token valueobject.AuthorizationToken, ctx input.Context) {
	user, err := useCase.userFactory.Find(token)
	if err != nil {
		useCase.errorPresenter.OutputError(ctx, entity.NewUnAuthorizedError(&token, err))
		return
	}

	guitarString, err := useCase.stringFactory.NewString(
		user,
		data.ThinGauge,
		data.ThickGauge,
		data.Name,
		data.Description,
		data.Maker,
		data.Url,
	)
	if err != nil {
		useCase.errorPresenter.OutputError(ctx, entity.NewBadRequestError(err))
		return
	}

	err = useCase.stringRepository.Save(guitarString)
	if err != nil {
		useCase.errorPresenter.OutputError(ctx, entity.NewInternalServerError(err))
		return
	}

	useCase.presenter.OutputAddString(ctx)
}
