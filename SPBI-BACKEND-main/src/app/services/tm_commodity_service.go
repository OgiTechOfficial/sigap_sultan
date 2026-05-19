package services

import (
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/repositories"
	"sigap-sultan-be/src/common"
)

type TmCommodityService struct {
	TmCommodityRepository *repositories.TmCommodityRepository
}

func NewTmCommodityService(tmCommodityRepository *repositories.TmCommodityRepository) *TmCommodityService {
	return &TmCommodityService{
		TmCommodityRepository: tmCommodityRepository,
	}
}

func (r *TmCommodityService) Get(params domain.TmCommodityRequestParam) (interface{}, *common.ErrorDomain) {
	data, err := r.TmCommodityRepository.Get(params)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmCommodityService) GetGrouping() (interface{}, *common.ErrorDomain) {
	data, err := r.TmCommodityRepository.GetGrouping()
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmCommodityService) GetById(id int) (interface{}, *common.ErrorDomain) {
	var err error
	var data interface{}

	if id != 0 {
		data, err = r.TmCommodityRepository.GetById(int32(id))
		if err != nil {
			return nil, &common.ErrorDomain{
				Message: err.Error(),
			}
		}
	}

	return data, nil
}
