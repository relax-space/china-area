package controllers

import (
	"area-china-api/models"
	"fmt"
	"net/http"
	"nomni/utils/api"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/echoswagger"
)

type AreaApiController struct {
}

// localhost:8080/docs
func (d AreaApiController) Init(g echoswagger.ApiGroup) {
	g.SetSecurity("Authorization")
	g.GET("/:uid", d.GetByUid).
		AddParamPath("", "uid", "130402000000").
		AddParamQuery("", "with", "front_json(remand),front_list", false).
		AddParamQuery("", "leaf", "3, 4(default)", false)
}

func (d AreaApiController) GetByUid(c echo.Context) error {

	uid := c.Param("uid")
	has, area, err := models.Area{}.GetByUid(c.Request().Context(), uid)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, api.NotFoundError(), err)
	}
	if !has {
		err = fmt.Errorf("?uid=%v", uid)
		return ReturnApiFail(c, http.StatusBadRequest, api.NotFoundError(), err)
	}

	with := c.QueryParam("with")
	switch with {
	case "front_list":
		return d.uidFrontList(c, area)
	case "front_json":
		return d.uidFrontJson(c, area)
	case "child_list":

	default:
		return ReturnApiSucc(c, http.StatusOK, MoveArea(area))
	}
	return ReturnApiSucc(c, http.StatusOK, MoveArea(area))
}

func (d AreaApiController) uidFrontList(c echo.Context, area models.Area) error {

	status, result, apiError, err := d.uidFront(c, area)
	if err != nil {
		return ReturnApiFail(c, status, apiError, err)
	}
	return ReturnApiSucc(c, http.StatusOK, result)
}

func (d AreaApiController) uidFrontJson(c echo.Context, area models.Area) error {

	status, result, apiError, err := d.uidFront(c, area)
	if err != nil {
		return ReturnApiFail(c, status, apiError, err)
	}

	nest := NewNest(result)
	target := nest.GetByParentId("0")
	nest.SetChild(d.uidFrontParam(area), target)

	return ReturnApiSucc(c, http.StatusOK, target)
}

func (d AreaApiController) uidFront(c echo.Context, area models.Area) (int, []Nest, api.Error, error) {
	ids := append(d.uidFrontParam(area), "0")
	areas, err := models.Area{}.GetByParentId(c.Request().Context(), ids...)
	if err != nil {
		return http.StatusInternalServerError, nil, api.NotFoundError(), err
	}
	dtoList := make([]Nest, 0)
	for _, area := range areas {
		if area.Level == 3 && c.QueryParam("leaf") == "3" {
			area.IsLeaf = true
		}
		dtoList = append(dtoList, MoveArea(area))
	}
	return http.StatusOK, dtoList, api.Error{}, nil
}

func (d AreaApiController) uidFrontParam(area models.Area) []string {
	parentIds := make([]string, 0)
	switch area.Level {
	case 1:
		parentIds = append(parentIds, area.ProvinceId)
	case 2:
		parentIds = append(parentIds, area.ProvinceId, area.CityId)
	case 3:
		parentIds = append(parentIds, area.ProvinceId, area.CityId, area.Uid)
	}
	return parentIds
}
