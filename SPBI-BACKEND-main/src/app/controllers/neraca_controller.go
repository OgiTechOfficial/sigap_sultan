package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/helper/common_helper"
	"sigap-sultan-be/src/app/helper/csv_helper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/services"
	"sigap-sultan-be/src/common"
	"strconv"
)

type NeracaController struct {
	Validator     *common.XValidator
	NeracaService *services.NeracaService
}

func NewNeracaController(validator *common.XValidator, priceService *services.NeracaService) *NeracaController {
	return &NeracaController{
		validator,
		priceService,
	}
}

func (r NeracaController) Upload(c *fiber.Ctx) error {
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
	case "neraca-city":
		_, errorDomain = r.NeracaService.UploadNeracaCity(fileStructure)
		if errorDomain != nil {
			response := models.SetError().Details(500, errorDomain.Message, errorDomain.Details)
			return c.Status(500).JSON(response)
		}
	case "neraca-province":
		_, errorDomain = r.NeracaService.UploadNeracaProvince(fileStructure)
		if errorDomain != nil {
			response := models.SetError().Details(500, errorDomain.Message, errorDomain.Details)
			return c.Status(500).JSON(response)
		}
	case "neraca-national":
		_, errorDomain = r.NeracaService.UploadNeracaNational(fileStructure)
		if errorDomain != nil {
			response := models.SetError().Details(500, errorDomain.Message, errorDomain.Details)
			return c.Status(500).JSON(response)
		}
	default:
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "upload-type must be price-city, price-province, price-national", nil))
	}

	return c.SendStatus(201)
}

func (r NeracaController) GetStockAkhir(c *fiber.Ctx) error {
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
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.NeracaService.GetStockAkhir(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetStockAkhirByCommodityList(c *fiber.Ctx) error {
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
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.NeracaService.GetStockAkhirList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetStockAkhirByCommodityHistory(c *fiber.Ctx) error {
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

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		ProvinceId:   "73",
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
		Status:       status,
	}

	data, errorDomain = r.NeracaService.GetStockAkhirByCommodityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetStockAkhirByCommodityCityHistory(c *fiber.Ctx) error {
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

	qStartDate := c.Query("startDate")
	if qStartDate != "" {
		startDate = qStartDate
	}

	qEndDate := c.Query("endDate")
	if qEndDate != "" {
		endDate = qEndDate
	}

	cityId := c.Query("cityId")
	if cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId is required", nil))
	}

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		ProvinceId:  "73",
		CityId:      cityId,
		CommodityId: commodityId,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	data, errorDomain = r.NeracaService.GetStockAkhirByCommodityCityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKetesediaanMap(c *fiber.Ctx) error {
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
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.NeracaService.GetKetersediaanMap(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKetersediaanList(c *fiber.Ctx) error {
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
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.NeracaService.GetKetersediaanList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKetersediaanByCommodityHistory(c *fiber.Ctx) error {
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

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	qStatus := c.Query("status")
	if qStatus != "" {
		status = qStatus
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		ProvinceId:   "73",
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
		Status:       status,
	}

	data, errorDomain = r.NeracaService.GetKetersediaanByCommodityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKetersediaanByCommodityCityHistory(c *fiber.Ctx) error {
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

	qStartDate := c.Query("startDate")
	if qStartDate != "" {
		startDate = qStartDate
	}

	qEndDate := c.Query("endDate")
	if qEndDate != "" {
		endDate = qEndDate
	}

	cityId := c.Query("cityId")
	if cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId is required", nil))
	}

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		ProvinceId:  "73",
		CityId:      cityId,
		CommodityId: commodityId,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	data, errorDomain = r.NeracaService.GetKetersediaanByCommodityCityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKebutuhanMap(c *fiber.Ctx) error {
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
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.NeracaService.GetKebutuhanMap(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKebutuhanList(c *fiber.Ctx) error {
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
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.NeracaService.GetKebutuhanList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKebutuhanByCommodityHistory(c *fiber.Ctx) error {
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

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		ProvinceId:   "73",
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
		Status:       status,
	}

	data, errorDomain = r.NeracaService.GetKebutuhanByCommodityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKebutuhanByCommodityCityHistory(c *fiber.Ctx) error {
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

	qStartDate := c.Query("startDate")
	if qStartDate != "" {
		startDate = qStartDate
	}

	qEndDate := c.Query("endDate")
	if qEndDate != "" {
		endDate = qEndDate
	}

	cityId := c.Query("cityId")
	if cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId is required", nil))
	}

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		ProvinceId:  "73",
		CityId:      cityId,
		CommodityId: commodityId,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	data, errorDomain = r.NeracaService.GetKebutuhanByCommodityCityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetStockAkhirByCityAndCommodityId(c *fiber.Ctx) error {
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

	cityId := c.Params("cityId")
	if cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId is required", nil))
	}

	commodityId := c.Params("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		ProvinceId:  "73",
		CityId:      cityId,
		CommodityId: commodityId,
	}

	data, errorDomain = r.NeracaService.GetStockAkhirByCityAndCommodityId(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetStockAkhirByCityAndCommodityIdChart(c *fiber.Ctx) error {
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

	qStartDate := c.Query("startDate")
	if qStartDate != "" {
		startDate = qStartDate
	}

	qEndDate := c.Query("endDate")
	if qEndDate != "" {
		endDate = qEndDate
	}

	cityId := c.Query("cityId")
	if cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId is required", nil))
	}

	commodityId := c.Query("commodityId")
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		ProvinceId:  "73",
		CityId:      cityId,
		CommodityId: commodityId,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	data, errorDomain = r.NeracaService.GetStockAkhirByCityAndCommodityChart(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) CompareWithPriceCommodityHistory(c *fiber.Ctx) error {
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

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CityId:      cityId,
		ProvinceId:  "73",
		CommodityId: commodityId,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	data, errorDomain = r.NeracaService.CompareWithPriceCommodityHistory(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetStockAkhirByCity(c *fiber.Ctx) error {
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

	cityId := c.Query("cityId")
	if cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		SelectedDate: selectedDate,
		CityId:       cityId,
		ProvinceId:   "73",
	}

	data, errorDomain = r.NeracaService.GetStockAkhirByCity(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetStockAkhirByCityList(c *fiber.Ctx) error {
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
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.NeracaService.GetStockAkhirByCityList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKetersediaanByCity(c *fiber.Ctx) error {
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

	cityId := c.Query("cityId")
	if cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		SelectedDate: selectedDate,
		CityId:       cityId,
		ProvinceId:   "73",
	}

	data, errorDomain = r.NeracaService.GetKetersediaanByCity(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKetersediaanByCityList(c *fiber.Ctx) error {
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
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.NeracaService.GetKetersediaanByCityList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKebutuhanByCity(c *fiber.Ctx) error {
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

	cityId := c.Query("cityId")
	if cityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		SelectedDate: selectedDate,
		CityId:       cityId,
		ProvinceId:   "73",
	}

	data, errorDomain = r.NeracaService.GetKebutuhanByCity(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) GetKebutuhanByCityList(c *fiber.Ctx) error {
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
	if commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "commodityId is required", nil))
	}

	selectedDate := c.Query("selectedDate")
	if selectedDate == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "selectedDate is required", nil))
	}

	params := domain.NeracaStokAkhirListRequestParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
		},
		CommodityId:  commodityId,
		SelectedDate: selectedDate,
	}

	data, errorDomain = r.NeracaService.GetKebutuhanByCityList(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) Exist(c *fiber.Ctx) error {
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

	params := domain.NeracaStokAkhirListRequestParams{
		ProvinceId:  "73",
		CommodityId: commodityId,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	data, errorDomain = r.NeracaService.Exist(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r NeracaController) LatestDateExist(c *fiber.Ctx) error {
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

	data, errorDomain = r.NeracaService.LatestDateExist(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}
