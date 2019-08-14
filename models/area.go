package models

import (
	"context"
	"time"

	"github.com/go-xorm/xorm"

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

func (d *Area) Create(ctx context.Context) (int64, error) {
	return factory.DB(ctx).Insert(d)
}

func (Area) GetById(ctx context.Context, id int64) (bool, Area, error) {
	area := Area{}
	has, err := factory.DB(ctx).Where("id=?", id).Get(&area)
	return has, area, err
}

func (Area) GetByCode(ctx context.Context, code string) (bool, Area, error) {
	area := Area{}
	has, err := factory.DB(ctx).Where("code=?", code).Get(&area)
	return has, area, err
}

func (Area) GetAll(ctx context.Context, sortby, order []string, offset, limit int) (int64, []Area, error) {
	queryBuilder := func() xorm.Interface {
		q := factory.DB(ctx)
		if err := setSortOrder(q, sortby, order); err != nil {
			factory.Logger(ctx).Error(err)
		}
		return q
	}
	var items []Area
	totalCount, err := queryBuilder().Limit(limit, offset).FindAndCount(&items)
	if err != nil {
		return totalCount, items, err
	}
	return totalCount, items, nil
}

func (d *Area) Update(ctx context.Context, id int64) (int64, error) {
	return factory.DB(ctx).Where("id=?", id).Update(d)

}

func (Area) Delete(ctx context.Context, id int64) (int64, error) {
	return factory.DB(ctx).Where("id=?", id).Delete(&Area{})
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
