package models

// TABLE HARGA MAP
type PriceTableHargaLevelMap struct {
	PriceLevel       *[]PriceTableHargaLevelMapData `json:"priceLevel" db:"priceLevel"`
	PriceMin         *int32                         `json:"priceMin" db:"priceMin"`
	PriceMax         *int32                         `json:"priceMax" db:"priceMax"`
	PriceDiff        *int32                         `json:"priceDiff" db:"priceDiff"`
	PriceBarCategory *int32                         `json:"priceBarCategory" db:"priceBarCategory"`
	Tier             []PriceTier                    `json:"priceTier"`
}
type PriceTableHargaLevelNew struct {
	PriceProvince    *PriceProvinceResponse         `json:"provincePrice"`
	CityPrice        *[]PriceTableHargaLevelMapData `json:"cityPrice"`
	PriceMin         *int32                         `json:"priceMin" db:"priceMin"`
	PriceMax         *int32                         `json:"priceMax" db:"priceMax"`
	PriceDiff        *int32                         `json:"priceDiff" db:"priceDiff"`
	PriceBarCategory *int32                         `json:"priceBarCategory" db:"priceBarCategory"`
	Tier             []PriceTier                    `json:"priceTier"`
}

type PriceTableHargaLevelMapData struct {
	Id
	ClientId          int32                       `json:"clientId" db:"clientId"`
	City              CityResponse                `json:"city" db:"city"`
	Commodity         CommodityPriceTableResponse `json:"commodity" db:"commodity"`
	Price             *int32                      `json:"price" db:"price"`
	PriceRupiahFormat string                      `json:"priceRupiahFormat" db:"priceRupiahFormat"`
}

type PriceTier struct {
	Title                string `json:"title"`
	PriceMin             int32  `json:"priceMin" db:"priceMin"`
	PriceMinRupiahFormat string `json:"priceMinRupiahFormat" db:"priceMinRupiahFormat"`
	PriceMax             int32  `json:"priceMax" db:"priceMax"`
	PriceMaxRupiahFormat string `json:"priceMaxRupiahFormat" db:"priceMaxRupiahFormat"`
	Color                string `json:"color"`
}

type PriceDiffLevelHarga struct {
	PriceMin         *int32 `json:"priceMin" db:"priceMin"`
	PriceMax         *int32 `json:"priceMax" db:"priceMax"`
	PriceDiff        *int32 `json:"priceDiff" db:"priceDiff"`
	PriceBarCategory *int32 `json:"priceBarCategory" db:"priceBarCategory"`
}

type PriceTierLevelHarga struct {
	SangatRendah SangatRendah `json:"sangat-rendah"`
	Rendah       Rendah       `json:"rendah"`
	Sedang       Sedang       `json:"sedang"`
	Tinggi       Tinggi       `json:"tinggi"`
	SangatTinggi SangatTinggi `json:"sangat-tinggi"`
}

type PriceTierCompare struct {
	Higher Higher `json:"higher"`
	Same   Same   `json:"same"`
	Lower  Lower  `json:"lower"`
}

type (
	Higher struct {
		TitleColor
	}
	Same struct {
		TitleColor
	}
	Lower struct {
		TitleColor
	}

	SangatRendah struct {
		TitleColor
	}
	Rendah struct {
		TitleColor
	}
	Sedang struct {
		TitleColor
	}
	Tinggi struct {
		TitleColor
	}
	SangatTinggi struct {
		TitleColor
	}
)

// PriceTableHargaLevelList TABLE HARGA LIST
type PriceTableHargaLevelList struct {
	PriceProvince *PriceProvinceResponse          `json:"priceProvince"`
	PriceCity     *[]PriceTableHargaLevelListData `json:"priceCity"`
}
type PriceTableHargaLevelListData struct {
	Id
	ClientId          *int32                       `json:"clientId" db:"clientId"`
	City              *CityLevelHargaResponse      `json:"city" db:"city"`
	Commodity         *CommodityPriceTableResponse `json:"commodity" db:"commodity"`
	Price             *int32                       `json:"price" db:"price"`
	PriceRupiahFormat *string                      `json:"priceRupiahFormat" db:"priceRupiahFormat"`
	Tier              *string                      `json:"tier"`
	PriceTier         PriceTierLevelHarga          `json:"priceTier"`
	PriceTierCode     []string                     `json:"priceTierCode"`
}

// PriceLast5DaysCityByIdResponse TABLE HARGA LAST 5 DAYS BY CITY
type PriceLast5DaysCityByIdResponse struct {
	City        *CityLevelHargaResponse         `json:"city" db:"city"`
	Commodities *CommodityWithPriceListResponse `json:"commodities"`
}

// PriceLast5DaysCommodityByIdResponse TABLE HARGA LAST 5 DAYS BY COMMUNITY ID
type PriceLast5DaysCommodityByIdResponse struct {
	Commodities *CommodityResponse         `json:"commodity"`
	City        *CityWithPriceListResponse `json:"cities"`
}

type PriceListByCommodityAndCityResponse struct {
	Commodities     *CommodityResponse      `json:"commodity"`
	City            *CityLevelHargaResponse `json:"city" db:"city"`
	PriceDiff       *int                    `json:"priceDiff"`
	PriceDiffFormat *string                 `json:"priceDiffFormat"`
	PriceList       *PriceListResponse      `json:"price"`
}

type PriceListResponse []struct {
	Date              string  `json:"date"`
	Price             *int32  `json:"price" db:"price"`
	PriceRupiahFormat *string `json:"priceRupiahFormat" db:"priceRupiahFormat"`
}

type PriceDiffResponse struct {
	StablePriceCity *int32 `json:"stablePriceCity" db:"stablePriceCity"`
	LowPriceCity    *int32 `json:"lowPriceCity" db:"lowPriceCity"`
	HighPriceCity   *int32 `json:"highPriceCity" db:"highPriceCity"`
}

type PriceCountDaerahResponse struct {
	Stable int16 `json:"stable"`
	Low    int16 `json:"low"`
	High   int16 `json:"high"`
}

// DIBANDINGKAN SAMA SULSEL
type PriceTableCompareProvinceSummaryDaerah struct {
	Higher int16 `json:"higher"`
	Same   int16 `json:"same"`
	Lower  int16 `json:"lower"`
}

type PriceTableCompareProvinceMapData []struct {
	Id
	ClientId          int32        `json:"client_id"`
	City              CityResponse `json:"city" db:"city"`
	Price             *int32       `json:"price" db:"price"`
	PriceRupiahFormat *string      `json:"priceRupiahFormat" db:"priceRupiahFormat"`
	Tier              *string      `json:"tier" db:"tier"`
}

type PriceCompareListResponse []struct {
	Id
	ClientId            int32        `json:"client_id"`
	City                CityResponse `json:"city" db:"city"`
	Price               *int32       `json:"price" db:"price"`
	PriceRupiahFormat   *string      `json:"priceRupiahFormat" db:"priceRupiahFormat"`
	PriceDiff           *float32     `json:"priceDiff" db:"priceDiff"`
	PriceDiffFormat     *string      `json:"priceDiffFormat" db:"priceDiffFormat"`
	PriceDiffPercentage *float32     `json:"priceDiffPercentage" db:"priceDiffPercentage"`
}

type PriceTableCompareProvinceListData []struct {
	//Id
	//ClientId          int32        `json:"client_id"`
	Date              string       `json:"date"`
	City              CityResponse `json:"city" db:"city"`
	Price             *int32       `json:"price" db:"price"`
	PriceRupiahFormat *string      `json:"priceRupiahFormat" db:"priceRupiahFormat"`
	PriceDiff         *int32       `json:"priceDiff" db:"priceDiff"`
	PriceDiffFormat   *string      `json:"priceDiffFormat" db:"priceDiffFormat"`
}

type PriceTableCompareProvinceMap struct {
	Summary       *PriceTableCompareProvinceSummaryDaerah `json:"summary" db:"summary"`
	Commodity     *CommodityPriceTableResponse            `json:"commodity" db:"commodity"`
	PriceProvince *PriceProvinceResponse                  `json:"provincePrice"`
	PriceLevel    *PriceTableCompareProvinceMapData       `json:"priceLevel"`
	PriceTier     PriceTierCompare                        `json:"priceTier"`
	PriceTierCode []string                                `json:"priceTierCode"`
}

type PriceTableCompareProvinceList struct {
	Commodity     *CommodityPriceTableResponse       `json:"commodity" db:"commodity"`
	PriceProvince *PriceProvinceResponse             `json:"provincePrice"`
	PriceCity     *PriceTableCompareProvinceListData `json:"priceCity"`
}

type PriceTableCompareProvinceCityHistory struct {
	Commodity *CommodityPriceTableResponse         `json:"commodity" db:"commodity"`
	City      *CityLevelHargaResponse              `json:"city" db:"city"`
	PriceDiff *PriceDiffCompareProvinceCityHistory `json:"priceDiff"`
}

type PriceTableCompareProvinceCommodityHistory struct {
	Commodity *CommodityPriceTableResponse              `json:"commodity" db:"commodity"`
	PriceDiff *PriceDiffCompareProvinceCommodityHistory `json:"priceDiff"`
}

type PriceDiffCompareProvinceCityHistory []struct {
	Date              string  `json:"date"`
	Price             *int32  `json:"price" db:"price"`
	PriceRupiahFormat *string `json:"priceRupiahFormat" db:"priceRupiahFormat"`
	PriceDiff         *int32  `json:"priceDiff" db:"priceDiff"`
	PriceDiffFormat   *string `json:"priceDiffFormat" db:"priceDiffFormat"`
}

type PriceDiffCompareProvinceCommodityHistory []struct {
	City              *CityLevelHargaResponse `json:"city" db:"city"`
	Price             *int32                  `json:"price" db:"price"`
	PriceRupiahFormat *string                 `json:"priceRupiahFormat" db:"priceRupiahFormat"`
	PriceDiff         *int32                  `json:"priceDiff" db:"priceDiff"`
	PriceDiffFormat   *string                 `json:"priceDiffFormat" db:"priceDiffFormat"`
}

type PriceTableCompareNationalMap struct {
	Summary       *PriceTableCompareProvinceSummaryDaerah `json:"summary" db:"summary"`
	Commodity     *CommodityPriceTableResponse            `json:"commodity" db:"commodity"`
	PriceNational *PriceNationalResponse                  `json:"nationalPrice"`
	PriceProvince *PriceProvinceResponse                  `json:"provincePrice"`
	PriceLevel    *PriceTableCompareProvinceMapData       `json:"priceLevel"`
	PriceTier     PriceTierCompare                        `json:"priceTier"`
	PriceTierCode []string                                `json:"priceTierCode"`
}

type PriceTableCompareNationalList struct {
	Commodity     *CommodityPriceTableResponse       `json:"commodity" db:"commodity"`
	PriceNational *PriceNationalResponse             `json:"nationalPrice"`
	PriceProvince *PriceProvinceResponse             `json:"provincePrice"`
	CityPrice     *PriceTableCompareProvinceListData `json:"cityPrice"`
}

type PriceProvinceResponse struct {
	Id
	ClientId          int32               `json:"clientId" db:"clientId" json:"clientId,omitempty"`
	Province          *TmProvinceResponse `json:"province"`
	Price             *int32              `json:"price" db:"price"`
	PriceRupiahFormat string              `json:"priceRupiahFormat" db:"priceRupiahFormat"`
	PriceDiff         *int                `json:"priceDiff"`
	PriceDiffFormat   *string             `json:"priceDiffFormat"`
	Tier              string              `json:"tier"`
}

type PriceCompareResponse struct {
	Id
	ClientId            int32               `json:"clientId" db:"clientId" json:"clientId,omitempty"`
	Province            *TmProvinceResponse `json:"province"`
	Price               *int32              `json:"price" db:"price"`
	PriceRupiahFormat   *string             `json:"priceRupiahFormat" db:"priceRupiahFormat"`
	PriceDiffPercentage *float32            `json:"priceDiffPercentage" db:"priceDiffPercentage"`
}

type PriceNationalResponse struct {
	Id
	ClientId          int32               `json:"clientId" db:"clientId" json:"clientId,omitempty"`
	National          *TmNationalResponse `json:"country"`
	Price             *int32              `json:"price" db:"price"`
	PriceRupiahFormat string              `json:"priceRupiahFormat" db:"priceRupiahFormat"`
}

type PriceMtmResponse struct {
	Summary       *PriceTableCompareProvinceSummaryDaerah `json:"summary" db:"summary"`
	Commodity     *CommodityPriceTableResponse            `json:"commodity" db:"commodity"`
	PriceProvince *PriceProvinceResponse                  `json:"provincePrice"`
	PriceLevel    *PriceTableCompareProvinceMapData       `json:"priceLevel"`
	PriceTier     PriceTierCompare                        `json:"priceTier"`
	PriceTierCode []string                                `json:"priceTierCode"`
}

type PriceMtmListResponse struct {
	Commodity     *CommodityPriceTableResponse `json:"commodity" db:"commodity"`
	PriceProvince *PriceCompareResponse        `json:"provincePrice"`
	PriceCity     *PriceCompareListResponse    `json:"priceCity"`
}

type InflationMtMHCityistoryResponse []struct {
	Date                  string   `json:"date"`
	Inflation             *float32 `json:"inflation" db:"inflation"`
	InflationRupiahFormat *string  `json:"inflationRupiahFormat" db:"inflationRupiahFormat"`
}

type InflationMtMCommodityHistoryResponse []struct {
	City                  *CityLevelHargaResponse `json:"city" db:"city"`
	Inflation             *float32                `json:"inflation" db:"inflation"`
	InflationRupiahFormat *string                 `json:"inflationRupiahFormat" db:"inflationRupiahFormat"`
}

type PriceMtmCityHistoryResponse struct {
	Commodity           *CommodityPriceTableResponse     `json:"commodity" db:"commodity"`
	City                *CityLevelHargaResponse          `json:"city" db:"city"`
	CommodityInflations *InflationMtMHCityistoryResponse `json:"commodityInflations"`
}

type PriceMtmProvinceHistoryResponse struct {
	Commodity           *CommodityPriceTableResponse     `json:"commodity" db:"commodity"`
	Province            *TmProvinceResponse              `json:"province"`
	CommodityInflations *InflationMtMHCityistoryResponse `json:"commodityInflations"`
}

type PriceMtmCommodityHistoryResponse struct {
	Commodity      *CommodityPriceTableResponse          `json:"commodity" db:"commodity"`
	CityInflations *InflationMtMCommodityHistoryResponse `json:"cityInflations" db:"cityInflations"`
}
