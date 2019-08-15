package controllers

import "area-china-api/models"

const (
	DefaultMaxResultCount = 50
)

type SearchInput struct {
	Sortby         []string `query:"sortby"`
	Order          []string `query:"order"`
	SkipCount      int      `query:"skipCount"`
	MaxResultCount int      `query:"maxResultCount"`
}

func MoveArea(source models.Area) Nest {
	return Nest{
		Id:       source.Uid,
		Name:     source.WholeName,
		ParentId: source.ParentId,
		IsLeaf:   source.IsLeaf,
		Level:    source.Level,
	}
}
