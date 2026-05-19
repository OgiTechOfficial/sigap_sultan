package domain

import "sigap-sultan-be/src/common"

// PriceGetRequestParams - FOR CONTROLLERS
type PriceGetRequestParams struct {
	CommonParams  common.CommonParams
	CommodityType string `json:"commodityType"`
	PriceInfoType string `json:"priceInfoType"`
	DateRange     string `json:"dateRange"`
}

type PriceGetRequestParamsNew struct {
	CommonParams common.CommonParams
	CommodityId  string `json:"commodityId"`
	//JenisInformasi       string `json:"jenisInformasi"`
	//DetailJenisInformasi string `json:"detailJenisInformasi"`
	SelectedDate string `json:"selectedDate"`
}

type PriceGetHargaRequestParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams"`
	CommodityId      string                  `json:"commodityId" validate:"required"`
	SelectedDate     string                  `json:"selectedDate" validate:"required"`
}

// PriceGetRepoParams - FOR REPOSITORIES
type PriceGetRepoParams struct {
	CommodityType string `db:"commodityType"`
	StartDate     string `db:"startDate"`
	EndDate       string `db:"endDate"`
}

type PriceGetRepoParamsNew struct {
	CommodityId      string                  `json:"commodityId"`
	SelectedDate     string                  `json:"selectedDate"`
	PaginationParams common.PaginationParams `json:"paginationParams"`
}

type PriceLast5DaysRepoParams struct {
	CityId      string `json:"cityId"`
	CommodityId string `json:"commodityId"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

type PriceLast5Days struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams" validate:"required"`
	common.CityIdParam
	common.CommodityIdParam
}

type PriceListParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams" validate:"required"`
	common.ProvinceIdParam
	common.CityIdParam
	common.CommodityIdParam
	common.PeriodDateParam
}

type PriceListRepoParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams" validate:"required"`
	ProvinceId       string                  `json:"provinceId"`
	CityId           string                  `json:"cityId"`
	CommodityId      string                  `json:"commodityId"`
	StartDate        string                  `json:"startDate"`
	EndDate          string                  `json:"endDate"`
}

type PriceGetCompareProvinceParams struct {
	CommodityId      string                  `json:"commodityId"`
	ProvinceId       string                  `json:"provinceId"`
	SelectedDate     string                  `json:"selectedDate"`
	PaginationParams common.PaginationParams `json:"paginationParams"`
}

type PriceGetCompareProvinceCommodityHistoryParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams" validate:"required"`
	CommodityId      string                  `json:"commodityId"`
	SelectedDate     string                  `json:"selectedDate"`
	Status           string                  `json:"status"`
}

type PriceGetCompareNationalParams struct {
	CommodityId      string                  `json:"commodityId"`
	ProvinceId       string                  `json:"provinceId"`
	NationalId       string                  `json:"nationalId"`
	SelectedDate     string                  `json:"selectedDate"`
	PaginationParams common.PaginationParams `json:"paginationParams"`
}

type PriceDiffRequestParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams"`
	CommodityId      string                  `json:"commodityId"`
	CityId           string                  `json:"cityId"`
	StartDate        string                  `json:"startDate"`
	EndDate          string                  `json:"endDate"`
	Status           string                  `json:"status"`
}

type PriceDiffByCityAndCommodityParams struct {
	CommodityId      string                  `json:"commodityId"`
	ProvinceId       string                  `json:"provinceId"`
	CityId           string                  `json:"cityId"`
	StartDate        string                  `json:"startDate"`
	EndDate          string                  `json:"endDate"`
	PaginationParams common.PaginationParams `json:"paginationParams"`
	Status           string                  `json:"status"`
}

type PriceGetPerbandinganParams struct {
	CommodityId      string                  `json:"commodityId"`
	ProvinceId       string                  `json:"provinceId"`
	SelectedDate     string                  `json:"selectedDate"`
	PaginationParams common.PaginationParams `json:"paginationParams"`
}

type PricePerubahanParams struct {
	CommodityId      string                  `json:"commodityId"`
	ProvinceId       string                  `json:"provinceId"`
	SelectedDate     string                  `json:"selectedDate"`
	PaginationParams common.PaginationParams `json:"paginationParams"`
}

type PriceMtmCityHistoryParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams" validate:"required"`
	CityId           string                  `json:"cityId"`
	ProvinceId       string                  `json:"provinceId"`
	CommodityId      string                  `json:"commodityId"`
	common.PeriodDateParam
}

type PriceMtmCommodityHistoryParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams" validate:"required"`
	CommodityId      string                  `json:"commodityId"`
	SelectedDate     string                  `json:"selectedDate"`
	Status           string                  `json:"status"`
}

type PriceLastUpdateParams struct {
	NationalId   string `json:"nationalId"`
	ProvinceId   string `json:"provinceId"`
	CityId       string `json:"cityId"`
	CommodityId  string `json:"commodityId"`
	SelectedDate string `json:"selectedDate"`
}

type HistoryRequestParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams"`
	Module           string                  `json:"module" validate:"required"`
	Search           *string                 `json:"search"`
}
