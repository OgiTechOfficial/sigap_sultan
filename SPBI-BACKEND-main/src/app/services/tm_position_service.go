package services

import (
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories"
	"sigap-sultan-be/src/common"
)

type TmPositionService struct {
	TmPositionRepository *repositories.TmPositionRepository
}

func NewTmPositionService(tmPositionRepository *repositories.TmPositionRepository) *TmPositionService {
	return &TmPositionService{
		TmPositionRepository: tmPositionRepository,
	}
}

func (r *TmPositionService) Get(paginationParams common.PaginationParams) (interface{}, *common.ErrorDomain) {
	data, err := r.TmPositionRepository.Get(paginationParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmPositionService) Count() (*int, *common.ErrorDomain) {
	data, err := r.TmPositionRepository.Count()
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmPositionService) GetById(id int) (interface{}, *common.ErrorDomain) {
	data, err := r.TmPositionRepository.GetById(id)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmPositionService) GetByName(name string) (interface{}, *common.ErrorDomain) {
	data, err := r.TmPositionRepository.GetByName(name)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmPositionService) CountByName(name string) (*int, *common.ErrorDomain) {
	data, err := r.TmPositionRepository.CountByName(name)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmPositionService) Insert(request models.PrivilegesRequest) (interface{}, *common.ErrorDomain) {
	data, err := r.TmPositionRepository.Insert(request)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmPositionService) Delete(id int) (interface{}, *common.ErrorDomain) {
	data, err := r.TmPositionRepository.Delete(id)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}
