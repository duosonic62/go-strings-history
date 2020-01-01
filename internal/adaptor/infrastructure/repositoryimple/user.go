package repositoryimple

import (
	"context"
	"fmt"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/repositoryimple/dtoconverter"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
)

type UserRepositoryImpl struct {
}

// コンストラクタ
func NewUserRpository() repository.UserRepository {
	return UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) Find(token valueobject.AuthorizationToken) (output.UserOutput, error) {
	users, err := models.Members(
		qm.Where("token = ?", token.GetToken()),
	).All(context.Background(), boil.GetContextDB())

	fmt.Println(token)

	if err != nil {
		return output.UserOutput{}, err
	}

	if len(users) != 1 {
		return output.UserOutput{}, entity.NewApplicationError(500, "token duplicated", "Internal Server Error", nil)
	}

	return dtoconverter.ConvertUser(users[0]), nil
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
