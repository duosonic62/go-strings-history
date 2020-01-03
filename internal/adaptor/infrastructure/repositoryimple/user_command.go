package repositoryimple

import (
	"context"
	"errors"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple/dtoconverter"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type UserCommandRepositoryImpl struct {
}

// コンストラクタ
func NewUserCommandRepository() repository.UserCommandRepository {
	return UserCommandRepositoryImpl{}
}

func (repository UserCommandRepositoryImpl) Find(token valueobject.AuthorizationToken) (*entity.User, error) {
	users, err := models.Members(
		qm.Where("token = ?", token.GetToken()),
	).All(context.Background(), boil.GetContextDB())

	if err != nil {
		return nil, err
	}

	if len(users) != 1 {
		return nil, entity.NewApplicationError(500, "token duplicated", "Internal Server Error", nil)
	}

	user, err := dtoconverter.ToUserEntity(users[0])
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
