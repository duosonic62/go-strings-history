package presenter

import (
	"github.com/duosonic/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic/go-strings-history/pkg/usecase/input"
	"github.com/duosonic/go-strings-history/pkg/usecase/output"
)

type UserPresenterImpl struct{}

// コンストラクタ
func NewUserCommandPresenter() outputboundary.UserCommandPresenter {
	return UserPresenterImpl{}
}

func (presenter UserPresenterImpl) OutputAddUser(data output.UserAddOutputData, ctx input.Context) {
	ctx.JSON(200, data)
}