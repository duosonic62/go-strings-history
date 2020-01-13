package entity

import "github.com/duosonic62/go-strings-history/internal/domain/valueobject"

type GuitarString struct {
	id          string
	name        string
	description string
	maker       string
	gauge       *valueobject.StringGauge
	url         string
	user        *User
}

func NewGuitarString(
	id string,
	name string,
	description string,
	maker string,
	gauge *valueobject.StringGauge,
	url string,
	user *User,
) (*GuitarString, error) {
	return &GuitarString{
		id:          id,
		name:        name,
		description: description,
		maker:       maker,
		gauge:       gauge,
		url:         url,
		user:        user,
	}, nil
}

func (guitarString GuitarString) GetID() string {
	return guitarString.id
}

func (guitarString GuitarString) GetName() string {
	return guitarString.name
}

func (guitarString GuitarString) GetDescription() string {
	return guitarString.description
}

func (guitarString GuitarString) GetMaker() string {
	return guitarString.maker
}

func (guitarString GuitarString) GetGauge() *valueobject.StringGauge {
	return guitarString.gauge
}

func (guitarString GuitarString) GetUrl() string {
	return guitarString.url
}

func (guitarString GuitarString) GetUser() *User {
	return guitarString.user
}