package controller

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
)

type UserCommandController interface {
	CreateUser(ctx input.Context)
}

type UserControllerImpl struct {
	useCase      inputboundary.UserCommandUseCase
	errorUseCase inputboundary.BadRequestUseCase
}

// コンストラクタ
func NewUserController(useCase inputboundary.UserCommandUseCase, errorUseCase inputboundary.BadRequestUseCase) UserCommandController {
	return UserControllerImpl{
		useCase:      useCase,
		errorUseCase: errorUseCase,
	}
}

// ユーザ作成
func (controller UserControllerImpl) CreateUser(ctx input.Context) {
	// コンテキストからコマンドを復元
	data := command.UserAddInputData{}
	// 入力のバインド & バリデーションチェック
	if err := ctx.Bind(&data); err != nil {
		controller.errorUseCase.BadRequestError(ctx, err)
		return
	}

	controller.useCase.AddUser(data, ctx)
}
