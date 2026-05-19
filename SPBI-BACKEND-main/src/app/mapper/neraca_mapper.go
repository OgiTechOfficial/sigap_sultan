package mapper

import (
	"github.com/gofiber/fiber/v2/log"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories"
	"strings"
)

type NeracaMapper struct {
	TmCityRepository      *repositories.TmCityRepository
	TmProvinceRepository  *repositories.TmProvinceRepository
	TmCommodityRepository *repositories.TmCommodityRepository
}

func NewNeracaMapper(tmCityRepository *repositories.TmCityRepository, tmCommodityRepository *repositories.TmCommodityRepository, tmProvinceRepository *repositories.TmProvinceRepository) *NeracaMapper {
	return &NeracaMapper{
		TmCityRepository:      tmCityRepository,
		TmCommodityRepository: tmCommodityRepository,
		TmProvinceRepository:  tmProvinceRepository,
	}
}

func (r *NeracaMapper) StringArrToNeracaCityModel(params []string) (*models.NeracaCity, error) {
	var err error
	var neracaData models.NeracaCity
	var tmCommodity *models.TmCommodity
	var currCommodityName string
	var currCityName string
	var tmCity *models.TmCity
	var tmCommodityFromRepo interface{}

	for i := 0; i < len(params); i++ {
		if i == 0 {
			lastUpdate := params[i] + " 16:00:00"
			neracaData.LastUpdate = &lastUpdate
		} else if i == 1 {
			if tmCity == nil {
				tmCity, err = r.TmCityRepository.GetByName(strings.ToLower(params[i]))
				if err != nil {
					log.Error("NeracaMapper.StringArrToNeracaCityModel.Error.Params: %s\n", params[i])
					return nil, err
				}

				currCityName = tmCity.Name
				neracaData.CityId = *tmCity.Id.Id
			} else {
				if currCityName != tmCity.Name {
					tmCity, err = r.TmCityRepository.GetByName(strings.ToLower(params[i]))
					if err != nil {
						log.Error("NeracaMapper.StringArrToNeracaCityModel.Error.Params: %s\n", params[i])
						return nil, err
					}

					currCityName = tmCity.Name
					neracaData.CityId = *tmCity.Id.Id
				} else {
					neracaData.CityId = *tmCity.Id.Id
				}
			}
		} else if i == 2 {
			if strings.Contains(params[i], "'") {
				params[i] = strings.ReplaceAll(params[i], "'", "''")
			}

			if tmCommodity == nil {
				tmCommodity, err = r.TmCommodityRepository.GetByEqualName(strings.ToLower(params[i]))
				if err != nil {
					if err.Error() != "no rows in result set" {
						log.Error("NeracaMapper.StringArrToNeracaCityModel.Error.Params: %s\n", params[i])
						return nil, err
					}

					tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
					if err != nil {
						log.Error("NeracaMapper.StringArrToNeracaCityModel.Error.Params: %s\n", params[i])
						return nil, err
					}
					tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
					currCommodityName = tmCommodity.Name
					neracaData.CommodityId = *tmCommodity.Id.Id
				} else {
					currCommodityName = tmCommodity.Name
					neracaData.CommodityId = *tmCommodity.Id.Id
				}
			} else {
				if currCommodityName != tmCommodity.Name {
					tmCommodity, err = r.TmCommodityRepository.GetByEqualName(strings.ToLower(params[i]))
					if err != nil {
						if err.Error() != "no rows in result set" {
							log.Error("NeracaMapper.StringArrToNeracaCityModel.Error.Params: %s\n", params[i])
							return nil, err
						}

						tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
						if err != nil {
							log.Error("NeracaMapper.StringArrToNeracaCityModel.Error.Params: %s\n", params[i])
							return nil, err
						}
						tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
						currCommodityName = tmCommodity.Name
						neracaData.CommodityId = *tmCommodity.Id.Id
					} else {
						currCommodityName = tmCommodity.Name
						neracaData.CommodityId = *tmCommodity.Id.Id
					}
				} else {
					neracaData.CommodityId = *tmCommodity.Id.Id
				}
			}

			neracaData.CommodityName = params[i]
		} else if i == 3 {
			neracaData.Ketersediaan = params[i]
		} else if i == 4 {
			neracaData.Kebutuhan = params[i]
		} else if i == 5 {
			neracaData.Neraca = params[i]
		}
	}

	return &neracaData, nil
}

func (r *NeracaMapper) StringArrToNeracaProvinceModel(params []string) (*models.NeracaProvince, error) {
	var err error
	var neracaData models.NeracaProvince
	var currProvinceName string
	var currCommodityName string
	var tmProvince *models.TmProvince
	var tmCommodity *models.TmCommodity
	var tmCommodityFromRepo interface{}

	for i := 0; i < len(params); i++ {
		if i == 0 {
			lastUpdate := params[i] + " 16:00:00"
			neracaData.LastUpdate = &lastUpdate
		} else if i == 1 {
			if tmProvince == nil {
				tmProvince, err = r.TmProvinceRepository.GetByName(strings.ToLower(params[i]))
				if err != nil {
					log.Error("NeracaMapper.StringArrToNeracaProvinceModel.Error.Params: %s\n", params[i])
					return nil, err
				}

				currProvinceName = tmProvince.Name
				neracaData.ProvinceId = *tmProvince.Id.Id
			} else {
				if currProvinceName != tmProvince.Name {
					tmProvince, err = r.TmProvinceRepository.GetByName(strings.ToLower(params[i]))
					if err != nil {
						log.Error("NeracaMapper.StringArrToNeracaProvinceModel.Error.Params: %s\n", params[i])
						return nil, err
					}

					currProvinceName = tmProvince.Name
					neracaData.ProvinceId = *tmProvince.Id.Id
				} else {
					neracaData.ProvinceId = *tmProvince.Id.Id
				}
			}
		} else if i == 2 {
			if strings.Contains(params[i], "'") {
				params[i] = strings.ReplaceAll(params[i], "'", "''")
			}

			if tmCommodity == nil {
				tmCommodity, err = r.TmCommodityRepository.GetByEqualName(strings.ToLower(params[i]))
				if err != nil {
					if err.Error() != "no rows in result set" {
						log.Error("NeracaMapper.StringArrToNeracaCityModel.Error.Params: %s\n", params[i])
						return nil, err
					}

					tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
					if err != nil {
						log.Error("NeracaMapper.StringArrToNeracaCityModel.Error.Params: %s\n", params[i])
						return nil, err
					}
					tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
					currCommodityName = tmCommodity.Name
					neracaData.CommodityId = *tmCommodity.Id.Id
				} else {
					currCommodityName = tmCommodity.Name
					neracaData.CommodityId = *tmCommodity.Id.Id
				}
			} else {
				if currCommodityName != tmCommodity.Name {
					tmCommodity, err = r.TmCommodityRepository.GetByEqualName(strings.ToLower(params[i]))
					if err != nil {
						if err.Error() != "no rows in result set" {
							log.Error("NeracaMapper.StringArrToNeracaCityModel.Error.Params: %s\n", params[i])
							return nil, err
						}

						tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
						if err != nil {
							log.Error("NeracaMapper.StringArrToNeracaCityModel.Error.Params: %s\n", params[i])
							return nil, err
						}
						tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
						currCommodityName = tmCommodity.Name
						neracaData.CommodityId = *tmCommodity.Id.Id
					} else {
						currCommodityName = tmCommodity.Name
						neracaData.CommodityId = *tmCommodity.Id.Id
					}
				} else {
					neracaData.CommodityId = *tmCommodity.Id.Id
				}
			}

			neracaData.CommodityName = params[i]
		} else if i == 3 {
			neracaData.Ketersediaan = params[i]
		} else if i == 4 {
			neracaData.Kebutuhan = params[i]
		} else if i == 5 {
			neracaData.Neraca = params[i]
		}
	}

	return &neracaData, nil
}

func (r *NeracaMapper) StringArrToNeracaNationalModel(params []string) (*models.NeracaNational, error) {
	var err error
	var neracaData models.NeracaNational
	var tmCommodity *models.TmCommodity
	var currCommodityName string
	var tmCommodityFromRepo interface{}

	for i := 0; i < len(params); i++ {
		if i == 0 {
			lastUpdate := params[i] + " 16:00:00"
			neracaData.LastUpdate = &lastUpdate
		} else if i == 1 {
			neracaData.NationalId = 1
			neracaData.NatinoalName = "INDONESIA"
		} else if i == 2 {
			if strings.Contains(params[i], "'") {
				params[i] = strings.ReplaceAll(params[i], "'", "''")
			}

			if tmCommodity == nil {
				tmCommodity, err = r.TmCommodityRepository.GetByEqualName(strings.ToLower(params[i]))
				if err != nil {
					if err.Error() != "no rows in result set" {
						log.Error("NeracaMapper.StringArrToNeracaNationalModel.Error.Params: %s\n", params[i])
						return nil, err
					}

					tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
					if err != nil {
						log.Error("NeracaMapper.StringArrToNeracaNationalModel.Error.Params: %s\n", params[i])
						return nil, err
					}
					tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
					currCommodityName = tmCommodity.Name
					neracaData.CommodityId = *tmCommodity.Id.Id
				} else {
					currCommodityName = tmCommodity.Name
					neracaData.CommodityId = *tmCommodity.Id.Id
				}
			} else {
				if currCommodityName != tmCommodity.Name {
					tmCommodity, err = r.TmCommodityRepository.GetByEqualName(strings.ToLower(params[i]))
					if err != nil {
						if err.Error() != "no rows in result set" {
							log.Error("NeracaMapper.StringArrToNeracaNationalModel.Error.Params: %s\n", params[i])
							return nil, err
						}

						tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
						if err != nil {
							log.Error("NeracaMapper.StringArrToNeracaNationalModel.Error.Params: %s\n", params[i])
							return nil, err
						}
						tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
						currCommodityName = tmCommodity.Name
						neracaData.CommodityId = *tmCommodity.Id.Id
					} else {
						currCommodityName = tmCommodity.Name
						neracaData.CommodityId = *tmCommodity.Id.Id
					}
				} else {
					neracaData.CommodityId = *tmCommodity.Id.Id
				}
			}

			neracaData.CommodityName = params[i]
		} else if i == 3 {
			neracaData.Ketersediaan = params[i]
		} else if i == 4 {
			neracaData.Kebutuhan = params[i]
		} else if i == 5 {
			neracaData.Neraca = params[i]
		}
	}

	return &neracaData, nil
}
