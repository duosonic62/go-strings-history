package controller

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
)

type StringCommandController interface {
	Create(ctx input.Context)
	Update(ctx input.Context)
}

type StringCommandControllerImpl struct {
	useCase      inputboundary.StringCommandUseCase
	errorUseCase inputboundary.ErrorUseCase
}

// コンストラクタ
func NewStringCommandController(useCase inputboundary.StringCommandUseCase, errorUseCase inputboundary.ErrorUseCase) StringCommandController {
	return StringCommandControllerImpl{
		useCase:      useCase,
		errorUseCase: errorUseCase,
	}
}

// ギター弦作成
func (controller StringCommandControllerImpl) Create(ctx input.Context) {
	// コンテキストからコマンドを復元
	data := command.StringRegisterInputData{}
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

	controller.useCase.Add(data, authToken, ctx)
}

// ギター弦情報変更
func (controller StringCommandControllerImpl) Update(ctx input.Context) {
	// コンテキストからコマンドを復元
	data := command.StringRegisterInputData{}
	// 入力のバインド & バリデーションチェック
	if err := ctx.Bind(&data); err != nil {
		controller.errorUseCase.BadRequestError(ctx, err)
		return
	}

	// パスパラメータからIDを取得
	id := ctx.Param("id")

	authToken, err := getAuthorizationToken(ctx)
	if err != nil {
		controller.errorUseCase.UnauthorizedError(ctx, err)
		return
	}

	controller.useCase.Update(id, data, authToken, ctx)
}