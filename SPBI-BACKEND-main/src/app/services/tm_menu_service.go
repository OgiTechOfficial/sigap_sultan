package services

import (
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories"
	"sigap-sultan-be/src/common"
)

type TmMenuService struct {
	TmMenuRepository *repositories.TmMenuRepository
}

func NewTmMenuService(tmMenuRepository *repositories.TmMenuRepository) *TmMenuService {
	return &TmMenuService{
		TmMenuRepository: tmMenuRepository,
	}
}

func (r *TmMenuService) Get(paginationParams common.PaginationParams) (interface{}, *common.ErrorDomain) {
	data, err := r.TmMenuRepository.Get(paginationParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmMenuService) Count() (*int, *common.ErrorDomain) {
	data, err := r.TmMenuRepository.GetCount()
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmMenuService) GetById(id int) (interface{}, *common.ErrorDomain) {
	data, err := r.TmMenuRepository.GetById(id)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmMenuService) GetByName(name string, paginationParams common.PaginationParams) (interface{}, *common.ErrorDomain) {
	data, err := r.TmMenuRepository.GetByName(name, paginationParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmMenuService) CountByName(name string, paginationParams common.PaginationParams) (*int, *common.ErrorDomain) {
	data, err := r.TmMenuRepository.GetByNameCount(name, paginationParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmMenuService) Insert(request models.TmMenu) (interface{}, *common.ErrorDomain) {
	request.ClientId = 1
	data, err := r.TmMenuRepository.Insert(request)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmMenuService) Update(id int, request models.TmMenu) (interface{}, *common.ErrorDomain) {
	data, err := r.TmMenuRepository.Update(id, request)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmMenuService) Delete(id int) (interface{}, *common.ErrorDomain) {
	data, err := r.TmMenuRepository.Delete(id)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}
