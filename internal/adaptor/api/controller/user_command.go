package controller

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
)

type UserCommandController interface {
	Create(ctx input.Context)
	Edit(ctx input.Context)
	Delete(ctx input.Context)
}

type UserCommandControllerImpl struct {
	useCase      inputboundary.UserCommandUseCase
	errorUseCase inputboundary.ErrorUseCase
}

// コンストラクタ
func NewUserController(useCase inputboundary.UserCommandUseCase, errorUseCase inputboundary.ErrorUseCase) UserCommandController {
	return UserCommandControllerImpl{
		useCase:      useCase,
		errorUseCase: errorUseCase,
	}
}

// ユーザ作成
func (controller UserCommandControllerImpl) Create(ctx input.Context) {
	// コンテキストからコマンドを復元
	data := command.UserAddInputData{}
	// 入力のバインド & バリデーションチェック
	if err := ctx.Bind(&data); err != nil {
		controller.errorUseCase.BadRequestError(ctx, err)
		return
	}

	controller.useCase.Add(data, ctx)
}

func (controller UserCommandControllerImpl) Edit(ctx input.Context) {
	// コンテキストからコマンドを復元
	data := command.UserEditInputData{}
	// 入力のバインド & バリデーションチェック
	if err := ctx.Bind(&data); err != nil {
		controller.errorUseCase.BadRequestError(ctx, err)
		return
	}

	authToken, err := getAuthorizationToken(ctx)
	if err != nil {
		controller.errorUseCase.UnauthorizedError(ctx, err)
		return
	}

	controller.useCase.Edit(&authToken, data, ctx)
}

func (controller UserCommandControllerImpl) Delete(ctx input.Context) {
	authToken, err := getAuthorizationToken(ctx)
	if err != nil {
		controller.errorUseCase.UnauthorizedError(ctx, err)
		return
	}

	controller.useCase.Delete(&authToken, ctx)
}