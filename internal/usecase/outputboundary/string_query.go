package outputboundary

import (
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type StringQueryPresenter interface {
	OutputGuitarString(output *output.GuitarStringOutput, ctx input.Context)
}
