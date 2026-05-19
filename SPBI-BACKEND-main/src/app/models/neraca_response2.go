package models

type NeracaCountDaerahResponse struct {
	Aman    int16 `json:"aman"`
	Waspada int16 `json:"waspada"`
	Rentan  int16 `json:"rentan"`
}

type NeracaTableResponse struct {
	StockCommodity *StockCommodityByLocationResponse `json:"stockCommodity" db:"stockCommodity"`
	StockDiff      *StockLocationDiffResponse        `json:"stockDiff" db:"stockDiff"`
}

type StockCommodityByLocationResponse []struct {
	City         *string  `json:"city" db:"city"`
	Ketersediaan *int32   `json:"ketersediaan" db:"ketersediaan"`
	Kebutuhan    *float32 `json:"kebutuhan" db:"kebutuhan"`
	Neraca       *int32   `json:"neraca" db:"neraca"`
}

type StockCommodityResponse []struct {
	CommodityName *string  `json:"commodity_name" db:"commodity_name"`
	Ketersediaan  *int32   `json:"ketersediaan" db:"ketersediaan"`
	Kebutuhan     *float32 `json:"kebutuhan" db:"kebutuhan"`
	Neraca        *int32   `json:"neraca" db:"neraca"`
}

type StockLocationDiffResponse struct {
	AmanStockCity    *int32 `json:"amanStockCity" db:"amanStockCity"`
	WaspadaStockCity *int32 `json:"waspadaStockCity" db:"waspadaStockCity"`
	RentanStockCity  *int32 `json:"rentanStockCity" db:"rentanStockCity"`
}
