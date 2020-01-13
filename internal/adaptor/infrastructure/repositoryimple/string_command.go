package repositoryimple

import (
	"context"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple/dtoconverter"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/volatiletech/sqlboiler/boil"
)

type StringCommandRepositoryImpl struct {}

func NewStringCommandRepository() repository.StringCommandRepository {
	return StringCommandRepositoryImpl{}
}

func (repository StringCommandRepositoryImpl) Save(guitarString *entity.GuitarString) error {
	stringModel := dtoconverter.ToStringModel(guitarString, 1)
	err := stringModel.Insert(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		return err
	}
	return nil
}