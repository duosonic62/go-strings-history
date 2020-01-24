package injector

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/api/controller"
	"github.com/duosonic62/go-strings-history/internal/adaptor/api/presenter"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/factoryimple"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple"
	"github.com/duosonic62/go-strings-history/internal/domain/service"
	"github.com/duosonic62/go-strings-history/internal/usecase/interactor"
)

type WebApp struct {
	UserCommandController   controller.UserCommandController
	UserQueryController     controller.UserQueryController
	StringCommandController controller.StringCommandController
}

func Initialize() *WebApp {
	// repository
	var userCommandRepository = repositoryimple.NewUserCommandRepository()
	var userQueryRepository = repositoryimple.NewUserQueryRpository()
	var stringCommandRepository = repositoryimple.NewStringCommandRepository()

	// factory
	var idFactory = factoryimple.NewIdFactory()
	var tokenFactory = factoryimple.NewTokenFactory()
	var userFactory = factoryimple.NewUserFactory(idFactory, tokenFactory, userCommandRepository)
	var stringFactory = factoryimple.NewStringFactory(idFactory)

	// domain service
	var authorizationService = service.NewAuthorizationService(userFactory)

	// use case
	var errorPresenter = presenter.NewErrorPresenter()
	var userCommandPresenter = presenter.NewUserCommandPresenter()
	var userQueryPresenter = presenter.NewUserQueryPresenter()
	var stringCommandPresenter = presenter.NewStringCommandPresenter()

	var errorUseCase = interactor.NewErrorUseCase(errorPresenter)
	var userCommandUseCase = interactor.NewUserCommandUseCase(
		userCommandPresenter,
		errorPresenter,
		userCommandRepository,
		userFactory,
	)
	var userQueryUseCase = interactor.NewUserQueryUseCase(
		userQueryRepository,
		userQueryPresenter,
		errorPresenter,
	)
	var stringCommandUseCase = interactor.NewStringCommandUseCase(
		authorizationService,
		stringCommandPresenter,
		errorPresenter,
		stringFactory,
		stringCommandRepository,
	)

	// controller
	var userCommandController = controller.NewUserController(userCommandUseCase, errorUseCase)
	var userQueryController = controller.NewUserQueryController(userQueryUseCase, errorUseCase)
	var stringCommandController = controller.NewStringCommandController(stringCommandUseCase, errorUseCase)

	var webApp = WebApp{
		UserCommandController:   userCommandController,
		UserQueryController:     userQueryController,
		StringCommandController: stringCommandController,
	}

	return &webApp
}
