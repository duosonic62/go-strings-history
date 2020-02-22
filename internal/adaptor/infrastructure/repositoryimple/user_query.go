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
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type UserQueryRepositoryImpl struct {
}

// コンストラクタ
func NewUserQueryRepository() repository.UserQueryRepository {
	return UserQueryRepositoryImpl{}
}

func (repository UserQueryRepositoryImpl) Find(token *valueobject.AuthorizationToken) (output.UserOutput, error) {
	user, err := findUserByToken(token)
	if err != nil {
		return output.UserOutput{}, err
	}
	return dtoconverter.ConvertUser(user), nil
}

func findUserByToken(token *valueobject.AuthorizationToken) (*models.Member, error) {
	users, err := models.Members(
		qm.Where("token = ?", token.GetToken()),
		qm.Where("is_deleted = ?", false),
	).All(context.Background(), boil.GetContextDB())

	if err != nil {
		return nil, err
	}

	if len(users) == 0  {
		return nil, entity.NewApplicationError(401, "user not found", "This Token is Invalid.", nil)
	}

	if len(users) != 1 {
		return nil, entity.NewApplicationError(500, "token duplicated", "Internal Server Error", nil)
	}
	return users[0], nil
}