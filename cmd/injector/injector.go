package injector

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/api/controller"
	"github.com/duosonic62/go-strings-history/internal/adaptor/api/presenter"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/factoryimple"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple"
	"github.com/duosonic62/go-strings-history/internal/domain/factory"
	"github.com/duosonic62/go-strings-history/internal/usecase/interactor"
)

type WebApp struct {
	UserCommandController controller.UserCommandController
	UserQueryController   controller.UserQueryController
}

func Initialize() *WebApp {
	// repository
	var userCommandRepository = repositoryimple.NewUserCommandRepository()
	var userQueryRepository = repositoryimple.NewUserQueryRpository()

	// factory
	var idFactory = factoryimple.NewIdFactory()
	var tokenFactory = factoryimple.NewTokenFactory()
	var userFactory = factory.NewUserFactory(idFactory, tokenFactory, userCommandRepository)

	// use case
	var errorPresenter = presenter.NewErrorPresenter()
	var userCommandPresenter = presenter.NewUserCommandPresenter()
	var userQueryPresenter = presenter.NewUserQueryPresenter()

	var errorUseCase = interactor.NewErrorUseCase(errorPresenter)
	var userCommandUseCase = interactor.NewUserCommandUseCase(
		userCommandPresenter,
		errorPresenter,
		userCommandRepository,
		userFactory,
	)
	var userQueryUseCase = interactor.NewUserQueyUseCase(
		userQueryRepository,
		userQueryPresenter,
		errorPresenter,
	)

	// controller
	var userCommandController = controller.NewUserController(userCommandUseCase, errorUseCase)
	var userQueryController = controller.NewUserQueryController(userQueryUseCase, errorUseCase)

	var webApp = WebApp{
		UserCommandController: userCommandController,
		UserQueryController:   userQueryController,
	}

	return &webApp
}
