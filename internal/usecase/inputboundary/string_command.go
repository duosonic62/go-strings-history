package inputboundary

import (
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
)

type StringCommandUseCase interface {
	Add(data command.StringAddInputData, ctx input.Context)
}