package services

import (
	"encoding/json"
	"fmt"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/helper/csv_helper"
	"sigap-sultan-be/src/app/helper/string_helper"
	"sigap-sultan-be/src/app/mapper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories"
	"sigap-sultan-be/src/common"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/redis/go-redis/v9"
)

type PriceService struct {
	PriceRepository               *repositories.PriceRepository
	TmCommodityRepository         *repositories.TmCommodityRepository
	TmCityRepository              *repositories.TmCityRepository
	TxFileUploadHistoryRepository *repositories.TxFileUploadHistoryRepository
	TmProvinceRepository          *repositories.TmProvinceRepository
	RedisClient                   *redis.Client
}

func NewPriceService(priceRepository *repositories.PriceRepository, tmCommodityRepository *repositories.TmCommodityRepository, tmCityRepository *repositories.TmCityRepository, txFileUploadHistoryRepository *repositories.TxFileUploadHistoryRepository, TmProvinceRepository *repositories.TmProvinceRepository, redisClient *redis.Client) *PriceService {
	return &PriceService{
		priceRepository,
		tmCommodityRepository,
		tmCityRepository,
		txFileUploadHistoryRepository,
		TmProvinceRepository,
		redisClient,
	}
}

func (r *PriceService) UploadPriceCity(params common.FileStructure) (interface{}, *common.ErrorDomain) {
	var errDomain *common.ErrorDomain
	var priceDatas []*models.PriceCity
	var result interface{}
	var isConvertOk bool

	csvValues, errDomain := csv_helper.CsvFileProcess(params, common.ModuleType{Price: true})
	if errDomain != nil {
		totalErrors := errDomain.Details.([]map[string]interface{})
		txFileUploadHistory := models.TxFileUploadHistory{
			FileName:   params.FileName,
			RowTotal:   len(totalErrors),
			Status:     0,
			ModuleType: common.MODUL_TYPE_PRICE,
		}
		r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	result, errDomain = r.ProcessCsvRows(csvValues, &common.ModuleReportType{PriceCity: true})
	if errDomain != nil {
		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	priceDatas, isConvertOk = result.([]*models.PriceCity)
	if !isConvertOk {
		return nil, &common.ErrorDomain{
			Message: "There's a problem, please contact your administrator!",
			Details: "Convertion interface{} to []*models.PriceCity is failed",
		}
	}

	r.PriceRepository.SavePriceCity(priceDatas)

	txFileUploadHistory := models.TxFileUploadHistory{
		FileName:   params.FileName,
		RowTotal:   len(priceDatas),
		Status:     1,
		ModuleType: common.MODUL_TYPE_PRICE,
	}

	r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

	return nil, nil
}

func (r *PriceService) UploadPriceProvince(params common.FileStructure) (interface{}, *common.ErrorDomain) {
	var errDomain *common.ErrorDomain
	var priceDatas []*models.PriceProvince
	var result interface{}
	var isConvertOk bool

	csvValues, errDomain := csv_helper.CsvFileProcess(params, common.ModuleType{Price: true})
	if errDomain != nil {
		totalErrors := errDomain.Details.([]map[string]interface{})
		txFileUploadHistory := models.TxFileUploadHistory{
			FileName:   params.FileName,
			RowTotal:   len(totalErrors),
			Status:     0,
			ModuleType: common.MODUL_TYPE_PRICE,
		}
		r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	result, errDomain = r.ProcessCsvRows(csvValues, &common.ModuleReportType{PriceProvince: true})
	if errDomain != nil {
		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	priceDatas, isConvertOk = result.([]*models.PriceProvince)
	if !isConvertOk {
		return nil, &common.ErrorDomain{
			Message: "There's a problem, please contact your administrator!",
			Details: "Convertion interface{} to []*models.PriceProvince is failed",
		}
	}

	r.PriceRepository.SavePriceProvince(priceDatas)

	txFileUploadHistory := models.TxFileUploadHistory{
		FileName:   params.FileName,
		RowTotal:   len(priceDatas),
		Status:     1,
		ModuleType: common.MODUL_TYPE_PRICE,
	}

	r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

	return nil, nil
}

func (r *PriceService) UploadPriceNational(params common.FileStructure) (interface{}, *common.ErrorDomain) {
	var errDomain *common.ErrorDomain
	var priceDatas []*models.PriceNational
	var result interface{}
	var isConvertOk bool

	csvValues, errDomain := csv_helper.CsvFileProcess(params, common.ModuleType{Price: true})
	if errDomain != nil {
		totalErrors := errDomain.Details.([]map[string]interface{})
		txFileUploadHistory := models.TxFileUploadHistory{
			FileName:   params.FileName,
			RowTotal:   len(totalErrors),
			Status:     0,
			ModuleType: common.MODUL_TYPE_PRICE,
		}
		r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	result, errDomain = r.ProcessCsvRows(csvValues, &common.ModuleReportType{PriceNational: true})
	if errDomain != nil {
		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	priceDatas, isConvertOk = result.([]*models.PriceNational)
	if !isConvertOk {
		return nil, &common.ErrorDomain{
			Message: "There's a problem, please contact your administrator!",
			Details: "Convertion interface{} to []*models.PriceProvince is failed",
		}
	}
	r.PriceRepository.SavePriceNational(priceDatas)

	txFileUploadHistory := models.TxFileUploadHistory{
		FileName:   params.FileName,
		RowTotal:   len(priceDatas),
		Status:     1,
		ModuleType: common.MODUL_TYPE_PRICE,
	}

	r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

	return nil, nil
}

func (r *PriceService) Delete(params any) (interface{}, *common.ErrorDomain) {
	panic("implement me")
}

func (r *PriceService) GetLevelHarga(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PriceGetRepoParamsNew
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = params.SelectedDate

	data, err := r.PriceRepository.GetLevelHarga(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetLevelHargaNew(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PriceGetRepoParamsNew
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = params.SelectedDate

	data, err := r.PriceRepository.GetLevelHargaNew(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetLevelHargaList(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PriceGetRepoParamsNew
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = params.SelectedDate
	queryParams.PaginationParams.Page = params.PaginationParams.Page
	queryParams.PaginationParams.Limit = params.PaginationParams.Limit
	queryParams.PaginationParams.SortBy = params.PaginationParams.SortBy

	data, err := r.PriceRepository.GetLevelHargaList(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetLast5Days(params domain.PriceLast5Days) (interface{}, *common.ErrorDomain) {
	var err error
	var data interface{}
	var latestDateExist interface{}
	var dateFormat = "2006-01-02"

	var queryParams domain.PriceLast5DaysRepoParams
	queryParams.CityId = params.CityId
	queryParams.CommodityId = params.CommodityId

	latestDateExist, err = r.PriceRepository.LatestDateExist(domain.PriceListRepoParams{
		ProvinceId:  "73",
		CommodityId: queryParams.CommodityId,
	})
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	var startDate time.Time
	startDate, err = time.Parse(dateFormat, *latestDateExist.(*string))
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.StartDate = startDate.AddDate(0, 0, -4).Format(dateFormat)
	queryParams.EndDate = *latestDateExist.(*string)

	if params.CityId != "" {
		data, err = r.PriceRepository.GetLast5DaysPriceByCityId(queryParams)
		if err != nil {
			fmt.Println("err", err)
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	} else {
		data, err = r.PriceRepository.GetLast5DaysPriceByCommodityId(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}

func (r *PriceService) GetLast5DaysCount(params domain.PriceLast5Days) (*int, *common.ErrorDomain) {
	var data *int
	var err error
	var queryParams domain.PriceLast5DaysRepoParams
	var dateFormat = "2006-01-02"
	var startDate time.Time

	queryParams.CityId = params.CityId
	queryParams.CommodityId = params.CommodityId

	queryParams.StartDate = startDate.AddDate(0, 0, -4).Format(dateFormat)
	queryParams.EndDate = common.GetDateNow()

	if params.CityId != "" {
		data, err = r.PriceRepository.GetLast5DaysPriceByCityIdCount(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	} else {
		data, err = r.PriceRepository.GetLast5DaysPriceByCommodityIdCount(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}

func (r *PriceService) GetDibandingkanSulsel(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PriceGetCompareProvinceParams
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = params.SelectedDate
	queryParams.ProvinceId = "73"

	data, err := r.PriceRepository.GetCompareBySulsel(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetDibandingkanSulselList(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PriceGetCompareProvinceParams
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = params.SelectedDate
	queryParams.ProvinceId = "73"
	queryParams.PaginationParams.Page = params.PaginationParams.Page
	queryParams.PaginationParams.Limit = params.PaginationParams.Limit
	queryParams.PaginationParams.SortBy = params.PaginationParams.SortBy

	data, err := r.PriceRepository.GetCompareBySulselList(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetCompareProvinceCityHistory(params domain.PriceDiffRequestParams) (interface{}, *common.ErrorDomain) {
	var err error
	var queryParams domain.PriceDiffByCityAndCommodityParams
	queryParams.CommodityId = strings.ToLower(params.CommodityId)
	queryParams.CityId = strings.ToLower(params.CityId)

	startDate := params.StartDate
	_, err = time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	endDate := params.StartDate
	_, err = time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.StartDate = params.StartDate
	queryParams.EndDate = params.EndDate
	queryParams.ProvinceId = "73"
	queryParams.Status = params.Status

	data, err := r.PriceRepository.GetCompareProvinceCityHistory(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetCompareProvinceCommodityHistory(params domain.PriceGetCompareProvinceCommodityHistoryParams) (interface{}, *common.ErrorDomain) {
	var err error
	_, err = time.Parse("2006-01-02", params.SelectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	data, err := r.PriceRepository.GetCompareProvinceCommodityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetDibandingkanNational(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PriceGetCompareNationalParams
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = params.SelectedDate
	queryParams.ProvinceId = "73"
	queryParams.NationalId = "1"

	data, err := r.PriceRepository.GetCompareByNational(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetDibandingkanNationalList(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PriceGetCompareNationalParams
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = params.SelectedDate
	queryParams.ProvinceId = "73"
	queryParams.PaginationParams.Page = params.PaginationParams.Page
	queryParams.PaginationParams.Limit = params.PaginationParams.Limit
	queryParams.PaginationParams.SortBy = params.PaginationParams.SortBy
	queryParams.NationalId = "1"

	data, err := r.PriceRepository.GetCompareByNationalList(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetCompareNationalCityHistory(params domain.PriceDiffRequestParams) (interface{}, *common.ErrorDomain) {
	var err error
	var queryParams domain.PriceDiffByCityAndCommodityParams
	queryParams.CommodityId = strings.ToLower(params.CommodityId)
	queryParams.CityId = strings.ToLower(params.CityId)

	startDate := params.StartDate
	_, err = time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	endDate := params.EndDate
	_, err = time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.StartDate = params.StartDate
	queryParams.EndDate = params.EndDate
	queryParams.ProvinceId = "73"

	data, err := r.PriceRepository.GetCompareNationalCityHistory(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetCompareNationalCommodityHistory(params domain.PriceGetCompareProvinceCommodityHistoryParams) (interface{}, *common.ErrorDomain) {
	var err error
	_, err = time.Parse("2006-01-02", params.SelectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	data, err := r.PriceRepository.GetCompareNationalCommodityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) Get(params domain.PriceListParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error
	var queryParams domain.PriceListRepoParams

	queryParams.PaginationParams = params.PaginationParams
	queryParams.ProvinceId = params.ProvinceId
	queryParams.CityId = params.CityId
	queryParams.CommodityId = params.CommodityId
	queryParams.StartDate = params.StartDate
	queryParams.EndDate = params.EndDate

	if queryParams.ProvinceId != "" && queryParams.ProvinceId != "false" {
		data, err = r.PriceRepository.GetForProvince(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	} else if queryParams.CityId != "" && queryParams.CityId != "false" {
		data, err = r.PriceRepository.Get(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}

func (r *PriceService) GetCount(params domain.PriceListParams) (*int, *common.ErrorDomain) {
	var data interface{}
	var err error
	var queryParams domain.PriceListRepoParams

	queryParams.PaginationParams = params.PaginationParams
	queryParams.ProvinceId = params.ProvinceId
	queryParams.CityId = params.CityId
	queryParams.CommodityId = params.CommodityId
	queryParams.StartDate = params.StartDate
	queryParams.EndDate = params.EndDate

	if queryParams.ProvinceId != "" && queryParams.ProvinceId != "false" {
		data, err = r.PriceRepository.GetCountProvince(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	} else if queryParams.CityId != "" && queryParams.CityId != "false" {
		data, err = r.PriceRepository.GetCount(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data.(*int), nil
}

func (r *PriceService) ProcessCsvRows(csvValues [][]interface{}, moduleReportType *common.ModuleReportType) (interface{}, *common.ErrorDomain) {
	var err error
	var priceNationals []*models.PriceNational
	var priceProvinces []*models.PriceProvince
	var priceCities []*models.PriceCity

	priceMapper := mapper.NewPriceMapper(r.TmCityRepository, r.TmCommodityRepository, r.TmProvinceRepository)
	for idx, csvValue := range csvValues {
		var priceProvince *models.PriceProvince
		csvValueStringArr := string_helper.InterfaceArrToStringArr(csvValue)
		if moduleReportType.PriceProvince {
			priceProvince, err = priceMapper.StringArrToPriceProvinceModel(csvValueStringArr)
			if err != nil {
				detailsErr := make([]interface{}, 1)
				detailsErr[0] = err.Error()

				if detailsErr[0] == "no rows in result set" {
					return nil, &common.ErrorDomain{
						Message: fmt.Sprintf("error read CSV file, row: %d", idx),
						Details: fmt.Sprintf("There's an error: %s, row: %d", err.Error(), idx),
					}
				} else {
					return nil, &common.ErrorDomain{
						Message: fmt.Sprintf("error read CSV file: %s. row: %d", err.Error(), idx),
						Details: detailsErr,
					}
				}
			}

			priceProvinces = append(priceProvinces, priceProvince)
		} else if moduleReportType.PriceNational {
			var priceNational *models.PriceNational
			priceNational, err = priceMapper.StringArrToPriceNationalModel(csvValueStringArr)
			if err != nil {
				detailsErr := make([]interface{}, 1)
				detailsErr[0] = err.Error()

				if detailsErr[0] == "no rows in result set" {
					return nil, &common.ErrorDomain{
						Message: fmt.Sprintf("error read CSV file, row: %d", idx),
						Details: fmt.Sprintf("There's an error: %s, row: %d", err.Error(), idx),
					}
				} else {
					return nil, &common.ErrorDomain{
						Message: fmt.Sprintf("error read CSV file: %s. row: %d", err.Error(), idx),
						Details: detailsErr,
					}
				}
			}

			priceNationals = append(priceNationals, priceNational)
		} else if moduleReportType.PriceCity {
			var priceCity *models.PriceCity
			priceCity, err = priceMapper.StringArrToPriceCityModel(csvValueStringArr)
			if err != nil {
				detailsErr := make([]interface{}, 1)
				detailsErr[0] = err.Error()

				if detailsErr[0] == "no rows in result set" {
					return nil, &common.ErrorDomain{
						Message: fmt.Sprintf("error read CSV file, row: %d", idx),
						Details: fmt.Sprintf("There's an error: %s, row: %d", err.Error(), idx),
					}
				} else {
					return nil, &common.ErrorDomain{
						Message: fmt.Sprintf("error read CSV file: %s. row: %d", err.Error(), idx),
						Details: detailsErr,
					}
				}
			}

			priceCities = append(priceCities, priceCity)
		} else {
			return nil, &common.ErrorDomain{
				Message: "One of Module Report Type must be filled",
				Details: nil,
			}
		}
	}

	if moduleReportType.PriceProvince {
		return priceProvinces, nil
	} else if moduleReportType.PriceNational {
		return priceNationals, nil
	} else if moduleReportType.PriceCity {
		return priceCities, nil
	} else {
		return nil, &common.ErrorDomain{
			Message: "One of Module Report Type must be filled",
			Details: nil,
		}
	}
}

func (r *PriceService) GetMtm(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PricePerubahanParams
	var redisResult *models.PriceMtmResponse
	var data interface{}

	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = selectedDate
	queryParams.ProvinceId = "73"

	data, err = r.PriceRepository.GetMtm(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}
	return data, nil

	queryString, queryStringErr := query.Values(queryParams)
	if queryStringErr != nil {
		return nil, &common.ErrorDomain{
			Message: queryStringErr.Error(),
		}
	}

	redisKey := "sigap_sultan_api_database_" + "1" + ":Price:WEB:GetMtm:" + queryString.Encode()
	redisGet := r.RedisClient.Get(params.CommonParams.Ctx.Context(), redisKey)

	if err = redisGet.Err(); err != nil {
		fmt.Printf("usecase.caleg.GetList.Redis.Get: %s\n", err.Error())
		data, err = r.PriceRepository.GetMtm(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		dataMarshal, errMarshal := json.Marshal(data)
		if errMarshal != nil {
			return nil, &common.ErrorDomain{
				Message: errMarshal.Error(),
			}
		}

		err = r.RedisClient.Set(params.CommonParams.Ctx.Context(), redisKey, dataMarshal, time.Duration(0)).Err()
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return data, nil
	} else {
		result, err := redisGet.Result()
		if err = json.Unmarshal([]byte(result), &redisResult); err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return redisResult, nil

		//	fmt.Println("usecase.caleg.GetList.Redis.Get is success")
		//	return redisResult, nil
		//	//var calegListResponseDomain []*domain.CalegListResponseDomain
		//	//marshalErr := json.Unmarshal([]byte(redisResult), &calegListResponseDomain)
		//	//if marshalErr != nil {
		//	//	existErr := errors.New("Redis marshal value err: " + marshalErr.Error())
		//	//	customErr := phastoserr.New(existErr).SetCode(400)
		//	//
		//	//	return nil, errors.Wrap(customErr, "usecase.caleg.GetList.Redis.MarshalValue")
		//	//}
		//	//
		//	//response := &phastosDb.SelectResponse{
		//	//	Data:          calegListResponseDomain,
		//	//	RequestParam:  requestData,
		//	//	TotalData:     int64(1),
		//	//	TotalFiltered: int64(0),
		//	//}
		//	//
		//	//return response, nil
	}

	//redisResult, err := s.repo.GetList(ctx, jwtData.Partai.Id, requestData)
	//if err != nil {
	//	return nil, err
	//}
	//
	//list := redisResult.([]*domain.CalegListResponseDomain)
	//
	//redisSetErr := s.redisClient.Set(ctx, redisKey, list)
	//if redisSetErr != nil {
	//	existErr := errors.New("Redis set value err: " + redisSetErr.Error())
	//	customErr := phastoserr.New(existErr).SetCode(400)
	//
	//	return nil, errors.Wrap(customErr, "usecase.caleg.GetList.Redis.Set")
	//}
}

func (r *PriceService) GetMtmList(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PricePerubahanParams
	var redisResult *models.PriceMtmListResponse
	var data interface{}

	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = selectedDate
	queryParams.ProvinceId = "73"
	queryParams.PaginationParams.SortBy = params.PaginationParams.SortBy

	data, err = r.PriceRepository.GetMtmList(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}
	return data, nil

	queryString, queryStringErr := query.Values(queryParams)
	if queryStringErr != nil {
		return nil, &common.ErrorDomain{
			Message: queryStringErr.Error(),
		}
	}

	redisKey := "sigap_sultan_api_database_" + "1" + ":Price:WEB:GetMtmList:" + queryString.Encode()
	redisGet := r.RedisClient.Get(params.CommonParams.Ctx.Context(), redisKey)

	if err = redisGet.Err(); err != nil {
		fmt.Printf("usecase.caleg.GetList.Redis.Get: %s\n", err.Error())
		data, err = r.PriceRepository.GetMtmList(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		dataMarshal, errMarshal := json.Marshal(data)
		if errMarshal != nil {
			return nil, &common.ErrorDomain{
				Message: errMarshal.Error(),
			}
		}

		err = r.RedisClient.Set(params.CommonParams.Ctx.Context(), redisKey, dataMarshal, time.Duration(0)).Err()
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return data, nil
	} else {
		result, err := redisGet.Result()
		if err = json.Unmarshal([]byte(result), &redisResult); err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return redisResult, nil
	}
}

func (r *PriceService) GetMtmCityHistory(params domain.PriceMtmCityHistoryParams) (interface{}, *common.ErrorDomain) {
	var err error
	params.CommodityId = strings.ToLower(params.CommodityId)

	_, err = time.Parse("2006-01-02", params.StartDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	_, err = time.Parse("2006-01-02", params.EndDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	data, err := r.PriceRepository.GetMtmCityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetMtmCommodityHistory(params domain.PriceMtmCommodityHistoryParams) (interface{}, *common.ErrorDomain) {
	var err error
	params.CommodityId = strings.ToLower(params.CommodityId)

	_, err = time.Parse("2006-01-02", params.SelectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	data, err := r.PriceRepository.GetMtmCommodityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetYtd(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PricePerubahanParams
	var redisResult *models.PriceMtmResponse
	var data interface{}

	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = selectedDate
	queryParams.ProvinceId = "73"

	queryString, queryStringErr := query.Values(queryParams)
	if queryStringErr != nil {
		return nil, &common.ErrorDomain{
			Message: queryStringErr.Error(),
		}
	}

	redisKey := "sigap_sultan_api_database_" + "1" + ":Price:WEB:GetYtd:" + queryString.Encode()
	redisGet := r.RedisClient.Get(params.CommonParams.Ctx.Context(), redisKey)

	if err = redisGet.Err(); err != nil {
		fmt.Printf("usecase.caleg.GetList.Redis.Get: %s\n", err.Error())
		data, err = r.PriceRepository.GetYtd(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		dataMarshal, errMarshal := json.Marshal(data)
		if errMarshal != nil {
			return nil, &common.ErrorDomain{
				Message: errMarshal.Error(),
			}
		}

		err = r.RedisClient.Set(params.CommonParams.Ctx.Context(), redisKey, dataMarshal, time.Duration(0)).Err()
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return data, nil
	} else {
		result, err := redisGet.Result()
		if err = json.Unmarshal([]byte(result), &redisResult); err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return redisResult, nil
	}
}

func (r *PriceService) GetYtdList(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PricePerubahanParams
	var redisResult *models.PriceMtmListResponse
	var data interface{}

	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = selectedDate
	queryParams.ProvinceId = "73"
	queryParams.PaginationParams.SortBy = params.PaginationParams.SortBy

	queryString, queryStringErr := query.Values(queryParams)
	if queryStringErr != nil {
		return nil, &common.ErrorDomain{
			Message: queryStringErr.Error(),
		}
	}

	//data, err = r.PriceRepository.GetYtdList(queryParams)
	//if err != nil {
	//	return nil, &common.ErrorDomain{
	//		Message: err.Error(),
	//	}
	//}
	//
	//return data, nil

	redisKey := "sigap_sultan_api_database_" + "1" + ":Price:WEB:GetYtdList:" + queryString.Encode()
	redisGet := r.RedisClient.Get(params.CommonParams.Ctx.Context(), redisKey)

	if err = redisGet.Err(); err != nil {
		fmt.Printf("usecase.caleg.GetList.Redis.Get: %s\n", err.Error())
		data, err = r.PriceRepository.GetYtdList(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		dataMarshal, errMarshal := json.Marshal(data)
		if errMarshal != nil {
			return nil, &common.ErrorDomain{
				Message: errMarshal.Error(),
			}
		}

		err = r.RedisClient.Set(params.CommonParams.Ctx.Context(), redisKey, dataMarshal, time.Duration(0)).Err()
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return data, nil
	} else {
		result, err := redisGet.Result()
		if err = json.Unmarshal([]byte(result), &redisResult); err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return redisResult, nil
	}
}

func (r *PriceService) GetYtdCityHistory(params domain.PriceMtmCityHistoryParams) (interface{}, *common.ErrorDomain) {
	var err error
	params.CommodityId = strings.ToLower(params.CommodityId)

	_, err = time.Parse("2006-01-02", params.StartDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	_, err = time.Parse("2006-01-02", params.EndDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	data, err := r.PriceRepository.GetYtdCityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetYtdCommodityHistory(params domain.PriceMtmCommodityHistoryParams) (interface{}, *common.ErrorDomain) {
	var err error
	params.CommodityId = strings.ToLower(params.CommodityId)

	_, err = time.Parse("2006-01-02", params.SelectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	data, err := r.PriceRepository.GetYtdCommodityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetYty(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PricePerubahanParams
	var redisResult *models.PriceMtmResponse
	var data interface{}

	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = selectedDate
	queryParams.ProvinceId = "73"

	queryString, queryStringErr := query.Values(queryParams)
	if queryStringErr != nil {
		return nil, &common.ErrorDomain{
			Message: queryStringErr.Error(),
		}
	}

	redisKey := "sigap_sultan_api_database_" + "1" + ":Price:WEB:GetYty:" + queryString.Encode()
	redisGet := r.RedisClient.Get(params.CommonParams.Ctx.Context(), redisKey)

	if err = redisGet.Err(); err != nil {
		fmt.Printf("usecase.caleg.GetList.Redis.Get: %s\n", err.Error())
		data, err = r.PriceRepository.GetYty(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		dataMarshal, errMarshal := json.Marshal(data)
		if errMarshal != nil {
			return nil, &common.ErrorDomain{
				Message: errMarshal.Error(),
			}
		}

		err = r.RedisClient.Set(params.CommonParams.Ctx.Context(), redisKey, dataMarshal, time.Duration(0)).Err()
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return data, nil
	} else {
		result, err := redisGet.Result()
		if err = json.Unmarshal([]byte(result), &redisResult); err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return redisResult, nil
	}
}

func (r *PriceService) GetYtyList(params domain.PriceGetHargaRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PricePerubahanParams
	var redisResult *models.PriceMtmListResponse
	var data interface{}

	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryParams.SelectedDate = selectedDate
	queryParams.ProvinceId = "73"
	queryParams.PaginationParams.SortBy = params.PaginationParams.SortBy

	queryString, queryStringErr := query.Values(queryParams)
	if queryStringErr != nil {
		return nil, &common.ErrorDomain{
			Message: queryStringErr.Error(),
		}
	}

	//data, err = r.PriceRepository.GetYtyList(queryParams)
	//if err != nil {
	//	return nil, &common.ErrorDomain{
	//		Message: err.Error(),
	//	}
	//}
	//
	//return data, nil

	redisKey := "sigap_sultan_api_database_" + "1" + ":Price:WEB:GetYtyList:" + queryString.Encode()
	redisGet := r.RedisClient.Get(params.CommonParams.Ctx.Context(), redisKey)

	if err = redisGet.Err(); err != nil {
		fmt.Printf("usecase.caleg.GetList.Redis.Get: %s\n", err.Error())
		data, err = r.PriceRepository.GetYtyList(queryParams)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		dataMarshal, errMarshal := json.Marshal(data)
		if errMarshal != nil {
			return nil, &common.ErrorDomain{
				Message: errMarshal.Error(),
			}
		}

		err = r.RedisClient.Set(params.CommonParams.Ctx.Context(), redisKey, dataMarshal, time.Duration(0)).Err()
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return data, nil
	} else {
		result, err := redisGet.Result()
		if err = json.Unmarshal([]byte(result), &redisResult); err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return redisResult, nil
	}
}

func (r *PriceService) GetYtyCityHistory(params domain.PriceMtmCityHistoryParams) (interface{}, *common.ErrorDomain) {
	var err error
	params.CommodityId = strings.ToLower(params.CommodityId)

	_, err = time.Parse("2006-01-02", params.StartDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	_, err = time.Parse("2006-01-02", params.EndDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	data, err := r.PriceRepository.GetYtyCityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetYtyCommodityHistory(params domain.PriceMtmCommodityHistoryParams) (interface{}, *common.ErrorDomain) {
	var err error
	params.CommodityId = strings.ToLower(params.CommodityId)

	_, err = time.Parse("2006-01-02", params.SelectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	data, err := r.PriceRepository.GetYtyCommodityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) Exist(params domain.PriceListRepoParams) (interface{}, *common.ErrorDomain) {
	var err error
	var redisResult *map[string]bool
	var data interface{}
	params.CommodityId = strings.ToLower(params.CommodityId)

	_, err = time.Parse("2006-01-02", params.StartDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	_, err = time.Parse("2006-01-02", params.EndDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	queryString, queryStringErr := query.Values(params)
	if queryStringErr != nil {
		return nil, &common.ErrorDomain{
			Message: queryStringErr.Error(),
		}
	}

	data, err = r.PriceRepository.Exist(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil

	redisKey := "sigap_sultan_api_database_" + "1" + ":Price:WEB:Exist:" + queryString.Encode()
	redisGet := r.RedisClient.Get(params.CommonParams.Ctx.Context(), redisKey)

	if err = redisGet.Err(); err != nil {
		fmt.Printf(redisKey+": %s\n", err.Error())
		data, err = r.PriceRepository.Exist(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		dataMarshal, errMarshal := json.Marshal(data)
		if errMarshal != nil {
			return nil, &common.ErrorDomain{
				Message: errMarshal.Error(),
			}
		}

		err = r.RedisClient.Set(params.CommonParams.Ctx.Context(), redisKey, dataMarshal, time.Duration(0)).Err()
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return data, nil
	} else {
		result, err := redisGet.Result()
		if err = json.Unmarshal([]byte(result), &redisResult); err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}

		return redisResult, nil
	}
}

func (r *PriceService) LatestDateExist(params domain.PriceListRepoParams) (interface{}, *common.ErrorDomain) {
	var err error
	params.CommodityId = strings.ToLower(params.CommodityId)

	data, err := r.PriceRepository.LatestDateExist(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetHistory(params domain.HistoryRequestParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.HistoryRequestParams

	queryParams.PaginationParams = params.PaginationParams
	queryParams.Module = params.Module
	queryParams.Search = params.Search

	data, err := r.PriceRepository.GetHistory(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *PriceService) GetHistoryCount(params domain.HistoryRequestParams) (*int, *common.ErrorDomain) {
	var data interface{}
	var err error
	var queryParams domain.HistoryRequestParams

	queryParams.PaginationParams = params.PaginationParams
	queryParams.Module = params.Module
	queryParams.Search = params.Search

	data, err = r.PriceRepository.GetCountHistory(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data.(*int), nil
}
