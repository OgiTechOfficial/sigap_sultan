package domain

import "sigap-sultan-be/src/common"

type ProvinceRequestParams struct {
	CommonParams     common.CommonParams
	PaginationParams common.PaginationParams `json:"paginationParams"`
	Id               int                     `json:"id" validate:"required"`
}
