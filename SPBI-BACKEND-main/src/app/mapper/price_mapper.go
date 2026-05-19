package mapper

import (
	"github.com/gofiber/fiber/v2/log"
	"sigap-sultan-be/src/app/domain"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/repositories"
	"strconv"
	"strings"
)

type PriceMapper struct {
	TmCityRepository      *repositories.TmCityRepository
	TmProvinceRepository  *repositories.TmProvinceRepository
	TmCommodityRepository *repositories.TmCommodityRepository
}

func NewPriceMapper(tmCityRepository *repositories.TmCityRepository, tmCommodityRepository *repositories.TmCommodityRepository, tmProvinceRepository *repositories.TmProvinceRepository) *PriceMapper {
	return &PriceMapper{
		TmCityRepository:      tmCityRepository,
		TmCommodityRepository: tmCommodityRepository,
		TmProvinceRepository:  tmProvinceRepository,
	}
}

func (r *PriceMapper) StringArrToPriceCityModel(params []string) (*models.PriceCity, error) {
	var err error
	var priceData models.PriceCity
	var tmCommodity *models.TmCommodity
	var currCommodityName string
	var currCityName string
	var tmCity *models.TmCity
	var tmCommodityFromRepo interface{}

	for i := 0; i < len(params); i++ {
		if i == 0 {
			lastUpdate := params[i] + " 16:00:00"
			priceData.LastUpdate = &lastUpdate
		} else if i == 1 {
			if tmCity == nil {
				tmCity, err = r.TmCityRepository.GetByName(strings.ToLower(params[i]))
				if err != nil {
					log.Error("PriceMapper.StringArrToPriceCityModel.Error.Params: %s\n", params[i])
					return nil, err
				}

				currCityName = tmCity.Name
				priceData.CityId = *tmCity.Id.Id
			} else {
				if currCityName != tmCity.Name {
					tmCity, err = r.TmCityRepository.GetByName(strings.ToLower(params[i]))
					if err != nil {
						log.Error("PriceMapper.StringArrToPriceCityModel.Error.Params: %s\n", params[i])
						return nil, err
					}

					currCityName = tmCity.Name
					priceData.CityId = *tmCity.Id.Id
				} else {
					priceData.CityId = *tmCity.Id.Id
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
						log.Error("PriceMapper.StringArrToPriceCityModel.Error.Params: %s\n", params[i])
						return nil, err
					}

					tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
					if err != nil {
						log.Error("PriceMapper.StringArrToPriceCityModel.Error.Params: %s\n", params[i])
						return nil, err
					}
					tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
					currCommodityName = tmCommodity.Name
					priceData.CommodityId = *tmCommodity.Id.Id
				} else {
					currCommodityName = tmCommodity.Name
					priceData.CommodityId = *tmCommodity.Id.Id
				}
			} else {
				if currCommodityName != tmCommodity.Name {
					tmCommodity, err = r.TmCommodityRepository.GetByEqualName(strings.ToLower(params[i]))
					if err != nil {
						if err.Error() != "no rows in result set" {
							log.Error("PriceMapper.StringArrToPriceCityModel.Error.Params: %s\n", params[i])
							return nil, err
						}

						tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
						if err != nil {
							log.Error("PriceMapper.StringArrToPriceCityModel.Error.Params: %s\n", params[i])
							return nil, err
						}
						tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
						currCommodityName = tmCommodity.Name
						priceData.CommodityId = *tmCommodity.Id.Id
					} else {
						currCommodityName = tmCommodity.Name
						priceData.CommodityId = *tmCommodity.Id.Id
					}
				} else {
					priceData.CommodityId = *tmCommodity.Id.Id
				}
			}

			priceData.CommodityName = params[i]
		} else if i == 3 {
			_, atoiErr := strconv.Atoi(params[i])
			if atoiErr != nil {
				priceData.Price = "0"
			} else {
				priceData.Price = params[i]
			}
			if atoiErr != nil {
				priceData.Price = "0"
			} else {
				priceData.Price = params[i]
			}
		}
	}

	return &priceData, nil
}

func (r *PriceMapper) StringArrToPriceProvinceModel(params []string) (*models.PriceProvince, error) {
	var err error
	var priceData models.PriceProvince
	var currProvinceName string
	var currCommodityName string
	var tmProvince *models.TmProvince
	var tmCommodity *models.TmCommodity
	var tmCommodityFromRepo interface{}

	for i := 0; i < len(params); i++ {
		if i == 0 {
			lastUpdate := params[i] + " 16:00:00"
			priceData.LastUpdate = &lastUpdate
		} else if i == 1 {
			if tmProvince == nil {
				tmProvince, err = r.TmProvinceRepository.GetByName(strings.ToLower(params[i]))
				if err != nil {
					log.Error("PriceMapper.StringArrToPriceProvinceModel.Error.Params: %s\n", params[i])
					return nil, err
				}

				currProvinceName = tmProvince.Name
				priceData.ProvinceId = *tmProvince.Id.Id
			} else {
				if currProvinceName != tmProvince.Name {
					tmProvince, err = r.TmProvinceRepository.GetByName(strings.ToLower(params[i]))
					if err != nil {
						log.Error("PriceMapper.StringArrToPriceProvinceModel.Error.Params: %s\n", params[i])
						return nil, err
					}

					currProvinceName = tmProvince.Name
					priceData.ProvinceId = *tmProvince.Id.Id
				} else {
					priceData.ProvinceId = *tmProvince.Id.Id
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
						log.Error("PriceMapper.StringArrToPriceProvinceModel.Error.Params: %s\n", params[i])
						return nil, err
					}

					tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
					if err != nil {
						log.Error("PriceMapper.StringArrToPriceProvinceModel.Error.Params: %s\n", params[i])
						return nil, err
					}
					tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
					currCommodityName = tmCommodity.Name
					priceData.CommodityId = *tmCommodity.Id.Id
				} else {
					currCommodityName = tmCommodity.Name
					priceData.CommodityId = *tmCommodity.Id.Id
				}
			} else {
				if currCommodityName != tmCommodity.Name {
					tmCommodity, err = r.TmCommodityRepository.GetByEqualName(strings.ToLower(params[i]))
					if err != nil {
						if err.Error() != "no rows in result set" {
							log.Error("PriceMapper.StringArrToPriceProvinceModel.Error.Params: %s\n", params[i])
							return nil, err
						}

						tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
						if err != nil {
							log.Error("PriceMapper.StringArrToPriceProvinceModel.Error.Params: %s\n", params[i])
							return nil, err
						}
						tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
						currCommodityName = tmCommodity.Name
						priceData.CommodityId = *tmCommodity.Id.Id
					} else {
						currCommodityName = tmCommodity.Name
						priceData.CommodityId = *tmCommodity.Id.Id
					}
				} else {
					priceData.CommodityId = *tmCommodity.Id.Id
				}
			}

			priceData.CommodityName = params[i]
		} else if i == 3 {
			_, atoiErr := strconv.Atoi(params[i])
			if atoiErr != nil {
				priceData.Price = "0"
			} else {
				priceData.Price = params[i]
			}
		}
	}

	return &priceData, nil
}

func (r *PriceMapper) StringArrToPriceNationalModel(params []string) (*models.PriceNational, error) {
	var err error
	var priceData models.PriceNational
	var tmCommodity *models.TmCommodity
	var currCommodityName string
	var tmCommodityFromRepo interface{}

	for i := 0; i < len(params); i++ {
		if i == 0 {
			lastUpdate := params[i] + " 16:00:00"
			priceData.LastUpdate = &lastUpdate
		} else if i == 1 {
			priceData.NationalId = 1
		} else if i == 2 {
			if strings.Contains(params[i], "'") {
				params[i] = strings.ReplaceAll(params[i], "'", "''")
			}

			if tmCommodity == nil {
				tmCommodity, err = r.TmCommodityRepository.GetByEqualName(strings.ToLower(params[i]))
				if err != nil {
					if err.Error() != "no rows in result set" {
						log.Error("PriceMapper.StringArrToPriceNationalModel.Error.Params: %s\n", params[i])
						return nil, err
					}

					tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
					if err != nil {
						log.Error("PriceMapper.StringArrToPriceNationalModel.Error.Params: %s\n", params[i])
						return nil, err
					}
					tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
					currCommodityName = tmCommodity.Name
					priceData.CommodityId = *tmCommodity.Id.Id
				} else {
					currCommodityName = tmCommodity.Name
					priceData.CommodityId = *tmCommodity.Id.Id
				}
			} else {
				if currCommodityName != tmCommodity.Name {
					tmCommodity, err = r.TmCommodityRepository.GetByEqualName(strings.ToLower(params[i]))
					if err != nil {
						if err.Error() != "no rows in result set" {
							log.Error("PriceMapper.StringArrToPriceNationalModel.Error.Params: %s\n", params[i])
							return nil, err
						}

						tmCommodityFromRepo, err = r.TmCommodityRepository.Insert(domain.TmCommodityParam{Name: params[i]})
						if err != nil {
							log.Error("PriceMapper.StringArrToPriceNationalModel.Error.Params: %s\n", params[i])
							return nil, err
						}
						tmCommodity = tmCommodityFromRepo.(*models.TmCommodity)
						currCommodityName = tmCommodity.Name
						priceData.CommodityId = *tmCommodity.Id.Id
					} else {
						currCommodityName = tmCommodity.Name
						priceData.CommodityId = *tmCommodity.Id.Id
					}
				} else {
					priceData.CommodityId = *tmCommodity.Id.Id
				}
			}

			priceData.CommodityName = params[i]
		} else if i == 3 {
			_, atoiErr := strconv.Atoi(params[i])
			if atoiErr != nil {
				priceData.Price = "0"
			} else {
				priceData.Price = params[i]
			}
		}
	}

	return &priceData, nil
}
