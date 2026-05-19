package services

import (
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
)

type NeracaService struct {
	NeracaRepository              *repositories.NeracaRepository
	TmCommodityRepository         *repositories.TmCommodityRepository
	TmCityRepository              *repositories.TmCityRepository
	TxFileUploadHistoryRepository *repositories.TxFileUploadHistoryRepository
	TmProvinceRepository          *repositories.TmProvinceRepository
}

func NewNeracaService(neracaRepository *repositories.NeracaRepository, tmCommodityRepository *repositories.TmCommodityRepository, tmCityRepository *repositories.TmCityRepository, txFileUploadHistoryRepository *repositories.TxFileUploadHistoryRepository, TmProvinceRepository *repositories.TmProvinceRepository) *NeracaService {
	return &NeracaService{
		neracaRepository,
		tmCommodityRepository,
		tmCityRepository,
		txFileUploadHistoryRepository,
		TmProvinceRepository,
	}
}

func (r *NeracaService) UploadNeracaCity(params common.FileStructure) (interface{}, *common.ErrorDomain) {
	var errDomain *common.ErrorDomain
	var neracaDatas []*models.NeracaCity
	var result interface{}
	var isConvertOk bool

	csvValues, errDomain := csv_helper.CsvFileProcess(params, common.ModuleType{Neraca: true})
	if errDomain != nil {
		totalErrors := errDomain.Details.([]map[string]interface{})
		txFileUploadHistory := models.TxFileUploadHistory{
			FileName:   params.FileName,
			RowTotal:   len(totalErrors),
			Status:     0,
			ModuleType: common.MODUL_TYPE_NERACA,
		}
		r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	result, errDomain = r.ProcessCsvRows(csvValues, &common.ModuleReportType{NeracaCity: true})
	if errDomain != nil {
		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	neracaDatas, isConvertOk = result.([]*models.NeracaCity)
	if !isConvertOk {
		return nil, &common.ErrorDomain{
			Message: "There's a problem, please contact your administrator!",
			Details: "Convertion interface{} to []*models.NeracaCity is failed",
		}
	}

	r.NeracaRepository.SaveNeracaCity(neracaDatas)

	txFileUploadHistory := models.TxFileUploadHistory{
		FileName:   params.FileName,
		RowTotal:   len(neracaDatas),
		Status:     1,
		ModuleType: common.MODUL_TYPE_NERACA,
	}

	r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

	return nil, nil
}

func (r *NeracaService) UploadNeracaProvince(params common.FileStructure) (interface{}, *common.ErrorDomain) {
	var errDomain *common.ErrorDomain
	var neracaDatas []*models.NeracaProvince
	var result interface{}
	var isConvertOk bool

	csvValues, errDomain := csv_helper.CsvFileProcess(params, common.ModuleType{Neraca: true})
	if errDomain != nil {
		totalErrors := errDomain.Details.([]map[string]interface{})
		txFileUploadHistory := models.TxFileUploadHistory{
			FileName:   params.FileName,
			RowTotal:   len(totalErrors),
			Status:     0,
			ModuleType: common.MODUL_TYPE_NERACA,
		}
		r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	result, errDomain = r.ProcessCsvRows(csvValues, &common.ModuleReportType{NeracaProvince: true})
	if errDomain != nil {
		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	neracaDatas, isConvertOk = result.([]*models.NeracaProvince)
	if !isConvertOk {
		return nil, &common.ErrorDomain{
			Message: "There's a problem, please contact your administrator!",
			Details: "Convertion interface{} to []*models.NeracaCity is failed",
		}
	}

	r.NeracaRepository.SaveNeracaProvince(neracaDatas)

	txFileUploadHistory := models.TxFileUploadHistory{
		FileName:   params.FileName,
		RowTotal:   len(neracaDatas),
		Status:     1,
		ModuleType: common.MODUL_TYPE_NERACA,
	}

	r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

	return nil, nil
}

func (r *NeracaService) UploadNeracaNational(params common.FileStructure) (interface{}, *common.ErrorDomain) {
	var errDomain *common.ErrorDomain
	var neracaDatas []*models.NeracaNational
	var result interface{}
	var isConvertOk bool

	csvValues, errDomain := csv_helper.CsvFileProcess(params, common.ModuleType{Neraca: true})
	if errDomain != nil {
		totalErrors := errDomain.Details.([]map[string]interface{})
		txFileUploadHistory := models.TxFileUploadHistory{
			FileName:   params.FileName,
			RowTotal:   len(totalErrors),
			Status:     0,
			ModuleType: common.MODUL_TYPE_NERACA,
		}
		r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	result, errDomain = r.ProcessCsvRows(csvValues, &common.ModuleReportType{NeracaNational: true})
	if errDomain != nil {
		return nil, &common.ErrorDomain{
			Message: errDomain.Message,
			Details: errDomain.Details,
		}
	}

	neracaDatas, isConvertOk = result.([]*models.NeracaNational)
	if !isConvertOk {
		return nil, &common.ErrorDomain{
			Message: "There's a problem, please contact your administrator!",
			Details: "Convertion interface{} to []*models.NeracaCity is failed",
		}
	}

	r.NeracaRepository.SaveNeracaNational(neracaDatas)

	txFileUploadHistory := models.TxFileUploadHistory{
		FileName:   params.FileName,
		RowTotal:   len(neracaDatas),
		Status:     1,
		ModuleType: common.MODUL_TYPE_NERACA,
	}

	r.TxFileUploadHistoryRepository.Save(txFileUploadHistory)

	return nil, nil
}

func (r *NeracaService) ProcessCsvRows(csvValues [][]interface{}, moduleReportType *common.ModuleReportType) (interface{}, *common.ErrorDomain) {
	var err error
	var neracaNationals []*models.NeracaNational
	var neracaProvinces []*models.NeracaProvince
	var neracaCities []*models.NeracaCity

	neracaMapper := mapper.NewNeracaMapper(r.TmCityRepository, r.TmCommodityRepository, r.TmProvinceRepository)
	for idx, csvValue := range csvValues {
		var neracaProvince *models.NeracaProvince
		csvValueStringArr := string_helper.InterfaceArrToStringArr(csvValue)
		if moduleReportType.NeracaProvince {
			neracaProvince, err = neracaMapper.StringArrToNeracaProvinceModel(csvValueStringArr)
			if err != nil {
				detailsErr := make([]interface{}, 1)
				detailsErr[0] = err.Error()

				if detailsErr[0] == "no rows in result set" {
					return nil, &common.ErrorDomain{
						Message: "error read CSV file",
						Details: fmt.Sprintf("Komoditas is not found, row: %d", idx),
					}
				} else {
					return nil, &common.ErrorDomain{
						Message: "error read CSV file",
						Details: detailsErr,
					}
				}
			}

			neracaProvinces = append(neracaProvinces, neracaProvince)
		} else if moduleReportType.NeracaNational {
			var neracaNational *models.NeracaNational
			neracaNational, err = neracaMapper.StringArrToNeracaNationalModel(csvValueStringArr)
			if err != nil {
				detailsErr := make([]interface{}, 1)
				detailsErr[0] = err.Error()

				if detailsErr[0] == "no rows in result set" {
					return nil, &common.ErrorDomain{
						Message: "error read CSV file",
						Details: fmt.Sprintf("Komoditas is not found, row: %d", idx),
					}
				} else {
					return nil, &common.ErrorDomain{
						Message: "error read CSV file",
						Details: detailsErr,
					}
				}
			}

			neracaNationals = append(neracaNationals, neracaNational)
		} else if moduleReportType.NeracaCity {
			var neracaCity *models.NeracaCity
			neracaCity, err = neracaMapper.StringArrToNeracaCityModel(csvValueStringArr)
			if err != nil {
				detailsErr := make([]interface{}, 1)
				detailsErr[0] = err.Error()

				if detailsErr[0] == "no rows in result set" {
					return nil, &common.ErrorDomain{
						Message: "error read CSV file",
						Details: fmt.Sprintf("Komoditas is not found, row: %d", idx),
					}
				} else {
					return nil, &common.ErrorDomain{
						Message: "error read CSV file",
						Details: detailsErr,
					}
				}
			}

			neracaCities = append(neracaCities, neracaCity)
		} else {
			return nil, &common.ErrorDomain{
				Message: "One of Module Report Type must be filled",
				Details: nil,
			}
		}
	}

	if moduleReportType.NeracaProvince {
		return neracaProvinces, nil
	} else if moduleReportType.NeracaNational {
		return neracaNationals, nil
	} else if moduleReportType.NeracaCity {
		return neracaCities, nil
	} else {
		return nil, &common.ErrorDomain{
			Message: "One of Module Report Type must be filled",
			Details: nil,
		}
	}
}

func (r *NeracaService) GetStockAkhir(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetStockAkhir(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetStockAkhirList(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error
	var queryParams domain.PriceGetRepoParamsNew
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err = time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}
	queryParams.SelectedDate = params.SelectedDate

	if params.CityId != "" && params.CommodityId != "" {

	} else if params.CommodityId != "" {
		data, err = r.NeracaRepository.GetStockAkhirListByCommodity(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	} else if params.CityId != "" {
		data, err = r.NeracaRepository.GetStockAkhir(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}

func (r *NeracaService) GetStockAkhirByCommodityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetStockAkhirByCommodityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetStockAkhirByCommodityCityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetStockAkhirByCommodityCityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetKetersediaanMap(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetKetersediaanMap(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetKetersediaanList(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	selectedDate := params.SelectedDate
	_, err = time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	if params.CityId != "" && params.CommodityId != "" {

	} else if params.CommodityId != "" {
		data, err = r.NeracaRepository.GetKetersediaanListByCommodity(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	} else if params.CityId != "" {
		data, err = r.NeracaRepository.GetStockAkhir(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}

func (r *NeracaService) GetKetersediaanByCommodityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetKetersediaanByCommodityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetKetersediaanByCommodityCityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetKetersediaanByCommodityCityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetKebutuhanMap(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetKebutuhanMap(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetKebutuhanList(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	selectedDate := params.SelectedDate
	_, err = time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	if params.CityId != "" && params.CommodityId != "" {

	} else if params.CommodityId != "" {
		data, err = r.NeracaRepository.GetKebutuhanListByCommodity(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	} else if params.CityId != "" {
		data, err = r.NeracaRepository.GetStockAkhir(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}

func (r *NeracaService) GetKebutuhanByCommodityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetKebutuhanByCommodityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetKebutuhanByCommodityCityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetKebutuhanByCommodityCityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetStockAkhirByCityAndCommodityId(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetStockAkhirByCityAndCommodityId(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetStockAkhirByCityAndCommodityChart(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetStockAkhirByCityAndCommodityChart(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) CompareWithPriceCommodityHistory(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.CompareWithPriceCommodityHistory(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetStockAkhirByCity(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetStockAkhirByCity(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetStockAkhirByCityList(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error
	var queryParams domain.PriceGetRepoParamsNew
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err = time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}
	queryParams.SelectedDate = params.SelectedDate

	data, err = r.NeracaRepository.GetStockAkhirListByCity(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetKetersediaanByCity(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetKetersediaanByCity(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetKetersediaanByCityList(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error
	var queryParams domain.PriceGetRepoParamsNew
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err = time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}
	queryParams.SelectedDate = params.SelectedDate

	data, err = r.NeracaRepository.GetKetersediaanListByCity(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetKebutuhanByCity(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error

	data, err = r.NeracaRepository.GetKebutuhanByCity(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) GetKebutuhanByCityList(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
	var data interface{}
	var err error
	var queryParams domain.PriceGetRepoParamsNew
	queryParams.CommodityId = strings.ToLower(params.CommodityId)

	selectedDate := params.SelectedDate
	_, err = time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}
	queryParams.SelectedDate = params.SelectedDate

	data, err = r.NeracaRepository.GetKebutuhanListByCity(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) Exist(params domain.NeracaStokAkhirListRequestParams) (interface{}, *common.ErrorDomain) {
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

	data, err := r.NeracaRepository.Exist(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *NeracaService) LatestDateExist(params domain.PriceListRepoParams) (interface{}, *common.ErrorDomain) {
	var err error
	params.CommodityId = strings.ToLower(params.CommodityId)

	data, err := r.NeracaRepository.LatestDateExist(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}
