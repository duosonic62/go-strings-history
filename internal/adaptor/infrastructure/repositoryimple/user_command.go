package repositoryimple

import (
	"context"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple/dtoconverter"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
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
	dbModelUser := models.Member{
		ID:           user.GetID(),
		UID:          user.GetUID(),
		Name:         user.GetName(),
		Token:        user.GetToken().GetToken(),
		TokenExpired: time.Now().Add(time.Duration(24 * time.Hour)),
		IsDeleted:    false,
		Version:      1,
	}
	err := dbModelUser.Insert(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		return entity.NewApplicationError(500, "db error", "Internal Server Error", err)
	}
	return nil
}
