package controllers

import (
	"area-china-api/models"
	"fmt"
	"net/http"
	"nomni/utils/api"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/echoswagger"
)

type AreaApiController struct {
}

// localhost:8080/docs
func (d AreaApiController) Init(g echoswagger.ApiGroup) {
	g.SetSecurity("Authorization")
	s := SimpleApiController{}
	g.GET("/:id", s.Get).
		AddParamPath("", "id", "350000000000").AddParamQuery("", "simple", "true", false)

	g.GET("/:uid/front", d.GetByUid).
		AddParamPath("", "uid", "130402000000").
		AddParamQuery("", "format", "json(remand),list", false).
		AddParamQuery("", "leaf", "1,2,3, 4(default)", false).
		AddParamQuery("", "fix_level", "1,2,3, 4", false)
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
	fixLevel, _ := strconv.ParseInt(c.QueryParam("fix_level"), 10, 64)

	format := c.QueryParam("format")
	switch format {
	case "list":
		return d.uidFrontList(c, area, int(fixLevel))
	case "json":
		return d.uidFrontJson(c, area, int(fixLevel))
	case "child_list":

	default:
		return ReturnApiSucc(c, http.StatusOK, MoveArea(area))
	}
	return ReturnApiSucc(c, http.StatusOK, MoveArea(area))
}

func (d AreaApiController) uidFrontList(c echo.Context, area models.Area, fixLevel int) error {

	status, result, apiError, err := d.uidFront(c, area)
	if err != nil {
		return ReturnApiFail(c, status, apiError, err)
	}
	return ReturnApiSucc(c, http.StatusOK, result)
}

func (d AreaApiController) uidFrontJson(c echo.Context, area models.Area, fixLevel int) error {

	status, result, apiError, err := d.uidFront(c, area)
	if err != nil {
		return ReturnApiFail(c, status, apiError, err)
	}

	nest := NewNest(result)
	target := nest.GetByParentId("0")
	if fixLevel != 0 {
		if fixLevel == 1 {
			target = append(make([]Nest, 0), nest.GetById(area.ProvinceId))
		}
		current := MoveArea(area)
		nest.WithFilter(fixLevel).WithObject(current)
	}
	nest.CallSetChild(d.uidFrontParam(area), target)
	return ReturnApiSucc(c, http.StatusOK, target)
}

func (d AreaApiController) uidFront(c echo.Context, area models.Area) (int, []Nest, api.Error, error) {
	ids := append(d.uidFrontParam(area), "0")
	areas, err := models.Area{}.GetByParentIds(c.Request().Context(), ids...)
	if err != nil {
		return http.StatusInternalServerError, nil, api.NotFoundError(), err
	}
	dtoList := make([]Nest, 0)
	for _, area := range areas {
		if c.QueryParam("leaf") == fmt.Sprint(area.Level) {
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

func (d SimpleApiController) Get(c echo.Context) error {

	id := c.Param("id")
	var simple bool
	if len(c.QueryParam("simple")) != 0 {
		var err error
		simple, err = strconv.ParseBool(c.QueryParam("simple"))
		if err != nil {
			return ReturnApiFail(c, http.StatusBadRequest, api.InvalidParamError("simple", c.QueryParam("simple"), err))
		}
	}
	if simple == false {
		return d.GetOne(c, id)
	}

	return d.GetOneSimple(c, id)
}

func (d SimpleApiController) GetOne(c echo.Context, id string) error {
	has, simple, err := models.Area{}.GetByUid(c.Request().Context(), id)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, api.NotFoundError(), err)
	}
	if !has {
		err = fmt.Errorf("?id=%v", id)
		return ReturnApiFail(c, http.StatusBadRequest, api.NotFoundError(), err)
	}
	areas, err := models.Area{}.GetByParentId(c.Request().Context(), id)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, api.NotFoundError(), err)
	}
	nests := make([]Nest, 0)
	for _, area := range areas {
		nests = append(nests, MoveArea(area))
	}
	p := MoveArea(simple)
	p.Children = nests
	return ReturnApiSucc(c, http.StatusOK, p)
}

func (d SimpleApiController) GetOneSimple(c echo.Context, id string) error {
	has, simple, err := models.Area{}.GetByUid(c.Request().Context(), id)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, api.NotFoundError(), err)
	}
	if !has {
		err = fmt.Errorf("?id=%v", id)
		return ReturnApiFail(c, http.StatusBadRequest, api.NotFoundError(), err)
	}
	return ReturnApiSucc(c, http.StatusOK, MoveArea(simple))
}
