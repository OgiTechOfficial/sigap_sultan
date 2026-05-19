package controllers

import (
	"github.com/gofiber/fiber/v2"
	"sigap-sultan-be/src/app/helper/common_helper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/common"
)

type AssetsController struct {
	Validator *common.XValidator
}

func NewAssetsController(validator *common.XValidator) *AssetsController {
	return &AssetsController{
		validator,
	}
}

func (r *AssetsController) GetAssets(c *fiber.Ctx) error {
	params := models.AssetsRequestParams{
		AssetsLocation: c.Query("assets_location"),
	}

	if errs := r.Validator.Validate(params); len(errs) > 0 && errs[0].Error {
		errorFields := make([]common.ErrorFieldValidationResponse, 0)

		for _, err := range errs {
			errorFields = append(errorFields, common.ErrorFieldValidationResponse{
				Field:   err.FailedField,
				Message: "Must be " + err.Tag,
			})
		}

		return common_helper.GenerateResponseJSON(c, fiber.ErrBadRequest.Code, models.SetError().Details(fiber.ErrBadRequest.Code, "Bad Request", errorFields))
	}

	c.Set("Content-Type", "image/png")
	return c.SendFile(params.AssetsLocation)
}

func (r *AssetsController) Get404(c *fiber.Ctx) error {
	c.Set("Content-Type", "image/png")
	return c.SendFile("resources/common-logo/404.png")
}
