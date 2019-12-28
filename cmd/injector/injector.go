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
	UserController controller.UserCommandController
}

func Initialize() *WebApp {
	// factory
	var idFactory = factoryimple.NewIdFactory()
	var tokenFactory = factoryimple.NewTokenFactory()
	var userFactory = factory.NewUserFactory(idFactory, tokenFactory)

	// repository
	var userRepository = repositoryimple.NewUserRpository()

	// usecase
	var userCommandPresenter = presenter.NewUserCommandPresenter()
	var userCommandUseCase = interactor.NewUserCommandUseCase(
		userCommandPresenter,
		userRepository,
		idFactory,
		userFactory,
	)

	// controller
	var userCommandController = controller.NewUserController(userCommandUseCase)

	var webApp = WebApp{
		UserController: userCommandController,
	}

	return &webApp
}

