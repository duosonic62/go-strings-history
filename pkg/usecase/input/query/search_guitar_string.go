package query

import "github.com/volatiletech/null"

type SearchGuitarString struct {
	Name       null.String
	Maker      null.String
	ThinGauge  null.Int
	ThickGauge null.Int
}
