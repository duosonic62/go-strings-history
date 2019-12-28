package presenter

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type ErrorPresenter struct{}

func (presenter ErrorPresenter) OutputError(ctx input.Context, err error) {
	switch e := err.(type) {
	case entity.ApplicationError:
		ctx.JSON(e.Code, output.ErrorOutput{Code: e.Code, Message: e.Front})
	default:
		ctx.JSON(500, output.ErrorOutput{Code: 500, Message: "Internal Server Error."})
	}
}
