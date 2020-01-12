package factory

import "github.com/duosonic62/go-strings-history/internal/domain/entity"

type StringFactory interface {
	NewString(user *entity.User, thinGauge int, thickGauge int, name, description, maker, url string) entity.GuitarString
}