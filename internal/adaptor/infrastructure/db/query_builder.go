package db

import "github.com/volatiletech/sqlboiler/queries/qm"

type QueryBuilder struct {
	queries map[string]string
}

func NewQueryBuilder() QueryBuilder {
	return QueryBuilder{
		queries: make(map[string]string),
	}
}

func (builder QueryBuilder) AddWhere(column string, arg string) {
	builder.queries[column] = arg
}

func (builder QueryBuilder) Build() []qm.QueryMod {
	var queryMods []qm.QueryMod
	for column, arg := range builder.queries {
		queryMods = append(queryMods, qm.And(column+ "=?", arg))
	}

	return queryMods
}