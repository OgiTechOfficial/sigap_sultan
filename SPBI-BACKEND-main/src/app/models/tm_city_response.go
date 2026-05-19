package models

type CityResponse struct {
	Id
	Name   string          `json:"name" db:"name"`
	Assets *AssetsResponse `json:"assets" db:"assets"`
}

type CityLevelHargaResponse struct {
	Id
	Name   string          `json:"name" db:"name"`
	Assets *AssetsResponse `json:"assets" db:"assets"`
	AuditRail
}

type CityWithPriceListResponse []struct {
	Id
	Name   string          `json:"name" db:"name"`
	Assets *AssetsResponse `json:"assets" db:"assets"`
	AuditRail
	Period                   *PeriodResponse    `json:"period"`
	CurrentPrice             int                `json:"currentPrice"`
	CurrentPriceRupiahFormat string             `json:"currentPriceRupiahFormat"`
	PriceDiffLast5Days       int                `json:"priceDiffLast5Days"`
	PriceDiffLast5DaysFormat string             `json:"priceDiffLast5DaysFormat"`
	PriceList                *PriceListResponse `json:"price"`
}
