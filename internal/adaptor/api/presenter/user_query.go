package presenter

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type UserQueryPresenterImpl struct{}

// コンストラクタ
func NewUserQueryPresenter() outputboundary.UserQueryPresenter {
	return UserQueryPresenterImpl{}
}

func (presenter UserQueryPresenterImpl) OutputUser(output output.UserOutput,ctx input.Context) {
	ctx.JSON(200, output)
}
