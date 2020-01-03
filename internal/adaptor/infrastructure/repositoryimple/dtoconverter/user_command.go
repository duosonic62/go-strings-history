package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
)

func ToUserEntity(model *models.Member) (*entity.User, error) {
	token, err := valueobject.NewAuthorizationToken(model.Token)
	if err != nil {
		return nil, err
	}
	user, err := entity.NewUser(model.ID, model.Name, model.UID, token)
	if err != nil {
		return nil, err
	}
	return user, nil
}