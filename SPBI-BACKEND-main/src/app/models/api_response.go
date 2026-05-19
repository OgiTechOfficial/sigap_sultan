package models

import "sigap-sultan-be/src/common"

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ApiResponsePagination struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	common.PaginationRespnose
}

type ApiResponseError struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func Set() *ApiResponse {
	return &ApiResponse{}
}

func SetPagination() *ApiResponsePagination {
	return &ApiResponsePagination{}
}

func SetPagination2() *ApiResponse {
	return &ApiResponse{}
}

func SetError() *ApiResponseError {
	return &ApiResponseError{}
}

func (r *ApiResponse) Success(message string, data interface{}) *ApiResponse {
	r.Status = 200
	if message != "" {
		r.Message = message
	} else {
		r.Message = "Success"
	}

	r.Data = data

	return r
}

func (r *ApiResponsePagination) Success(message string, data interface{}, page int, totalPage int, totalData int) *ApiResponsePagination {
	r.Status = 200
	if message != "" {
		r.Message = message
	} else {
		r.Message = "Success"
	}

	r.Data = data
	r.Page = page
	r.TotalPage = totalPage
	r.TotalData = totalData

	return r
}

func (r *ApiResponse) SuccessPagination(message string, data interface{}) *ApiResponse {
	r.Status = 200
	if message != "" {
		r.Message = message
	} else {
		r.Message = "Success"
	}

	r.Data = data

	return r
}

func (r *ApiResponseError) Details(status int, message string, err interface{}) *ApiResponseError {
	r.Status = status
	r.Message = message
	r.Errors = err

	return r
}
