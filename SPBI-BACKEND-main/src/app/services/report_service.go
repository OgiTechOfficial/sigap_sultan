package services

import (
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/repositories"
	"sigap-sultan-be/src/common"
)

type ReportService struct {
	PriceRepository               *repositories.PriceRepository
	TmCommodityRepository         *repositories.TmCommodityRepository
	TmCityRepository              *repositories.TmCityRepository
	TxFileUploadHistoryRepository *repositories.TxFileUploadHistoryRepository
	TmProvinceRepository          *repositories.TmProvinceRepository
	NeracaRepository              *repositories.NeracaRepository
}

func NewReportService(priceRepository *repositories.PriceRepository, tmCommodityRepository *repositories.TmCommodityRepository, tmCityRepository *repositories.TmCityRepository, txFileUploadHistoryRepository *repositories.TxFileUploadHistoryRepository, TmProvinceRepository *repositories.TmProvinceRepository, neracaRepository *repositories.NeracaRepository) *ReportService {
	return &ReportService{
		priceRepository,
		tmCommodityRepository,
		tmCityRepository,
		txFileUploadHistoryRepository,
		TmProvinceRepository,
		neracaRepository,
	}
}

func (r ReportService) GetReportPrice(params domain.PriceListParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PriceListRepoParams
	queryParams.PaginationParams = params.PaginationParams
	queryParams.CityId = params.CityId
	queryParams.CommodityId = params.CommodityId
	queryParams.StartDate = params.StartDate
	queryParams.EndDate = params.EndDate

	data, err := r.PriceRepository.GetPriceReport(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r ReportService) GetReportPriceDownload(params domain.PriceListParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.PriceListRepoParams
	queryParams.PaginationParams = params.PaginationParams
	queryParams.CityId = params.CityId
	queryParams.CommodityId = params.CommodityId
	queryParams.StartDate = params.StartDate
	queryParams.EndDate = params.EndDate

	data, err := r.PriceRepository.GetPriceReportDownload(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r ReportService) GetReportCount(params domain.PriceListParams) (*int, *common.ErrorDomain) {
	var queryParams domain.PriceListRepoParams
	queryParams.PaginationParams = params.PaginationParams
	queryParams.CityId = params.CityId
	queryParams.CommodityId = params.CommodityId
	queryParams.StartDate = params.StartDate
	queryParams.EndDate = params.EndDate

	data, err := r.PriceRepository.GetPriceReportCount(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r ReportService) GetReportNeraca(params domain.NeracaListParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.NeracaListParams
	queryParams.PaginationParams = params.PaginationParams
	queryParams.CityId = params.CityId
	queryParams.CommodityId = params.CommodityId
	queryParams.StartDate = params.StartDate
	queryParams.EndDate = params.EndDate

	data, err := r.NeracaRepository.GetNeracaReport(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r ReportService) GetReportNeracaDownload(params domain.NeracaListParams) (interface{}, *common.ErrorDomain) {
	var queryParams domain.NeracaListParams
	queryParams.PaginationParams = params.PaginationParams
	queryParams.CityId = params.CityId
	queryParams.CommodityId = params.CommodityId
	queryParams.StartDate = params.StartDate
	queryParams.EndDate = params.EndDate

	data, err := r.NeracaRepository.GetNeracaReportDownload(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r ReportService) GetReportNeracaCount(params domain.NeracaListParams) (*int, *common.ErrorDomain) {
	var queryParams domain.NeracaListParams
	queryParams.PaginationParams = params.PaginationParams
	queryParams.CityId = params.CityId
	queryParams.CommodityId = params.CommodityId
	queryParams.StartDate = params.StartDate
	queryParams.EndDate = params.EndDate

	data, err := r.NeracaRepository.GetNeracaReportCount(queryParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}
