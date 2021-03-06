package repositoryimple

import (
	"context"
	"errors"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple/dtoconverter"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/volatiletech/sqlboiler/boil"
)

type UserCommandRepositoryImpl struct {
}

// コンストラクタ
func NewUserCommandRepository() repository.UserCommandRepository {
	return UserCommandRepositoryImpl{}
}

func (repository UserCommandRepositoryImpl) Find(token *valueobject.AuthorizationToken) (*entity.User, error) {
	dbModelUser, err := findUserByToken(token)
	if err != nil {
		return nil, err
	}

	user, err := dtoconverter.ToUserEntity(dbModelUser)
	if err != nil {
		return nil, entity.NewApplicationError(500, err.Error(), "Internal Server Error", err)
	}

	return user, nil
}

func (repository UserCommandRepositoryImpl) Save(user *entity.User) error {
	dbModelUser := dtoconverter.ToUserModel(user, 1)
	err := dbModelUser.Insert(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		return entity.NewApplicationError(500, "db error", "Internal Server Error", err)
	}
	return nil
}

func (repository UserCommandRepositoryImpl) Edit(user *entity.User) error {
	dbModelUser := dtoconverter.ToUserModel(user, 1) // FIXME
	updateCount, err := dbModelUser.Update(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		return err
	}
	if updateCount != 1 {
		// TODO: transaction
		return errors.New("db error")
	}
	return nil
}

func (repository UserCommandRepositoryImpl) Delete(token *valueobject.AuthorizationToken) error {
	dbModelUser, err := findUserByToken(token)
	if err != nil {
		return err
	}

	dbModelUser.IsDeleted = true
	deleteCount, err := dbModelUser.Update(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		return errors.New("db error")
	}
	if deleteCount != 1 {
		// TODO: transaction
		return errors.New("db error")
	}

	return nil
}
