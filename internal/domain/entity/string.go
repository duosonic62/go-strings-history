package entity

import "github.com/duosonic62/go-strings-history/internal/domain/valueobject"

type GuitarString struct {
	id          string
	name        string
	description string
	maker       string
	gauge       *valueobject.StringGauge
	url         string
}

func NewGuitarString(
	id string,
	name string,
	description string,
	maker string,
	gauge *valueobject.StringGauge,
	url string,
) (*GuitarString, error) {
	return &GuitarString{
		id:          id,
		name:        name,
		description: description,
		maker:       maker,
		gauge:       gauge,
		url:         url,
	}, nil
}

func (guitarString GuitarString) GetID() string {
	return guitarString.id
}

func (guitarString GuitarString) GetName() string {
	return guitarString.name
}

func (guitarString GuitarString) ChangeName(name string) {
	guitarString.name = name
}

func (guitarString GuitarString) GetDescription() string {
	return guitarString.description
}

func (guitarString GuitarString) ChangeDescription(description string) {
	guitarString.description = description
}

func (guitarString GuitarString) GetMaker() string {
	return guitarString.maker
}

func (guitarString GuitarString) ChangeMaker(maker string) {
	guitarString.maker = maker
}

func (guitarString GuitarString) GetGauge() *valueobject.StringGauge {
	return guitarString.gauge
}

func (guitarString GuitarString) ChangeGauge(thinGauge, thickGauge int) error {
	gauge, err := valueobject.NewStringGauge(thinGauge, thickGauge)
	if err != nil {
		return err
	}
	guitarString.gauge = gauge
	return nil
}

func (guitarString GuitarString) GetUrl() string {
	return guitarString.url
}

func (guitarString GuitarString) ChangeUrl(url string) {
	guitarString.url = url
}