package presenter

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type UserPresenterImpl struct{}

// コンストラクタ
func NewUserCommandPresenter() outputboundary.UserCommandPresenter {
	return UserPresenterImpl{}
}

func (presenter UserPresenterImpl) OutputAddUser(output output.UserAddOutputData, ctx input.Context) {
	ctx.JSON(200, output)
}