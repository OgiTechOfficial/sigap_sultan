package services

import (
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/repositories"
	"sigap-sultan-be/src/common"
)

type TmCityService struct {
	TmCityRepository *repositories.TmCityRepository
}

func NewTmCityService(tmCityRepository *repositories.TmCityRepository) *TmCityService {
	return &TmCityService{
		TmCityRepository: tmCityRepository,
	}
}

func (r *TmCityService) Get(params domain.CityRequestParams) (interface{}, *common.ErrorDomain) {
	var err error
	var data interface{}

	if params.ProvinceId != "" {
		data, err = r.TmCityRepository.GetByProvince(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	} else {
		data, err = r.TmCityRepository.Get(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}

func (r *TmCityService) Count(params domain.CityRequestParams) (*int, *common.ErrorDomain) {
	var err error
	var data *int

	if params.ProvinceId != "" {
		data, err = r.TmCityRepository.CountByProvinceId(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	} else {
		data, err = r.TmCityRepository.Count()
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}

func (r *TmCityService) GetById(id int) (interface{}, *common.ErrorDomain) {
	var err error
	var data interface{}

	if id != 0 {
		data, err = r.TmCityRepository.GetById(id)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}
