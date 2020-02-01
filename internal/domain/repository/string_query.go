package repository

import "github.com/duosonic62/go-strings-history/pkg/usecase/output"

type StringQueryRepository interface {
	Find(stringID string) (*output.GuitarStringOutput, error)
}