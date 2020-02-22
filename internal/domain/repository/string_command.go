package repository

import (
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
)

type StringCommandRepository interface {
	Save(guitarString *entity.GuitarString) error
	Update(guitarString *entity.GuitarString) error
	FindByID(iD string) (*entity.GuitarString, error)
}
