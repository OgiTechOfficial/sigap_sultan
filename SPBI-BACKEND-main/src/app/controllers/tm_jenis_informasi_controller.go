package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"net/http"
	"sigap-sultan-be/src/app/helper/common_helper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/services"
	"sigap-sultan-be/src/common"
	"strconv"
)

type TmJenisInformasiController struct {
	TmJenisInformasiService *services.TmJenisInformasiService
}

func NewTmJenisInformasiController(priceService *services.TmJenisInformasiService) *TmJenisInformasiController {
	return &TmJenisInformasiController{
		priceService,
	}
}

func (r TmJenisInformasiController) Get(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

	qPage := c.Query("page")
	if qPage == "" {
		page = 0
	} else {
		page, err = strconv.Atoi(qPage)
		if err != nil {
			return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
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

	qName := c.Query("name")
	if qName != "" {
		data, errorDomain = r.TmJenisInformasiService.GetByName(qName)
		if errorDomain != nil {
			return common_helper.ProcessErrorDomain(c, errorDomain)
		}
	} else {
		data, errorDomain = r.TmJenisInformasiService.Get(common.PaginationParams{
			page,
			limit,
			sortBy,
		})
		if errorDomain != nil {
			return common_helper.ProcessErrorDomain(c, errorDomain)
		}
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r TmJenisInformasiController) GetById(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var id int

	paramId := utils.CopyString(c.Params("id"))
	if paramId != "" {
		id, err = strconv.Atoi(paramId)
		if err != nil {
			return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "ID is required", nil))
		}
	}

	data, errorDomain = r.TmJenisInformasiService.GetById(id)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

//func (r TmJenisInformasiController) Insert(c *fiber.Ctx) error {
//	var data interface{}
//	var errorDomain *common.ErrorDomain
//	var err error
//
//	request := new(models.JenisInformasiRequest)
//
//	err = c.BodyParser(request)
//	if err != nil {
//		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, err.Error(), nil))
//	}
//
//	data, errorDomain = r.TmJenisInformasiService.Insert(*request)
//	if errorDomain != nil {
//		return common_helper.ProcessErrorDomain(c, errorDomain)
//	}
//
//	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
//}
