package models

type ReportPriceModel struct {
	Title  string `json:"title"`
	Prices []map[string]string
}

type ReportNeracaModel struct {
	Title  string `json:"title"`
	Stocks map[string][]map[string]int
}

type StocksReport struct {
	Ketersediaan []map[string]string
	Kebutuhan    []map[string]string
	Neraca       []map[string]string
}
