package entity

import "github.com/duosonic62/go-strings-history/internal/domain/valueobject"

type GuitarString struct {
	name        string
	description string
	maker       string
	gauge       *valueobject.StringGauge
	url         string
	user        *User
}

func NewGuitarString(
	name string,
	description string,
	maker string,
	gauge *valueobject.StringGauge,
	url string,
	user *User,
) (*GuitarString, error) {
	return &GuitarString{
		name:        name,
		description: description,
		maker:       maker,
		gauge:       gauge,
		url:         url,
		user:        user,
	}, nil
}
