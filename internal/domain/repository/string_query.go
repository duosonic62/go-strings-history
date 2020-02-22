package repository

import (
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"github.com/volatiletech/null"
)

type StringQueryRepository interface {
	Find(stringID string) (*output.GuitarStringOutput, error)
	Search(name null.String, maker null.String, thinGauge null.Int, thickGauge null.Int) (*[]output.GuitarStringOutput, error)
}