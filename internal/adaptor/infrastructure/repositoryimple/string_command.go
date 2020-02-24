package repositoryimple

import (
	"context"
	"errors"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple/dtoconverter"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type stringCommandRepository struct{}

func NewStringCommandRepository() repository.StringCommandRepository {
	return stringCommandRepository{}
}

func (repository stringCommandRepository) Save(guitarString *entity.GuitarString) error {
	stringModel := dtoconverter.ToStringModel(guitarString, 1)
	err := stringModel.Insert(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (repository stringCommandRepository) FindByID(iD string) (*entity.GuitarString, error) {
	stringModel, err := repository.findByID(iD)
	if err != nil {
		return nil, err
	}

	return dtoconverter.ToStringEntity(stringModel)
}

func (repository stringCommandRepository) Update(guitarString *entity.GuitarString) error {
	guitarModel, err := repository.findByID(guitarString.GetID())
	if err != nil {
		return err
	}

	// 現在のバージョン
	currentVersion := guitarModel.Version

	// 更新内容を変更する
	updateModel := models.M{
		models.GuitarStringColumns.Name: guitarString.GetName(),
		models.GuitarStringColumns.Description: guitarString.GetDescription(),
		models.GuitarStringColumns.Maker: guitarString.GetMaker(),
		models.GuitarStringColumns.URL: guitarString.GetUrl(),
		models.GuitarStringColumns.ThinGauge: guitarString.GetGauge().ThinGauge(),
		models.GuitarStringColumns.ThickGauge: guitarString.GetGauge().ThickGauge(),
		models.GuitarStringColumns.Version: currentVersion + 1,
	}

	queries := []qm.QueryMod{
		qm.Where(models.GuitarStringColumns.ID+"=?", guitarString.GetID()),
		qm.Where(models.GuitarStringColumns.Version+"=?", currentVersion),
		qm.Where(models.GuitarStringColumns.IsDeleted+"=?", false),
	}

	rowsAff, err := models.GuitarStrings(queries...).UpdateAll(
		context.Background(),
		boil.GetContextDB(),
		updateModel,
	)

	if err != nil {
		return err
	}

	if rowsAff != 1 {
		return errors.New("optimistic lock error")
	}

	return nil
}

func (repository stringCommandRepository) findByID(iD string) (*models.GuitarString, error) {
	queries := []qm.QueryMod{
		qm.Where(models.GuitarStringColumns.ID+"=?", iD),
		qm.Where(models.GuitarStringColumns.IsDeleted+"=?", false),
	}
	stringModel, err := models.GuitarStrings(queries...).One(context.Background(), boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	if stringModel == nil {
		return nil, errors.New("invalid string id")
	}

	return stringModel, nil
}
