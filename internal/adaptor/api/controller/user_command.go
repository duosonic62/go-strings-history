package controller

import (
	"github.com/duosonic/go-strings-history/pkg/usecase/input"
	"github.com/duosonic/go-strings-history/pkg/usecase/input/command"
)

type UserCommandController interface {
	CreateUser(c input.Context)
}

type UserControllerImpl struct {
	useCase inputboundary.UserUseCase
}

// コンストラクタ
func NewUserController(useCase inputboundary.UserUseCase) UserCommandController {
	return UserControllerImpl{
		useCase: useCase,
	}
}

// ユーザ作成
func (controller UserControllerImpl) CreateUser(ctx input.Context) {
	// コンテキストからコマンドを復元
	data := command.UserAddInputData{}
	ctx.Bind(&data)

	controller.useCase.AddUser(data, ctx)
}