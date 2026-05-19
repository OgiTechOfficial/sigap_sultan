package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/helper/common_helper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/services"
	"sigap-sultan-be/src/common"
	"strconv"
	"strings"
)

type TmCityController struct {
	TmCityService *services.TmCityService
}

func NewTmCityController(priceService *services.TmCityService) *TmCityController {
	return &TmCityController{
		priceService,
	}
}

func (r TmCityController) Get(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var sortBy string
	var page int
	var limit int
	var dataCount *int

	provinceId := c.Query("provinceId")

	qPage := c.Query("page")
	if qPage == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(qPage)
		if err != nil {
			return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
		}

		if page == 0 {
			page = 1
		}
	}

	qLimit := c.Query("limit")
	if qLimit == "" {
		limit = 100
	} else {
		limit, err = strconv.Atoi(qLimit)
		if err != nil {
			return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
		}
	}

	paginationParams := common.PaginationParams{
		common.GenerateOffset(page, limit),
		limit,
		sortBy,
	}

	qSortBy := c.Query("sortBy")
	if qSortBy != "" {
		sortBy = qSortBy
		contains := strings.Contains(sortBy, ":")
		if !contains {
			return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "Sort by must be format: id:asc", nil))
		}
	}

	params := domain.CityRequestParams{
		CommonParams:     common.CommonParams{Ctx: c},
		PaginationParams: paginationParams,
		ProvinceId:       provinceId,
	}

	data, errorDomain = r.TmCityService.Get(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	dataCount, errorDomain = r.TmCityService.Count(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.SetPagination().Success("", data, page, limit, *dataCount))
}
