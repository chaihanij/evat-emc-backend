package dtos

import (
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type VisitResponseJSON struct{
	TotalVisited int `json:"total_visited"`
	TodayVisit   int `json:"today_visit"`
}

func (m *VisitResponseJSON) Parse(data *entities.Visited) *VisitResponseJSON {

	_ = copier.Copy(&m, data)
	m.TodayVisit = data.TodayVisit
	m.TotalVisited = data.TotalVisited

	return m
}