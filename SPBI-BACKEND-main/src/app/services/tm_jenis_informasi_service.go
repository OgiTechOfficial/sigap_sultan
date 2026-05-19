package services

import (
	"sigap-sultan-be/src/app/repositories"
	"sigap-sultan-be/src/common"
)

type TmJenisInformasiService struct {
	TmJenisInformasiRepository *repositories.TmJenisInformasiRepository
}

func NewTmJenisInformasiService(tmJenisInformasiRepository *repositories.TmJenisInformasiRepository) *TmJenisInformasiService {
	return &TmJenisInformasiService{
		TmJenisInformasiRepository: tmJenisInformasiRepository,
	}
}

func (r *TmJenisInformasiService) Get(paginationParams common.PaginationParams) (interface{}, *common.ErrorDomain) {
	data, err := r.TmJenisInformasiRepository.Get(paginationParams)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmJenisInformasiService) GetById(id int) (interface{}, *common.ErrorDomain) {
	data, err := r.TmJenisInformasiRepository.GetById(id)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

func (r *TmJenisInformasiService) GetByName(name string) (interface{}, *common.ErrorDomain) {
	data, err := r.TmJenisInformasiRepository.GetByName(name)
	if err != nil {
		return nil, &common.ErrorDomain{
			Message: err.Error(),
		}
	}

	return data, nil
}

//func (r *TmJenisInformasiService) Insert(request models.JenisInformasiRequest) (interface{}, *common.ErrorDomain) {
//	data, err := r.TmJenisInformasiRepository.Insert(request)
//	if err != nil {
//		return nil, &common.ErrorDomain{
//			Message: err.Error(),
//		}
//	}
//
//	return data, nil
//}
