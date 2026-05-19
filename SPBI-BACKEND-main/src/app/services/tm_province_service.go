package services

import (
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/repositories"
	"sigap-sultan-be/src/common"
)

type TmProvinceService struct {
	TmProvinceRepository *repositories.TmProvinceRepository
}

func NewTmProvinceService(tmProvinceRepository *repositories.TmProvinceRepository) *TmProvinceService {
	return &TmProvinceService{
		TmProvinceRepository: tmProvinceRepository,
	}
}

func (r *TmProvinceService) Get(params domain.ProvinceRequestParams) (interface{}, *common.ErrorDomain) {
	var err error
	var data interface{}

	if params.Id != 0 {
		data, err = r.TmProvinceRepository.GetByProvince(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	} else {
		data, err = r.TmProvinceRepository.Get(params)
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}

func (r *TmProvinceService) Count(params domain.ProvinceRequestParams) (*int, *common.ErrorDomain) {
	var err error
	var data *int

	if params.Id != 0 {
		//data, err = r.TmProvinceRepository.CountByProvinceId(params)
		//if err != nil {
		//	return nil, &common.ErrorDomain{
		//		Message: err.Error(),
		//	}
		//}
	} else {
		data, err = r.TmProvinceRepository.Count()
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}

func (r *TmProvinceService) GetById(params domain.ProvinceRequestParams) (interface{}, *common.ErrorDomain) {
	var err error
	var data interface{}

	data, err = r.TmProvinceRepository.GetByProvince(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}
