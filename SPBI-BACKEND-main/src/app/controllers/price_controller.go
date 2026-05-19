package controllers

import (
	"fmt"
	"net/http"
	"os"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/helper/common_helper"
	"sigap-sultan-be/src/app/helper/csv_helper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/services"
	"sigap-sultan-be/src/common"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type PriceController struct {
	Validator    *common.XValidator
	PriceService *services.PriceService
}

func NewPriceController(validator *common.XValidator, priceService *services.PriceService) *PriceController {
	return &PriceController{
		validator,
		priceService,
	}
}

func (r PriceController) Upload(c *fiber.Ctx) error {
	var err error
	var errorDomain *common.ErrorDomain

	uploadType := c.FormValue("upload-type")
	if uploadType == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "upload-type is required, and must be price-city, price-province, price-national", nil))
	}

	file, err := c.FormFile("file")
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "Form is required! "+err.Error(), nil))
	}

	if file.Header.Get("Content-Type") != "text/csv" {
		response := models.SetError().Details(400, "Invalid file type", nil)
		return c.Status(400).JSON(response)
	}

	temp, err := os.MkdirTemp(".", "temp-files")
	if err != nil {
		return err
	}
	defer csv_helper.DeleteTempDir(c, temp)

	filePath := fmt.Sprintf("%s/%s", temp, file.Filename)
	err = c.SaveFile(file, filePath)
	if err != nil {
		response := models.SetError().Details(500, err.Error(), nil)
		return c.Status(500).JSON(response)
	}

	fileStructure := common.FileStructure{
		FilePath: filePath,
		FileName: file.Filename,
		FileSize: strconv.FormatInt(file.Size, 10),
		FileType: file.Header.Get("Content-Type"),
	}

	switch uploadType {
	case "price-city":
		_, errorDomain = r.PriceService.UploadPriceCity(fileStructure)
		if errorDomain != nil {
			response := models.SetError().Details(500, errorDomain.Message, errorDomain.Details)
			return c.Status(500).JSON(response)
		}
	case "price-province":
		_, errorDomain = r.PriceService.UploadPriceProvince(fileStructure)
		if errorDomain != nil {
			response := models.SetError().Details(500, errorDomain.Message, errorDomain.Details)
			return c.Status(500).JSON(response)
		}
	case "price-national":
		_, errorDomain = r.PriceService.UploadPriceNational(fileStructure)
		if errorDomain != nil {
			response := models.SetError().Details(500, errorDomain.Message, errorDomain.Details)
			return c.Status(500).JSON(response)
		}
	default:
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "upload-type must be price-city, price-province, price-national", nil))
	}

	return c.SendStatus(201)
}

func (r PriceController) UploadHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var dataCount *int

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

	module := c.Query("module")
	search := c.Query("search")
	params := domain.HistoryRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		Module: module,
		Search: &search,
	}

	data, errorDomain = r.PriceService.GetHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	dataCount, errorDomain = r.PriceService.GetHistoryCount(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.SetPagination().Success("", data, page, limit, *dataCount))
}

func (r PriceController) GetLevelHarga(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")
	//if commodityId == "" {
	//	return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	//}

	selectedDate := c.Query("selectedDate")
	//if selectedDate == "" {
	//	return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	//}

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetLevelHarga(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetLevelHargaNew(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")
	//if commodityId == "" {
	//	return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	//}

	selectedDate := c.Query("selectedDate")
	//if selectedDate == "" {
	//	return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	//}

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetLevelHargaNew(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetLevelHargaList(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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
	if qSortBy != "" {
		split := strings.Split(qSortBy, ":")
		if len(split) == 0 {
			qSortBy = "price:desc"
			//return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "sortBy must be formatted: price:desc", nil))
		} else {
			if len(split) > 1 {
				if split[0] != "price" {
					qSortBy = "price:desc"
					//return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "sortBy column order options must be: price", nil))
				}

				if split[1] != "" {
					if split[1] != "asc" && split[1] != "desc" {
						qSortBy = "price:desc"
						//return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "sortBy order by options must be: asc, desc or blank", nil))
					} else if split[1] == "asc" {
						qSortBy = "price:asc"
					} else {
						qSortBy = "price:desc"
					}
				}
			} else {
				if split[0] != "price" {
					qSortBy = "price:desc"
					//return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "sortBy column order options must be: price", nil))
				}
			}
		}
		sortBy = qSortBy
	}

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  c.Query("commodityId"),
		SelectedDate: c.Query("selectedDate"),
	}

	validationErr := common.Validate(c, params)
	if validationErr != nil {
		return common_helper.GenerateResponseJSON(c, fiber.ErrBadRequest.Code, models.SetError().Details(fiber.ErrBadRequest.Code, "Bad Request", validationErr))
	}

	data, errorDomain = r.PriceService.GetLevelHargaList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetLastFiveDays(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var dataCount *int

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

	cityId := c.Query("cityId")
	commodityId := c.Query("commodityId")

	if cityId == "" && commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId or commodityId is required", nil))
	} else {
		if cityId != "" && commodityId != "" {
			return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "Please choose cityId or commodityId not both", nil))
		}
	}

	var params domain.PriceLast5Days
	if cityId != "" {
		params = domain.PriceLast5Days{
			CommonParams: common.CommonParams{Ctx: c},
			PaginationParams: common.PaginationParams{
				common.GenerateOffset(page, limit),
				limit,
				sortBy,
			},
			CityIdParam: common.CityIdParam{
				CityId: cityId,
			},
		}
	}

	if commodityId != "" {
		params = domain.PriceLast5Days{
			CommonParams: common.CommonParams{Ctx: c},
			PaginationParams: common.PaginationParams{
				common.GenerateOffset(page, limit),
				limit,
				sortBy,
			},
			CommodityIdParam: common.CommodityIdParam{
				CommodityId: commodityId,
			},
		}
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

	data, errorDomain = r.PriceService.GetLast5Days(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	dataCount, errorDomain = r.PriceService.GetLast5DaysCount(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.SetPagination().Success("", data, page, limit, *dataCount))
}

func (r PriceController) Get(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var startDate string
	var endDate string
	var dataCount *int

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
	if qSortBy != "" {
		sortBy = qSortBy
	}

	qStartDate := c.Query("startDate")
	if qStartDate != "" {
		startDate = qStartDate
	}

	qEndDate := c.Query("endDate")
	if qEndDate != "" {
		endDate = qEndDate
	}

	provinceid := c.Query("provinceId")
	cityId := c.Query("cityId")
	if provinceid == "" && cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "provinceId or cityId must be filled", nil))
	}

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	params := domain.PriceListParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		ProvinceIdParam: common.ProvinceIdParam{
			provinceid,
		},
		CityIdParam: common.CityIdParam{
			cityId,
		},

		CommodityIdParam: common.CommodityIdParam{
			commodityId,
		},
		PeriodDateParam: common.PeriodDateParam{
			StartDate: startDate,
			EndDate:   endDate,
		},
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

	data, errorDomain = r.PriceService.Get(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	dataCount, errorDomain = r.PriceService.GetCount(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.SetPagination().Success("", data, page, limit, *dataCount))
}

func (r PriceController) GetDibandingkanProvince(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetDibandingkanSulsel(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetDibandingkanProvinceList(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetDibandingkanSulselList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetCompareProvinceCityHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var startDate string
	var endDate string
	var status string

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

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	cityId := c.Query("cityId")
	if cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId is required", nil))
	}

	qStartDate := c.Query("startDate")
	if qStartDate != "" {
		startDate = qStartDate
	}

	qEndDate := c.Query("endDate")
	if qEndDate != "" {
		endDate = qEndDate
	}

	qStatus := c.Query("status")
	if qStatus != "" {
		status = qStatus
	}

	params := domain.PriceDiffRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId: commodityId,
		CityId:      cityId,
		StartDate:   startDate,
		EndDate:     endDate,
		Status:      status,
	}

	data, errorDomain = r.PriceService.GetCompareProvinceCityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetCompareProvinceCommodityHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var status string

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

	commodityId := c.Query("commodityId")
	selectedDate := c.Query("selectedDate")

	qStatus := c.Query("status")
	if qStatus != "" {
		status = qStatus
	}

	params := domain.PriceGetCompareProvinceCommodityHistoryParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
		Status:       status,
	}

	data, errorDomain = r.PriceService.GetCompareProvinceCommodityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetDibandingkanNasional(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetDibandingkanNational(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetDibandingkanNasionalList(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetDibandingkanNationalList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetCompareNationalCityHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var startDate string
	var endDate string

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

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	cityId := c.Query("cityId")
	if cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId is required", nil))
	}

	qStartDate := c.Query("startDate")
	if qStartDate != "" {
		startDate = qStartDate
	}

	qEndDate := c.Query("endDate")
	if qEndDate != "" {
		endDate = qEndDate
	}

	params := domain.PriceDiffRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId: commodityId,
		CityId:      cityId,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	data, errorDomain = r.PriceService.GetCompareNationalCityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetCompareNationalCommodityHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var status string

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

	qStatus := c.Query("status")
	if qStatus != "" {
		status = qStatus
	}

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetCompareProvinceCommodityHistoryParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
		Status:       status,
	}

	data, errorDomain = r.PriceService.GetCompareNationalCommodityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetByCommodity(c *fiber.Ctx) error {
	return nil
}

func (r PriceController) GetMtM(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetMtm(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetMtMList(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetMtmList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetMtMCityHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var startDate string
	var endDate string

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

	commodityId := c.Query("commodityId")
	cityId := c.Query("cityId")
	provinceId := c.Query("provinceId")

	qStartDate := c.Query("startDate")
	if qStartDate != "" {
		startDate = qStartDate
	}

	qEndDate := c.Query("endDate")
	if qEndDate != "" {
		endDate = qEndDate
	}

	params := domain.PriceMtmCityHistoryParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId: commodityId,
		CityId:      cityId,
		ProvinceId:  provinceId,
		PeriodDateParam: common.PeriodDateParam{
			StartDate: startDate,
			EndDate:   endDate,
		},
	}

	data, errorDomain = r.PriceService.GetMtmCityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetMtMCommodityHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var status string

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

	qStatus := c.Query("status")
	if qStatus != "" {
		status = qStatus
	}

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceMtmCommodityHistoryParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
		Status:       status,
	}

	data, errorDomain = r.PriceService.GetMtmCommodityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetYtd(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetYtd(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetYtdList(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetYtdList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetYtdCityHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var startDate string
	var endDate string

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

	commodityId := c.Query("commodityId")
	cityId := c.Query("cityId")
	provinceId := c.Query("provinceId")

	qStartDate := c.Query("startDate")
	if qStartDate != "" {
		startDate = qStartDate
	}

	qEndDate := c.Query("endDate")
	if qEndDate != "" {
		endDate = qEndDate
	}

	params := domain.PriceMtmCityHistoryParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId: commodityId,
		CityId:      cityId,
		ProvinceId:  provinceId,
		PeriodDateParam: common.PeriodDateParam{
			StartDate: startDate,
			EndDate:   endDate,
		},
	}

	data, errorDomain = r.PriceService.GetYtdCityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetYtdCommodityHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var status string

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

	qStatus := c.Query("status")
	if qStatus != "" {
		status = qStatus
	}

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceMtmCommodityHistoryParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
		Status:       status,
	}

	data, errorDomain = r.PriceService.GetYtdCommodityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetYty(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetYty(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetYtyList(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string

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

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceGetHargaRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.PriceService.GetYtyList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetYtYCityHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var startDate string
	var endDate string

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

	commodityId := c.Query("commodityId")
	cityId := c.Query("cityId")
	provinceId := c.Query("provinceId")

	qStartDate := c.Query("startDate")
	if qStartDate != "" {
		startDate = qStartDate
	}

	qEndDate := c.Query("endDate")
	if qEndDate != "" {
		endDate = qEndDate
	}

	params := domain.PriceMtmCityHistoryParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId: commodityId,
		CityId:      cityId,
		ProvinceId:  provinceId,
		PeriodDateParam: common.PeriodDateParam{
			StartDate: startDate,
			EndDate:   endDate,
		},
	}

	data, errorDomain = r.PriceService.GetYtyCityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) GetYtYCommodityHistory(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var status string

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

	qStatus := c.Query("status")
	if qStatus != "" {
		status = qStatus
	}

	commodityId := c.Query("commodityId")

	selectedDate := c.Query("selectedDate")

	params := domain.PriceMtmCommodityHistoryParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
		Status:       status,
	}

	data, errorDomain = r.PriceService.GetYtyCommodityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) Exist(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	startDate := c.Query("startDate")
	if startDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "startDate is required", nil))
	}

	endDate := c.Query("endDate")
	if endDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "startDate is required", nil))
	}

	params := domain.PriceListRepoParams{
		CommonParams: common.CommonParams{Ctx: c},
		ProvinceId:   "73",
		CommodityId:  commodityId,
		StartDate:    startDate,
		EndDate:      endDate,
	}

	data, errorDomain = r.PriceService.Exist(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r PriceController) LatestDateExist(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	params := domain.PriceListRepoParams{
		ProvinceId:  "73",
		CommodityId: commodityId,
	}

	data, errorDomain = r.PriceService.LatestDateExist(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}
