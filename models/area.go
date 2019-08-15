package models

import (
	"context"
	"time"

	"area-china-api/factory"
)

type Area struct {
	//required
	Id         int64  `json:"id" xorm:"pk autoincr 'id'"`
	ProvinceId string `json:"province_id"`
	CityId     string `json:"city_id"`
	CountyId   string `json:"county_id"`
	Uid        string `json:"uid"`

	ParentId  string `json:"parent_id"`
	AreaName  string `json:"area_name"`
	Level     int    `json:"level"`
	IsLeaf    bool   `json:"is_leaf"`
	WholeName string `json:"whole_name"`

	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`

	//option

	SimpleName string `json:"simple_name"`
	PinYin     string `json:"pin_yin"`
	PrePinYin  string `json:"pre_pin_yin"`
	SimplePy   string `json:"simple_py"`
	ZipCode    string `json:"zip_code"`
	AreaCode   string `json:"area_code"`
	Lat        string `json:"lat"`
	Lon        string `json:"lon"`
}

func (Area) GetByUid(ctx context.Context, uid string) (bool, Area, error) {
	area := Area{}
	has, err := factory.DB(ctx).Where("uid=?", uid).Get(&area)
	return has, area, err
}

func (Area) GetByParentId(ctx context.Context, pids ...string) ([]Area, error) {
	var areas []Area
	err := factory.DB(ctx).In("parent_id", pids).Find(&areas)
	return areas, err
}

type Province struct {
	Id         int
	Name       string
	ProvinceId string
}
type City struct {
	Id         int
	Name       string
	CityId     string
	ProvinceId string
}
type County struct {
	Id       int
	Name     string
	CountyId string
	CityId   string
}
type Town struct {
	Id       int
	Name     string
	TownId   string
	CountyId string
}

type Village struct {
	Id      int
	Name    string
	Village string
	TownId  string
}
