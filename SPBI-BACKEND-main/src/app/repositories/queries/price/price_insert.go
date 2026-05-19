package price

const (
	PriceTxCommodityCityInsert = `
			INSERT INTO tx_commodity_price(
				client_id,city_id,commodity_id,commodity_name,price,last_update,created_at
			) VALUES
	`

	PriceTxCommodityProvinceInsert = `
			INSERT INTO tx_commodity_price_province(
				client_id,province_id,commodity_id,commodity_name,price,last_update,created_at
			) VALUES
		`

	PriceTxCommodityNationalInsert = `
			INSERT INTO tx_commodity_price_national(
				client_id,national_id,commodity_id,commodity_name,price,last_update,created_at
			) VALUES
		`
)
