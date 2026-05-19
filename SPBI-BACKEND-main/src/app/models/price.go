package models

type PriceCity struct {
	Id
	ClientId      int32   `json:"clientId" db:"client_id"`
	CityId        int32   `json:"cityId" db:"city_id"`
	CommodityId   int32   `json:"commodityId" db:"commodity_id"`
	CommodityName string  `json:"commodityName" db:"commodity_name"`
	Price         string  `json:"price" db:"price"`
	LastUpdate    *string `json:"lastUpdate" db:"last_update"`
	AuditRail
}

type PriceProvince struct {
	Id
	ClientId      int32   `json:"clientId" db:"client_id"`
	ProvinceId    int32   `json:"provinceId" db:"province_id"`
	CommodityId   int32   `json:"commodityId" db:"commodity_id"`
	CommodityName string  `json:"commodityName" db:"commodity_name"`
	Price         string  `json:"price" db:"price"`
	LastUpdate    *string `json:"lastUpdate" db:"last_update"`
	AuditRail
}

type PriceNational struct {
	Id
	ClientId      int32   `json:"clientId" db:"client_id"`
	NationalId    int32   `json:"nationalId" db:"national_id"`
	CommodityId   int32   `json:"commodityId" db:"commodity_id"`
	CommodityName string  `json:"commodityName" db:"commodity_name"`
	Price         string  `json:"price" db:"price"`
	LastUpdate    *string `json:"lastUpdate" db:"last_update"`
	AuditRail
}
