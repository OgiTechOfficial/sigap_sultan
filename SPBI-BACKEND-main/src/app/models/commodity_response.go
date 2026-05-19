package models

type CommodityResponse struct {
	Id       int32           `json:"id" db:"id"`
	ParentId *int32          `json:"parentId" db:"parent_id"`
	Name     *string         `json:"name" db:"name"`
	Assets   *AssetsResponse `json:"assets" db:"assets"`
	AuditRail
}

type CommodityPriceTableResponse struct {
	Id
	Name string `json:"name" db:"name"`
}

type CommodityWithPriceListResponse []struct {
	Id       int32           `json:"id" db:"id"`
	ParentId *int32          `json:"parentId" db:"parent_id"`
	Name     *string         `json:"name" db:"name"`
	Assets   *AssetsResponse `json:"assets" db:"assets"`
	AuditRail
	Period                   *PeriodResponse    `json:"period"`
	CurrentPrice             int                `json:"currentPrice"`
	CurrentPriceRupiahFormat string             `json:"currentPriceRupiahFormat"`
	PriceDiffLast5Days       int                `json:"priceDiffLast5Days"`
	PriceDiffLast5DaysFormat string             `json:"priceDiffLast5DaysFormat"`
	PriceList                *PriceListResponse `json:"price"`
}

type PeriodResponse struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
