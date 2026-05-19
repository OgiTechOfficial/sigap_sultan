package models

type NeracaCity struct {
	Id
	ClientId      int32   `json:"clientId" db:"client_id"`
	CityId        int32   `json:"cityId" db:"city_id"`
	CityName      string  `json:"cityName" db:"city_name"`
	CommodityId   int32   `json:"commodityId" db:"commodity_id"`
	CommodityName string  `json:"commodityName" db:"commodity_name"`
	Ketersediaan  string  `json:"ketersediaan" db:"ketersediaan"`
	Kebutuhan     string  `json:"kebutuhan" db:"kebutuhan"`
	Neraca        string  `json:"neraca" db:"neraca"`
	LastUpdate    *string `json:"lastUpdate" db:"last_update"`
	AuditRail
}

type NeracaProvince struct {
	Id
	ClientId      int32   `json:"clientId" db:"client_id"`
	ProvinceId    int32   `json:"provinceId" db:"province_id"`
	ProvinceName  string  `json:"provinceName" db:"province_name"`
	CommodityId   int32   `json:"commodityId" db:"commodity_id"`
	CommodityName string  `json:"commodityName" db:"commodity_name"`
	Ketersediaan  string  `json:"ketersediaan" db:"ketersediaan"`
	Kebutuhan     string  `json:"kebutuhan" db:"kebutuhan"`
	Neraca        string  `json:"neraca" db:"neraca"`
	LastUpdate    *string `json:"lastUpdate" db:"last_update"`
	AuditRail
}

type NeracaNational struct {
	Id
	ClientId      int32   `json:"clientId" db:"client_id"`
	NationalId    int32   `json:"nationalId" db:"national_id"`
	NatinoalName  string  `json:"nationalName" db:"national_name"`
	CommodityId   int32   `json:"commodityId" db:"commodity_id"`
	CommodityName string  `json:"commodityName" db:"commodity_name"`
	Ketersediaan  string  `json:"ketersediaan" db:"ketersediaan"`
	Kebutuhan     string  `json:"kebutuhan" db:"kebutuhan"`
	Neraca        string  `json:"neraca" db:"neraca"`
	LastUpdate    *string `json:"lastUpdate" db:"last_update"`
	AuditRail
}
