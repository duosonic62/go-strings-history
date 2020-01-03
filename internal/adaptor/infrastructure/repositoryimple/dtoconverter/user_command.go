package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"time"
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

func ToUserModel(entity *entity.User, version int) models.Member {
	return models.Member{
		ID:           entity.GetID(),
		UID:          entity.GetUID(),
		Name:         entity.GetName(),
		Token:        entity.GetToken().GetToken(),
		TokenExpired: time.Now().Add(time.Duration(24 * time.Hour)),
		IsDeleted:    false,
		Version:      version,
	}
}
