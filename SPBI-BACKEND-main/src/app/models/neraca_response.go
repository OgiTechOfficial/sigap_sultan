package models

type NeracaStokAkhirSummary struct {
	Defisit int16 `json:"defisit"`
	Rentan  int16 `json:"rentan"`
	Waspada int16 `json:"waspada"`
	Aman    int16 `json:"aman"`
}

type NeracaKetersediaanKebutuhanSummary struct {
	Menurun   int16 `json:"menurun"`
	Stabil    int16 `json:"stabil"`
	Meningkat int16 `json:"meningkat"`
}

type NeracaStokAkhirMapResponse struct {
	Unit            string                       `json:"unit"`
	Commodity       *CommodityResponse           `json:"commodity" db:"commodity"`
	Summary         *NeracaStokAkhirSummary      `json:"summary"`
	ProvinceStock   *NeracaProvinceStockResponse `json:"provinceStock"`
	CityStock       *[]NeracaCityStockResponse   `json:"cityStock"`
	NeracaTierLevel *NeracaTierLevel             `json:"stockTier"`
	StockTierCode   []string                     `json:"stockTierCode"`
}

type NeracaStokAkhirListResponse struct {
	Unit            string                       `json:"unit"`
	Commodities     *CommodityResponse           `json:"commodity"`
	ProvinceStock   *NeracaProvinceStockResponse `json:"provinceStock"`
	CityStock       *[]NeracaCityStockResponse   `json:"cityStock"`
	NeracaTierLevel *NeracaTierLevel             `json:"stockTier"`
}

type NeracaStokAkhirCommodityHistoryResponse struct {
	Unit            string                     `json:"unit"`
	Commodities     *CommodityResponse         `json:"commodity"`
	CityStock       *[]NeracaCityStockResponse `json:"cityStock"`
	NeracaTierLevel *NeracaTierLevel           `json:"stockTier"`
}

type NeracaStokAkhirCityHistoryResponse struct {
	Unit        string             `json:"unit"`
	Commodities *CommodityResponse `json:"commodity"`
	City        *CityResponse      `json:"city"`
	//ProvinceStock   *NeracaProvinceStockResponse `json:"provinceStock"`
	StockDiff *float32 `json:"stockDiff"`
	//CityStock       *NeracaCityStockResponse `json:"cityStock"`
	NeracaListResponse *[]NeracaListResponse `json:"stock" db:"stock"`
	NeracaTierLevel    *NeracaTierLevel      `json:"stockTier"`
}

type NeracaKetersediaanMapResponse struct {
	Unit            string                              `json:"unit"`
	Commodities     *CommodityResponse                  `json:"commodity"`
	Summary         *NeracaKetersediaanKebutuhanSummary `json:"summary"`
	ProvinceStock   *NeracaProvinceStockResponse        `json:"provinceStock"`
	CityStock       *[]NeracaCityStockWithDiffResponse  `json:"cityStock"`
	NeracaTierLevel *NeracaTierKetersediaanLevel        `json:"stockTier"`
	StockTierCode   []string                            `json:"stockTierCode"`
}

type NeracaKetersediaanListResponse struct {
	Unit            string                              `json:"unit"`
	Commodities     *CommodityResponse                  `json:"commodity"`
	Summary         *NeracaKetersediaanKebutuhanSummary `json:"summary"`
	ProvinceStock   *NeracaProvinceStockResponse        `json:"provinceStock"`
	CityStock       *[]NeracaCityStockWithDiffResponse  `json:"cityStock"`
	NeracaTierLevel *NeracaTierKetersediaanLevel        `json:"stockTier"`
	StockTierCode   []string                            `json:"stockTierCode"`
}

type NeracaKetersediaanCommodityCityHistoryResponse struct {
	Unit        string             `json:"unit"`
	Commodities *CommodityResponse `json:"commodity"`
	//CityStock       *[]NeracaCitKetersediaanWithDiffResponse `json:"cityStock"`
	NeracaListResponse *[]NeracaListResponse `json:"stock" db:"stock"`
	NeracaTierLevel    *NeracaTierLevel      `json:"stockTier"`
}

type NeracaKetersediaanCommodityHistoryResponse struct {
	Unit            string                                   `json:"unit"`
	Commodities     *CommodityResponse                       `json:"commodity"`
	CityStock       *[]NeracaCitKetersediaanWithDiffResponse `json:"cityStock"`
	NeracaTierLevel *NeracaTierLevel                         `json:"stockTier"`
}

type NeracaKebutuhanCommodityHistoryResponse struct {
	Unit            string                                `json:"unit"`
	Commodities     *CommodityResponse                    `json:"commodity"`
	CityStock       *[]NeracaCitKebutuhanWithDiffResponse `json:"cityStock"`
	NeracaTierLevel *NeracaTierLevel                      `json:"stockTier"`
}

type NeracaStokAkhirCityCommodityResponse struct {
	Unit            string               `json:"unit"`
	UnitDiff        string               `json:"unitDiff"`
	Commodities     *CommodityResponse   `json:"commodity"`
	City            *CityResponse        `json:"city"`
	StockDiff       *float32             `json:"stockDiff"`
	Stock           *NeracaCityCommodity `json:"stock" db:"stock"`
	NeracaTierLevel *NeracaTierLevel     `json:"stockTier"`
}

type NeracaStokAkhirCityCommodityChartResponse struct {
	Unit            string               `json:"unit"`
	UnitDiff        string               `json:"unitDiff"`
	Commodities     *CommodityResponse   `json:"commodity"`
	City            *CityResponse        `json:"city"`
	StockDiff       *float32             `json:"stockDiff"`
	Stock           *NeracaCityCommodity `json:"stock" db:"stock"`
	NeracaTierLevel *NeracaTierLevel     `json:"stockTier"`
}

type NeracaCompareWithPriceCommodityHistory struct {
	Unit                        string                         `json:"unit"`
	UnitDiff                    string                         `json:"unitDiff"`
	Commodities                 *CommodityResponse             `json:"commodity"`
	CityStock                   *[]NeracaCityStockResponse     `json:"city"`
	NeracaListWithPriceResponse *[]NeracaListWithPriceResponse `json:"stock" db:"stock"`
	NeracaTierLevel             *NeracaTierLevel               `json:"stockTier"`
}

type NeracaStokAkhirByCityMapResponse struct {
	Unit            string                          `json:"unit"`
	City            *CityResponse                   `json:"city"`
	Summary         *NeracaStokAkhirSummary         `json:"summary"`
	ProvinceStock   *NeracaProvinceStockResponse    `json:"provinceStock"`
	CommodityStock  *[]NeracaCommodityStockResponse `json:"commodityStock"`
	NeracaTierLevel *NeracaTierLevel                `json:"stockTier"`
	StockTierCode   []string                        `json:"stockTierCode"`
}

type NeracaKetersediaanByCityMapResponse struct {
	Unit            string                              `json:"unit"`
	City            *CityResponse                       `json:"city"`
	Summary         *NeracaKetersediaanKebutuhanSummary `json:"summary"`
	ProvinceStock   *NeracaProvinceStockResponse        `json:"provinceStock"`
	CommodityStock  *[]NeracaCommodityStockResponse     `json:"commodityStock"`
	NeracaTierLevel *NeracaTierKetersediaanLevel        `json:"stockTier"`
	StockTierCode   []string                            `json:"stockTierCode"`
}

// NESTED
type NeracaProvinceStockResponse struct {
	Id
	ClientId int32               `json:"clientId" db:"clientId" json:"clientId,omitempty"`
	Province *TmProvinceResponse `json:"province"`
	Stock    *float32            `json:"stock" db:"stock"`
	Tier     string              `json:"tier"`
}

type NeracaCityStockResponse struct {
	Id
	ClientId int32        `json:"client_id"`
	City     CityResponse `json:"city" db:"city"`
	Stock    *float32     `json:"stock" db:"stock"`
	Tier     string       `json:"tier"`
}

type NeracaCommodityStockResponse struct {
	Id
	ClientId  int32             `json:"client_id"`
	Commodity CommodityResponse `json:"commodity" db:"commodity"`
	Stock     *float32          `json:"stock" db:"stock"`
	Tier      string            `json:"tier"`
}

type NeracaCityStockWithDiffResponse struct {
	Id
	ClientId  int32        `json:"client_id"`
	City      CityResponse `json:"city" db:"city"`
	Stock     *float32     `json:"stock" db:"stock"`
	StockDiff *float32     `json:"stockDiff" db:"stockDiff"`
	Tier      string       `json:"tier"`
}

type NeracaCitKetersediaanWithDiffResponse struct {
	Id
	ClientId         int32        `json:"client_id"`
	City             CityResponse `json:"city" db:"city"`
	Stock            *float32     `json:"ketersediaan" db:"ketersediaan"`
	KetersediaanDiff *float32     `json:"ketersediaanDiffPercentage" db:"ketersediaanDiffPercentage"`
	//StockDiff *float32     `json:"stockDiff" db:"stockDiff"`
	Tier string `json:"tier"`
}

type NeracaCitKebutuhanWithDiffResponse struct {
	Id
	ClientId  int32        `json:"client_id"`
	City      CityResponse `json:"city" db:"city"`
	Stock     *float32     `json:"kebutuhan" db:"kebutuhan"`
	StockDiff *float32     `json:"kebutuhanDiffPercentage" db:"kebutuhanDiffPercentage"`
	Tier      string       `json:"tier"`
}

type NeracaTierLevel struct {
	NeracaDefisit NeracaDefisit `json:"defisit"`
	NeracaRentan  NeracaRentan  `json:"rentan"`
	NeracaWaspada NeracaWaspada `json:"waspada"`
	NeracaAman    NeracaAman    `json:"aman"`
}

type NeracaTierKetersediaanLevel struct {
	NeracaMenurun   NeracaMenurun   `json:"menurun"`
	NeracaStabil    NeracaStabil    `json:"stabil"`
	NeracaMeningkat NeracaMeningkat `json:"meningkat"`
}

type (
	NeracaDefisit struct {
		NeracaTitleColor
	}
	NeracaRentan struct {
		NeracaTitleColor
	}
	NeracaWaspada struct {
		NeracaTitleColor
	}
	NeracaAman struct {
		NeracaTitleColor
	}

	NeracaMenurun struct {
		NeracaKetersediaanColor
	}
	NeracaStabil struct {
		NeracaKetersediaanColor
	}
	NeracaMeningkat struct {
		NeracaKetersediaanColor
	}
)

type NeracaCityCommodity []struct {
	Period       string  `json:"period"`
	Neraca       float32 `json:"neraca"`
	Ketersediaan float32 `json:"ketersediaan"`
	Produksi     float32 `json:"produksi"`
	Kebutuhan    float32 `json:"kebutuhan"`
	Tier         string  `json:"tier"`
}

type NeracaListResponse struct {
	Date              string   `json:"date"`
	Stock             *float32 `json:"stock" db:"stock"`
	StockRupiahFormat *string  `json:"stockFormat" db:"stockFormat"`
}

type NeracaListWithPriceResponse struct {
	Date              string   `json:"date"`
	Stock             *float32 `json:"stock" db:"stock"`
	StockRupiahFormat *string  `json:"stockFormat" db:"stockFormat"`
	Price             *float32 `json:"price" db:"price"`
	PriceRupiahFormat *string  `json:"priceRupiahFormat" db:"priceRupiahFormat"`
}

//type NeracaTier struct {
//	Title                string `json:"title"`
//	PriceMin             int32  `json:"priceMin" db:"priceMin"`
//	PriceMinRupiahFormat string `json:"priceMinRupiahFormat" db:"priceMinRupiahFormat"`
//	PriceMax             int32  `json:"priceMax" db:"priceMax"`
//	PriceMaxRupiahFormat string `json:"priceMaxRupiahFormat" db:"priceMaxRupiahFormat"`
//	Color                string `json:"color"`
//}
