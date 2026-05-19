package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/helper/common_helper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/services"
	"sigap-sultan-be/src/common"
)

type TmCommodityController struct {
	TmCommodityService *services.TmCommodityService
}

func NewTmCommodityController(priceService *services.TmCommodityService) *TmCommodityController {
	return &TmCommodityController{
		priceService,
	}
}

func (r TmCommodityController) Get(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var moduleType string

	qModuleType := c.Query("moduleType")
	if qModuleType != "" {
		moduleType = qModuleType
	}

	data, errorDomain = r.TmCommodityService.Get(
		domain.TmCommodityRequestParam{
			ModuleType: moduleType,
		},
	)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r TmCommodityController) GetGrouping(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain

	data, errorDomain = r.TmCommodityService.GetGrouping()
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}
