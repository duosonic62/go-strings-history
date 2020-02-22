package repositoryimple

import (
	"context"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple/dtoconverter"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"strconv"
)

type stringQueryRepository struct{}


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
	queryBuilder := db.NewQueryBuilder()
	if name.Valid {
		queryBuilder.AddWhere("name", name.String)
	}
	if maker.Valid {
		queryBuilder.AddWhere("maker", maker.String)
	}
	if thinGauge.Valid {
		queryBuilder.AddWhere("thin_gauge", strconv.Itoa(thinGauge.Int))
	}
	if thickGauge.Valid {
		queryBuilder.AddWhere("thick_gauge", strconv.Itoa(thickGauge.Int))
	}
	queries := queryBuilder.Build()
	guitarStringModels, err := models.GuitarStrings(queries...).All(context.Background(), boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	return dtoconverter.ToStringOutputs(&guitarStringModels), nil
}
