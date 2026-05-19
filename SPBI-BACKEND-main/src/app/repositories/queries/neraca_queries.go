package queries

const (
	NeracaTxCommodityStockCityInsert = `
			INSERT INTO tx_commodity_stock(
				client_id,city_id,city_name,commodity_id,commodity_name,ketersediaan,kebutuhan,neraca,last_update,created_at
			) VALUES
		`

	NeracaTxCommodityStockProvinceInsert = `
			INSERT INTO tx_commodity_stock_province(
				client_id,province_id,province_name,commodity_id,commodity_name,ketersediaan,kebutuhan,neraca,last_update,created_at
			) VALUES
		`

	NeracaTxCommodityStockNationalInsert = `
			INSERT INTO tx_commodity_stock_national(
				client_id,national_id,national_name,commodity_id,commodity_name,ketersediaan,kebutuhan,neraca,last_update,created_at
			) VALUES
		`
)
