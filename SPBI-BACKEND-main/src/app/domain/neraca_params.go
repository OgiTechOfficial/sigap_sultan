package domain

import "sigap-sultan-be/src/common"

type NeracaStokAkhirListRequestParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams"`
	CommodityId      string                  `json:"commodityId"`
	ProvinceId       string                  `json:"provinceId"`
	CityId           string                  `json:"cityId"`
	SelectedDate     string                  `json:"selectedDate" validate:"required"`
	StartDate        string                  `json:"startDate" validate:"required"`
	EndDate          string                  `json:"endDate" validate:"required"`
	Status           string                  `json:"status"`
}

type NeracaListRepoParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams" validate:"required"`
	ProvinceId       string                  `json:"provinceId"`
	CityId           string                  `json:"cityId"`
	CommodityId      string                  `json:"commodityId"`
	StartDate        string                  `json:"startDate"`
	EndDate          string                  `json:"endDate"`
}

type NeracaListParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams" validate:"required"`
	common.ProvinceIdParam
	common.CityIdParam
	common.CommodityIdParam
	common.PeriodDateParam
}
