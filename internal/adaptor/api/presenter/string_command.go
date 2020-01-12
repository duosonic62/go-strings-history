package presenter

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
)

type StringCommandPresenterImpl struct{}

// コンストラクタ
func NewStringCommandPresenter() outputboundary.StringCommandPresenter {
	return StringCommandPresenterImpl{}
}

func (presenter StringCommandPresenterImpl) OutputAddString(ctx input.Context) {
	ctx.JSON(200, nil)
}
