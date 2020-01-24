package presenter

import (
	"fmt"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type ErrorPresenter struct{}

func NewErrorPresenter() outputboundary.ErrorPresenter {
	return ErrorPresenter{}
}

func (presenter ErrorPresenter) OutputError(ctx input.Context, err error) {
	fmt.Println(err)
	switch e := err.(type) {
	case entity.ApplicationError:
		ctx.JSON(e.Code, output.ErrorOutput{Code: e.Code, Message: e.Front})
	default:
		ctx.JSON(500, output.ErrorOutput{Code: 500, Message: "Internal Server Error."})
	}
}
