package repositoryimple

import (
	"context"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple/dtoconverter"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type stringQueryRepository struct{}

type queries

func NewStringQueryRepository() repository.StringQueryRepository {
	return stringQueryRepository{}
}

func (repository stringQueryRepository) Find(stringID string) (*output.GuitarStringOutput, error) {
	guitarString, err := models.FindGuitarString(context.Background(), boil.GetContextDB(), stringID)
	if err != nil {
		return nil, err
	}

	return dtoconverter.ToStringOutput(guitarString), nil
}

func (repository stringQueryRepository) Search(
	name null.String,
	maker null.String,
	thinGauge null.Int,
	thickGauge null.Int,
) (*[]output.GuitarStringOutput, error) {
	var query string
	if name.Valid {
		query = "" + name.String
	}
	if maker.Valid {

	}
	guitarStringModels := models.GuitarStrings(qm.Where())
}
