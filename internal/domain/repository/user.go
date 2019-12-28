package repository

import "github.com/duosonic/go-strings-history/internal/domain/entity"

type UserRepository interface {
	Save(entity.User)
}