package presenter

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type UserCommandPresenterImpl struct{}

// コンストラクタ
func NewUserCommandPresenter() outputboundary.UserCommandPresenter {
	return UserCommandPresenterImpl{}
}

func (presenter UserCommandPresenterImpl) OutputAddUser(output output.UserAddOutputData, ctx input.Context) {
	ctx.JSON(200, output)
}

func (presenter UserCommandPresenterImpl) OutputEditUser(output output.UserEditOutputData, ctx input.Context) {
	ctx.JSON(200, output)
}