package valueobject

import "errors"

type StringGauge struct {
	thin  int
	thick int
}

func NewStringGauge(
	thin int,
	thick int,
) (*StringGauge, error) {
	if thin > thick {
		return nil, errors.New("thin gauge must not be bigger than thick gauge")
	}

	return &StringGauge{
		thin: thin,
		thick: thick,
	}, nil
}

func (gauge StringGauge) ThinGauge() int {
	return gauge.thin
}

func (gauge StringGauge) ThickGauge() int {
	return gauge.thick
}