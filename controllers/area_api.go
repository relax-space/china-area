package controllers

import (
	"github.com/pangpanglabs/echoswagger"
)

type AreaApiController struct {
}

// localhost:8080/docs
func (d AreaApiController) Init(g echoswagger.ApiGroup) {
	// g.SetSecurity("Authorization")
	// g.GET("", d.GetAll).
	// 	AddParamQueryNested(SearchInput{})
	// g.GET("/:id", d.GetOne).
	// 	AddParamPath("", "id", "id").AddParamQuery("", "with_store", "with_store", false)
	// g.PUT("/:id", d.Update).
	// 	AddParamPath("", "id", "id").
	// 	AddParamBody(models.Area{}, "area", "only can modify name,color,price", true)
	// g.POST("", d.Create).
	// 	AddParamBody(models.Area{}, "area", "new area", true)
	// g.DELETE("/:id", d.Delete).
	// 	AddParamPath("", "id", "id")
}

// /*
// localhost:8080/areas
// localhost:8080/areas?name=apple
// localhost:8080/areas?skipCount=0&maxResultCount=2
// localhost:8080/areas?skipCount=0&maxResultCount=2&sortby=store_code&order=desc
// */
// func (AreaApiController) GetAll(c echo.Context) error {
// 	var v SearchInput
// 	if err := c.Bind(&v); err != nil {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.ParameterParsingError(err))
// 	}
// 	if v.MaxResultCount == 0 {
// 		v.MaxResultCount = DefaultMaxResultCount
// 	}
// 	totalCount, items, err := models.Area{}.GetAll(c.Request().Context(), v.Sortby, v.Order, v.SkipCount, v.MaxResultCount)
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusInternalServerError, api.NotFoundError(), err)
// 	}
// 	if len(items) == 0 {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.NotFoundError())
// 	}
// 	return ReturnApiSucc(c, http.StatusOK, api.ArrayResult{
// 		TotalCount: totalCount,
// 		Items:      items,
// 	})
// }

// /*
// localhost:8080/areas/1?with_store=true
// localhost:8080/areas/1
// */
// func (d AreaApiController) GetOne(c echo.Context) error {

// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.InvalidParamError("id", c.Param("id"), err))
// 	}

// 	has, area, err := models.Area{}.GetById(c.Request().Context(), id)
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusInternalServerError, api.NotFoundError(), err)
// 	}
// 	if !has {
// 		param := fmt.Sprintf("?id=%v", id)
// 		return ReturnApiFail(c, http.StatusBadRequest, api.NotFoundError(), param)
// 	}
// 	return ReturnApiSucc(c, http.StatusOK, area)
// }

// /*
// localhost:8080/areas
//  {
//         "code": "AA01",
//         "name": "Apple",
//         "color": "",
//         "price": 2,
//         "store_code": ""
//     }
// */
// func (d AreaApiController) Create(c echo.Context) error {
// 	var v models.Area
// 	if err := c.Bind(&v); err != nil {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.ParameterParsingError(err))
// 	}
// 	if err := c.Validate(v); err != nil {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.ParameterParsingError(err))
// 	}
// 	has, _, err := models.Area{}.GetByCode(c.Request().Context(), v.Code)
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusInternalServerError, api.NotCreatedError(), err)
// 	}
// 	if has {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.NotCreatedError(), "code has exist")
// 	}
// 	affectedRow, err := v.Create(c.Request().Context())
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusInternalServerError, api.NotCreatedError(), err)
// 	}
// 	if affectedRow == int64(0) {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.NotCreatedError())
// 	}
// 	return ReturnApiSucc(c, http.StatusCreated, v)
// }

// /*
// localhost:8080/areas
//  {
//         "price": 21,
//     }
// */
// func (d AreaApiController) Update(c echo.Context) error {
// 	var v models.Area
// 	if err := c.Bind(&v); err != nil {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.ParameterParsingError(err))
// 	}
// 	if err := c.Validate(&v); err != nil {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.ParameterParsingError(err))
// 	}
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.InvalidParamError("id", c.Param("id"), err))
// 	}
// 	has, _, err := models.Area{}.GetById(c.Request().Context(), id)
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusInternalServerError, api.NotUpdatedError(), err)

// 	}
// 	if has == false {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.NotUpdatedError(), "id has not found")
// 	}
// 	affectedRow, err := v.Update(c.Request().Context(), id)
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusInternalServerError, api.NotUpdatedError(), err)
// 	}
// 	if affectedRow == int64(0) {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.NotUpdatedError())
// 	}
// 	return ReturnApiSucc(c, http.StatusOK, v)
// }

// /*
// localhost:8080/areas/45
// */
// func (d AreaApiController) Delete(c echo.Context) error {
// 	idStr := c.Param("id")
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.InvalidParamError("id", c.Param("id"), err))
// 	}
// 	has, v, err := models.Area{}.GetById(c.Request().Context(), id)
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusInternalServerError, api.NotDeletedError(), err)
// 	}
// 	if has == false {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.NotDeletedError(), "id has not found")
// 	}
// 	affectedRow, err := models.Area{}.Delete(c.Request().Context(), id)
// 	if err != nil {
// 		return ReturnApiFail(c, http.StatusInternalServerError, api.NotDeletedError(), err)
// 	}
// 	if affectedRow == int64(0) {
// 		return ReturnApiFail(c, http.StatusBadRequest, api.NotDeletedError())
// 	}
// 	return ReturnApiSucc(c, http.StatusOK, v)
// }
