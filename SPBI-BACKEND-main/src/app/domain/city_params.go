package domain

import "sigap-sultan-be/src/common"

type CityRequestParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams"`
	ProvinceId       string                  `json:"provinceId" validate:"required"`
}
