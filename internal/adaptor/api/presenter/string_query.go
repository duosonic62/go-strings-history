package presenter

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type stringQueryPresenter struct{}

func NewStringQueryPresenter() outputboundary.StringQueryPresenter {
	return stringQueryPresenter{}
}

func (presenter stringQueryPresenter) OutputGuitarString(output *output.GuitarStringOutput, ctx input.Context) {
	ctx.JSON(200, output)
}

func (presenter stringQueryPresenter) OutputGuitarStrings(output *[]output.GuitarStringOutput, ctx input.Context) {
	ctx.JSON(200, output)
}
