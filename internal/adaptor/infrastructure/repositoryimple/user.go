package repositoryimple

import (
	"context"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple/dtoconverter"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"github.com/volatiletech/sqlboiler/boil"
	"time"
)

type UserRepositoryImpl struct {
}

// コンストラクタ
func NewUserRpository() repository.UserRepository {
	return UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) Find(token valueobject.AuthorizationToken) (output.UserOutput, error) {
	user, err := models.FindMember(context.Background(), boil.GetContextDB(), token.GetToken())
	if err != nil {
		return output.UserOutput{}, err
	}
	return dtoconverter.ConvertUser(user), nil
}

func (repository UserRepositoryImpl) Save(user entity.User) error {
	dbModelUser := models.Member{
		ID:           user.ID,
		UID:          user.UID,
		Name:         user.Name,
		Token:        user.Token,
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
