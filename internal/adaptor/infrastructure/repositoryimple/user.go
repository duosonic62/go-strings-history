package repositoryimple

import (
	"context"
	"fmt"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/volatiletech/sqlboiler/boil"
	"time"
)

type UserRepositoryImpl struct {

}

// コンストラクタ
func NewUserRpository() repository.UserRepository  {
	return UserRepositoryImpl{}
}


func (repository UserRepositoryImpl) Save(user entity.User)  {
	dbModelUser := models.Member{
		ID: user.ID,
		UID: user.UID,
		Name: user.Name,
		Token: user.Token,
		TokenExpired: time.Now().Add(time.Duration(24 * time.Hour)),
		IsDeleted: false,
		Version: 1,
	}
	err := dbModelUser.Insert(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		fmt.Println(err)
	}
	return
}