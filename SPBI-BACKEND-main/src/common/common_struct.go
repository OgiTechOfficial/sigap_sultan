package common

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	CommonParams struct {
		Ctx *fiber.Ctx
	}

	EnvConfig struct {
		AddrListen string

		DbHost     string
		DbPort     int16
		DbUsername string
		DbPass     string
		DbName     string
		DbScheme   string
	}

	ProvinceIdParam struct {
		ProvinceId string `json:"provinceId"`
	}

	CityIdParam struct {
		CityId string `json:"cityId"`
	}

	CommodityIdParam struct {
		CommodityId string `json:"commodityId"`
	}

	SelectedDateParam struct {
		SelectedDate string `json:"selectedDate"`
	}

	PeriodDateParam struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}

	SqlColumn struct {
		LastUpdate *string `json:"lastUpdate" db:"last_update"`
		CreatedAt  *string `json:"createdAt" db:"created_at"`
		UpdatedAt  *string `json:"updatedAt" db:"updated_at"`
		DeletedAt  *string `json:"deletedAt" db:"deleted_at"`
	}

	FileStructure struct {
		FilePath string
		FileName string
		FileSize string
		FileType string
	}

	PaginationParams struct {
		Page   int    `db:"page"`
		Limit  int    `db:"limit"`
		SortBy string `db:"sortBy"`
	}

	PaginationRespnose struct {
		Page      int `json:"page" db:"page"`
		TotalPage int `json:"totalPage" db:"total_page"`
		TotalData int `json:"totalData" db:"total_data"`
	}

	ErrorValidationResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		Validator *validator.Validate
	}

	ErrorsValidationResponse struct {
		Errors ErrorFieldValidationResponse `json:"errors"`
	}

	ErrorFieldValidationResponse struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	GenerateNewDate struct {
		StartDate string
		EndDate   string
	}
)

func (r PaginationParams) countOffsetLI() PaginationParams {
	if r.Page > 0 {
		r.Page = r.Page * r.Limit
	} else {
		r.Page = 0 * r.Limit
	}

	return r
}

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorValidationResponse {
	validationErrors := []ErrorValidationResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorValidationResponse

			elem.FailedField = err.Field() // Struct field name
			elem.Tag = err.Tag()           // Struct tag
			elem.Value = err.Value()       // Struct field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
