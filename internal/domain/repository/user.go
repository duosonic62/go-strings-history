package repository

import "github.com/duosonic62/go-strings-history/internal/domain/entity"

type UserRepository interface {
	Save(entity.User)
}