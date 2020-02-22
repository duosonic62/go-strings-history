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
	stringFactory     factory.StringFactory
	stringRepository  repository.StringCommandRepository
}

func NewStringCommandUseCase(
	authorizedService service.AuthorizationService,
	presenter outputboundary.StringCommandPresenter,
	errorPresenter outputboundary.ErrorPresenter,
	stringFactory factory.StringFactory,
	stringRepository repository.StringCommandRepository,
) inputboundary.StringCommandUseCase {
	return StringCommandUseCaseImpl{
		authorizedService: authorizedService,
		presenter:         presenter,
		errorPresenter:    errorPresenter,
		stringFactory:     stringFactory,
		stringRepository:  stringRepository,
	}
}

func (useCase StringCommandUseCaseImpl) Add(data command.StringRegisterInputData, token *valueobject.AuthorizationToken, ctx input.Context) {
	if _, err := useCase.authorizedService.Authorized(token); err != nil {
		useCase.errorPresenter.OutputError(ctx, err)
		return
	}

	guitarString, err := useCase.stringFactory.NewString(
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

func (useCase StringCommandUseCaseImpl) Update(id string, data command.StringRegisterInputData, token *valueobject.AuthorizationToken, ctx input.Context) {
	if _, err := useCase.authorizedService.Authorized(token); err != nil {
		useCase.errorPresenter.OutputError(ctx, err)
		return
	}

	guitarString, err := useCase.stringRepository.FindByID(id)
	if err != nil {
		useCase.errorPresenter.OutputError(ctx, entity.NewBadRequestError(err))
		return
	}

	// 変更を詰め替え
	guitarString.ChangeName(data.Name)
	guitarString.ChangeDescription(data.Description)
	guitarString.ChangeMaker(data.Maker)
	guitarString.ChangeUrl(data.Url)
	err = guitarString.ChangeGauge(data.ThinGauge, data.ThickGauge)
	if err != nil {
		useCase.errorPresenter.OutputError(ctx, entity.NewBadRequestError(err))
		return
	}

	err = useCase.stringRepository.Update(&guitarString)

	useCase.presenter.OutputUpdateString(ctx)
}
