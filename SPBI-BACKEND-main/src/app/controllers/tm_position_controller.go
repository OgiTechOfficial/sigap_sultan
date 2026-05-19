package controllers

import (
	"net/http"
	"sigap-sultan-be/src/app/helper/common_helper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/services"
	"sigap-sultan-be/src/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type TmPositionController struct {
	TmPositionService *services.TmPositionService
}

func NewTmPositionController(priceService *services.TmPositionService) *TmPositionController {
	return &TmPositionController{
		priceService,
	}
}

func (r TmPositionController) Get(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var dataCount *int

	//qPage := c.Query("page")
	//if qPage == "" {
	//	page = 0
	//} else {
	//	page, err = strconv.Atoi(qPage)
	//	if err != nil {
	//		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	//	}
	//}
	//
	//qLimit := c.Query("limit")
	//if qLimit == "" {
	//	limit = 10
	//} else {
	//	limit, err = strconv.Atoi(qLimit)
	//	if err != nil {
	//		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	//	}
	//}
	//
	//qSortBy := c.Query("sortBy")
	//if qLimit != "" {
	//	sortBy = qSortBy
	//}

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
		limit = 10
	} else {
		limit, err = strconv.Atoi(qLimit)
		if err != nil {
			return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
		}
	}

	qSortBy := c.Query("sortBy")
	if qLimit != "" {
		sortBy = qSortBy
	}

	paginationParams := common.PaginationParams{
		common.GenerateOffset(page, limit),
		limit,
		sortBy,
	}

	qName := c.Query("name")
	if qName != "" {
		data, errorDomain = r.TmPositionService.GetByName(qName)
		if errorDomain != nil {
			return common_helper.ProcessErrorDomain(c, errorDomain)
		}

		dataCount, errorDomain = r.TmPositionService.CountByName(qName)
		if errorDomain != nil {
			return common_helper.ProcessErrorDomain(c, errorDomain)
		}
	} else {
		data, errorDomain = r.TmPositionService.Get(paginationParams)
		if errorDomain != nil {
			return common_helper.ProcessErrorDomain(c, errorDomain)
		}

		dataCount, errorDomain = r.TmPositionService.Count()
		if errorDomain != nil {
			return common_helper.ProcessErrorDomain(c, errorDomain)
		}
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.SetPagination().Success("", data, page, limit, *dataCount))
}

func (r TmPositionController) GetById(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var id int

	paramId := utils.CopyString(c.Params("id"))
	if paramId != "" {
		id, err = strconv.Atoi(paramId)
		if err != nil {
			return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
		}
	}

	data, errorDomain = r.TmPositionService.GetById(id)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r TmPositionController) Insert(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error

	request := new(models.PrivilegesRequest)

	err = c.BodyParser(request)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	data, errorDomain = r.TmPositionService.Insert(*request)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r TmPositionController) Delete(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var id int

	paramId := utils.CopyString(c.Params("id"))
	if paramId != "" {
		id, err = strconv.Atoi(paramId)
		if err != nil {
			return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
		}
	} else {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "ID is required", nil))
	}

	data, errorDomain = r.TmPositionService.Delete(id)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}
