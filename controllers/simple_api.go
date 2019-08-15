package controllers

import (
	"area-china-api/models"
	"net/http"
	"nomni/utils/api"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/echoswagger"
)

type SimpleApiController struct {
}

// localhost:8080/docs
func (d SimpleApiController) Init(g echoswagger.ApiGroup) {
	g.SetSecurity("Authorization")
	g.GET("/provinces", d.GetProvince).
		AddParamQueryNested(SearchInput{})
	g.GET("/citys", d.GetCity).
		AddParamQueryNested(SearchInput{})
	g.GET("/countys", d.GetCounty).
		AddParamQueryNested(SearchInput{})

}

func (d SimpleApiController) GetProvince(c echo.Context) error {

	return d.GetLevel(c, 1)
}

func (d SimpleApiController) GetCity(c echo.Context) error {

	return d.GetLevel(c, 2)
}

func (d SimpleApiController) GetCounty(c echo.Context) error {

	return d.GetLevel(c, 3)
}

func (d SimpleApiController) GetLevel(c echo.Context, level int) error {
	var v SearchInput
	if err := c.Bind(&v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, api.ParameterParsingError(err))
	}
	if v.MaxResultCount == 0 {
		v.MaxResultCount = DefaultMaxResultCount
	}
	totalCount, items, err := models.Area{}.GetByLevel(c.Request().Context(), v.Sortby, v.Order, v.SkipCount, v.MaxResultCount, level)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, api.NotFoundError(), err)
	}
	if len(items) == 0 {
		return ReturnApiFail(c, http.StatusBadRequest, api.NotFoundError())
	}
	nest := make([]Nest, 0)
	for _, area := range items {
		nest = append(nest, MoveArea(area))
	}
	return ReturnApiSucc(c, http.StatusOK, api.ArrayResult{
		TotalCount: totalCount,
		Items:      nest,
	})
}
