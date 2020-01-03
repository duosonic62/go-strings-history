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
func NewUserQueryRpository() repository.UserQueryRepository {
	return UserQueryRepositoryImpl{}
}

func (repository UserQueryRepositoryImpl) Find(token valueobject.AuthorizationToken) (output.UserOutput, error) {
	users, err := models.Members(
		qm.Where("token = ?", token.GetToken()),
	).All(context.Background(), boil.GetContextDB())

	if err != nil {
		return output.UserOutput{}, err
	}

	if len(users) != 1 {
		return output.UserOutput{}, entity.NewApplicationError(500, "token duplicated", "Internal Server Error", nil)
	}

	return dtoconverter.ConvertUser(users[0]), nil
}
