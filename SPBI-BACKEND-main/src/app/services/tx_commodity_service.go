package services

import (
	"sigap-sultan-be/src/app/repositories"
)

type TxCommodityService struct {
	PriceRepository       *repositories.PriceRepository
	TmCommodityRepository *repositories.TmCommodityRepository
	TmCityRepository      *repositories.TmCityRepository
}
