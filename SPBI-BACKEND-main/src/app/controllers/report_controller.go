package controllers

import (
	"fmt"
	"math"
	"net/http"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/helper/common_helper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/services"
	"sigap-sultan-be/src/common"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/xuri/excelize/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gofiber/fiber/v2"
)

type ReportController struct {
	Validator          *common.XValidator
	ReportService      *services.ReportService
	TmCityService      *services.TmCityService
	TmCommodityService *services.TmCommodityService
}

func NewReportController(validator *common.XValidator, reportService *services.ReportService, tmCity *services.TmCityService, tmCommodityService *services.TmCommodityService) *ReportController {
	return &ReportController{
		validator,
		reportService,
		tmCity,
		tmCommodityService,
	}
}

func (r ReportController) GetPriceReport(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var startDate string
	var endDate string
	var totalFiltered *int
	//var totalData *int

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

	cityId := c.Query("cityId")
	commodityId := c.Query("commodityId")

	if cityId == "" && commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId or commodityId is required", nil))
	} else {
		if (cityId != "" && cityId != "0") && (commodityId != "" && commodityId != "0") {
			return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "Please choose cityId or commodityId not both", nil))
		}
	}

	params := domain.PriceListParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
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

	data, errorDomain = r.ReportService.GetReportPrice(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	totalFiltered, errorDomain = r.ReportService.GetReportCount(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	totalPage := float64(*totalFiltered) / float64(limit)

	return common_helper.GenerateResponseJSON(
		c,
		http.StatusOK, models.SetPagination2().SuccessPagination(
			"",
			map[string]any{
				"prices":    data,
				"page":      page,
				"totalData": totalFiltered,
				"totalPage": math.Ceil(totalPage),
			},
		),
	)
}

func convertIfEmptyMap(data interface{}) interface{} {
	// Cek apakah data adalah map kosong
	if v, ok := data.(map[string]interface{}); ok && len(v) == 0 {
		return []interface{}{} // Ubah ke array kosong
	}
	return data // Jika bukan map kosong, kembalikan data asli
}

func (r ReportController) GetNeracaReport(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var startDate string
	var endDate string
	var totalFiltered *int
	//var totalData *int

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

	cityId := c.Query("cityId")
	commodityId := c.Query("commodityId")

	if cityId == "" && commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId or commodityId is required", nil))
	} else {
		if (cityId != "" && cityId != "0") && (commodityId != "" && commodityId != "0") {
			return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "Please choose cityId or commodityId not both", nil))
		}
	}

	params := domain.NeracaListParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
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

	data, errorDomain = r.ReportService.GetReportNeraca(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	data = convertIfEmptyMap(data)

	totalFiltered, errorDomain = r.ReportService.GetReportNeracaCount(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	totalPage := float64(*totalFiltered) / float64(limit)

	return common_helper.GenerateResponseJSON(
		c,
		http.StatusOK, models.SetPagination2().SuccessPagination(
			"",
			map[string]any{
				"informationTypes": []string{
					"ketersediaan",
					"kebutuhan",
					"neraca",
				},
				"data":      data,
				"page":      page,
				"totalData": totalFiltered,
				"totalPage": math.Ceil(totalPage),
			},
		),
	)
	//return common_helper.GenerateResponseJSON(
	//	c,
	//	200,
	//	map[string]any{
	//		"status":  200,
	//		"message": "Success",
	//		"informationTypes": []string{
	//			"ketersediaan",
	//			"kebutuhan",
	//			"neraca",
	//		},
	//		"data":      data,
	//		"page":      page,
	//		"totalData": totalFiltered,
	//		"totalPage": math.Ceil(totalPage),
	//	},
	//)
}

func (r ReportController) GetNeracaReportDownload(c *fiber.Ctx) error {
	f := excelize.NewFile()
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var startDate string
	var endDate string
	//var totalFiltered *int
	//var totalData *int
	var sheetDefault = "Sheet1"
	var headersDefaultByCity = []interface{}{
		"Nama Komoditas",
		"Jenis Informasi",
	}
	var headersDefaultByCommodity = []interface{}{
		"Nama Daerah",
		"Jenis Informasi",
	}
	var defaultJenisInformasi = []string{
		"ketersediaan",
		"kebutuhan",
		"neraca",
	}

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

	cityId := c.Query("cityId")
	commodityId := c.Query("commodityId")

	if cityId == "" && commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId or commodityId is required", nil))
	} else {
		if (cityId != "" && cityId != "0") && (commodityId != "" && commodityId != "0") {
			return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "Please choose cityId or commodityId not both", nil))
		}
	}

	params := domain.NeracaListParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
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

	var tmCommodity *models.TmCommodity
	var tmCity *models.TmCity
	var errDomain *common.ErrorDomain
	var queryResult interface{}
	if commodityId != "" && commodityId != "0" {
		if commodityId, err := strconv.Atoi(commodityId); err == nil {
			if queryResult, errDomain = r.TmCommodityService.GetById(commodityId); errDomain != nil {
				return common_helper.ProcessErrorDomain(c, errorDomain)
			} else {
				tmCommodity = queryResult.(*models.TmCommodity)
			}
		}
	} else {
		if cityId, err := strconv.Atoi(cityId); err == nil {
			if queryResult, errDomain = r.TmCityService.GetById(cityId); errDomain != nil {
				return common_helper.ProcessErrorDomain(c, errorDomain)
			} else {
				tmCity = queryResult.(*models.TmCity)
			}
		}
	}

	data, errorDomain = r.ReportService.GetReportNeracaDownload(params)
	if errorDomain != nil {
		return common_helper.GenerateResponseJSON(c, fiber.StatusInternalServerError, models.SetError().Details(fiber.StatusInternalServerError, "Data Tidak Ditemukan atau Kosong", errorDomain))
		// return common_helper.ProcessErrorDomain(c, errorDomain)
	}
	if data == nil {
		erDom := &common.ErrorDomain{
			Message: "No Data Found",
		}
		return common_helper.ProcessErrorDomain(c, erDom)
	}

	if data == nil {
		headerInformation := make([][]interface{}, 0)
		headerInformation = neracaCreateHeaderInformation(tmCommodity, tmCity, startDate, endDate)
		for i, row := range headerInformation {
			startCell, err := excelize.JoinCellName("A", i+1)
			if err != nil {
				log.Error(err)
			}
			if err := f.SetSheetRow(sheetDefault, startCell, &row); err != nil {
				log.Error(err)
			}
		}

		common.SetHeadersBold(sheetDefault, f)
		common.SetWidthColXlsx(sheetDefault, f)

		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		if commodityId != "" && commodityId != "0" {
			if err := f.SaveAs("./neraca-by-commodity.xlsx"); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := f.SaveAs("./neraca-by-city.xlsx"); err != nil {
				fmt.Println(err)
			}
		}

		var neracaBasedOnName string

		if commodityId != "" && commodityId != "0" {
			if commodityId, err := strconv.Atoi(commodityId); err == nil {
				if commodity, errDomain := r.TmCommodityService.GetById(commodityId); errDomain != nil {
					return common_helper.ProcessErrorDomain(c, errorDomain)
				} else {
					tmCommodity := commodity.(*models.TmCommodity)
					neracaBasedOnName = tmCommodity.Name
				}
			}

			c.Set(fiber.HeaderContentDisposition, `attachment; filename="REPORT_NERACA_KOMODITAS_`+strings.ToUpper(neracaBasedOnName)+`_PERIODE_`+startDate+`_`+endDate+`.xlsx"`)
			return c.SendFile("./neraca-by-commodity.xlsx")
		} else {
			if cityId, err := strconv.Atoi(cityId); err == nil {
				if city, errDomain := r.TmCityService.GetById(cityId); errDomain != nil {
					return common_helper.ProcessErrorDomain(c, errorDomain)
				} else {
					tmCity := city.(*models.TmCity)
					neracaBasedOnName = tmCity.Name
				}
			}

			c.Set(fiber.HeaderContentDisposition, `attachment; filename="REPORT_NERACA_WILAYAH_`+strings.ToUpper(neracaBasedOnName)+`_PERIODE_`+startDate+`_`+endDate+`.xlsx"`)
			return c.SendFile("./neraca-by-city.xlsx")
		}
	}

	dataConverted := *data.(*[]map[string]interface{})
	dataForXlsx := make([][]interface{}, 0)

	headerInformation := make([][]interface{}, 0)
	headerInformation = neracaCreateHeaderInformation(tmCommodity, tmCity, startDate, endDate)
	for i, row := range headerInformation {
		startCell, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			log.Error(err)
		}
		if err := f.SetSheetRow(sheetDefault, startCell, &row); err != nil {
			log.Error(err)
		}
	}

	dataForXlsx = [][]interface{}{}
	if commodityId != "" && commodityId != "0" {
		neracaAppendDateHeader(dataConverted, &headersDefaultByCommodity)
		dataForXlsx = append(dataForXlsx, headersDefaultByCommodity)
	} else {
		neracaAppendDateHeader(dataConverted, &headersDefaultByCity)
		dataForXlsx = append(dataForXlsx, headersDefaultByCity)
	}

	for _, dataMap := range dataConverted {
		stocks := dataMap["stocks"].(map[string]map[string]int)
		for _, jenisInformasi := range defaultJenisInformasi {
			valueStock := stocks[jenisInformasi]
			if jenisInformasi == "ketersediaan" {
				ketersediaan := []interface{}{
					dataMap["title"],
					cases.Title(language.Indonesian).String(jenisInformasi),
				}

				valueStockKeyToSlice := make([]string, 0, len(valueStock))
				for valueStockKey := range valueStock {
					valueStockKeyToSlice = append(valueStockKeyToSlice, valueStockKey)
				}
				sort.Strings(valueStockKeyToSlice)
				for _, valueStockKey := range valueStockKeyToSlice {
					ketersediaan = append(ketersediaan, common.ThousandFormat(int32(valueStock[valueStockKey])))
				}

				dataForXlsx = append(dataForXlsx, ketersediaan)
			} else if jenisInformasi == "kebutuhan" {
				kebutuhan := []interface{}{
					"",
					cases.Title(language.Indonesian).String(jenisInformasi),
				}

				valueStockKeyToSlice := make([]string, 0, len(valueStock))
				for valueStockKey := range valueStock {
					valueStockKeyToSlice = append(valueStockKeyToSlice, valueStockKey)
				}

				sort.Strings(valueStockKeyToSlice)
				for _, valueStockKey := range valueStockKeyToSlice {
					kebutuhan = append(kebutuhan, common.ThousandFormat(int32(valueStock[valueStockKey])))
				}

				dataForXlsx = append(dataForXlsx, kebutuhan)
			} else {
				neraca := []interface{}{
					"",
					cases.Title(language.Indonesian).String(jenisInformasi),
				}

				valueStockKeyToSlice := make([]string, 0, len(valueStock))
				for valueStockKey := range valueStock {
					valueStockKeyToSlice = append(valueStockKeyToSlice, valueStockKey)
				}

				sort.Strings(valueStockKeyToSlice)
				for _, valueStockKey := range valueStockKeyToSlice {
					neraca = append(neraca, common.ThousandFormat(int32(valueStock[valueStockKey])))
				}

				dataForXlsx = append(dataForXlsx, neraca)
			}
		}
	}

	mapForMerge := make(map[string]string)
	lenHeaderInformation := len(headerInformation)
	for _, row := range dataForXlsx {
		startCell, err := excelize.JoinCellName("A", lenHeaderInformation+2)
		if err != nil {
			log.Error(err)
		}
		if lenHeaderInformation > len(headerInformation) {
			getCellAbjad := startCell[0:1]
			getCellIdx := startCell[1:]
			if cellIdx, err := strconv.Atoi(getCellIdx); err != nil {
				log.Error(err)
			} else {
				endCellMerge := cellIdx + 2
				if row[0] != "" {
					mapForMerge[row[0].(string)] = startCell + ":" + getCellAbjad + strconv.Itoa(endCellMerge)
				}
			}

			if err := f.SetSheetRow(sheetDefault, startCell, &row); err != nil {
				log.Error(err)
			}
		} else {
			if err := f.SetSheetRow(sheetDefault, startCell, &row); err != nil {
				log.Error(err)
			}
		}

		lenHeaderInformation++
	}

	// SETUP MERGE CELL FOR TITLE AND SETUP VERTICAL CENTER
	// EXP: A2:A4
	//	===================
	//	|| NAMA KOMODITAS ||
	//	||                ||
	//	|| BAWANG MERAH   ||
	//	||				  ||
	//	==================||
	for rowTitle, val := range mapForMerge {
		cellSplit := strings.Split(val, ":")
		startCell := cellSplit[0]
		endCellAbjad := cellSplit[1]
		if err := f.MergeCell(sheetDefault, startCell, endCellAbjad); err != nil {
			log.Error(err)
		}
		if err := f.SetCellValue(sheetDefault, startCell, rowTitle); err != nil {
			log.Error(err)
		}
		if style, err := f.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Vertical: "center",
			},
		}); err != nil {
			log.Error(err)
		} else {
			if err = f.SetCellStyle(sheetDefault, startCell, startCell, style); err != nil {
				log.Error(err)
			}
		}
	}

	common.SetHeadersBold(sheetDefault, f)
	common.SetWidthColXlsx(sheetDefault, f)

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if commodityId != "" && commodityId != "0" {
		if err := f.SaveAs("./neraca-by-commodity.xlsx"); err != nil {
			fmt.Println(err)
		}
	} else {
		if err := f.SaveAs("./neraca-by-city.xlsx"); err != nil {
			fmt.Println(err)
		}
	}

	var neracaBasedOnName string

	if commodityId != "" && commodityId != "0" {
		if commodityId, err := strconv.Atoi(commodityId); err == nil {
			if commodity, errDomain := r.TmCommodityService.GetById(commodityId); errDomain != nil {
				return common_helper.ProcessErrorDomain(c, errorDomain)
			} else {
				tmCommodity := commodity.(*models.TmCommodity)
				neracaBasedOnName = tmCommodity.Name
			}
		}

		c.Set(fiber.HeaderContentDisposition, `attachment; filename="REPORT_NERACA_KOMODITAS_`+strings.ToUpper(neracaBasedOnName)+`_PERIODE_`+startDate+`_`+endDate+`.xlsx"`)
		return c.SendFile("./neraca-by-commodity.xlsx")
	} else {
		if cityId, err := strconv.Atoi(cityId); err == nil {
			if city, errDomain := r.TmCityService.GetById(cityId); errDomain != nil {
				return common_helper.ProcessErrorDomain(c, errorDomain)
			} else {
				tmCity := city.(*models.TmCity)
				neracaBasedOnName = tmCity.Name
			}
		}

		c.Set(fiber.HeaderContentDisposition, `attachment; filename="REPORT_NERACA_WILAYAH_`+strings.ToUpper(neracaBasedOnName)+`_PERIODE_`+startDate+`_`+endDate+`.xlsx"`)
		return c.SendFile("./neraca-by-city.xlsx")
	}
}

func neracaCreateHeaderInformation(tmCommodity *models.TmCommodity, tmCity *models.TmCity, startDate string, endDate string) [][]interface{} {
	if startDateTime, err := time.Parse("2006-01-02", startDate); err == nil {
		startDate = startDateTime.Format("Jan 2006")
	}

	if endDateTime, err := time.Parse("2006-01-02", endDate); err == nil {
		endDate = endDateTime.Format("Jan 2006")
	}

	startDateSplit := strings.Split(startDate, " ")
	startDate = common.BulanIndonesiaFromEnglishShortName[strings.ToLower(startDateSplit[0])] + " " + startDateSplit[1]

	endDateSplit := strings.Split(endDate, " ")
	endDate = common.BulanIndonesiaFromEnglishShortName[strings.ToLower(endDateSplit[0])] + " " + endDateSplit[1]

	result := [][]interface{}{
		{
			"Jenis Laporan", "Neraca", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		},
	}

	if tmCommodity != nil {
		result = append(result, []interface{}{
			"Komoditas", tmCommodity.Name,
		})
		result = append(result, []interface{}{
			"Daerah", "-",
		})
		result = append(result, []interface{}{
			"Periode Laporan", "Bulanan",
		})
		result = append(result, []interface{}{
			"Periode Mulai - Periode Akhir", startDate + " - " + endDate,
		})
	} else {
		result = append(result, []interface{}{
			"Komoditas", "-",
		})
		result = append(result, []interface{}{
			"Daerah", cases.Title(language.Indonesian).String(tmCity.Name),
		})
		result = append(result, []interface{}{
			"Periode Laporan", "Bulanan",
		})
		result = append(result, []interface{}{
			"Periode Mulai - Periode Akhir", startDate + " - " + endDate,
		})
	}

	return result
}

func (r ReportController) GetPriceReportDownload(c *fiber.Ctx) error {
	f := excelize.NewFile()
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error
	var page int
	var limit int
	var sortBy string
	var startDate string
	var endDate string
	var sheetDefault = "Sheet1"
	var priceHeadersDefaultByCity = []interface{}{
		"Nama Komoditas",
	}
	var priceHeadersDefaultByCommodity = []interface{}{
		"Nama Daerah",
	}

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

	cityId := c.Query("cityId")
	commodityId := c.Query("commodityId")

	if cityId == "" && commodityId == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "cityId or commodityId is required", nil))
	} else {
		if (cityId != "" && cityId != "0") && (commodityId != "" && commodityId != "0") {
			return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "Please choose cityId or commodityId not both", nil))
		}
	}

	params := domain.PriceListParams{
		CommonParams: common.CommonParams{Ctx: c},
		PaginationParams: common.PaginationParams{
			common.GenerateOffset(page, limit),
			limit,
			sortBy,
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

	var tmCommodity *models.TmCommodity
	var tmCity *models.TmCity
	var errDomain *common.ErrorDomain
	var queryResult interface{}
	if commodityId != "" && commodityId != "0" {
		if commodityId, err := strconv.Atoi(commodityId); err == nil {
			if queryResult, errDomain = r.TmCommodityService.GetById(commodityId); errDomain != nil {
				return common_helper.ProcessErrorDomain(c, errorDomain)
			} else {
				tmCommodity = queryResult.(*models.TmCommodity)
			}
		}
	} else {
		if cityId, err := strconv.Atoi(cityId); err == nil {
			if queryResult, errDomain = r.TmCityService.GetById(cityId); errDomain != nil {
				return common_helper.ProcessErrorDomain(c, errorDomain)
			} else {
				tmCity = queryResult.(*models.TmCity)
			}
		}
	}

	data, errorDomain = r.ReportService.GetReportPriceDownload(params)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	if data == nil {
		headerInformation := make([][]interface{}, 0)
		headerInformation = priceCreateHeaderInformation(tmCommodity, tmCity, startDate, endDate)
		for i, row := range headerInformation {
			startCell, err := excelize.JoinCellName("A", i+1)
			if err != nil {
				log.Error(err)
			}
			if err := f.SetSheetRow(sheetDefault, startCell, &row); err != nil {
				log.Error(err)
			}
		}

		common.SetHeadersBold(sheetDefault, f)
		common.SetWidthColXlsx(sheetDefault, f)

		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		if commodityId != "" && commodityId != "0" {
			if err := f.SaveAs("./price-by-commodity.xlsx"); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := f.SaveAs("./price-by-city.xlsx"); err != nil {
				fmt.Println(err)
			}
		}

		var priceBasedOnCityName string

		if commodityId != "" && commodityId != "0" {
			if commodityId, err := strconv.Atoi(commodityId); err == nil {
				if commodity, errDomain := r.TmCommodityService.GetById(commodityId); errDomain != nil {
					return common_helper.ProcessErrorDomain(c, errorDomain)
				} else {
					tmCommodity := commodity.(*models.TmCommodity)
					priceBasedOnCityName = tmCommodity.Name
				}
			}

			c.Set(fiber.HeaderContentDisposition, `attachment; filename="REPORT_PRICE_KOMODITAS_`+strings.ToUpper(priceBasedOnCityName)+`_PERIODE_`+startDate+`_`+endDate+`.xlsx"`)
			return c.SendFile("./price-by-commodity.xlsx")
		} else {
			if cityId, err := strconv.Atoi(cityId); err == nil {
				if city, errDomain := r.TmCityService.GetById(cityId); errDomain != nil {
					return common_helper.ProcessErrorDomain(c, errorDomain)
				} else {
					tmCity := city.(*models.TmCity)
					priceBasedOnCityName = tmCity.Name
				}
			}

			c.Set(fiber.HeaderContentDisposition, `attachment; filename="REPORT_PRICE_WILAYAH_`+strings.ToUpper(priceBasedOnCityName)+`_PERIODE_`+startDate+`_`+endDate+`.xlsx"`)
			return c.SendFile("./price-by-city.xlsx")
		}
	}

	dataConverted := *data.(*[]map[string]string)
	dataForXlsx := make([][]interface{}, 0)

	headerInformation := make([][]interface{}, 0)
	headerInformation = priceCreateHeaderInformation(tmCommodity, tmCity, startDate, endDate)
	for i, row := range headerInformation {
		startCell, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			log.Error(err)
		}
		if err := f.SetSheetRow(sheetDefault, startCell, &row); err != nil {
			log.Error(err)
		}
	}

	if commodityId != "" && commodityId != "0" {
		priceAppendDateHeader(dataConverted, &priceHeadersDefaultByCommodity)
		dataForXlsx = append(dataForXlsx, priceHeadersDefaultByCommodity)
	} else {
		priceAppendDateHeader(dataConverted, &priceHeadersDefaultByCity)
		dataForXlsx = append(dataForXlsx, priceHeadersDefaultByCity)
	}

	for _, dataMap := range dataConverted {
		valueStockKeyToSlice := make([]string, 0, len(dataMap))
		for valueStockKey := range dataMap {
			if valueStockKey != "title" {
				valueStockKeyToSlice = append(valueStockKeyToSlice, valueStockKey)
			}
		}
		sort.Strings(valueStockKeyToSlice)

		finalValueStockKey := make([]string, 0, len(dataMap))
		finalValueStockKey = append(finalValueStockKey, "title")
		for _, finalKey := range valueStockKeyToSlice {
			finalValueStockKey = append(finalValueStockKey, finalKey)
		}

		valueForDataXlsxSlice := make([]interface{}, 0, len(dataMap))
		for _, dataKey := range finalValueStockKey {
			valueForDataXlsxSlice = append(valueForDataXlsxSlice, dataMap[dataKey])
		}
		dataForXlsx = append(dataForXlsx, valueForDataXlsxSlice)
	}

	lenHeaderInformation := len(headerInformation)
	for _, row := range dataForXlsx {
		startCell, err := excelize.JoinCellName("A", lenHeaderInformation+2)
		if err != nil {
			log.Error(err)
		}
		if err := f.SetSheetRow(sheetDefault, startCell, &row); err != nil {
			log.Error(err)
		}

		lenHeaderInformation++
	}

	common.SetHeadersBold(sheetDefault, f)
	common.SetWidthColXlsx(sheetDefault, f)

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if commodityId != "" && commodityId != "0" {
		if err := f.SaveAs("./price-by-commodity.xlsx"); err != nil {
			fmt.Println(err)
		}
	} else {
		if err := f.SaveAs("./price-by-city.xlsx"); err != nil {
			fmt.Println(err)
		}
	}

	var priceBasedOnCityName string

	if commodityId != "" && commodityId != "0" {
		if commodityId, err := strconv.Atoi(commodityId); err == nil {
			if commodity, errDomain := r.TmCommodityService.GetById(commodityId); errDomain != nil {
				return common_helper.ProcessErrorDomain(c, errorDomain)
			} else {
				tmCommodity := commodity.(*models.TmCommodity)
				priceBasedOnCityName = tmCommodity.Name
			}
		}

		c.Set(fiber.HeaderContentDisposition, `attachment; filename="REPORT_PRICE_KOMODITAS_`+strings.ToUpper(priceBasedOnCityName)+`_PERIODE_`+startDate+`_`+endDate+`.xlsx"`)
		return c.SendFile("./price-by-commodity.xlsx")
	} else {
		if cityId, err := strconv.Atoi(cityId); err == nil {
			if city, errDomain := r.TmCityService.GetById(cityId); errDomain != nil {
				return common_helper.ProcessErrorDomain(c, errorDomain)
			} else {
				tmCity := city.(*models.TmCity)
				priceBasedOnCityName = tmCity.Name
			}
		}

		c.Set(fiber.HeaderContentDisposition, `attachment; filename="REPORT_PRICE_WILAYAH_`+strings.ToUpper(priceBasedOnCityName)+`_PERIODE_`+startDate+`_`+endDate+`.xlsx"`)
		return c.SendFile("./price-by-city.xlsx")
	}
}

func priceCreateHeaderInformation(tmCommodity *models.TmCommodity, tmCity *models.TmCity, startDate string, endDate string) [][]interface{} {
	if startDateTime, err := time.Parse("2006-01-02", startDate); err == nil {
		startDate = startDateTime.Format("Jan 2006")
	}

	if endDateTime, err := time.Parse("2006-01-02", endDate); err == nil {
		endDate = endDateTime.Format("Jan 2006")
	}

	startDateSplit := strings.Split(startDate, " ")
	startDate = common.BulanIndonesiaFromEnglishShortName[strings.ToLower(startDateSplit[0])] + " " + startDateSplit[1]

	endDateSplit := strings.Split(endDate, " ")
	endDate = common.BulanIndonesiaFromEnglishShortName[strings.ToLower(endDateSplit[0])] + " " + endDateSplit[1]

	result := [][]interface{}{
		{
			"Jenis Laporan", "Harga", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		},
	}

	if tmCommodity != nil {
		result = append(result, []interface{}{
			"Komoditas", tmCommodity.Name,
		})
		result = append(result, []interface{}{
			"Daerah", "-",
		})
		result = append(result, []interface{}{
			"Periode Laporan", "Bulanan",
		})
		result = append(result, []interface{}{
			"Periode Mulai - Periode Akhir", startDate + " - " + endDate,
		})
	} else {
		result = append(result, []interface{}{
			"Komoditas", "-",
		})
		result = append(result, []interface{}{
			"Daerah", cases.Title(language.Indonesian).String(tmCity.Name),
		})
		result = append(result, []interface{}{
			"Periode Laporan", "Bulanan",
		})
		result = append(result, []interface{}{
			"Periode Mulai - Periode Akhir", startDate + " - " + endDate,
		})
	}

	return result
}

func neracaAppendDateHeader(dataConverted []map[string]interface{}, headersDefaultByCity *[]interface{}) {
	dataMap := dataConverted[0]

	stocks := dataMap["stocks"].(map[string]map[string]int)

	mapOfPeriode := stocks["kebutuhan"]

	firstNames := make([]string, 0, len(mapOfPeriode))
	for k, _ := range mapOfPeriode {
		firstNames = append(firstNames, k)
	}
	sort.Strings(firstNames)
	for _, periodeKey := range firstNames {
		t, err := time.Parse("200601", periodeKey)
		if err != nil {
			panic(err)
		}

		periodeParsed := t.Format("January 2006")
		periodeParsedSplit := strings.Split(periodeParsed, " ")
		headersDefaultByCityAppended := append(*headersDefaultByCity, common.BulanIndonesiaFromEnglish[strings.ToLower(periodeParsedSplit[0])]+" "+periodeParsedSplit[1])
		*headersDefaultByCity = headersDefaultByCityAppended
	}
}

func priceAppendDateHeader(dataConverted []map[string]string, headersDefaultByCity *[]interface{}) {
	dataMap := dataConverted[0]

	firstNames := make([]string, 0, len(dataMap))
	for k := range dataMap {
		if k != "title" {
			firstNames = append(firstNames, k)
		}
	}
	sort.Strings(firstNames)
	for _, periodeKey := range firstNames {
		if periodeKey != "title" {
			t, err := time.Parse("012006", periodeKey)
			if err != nil {
				panic(err)
			}

			periodeParsed := t.Format("January 2006")
			periodeParsedSplit := strings.Split(periodeParsed, " ")
			headersDefaultByCityAppended := append(*headersDefaultByCity, common.BulanIndonesiaFromEnglish[strings.ToLower(periodeParsedSplit[0])]+" "+periodeParsedSplit[1])
			*headersDefaultByCity = headersDefaultByCityAppended
		}
	}
}

func generateHeaderInformationByCity(f *excelize.File, params domain.NeracaListParams) {

}
