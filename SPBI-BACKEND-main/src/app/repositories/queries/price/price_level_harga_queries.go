package price

const (
	PricePerubahanHargaMap = `
		SELECT
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'id', a.commodity_price_id,
					'clientId', a.client_id,
					'city', (
						SELECT
							JSON_BUILD_OBJECT(
									'id', b.id,
									'name', b.name
								)
						FROM tm_city b
						WHERE id = a.id
					),
					'commodity', (
						SELECT
							JSON_BUILD_OBJECT(
									'id', c.id,
									'name', c.name,
									'parent_id', c.parent_id,
									'created_at', c.created_at,
									'updated_at', c.updated_at,
									'deleted_at', c.deleted_at
								)
						FROM tm_commodity c
						WHERE id = a.commodity_id
					),
					'price', a.price,
					'priceRupiahFormat', rupiah_format(a.price)
				)
			) "priceLevel",
			(
				SELECT get_level_harga_min(@provinceId, @commodityId, @selectedDate)
			) "priceMin",
			(
				SELECT get_level_harga_max(@provinceId, @commodityId, @selectedDate)
			) "priceMax",
			(
				SELECT get_level_harga_price_range(@provinceId, @commodityId, @selectedDate)
			) "priceDiff",
			(
				(
					(
						SELECT get_level_harga_max(@provinceId, @commodityId, @selectedDate)
					) -
					(
						SELECT get_level_harga_min(@provinceId, @commodityId, @selectedDate)
					)
				) / 5
			) "priceBarCategory"
		FROM get_level_harga(@provinceId, @commodityId, @selectedDate) a
	`

	PricePerubahanHargaMapNew = `
		SELECT
		    (
		        SELECT
					JSON_BUILD_OBJECT(
						'id', commodity_price_province_id,
						'clientId', 1,
						'province', (
							SELECT * FROM province_object(province_id)
						),
						'price', price,
						'priceRupiahFormat', (
							SELECT rupiah_format(price::int)
						)
					)
				FROM price_province_cr(@commodityId, @provinceId, @selectedDate)
			) "provincePrice",
		    (
				 SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_price_id,
							'clientId', a.client_id,
							'city', (
								SELECT
									JSON_BUILD_OBJECT(
											'id', b.id,
											'name', b.name
										)
								FROM tm_city b
								WHERE id = a.id
							),
							'commodity', (
								SELECT
									JSON_BUILD_OBJECT(
											'id', c.id,
											'name', c.name,
											'parent_id', c.parent_id,
											'created_at', c.created_at,
											'updated_at', c.updated_at,
											'deleted_at', c.deleted_at
										)
								FROM tm_commodity c
								WHERE id = @commodityId
							),
							'price', a.price,
							'priceRupiahFormat', rupiah_format(a.price)
						)
					) "priceLevel"
				FROM get_level_harga(@provinceId, @commodityId, @selectedDate) a
			) "cityPrice",
		    (
				SELECT get_level_harga_min(@provinceId, @commodityId, @selectedDate)
			) "priceMin",
			(
				SELECT get_level_harga_max(@provinceId, @commodityId, @selectedDate)
			) "priceMax",
			(
				SELECT get_level_harga_price_range(@provinceId, @commodityId, @selectedDate)
			) "priceDiff",
			(
				(
					(
						SELECT get_level_harga_max(@provinceId, @commodityId, @selectedDate)
					) -
					(
						SELECT get_level_harga_min(@provinceId, @commodityId, @selectedDate)
					)
				) / 5
			) "priceBarCategory"
	`

	PricePerubahanHargaMapNewAvgChild = `
		SELECT
		    (
		        SELECT
					JSON_BUILD_OBJECT(
						'id', null,
						'clientId', 1,
						'province', (
							SELECT * FROM province_object(province_id)
						),
						'price', price,
						'priceRupiahFormat', (
							SELECT rupiah_format(price::int)
						)
					)
				FROM price_province_avg_child(@commodityId, @provinceId, @selectedDate)
			) "provincePrice",
		    (
				 SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', null,
							'clientId', a.client_id,
							'city', (
								SELECT
									JSON_BUILD_OBJECT(
											'id', b.id,
											'name', b.name
										)
								FROM tm_city b
								WHERE id = a.id
							),
							'commodity', (
								SELECT
									JSON_BUILD_OBJECT(
											'id', c.id,
											'name', c.name,
											'parent_id', c.parent_id,
											'created_at', c.created_at,
											'updated_at', c.updated_at,
											'deleted_at', c.deleted_at
										)
								FROM tm_commodity c
								WHERE id = @commodityId
							),
							'price', a.price,
							'priceRupiahFormat', rupiah_format(a.price)
						)
					) "priceLevel"
				FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate) a
			) "cityPrice",
		    (
				SELECT get_level_harga_min_avg_child(@provinceId, @commodityId, @selectedDate)
			) "priceMin",
			(
				SELECT get_level_harga_max_avg_child(@provinceId, @commodityId, @selectedDate)
			) "priceMax",
			(
				SELECT get_level_harga_price_range_avg_child(@provinceId, @commodityId, @selectedDate)
			) "priceDiff",
			(
				(
					(
						SELECT get_level_harga_max_avg_child(@provinceId, @commodityId, @selectedDate)
					) -
					(
						SELECT get_level_harga_min_avg_child(@provinceId, @commodityId, @selectedDate)
					)
				) / 5
			) "priceBarCategory"
	`

	PriceLevelHargaAvgChild = `
		SELECT
		    (
		        SELECT
					JSON_BUILD_OBJECT(
						'id', null,
						'clientId', 1,
						'province', (
							SELECT * FROM province_object(province_id)
						),
						'price', price,
						'priceRupiahFormat', (
							SELECT rupiah_format(price::int)
						)
					)
				FROM price_province_avg_child(@commodityId, @provinceId, @selectedDate)
			) "provincePrice",
		    (
				 SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_price_id,
							'clientId', a.client_id,
							'city', (
								SELECT
									JSON_BUILD_OBJECT(
											'id', b.id,
											'name', b.name
										)
								FROM tm_city b
								WHERE id = a.id
							),
							'commodity', (
								SELECT
									JSON_BUILD_OBJECT(
											'id', c.id,
											'name', c.name,
											'parent_id', c.parent_id,
											'created_at', c.created_at,
											'updated_at', c.updated_at,
											'deleted_at', c.deleted_at
										)
								FROM tm_commodity c
								WHERE id = @commodityId
							),
							'price', a.price,
							'priceRupiahFormat', rupiah_format(a.price)
						)
					) "priceLevel"
				FROM get_level_harga(@provinceId, @commodityId, @selectedDate) a
			) "cityPrice",
		    (
				SELECT get_level_harga_min(@provinceId, @commodityId, @selectedDate)
			) "priceMin",
			(
				SELECT get_level_harga_max(@provinceId, @commodityId, @selectedDate)
			) "priceMax",
			(
				SELECT get_level_harga_price_range(@provinceId, @commodityId, @selectedDate)
			) "priceDiff",
			(
				(
					(
						SELECT get_level_harga_max(@provinceId, @commodityId, @selectedDate)
					) -
					(
						SELECT get_level_harga_min(@provinceId, @commodityId, @selectedDate)
					)
				) / 5
			) "priceBarCategory"
	`

	PriceLevelHargaList = `
		SELECT
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', province_id,
						'clientId', client_id,
						'province', (
							SELECT * FROM province_object(@provinceId)
						),
						'price', c.price,
						'priceRupiahFormat', (
							SELECT rupiah_format(c.price)
						)
					)
				FROM get_level_harga_province(@commodityId, @selectedDate) c
				WHERE province_id = @provinceId
			) "priceProvince",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_price_id,
							'client_id', a.client_id,
							'city', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', b.id,
										'name', b.name,
										'assets', (
											SELECT 
												JSON_BUILD_OBJECT(
													'id', id,
													'assetsType', assets_type,
													'assetsLocation', assets_location,
													'assetsLocationType', assets_location_type,
													'assetsMediaType', assets_media_type,
													'assetsExt', assets_ext,
													'assetsName', assets_name,
													'assetsUrl', CONCAT(
														(
															SELECT
																name
															FROM settings s 
															WHERE s.parent_id = (
																SELECT
																	ID
																FROM SETTINGS S
																WHERE s."name" = 'BASE_URL'
															)
														),
														'/assets?assets_location=',
														CONCAT(assets_location, '/',assets_name)
													)
												)
											FROM assets
											WHERE id = b.assets_relation_id
										),
										'created_at', b.created_at,
										'updated_at', b.updated_at,
										'deleted_at', b.deleted_at
									)
								FROM tm_city b
								WHERE id = a.id
							),
							'commodity', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', c.id,
										'name', c.name,
										'parent_id', c.parent_id,
										'created_at', c.created_at,
										'updated_at', c.updated_at,
										'deleted_at', c.deleted_at
									)
								FROM tm_commodity c
								WHERE id = a.commodity_id
							),
							'price', a.price,
							'priceRupiahFormat', rupiah_format(a.price),
							'priceTier', (
								SELECT '{"sangat-rendah":{"title":"Sangat Rendah","color":"#208245"},"rendah":{"title":"Rendah","color":"#27AE65"},"sedang":{"title":"Sedang","color":"#E4B701"},"tinggi":{"title":"Tinggi","color":"#C0392B"},"sangat-tinggi":{"title":"Sangat Tinggi","color":"#7C2019"}}'::JSON
							),
							'priceTierCode',
							(
								SELECT '["sangat-rendah","rendah","sedang","tinggi","sangat-tinggi"]'::JSON
							)
						)
					)
				FROM get_level_harga(@provinceId, @commodityId, @selectedDate) a
				OFFSET @page
				LIMIT @limit
			) "priceCity"
	`

	PriceLevelHargaListAvgChild = `
		SELECT
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', null,
						'clientId', 1,
						'province', (
							SELECT * FROM province_object(province_id)
						),
						'price', price,
						'priceRupiahFormat', (
							SELECT rupiah_format(price::int)
						)
					)
				FROM price_province_avg_child(@commodityId, @provinceId, @selectedDate)
			) "priceProvince",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', null,
							'client_id', a.client_id,
							'city', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', b.id,
										'name', b.name,
										'assets', (
											SELECT 
												JSON_BUILD_OBJECT(
													'id', id,
													'assetsType', assets_type,
													'assetsLocation', assets_location,
													'assetsLocationType', assets_location_type,
													'assetsMediaType', assets_media_type,
													'assetsExt', assets_ext,
													'assetsUrl', CONCAT(
														(
															SELECT
																name
															FROM settings s 
															WHERE s.parent_id = (
																SELECT
																	ID
																FROM SETTINGS S
																WHERE s."name" = 'BASE_URL'
															)
														),
														'/assets?assets_location=',
														CONCAT(assets_location, '/',assets_name)
													)
												)
											FROM assets
											WHERE id = b.assets_relation_id
										),
										'created_at', b.created_at,
										'updated_at', b.updated_at,
										'deleted_at', b.deleted_at
									)
								FROM tm_city b
								WHERE id = a.id
							),
							'commodity', null,
							'price', a.price,
							'priceRupiahFormat', rupiah_format(a.price),
							'priceTier', (
								SELECT '{"sangat-rendah":{"title":"Sangat Rendah","color":"#208245"},"rendah":{"title":"Rendah","color":"#27AE65"},"sedang":{"title":"Sedang","color":"#E4B701"},"tinggi":{"title":"Tinggi","color":"#C0392B"},"sangat-tinggi":{"title":"Sangat Tinggi","color":"#7C2019"}}'::JSON
							),
							'priceTierCode',
							(
								SELECT '["sangat-rendah","rendah","sedang","tinggi","sangat-tinggi"]'::JSON
							)
						)
					)
				FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate) a
				OFFSET @page
				LIMIT @limit
			) "priceCity"
	`

	PriceLevelHargaListWithOrder = `
		SELECT
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', province_id,
						'clientId', client_id,
						'province', (
							SELECT * FROM province_object(@provinceId)
						),
						'price', c.price,
						'priceRupiahFormat', (
							SELECT rupiah_format(c.price)
						)
					)
				FROM get_level_harga_province(@commodityId, @selectedDate) c
				WHERE province_id = @provinceId
			) "priceProvince",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_price_id,
							'client_id', a.client_id,
							'city', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', b.id,
										'name', b.name,
										'assets', (
											SELECT 
												JSON_BUILD_OBJECT(
													'id', id,
													'assetsType', assets_type,
													'assetsLocation', assets_location,
													'assetsLocationType', assets_location_type,
													'assetsMediaType', assets_media_type,
													'assetsExt', assets_ext,
													'assetsUrl', CONCAT(
														(
															SELECT
																name
															FROM settings s 
															WHERE s.parent_id = (
																SELECT
																	ID
																FROM SETTINGS S
																WHERE s."name" = 'BASE_URL'
															)
														),
														'/assets?assets_location=',
														CONCAT(assets_location, '/',assets_name)
													)
												)
											FROM assets
											WHERE id = b.assets_relation_id
										),
										'created_at', b.created_at,
										'updated_at', b.updated_at,
										'deleted_at', b.deleted_at
									)
								FROM tm_city b
								WHERE id = a.id
							),
							'commodity', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', c.id,
										'name', c.name,
										'parent_id', c.parent_id,
										'created_at', c.created_at,
										'updated_at', c.updated_at,
										'deleted_at', c.deleted_at
									)
								FROM tm_commodity c
								WHERE id = a.commodity_id
							),
							'price', a.price,
							'priceRupiahFormat', rupiah_format(a.price),
							'priceTier', (
								SELECT '{"sangat-rendah":{"title":"Sangat Rendah","color":"#208245"},"rendah":{"title":"Rendah","color":"#27AE65"},"sedang":{"title":"Sedang","color":"#FFFF00"},"tinggi":{"title":"Tinggi","color":"#C0392B"},"sangat-tinggi":{"title":"Sangat Tinggi","color":"#7C2019"}}'::JSON
							),
							'priceTierCode',
							(
								SELECT '["sangat-rendah","rendah","sedang","tinggi","sangat-tinggi"]'::JSON
							)
						)
					)
				FROM (
				    SELECT *
				    FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
					ORDER BY price
					OFFSET @page
					LIMIT @limit
				) a
			) "priceCity"
	`

	PriceLevelHargaListWithOrderAvgChild = `
		SELECT
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', null,
						'clientId', 1,
						'province', (
							SELECT * FROM province_object(province_id)
						),
						'price', price,
						'priceRupiahFormat', (
							SELECT rupiah_format(price::int)
						)
					)
				FROM price_province_avg_child(@commodityId, @provinceId, @selectedDate)
			) "priceProvince",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', null,
							'client_id', a.client_id,
							'city', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', b.id,
										'name', b.name,
										'assets', (
											SELECT 
												JSON_BUILD_OBJECT(
													'id', id,
													'assetsType', assets_type,
													'assetsLocation', assets_location,
													'assetsLocationType', assets_location_type,
													'assetsMediaType', assets_media_type,
													'assetsExt', assets_ext,
													'assetsUrl', CONCAT(
														(
															SELECT
																name
															FROM settings s 
															WHERE s.parent_id = (
																SELECT
																	ID
																FROM SETTINGS S
																WHERE s."name" = 'BASE_URL'
															)
														),
														'/assets?assets_location=',
														CONCAT(assets_location, '/',assets_name)
													)
												)
											FROM assets
											WHERE id = b.assets_relation_id
										),
										'created_at', b.created_at,
										'updated_at', b.updated_at,
										'deleted_at', b.deleted_at
									)
								FROM tm_city b
								WHERE id = a.id
							),
							'commodity', null,
							'price', a.price,
							'priceRupiahFormat', rupiah_format(a.price),
							'priceTier', (
								SELECT '{"sangat-rendah":{"title":"Sangat Rendah","color":"#208245"},"rendah":{"title":"Rendah","color":"#27AE65"},"sedang":{"title":"Sedang","color":"#FFFF00"},"tinggi":{"title":"Tinggi","color":"#C0392B"},"sangat-tinggi":{"title":"Sangat Tinggi","color":"#7C2019"}}'::JSON
							),
							'priceTierCode',
							(
								SELECT '["sangat-rendah","rendah","sedang","tinggi","sangat-tinggi"]'::JSON
							)
						)
					)
				FROM (
				    SELECT *
				    FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
					ORDER BY price
					OFFSET @page
					LIMIT @limit
				) a
			) "priceCity"
	`

	PriceLevelHargaListWithOrderDescAvgChild = `
		SELECT
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', null,
						'clientId', 1,
						'province', (
							SELECT * FROM province_object(province_id)
						),
						'price', price,
						'priceRupiahFormat', (
							SELECT rupiah_format(price::int)
						)
					)
				FROM price_province_avg_child(@commodityId, @provinceId, @selectedDate)
			) "priceProvince",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', null,
							'client_id', a.client_id,
							'city', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', b.id,
										'name', b.name,
										'assets', (
											SELECT 
												JSON_BUILD_OBJECT(
													'id', id,
													'assetsType', assets_type,
													'assetsLocation', assets_location,
													'assetsLocationType', assets_location_type,
													'assetsMediaType', assets_media_type,
													'assetsExt', assets_ext,
													'assetsUrl', CONCAT(
														(
															SELECT
																name
															FROM settings s 
															WHERE s.parent_id = (
																SELECT
																	ID
																FROM SETTINGS S
																WHERE s."name" = 'BASE_URL'
															)
														),
														'/assets?assets_location=',
														CONCAT(assets_location, '/',assets_name)
													)
												)
											FROM assets
											WHERE id = b.assets_relation_id
										),
										'created_at', b.created_at,
										'updated_at', b.updated_at,
										'deleted_at', b.deleted_at
									)
								FROM tm_city b
								WHERE id = a.id
							),
							'commodity', null,
							'price', a.price,
							'priceRupiahFormat', rupiah_format(a.price),
							'priceTier', (
								SELECT '{"sangat-rendah":{"title":"Sangat Rendah","color":"#208245"},"rendah":{"title":"Rendah","color":"#27AE65"},"sedang":{"title":"Sedang","color":"#FFFF00"},"tinggi":{"title":"Tinggi","color":"#C0392B"},"sangat-tinggi":{"title":"Sangat Tinggi","color":"#7C2019"}}'::JSON
							),
							'priceTierCode',
							(
								SELECT '["sangat-rendah","rendah","sedang","tinggi","sangat-tinggi"]'::JSON
							)
						)
					)
				FROM (
				    SELECT *
				    FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
					ORDER BY price DESC
					OFFSET @page
					LIMIT @limit
				) a
			) "priceCity"
	`

	PriceLevelHargaListWithOrderDesc = `
		SELECT
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', province_id,
						'clientId', client_id,
						'province', (
							SELECT * FROM province_object(@provinceId)
						),
						'price', c.price,
						'priceRupiahFormat', (
							SELECT rupiah_format(c.price)
						)
					)
				FROM get_level_harga_province(@commodityId, @selectedDate) c
				WHERE province_id = @provinceId
			) "priceProvince",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_price_id,
							'client_id', a.client_id,
							'city', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', b.id,
										'name', b.name,
										'assets', (
											SELECT 
												JSON_BUILD_OBJECT(
													'id', id,
													'assetsType', assets_type,
													'assetsLocation', assets_location,
													'assetsLocationType', assets_location_type,
													'assetsMediaType', assets_media_type,
													'assetsExt', assets_ext,
													'assetsUrl', CONCAT(
														(
															SELECT
																name
															FROM settings s 
															WHERE s.parent_id = (
																SELECT
																	ID
																FROM SETTINGS S
																WHERE s."name" = 'BASE_URL'
															)
														),
														'/assets?assets_location=',
														CONCAT(assets_location, '/',assets_name)
													)
												)
											FROM assets
											WHERE id = b.assets_relation_id
										),
										'created_at', b.created_at,
										'updated_at', b.updated_at,
										'deleted_at', b.deleted_at
									)
								FROM tm_city b
								WHERE id = a.id
							),
							'commodity', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', c.id,
										'name', c.name,
										'parent_id', c.parent_id,
										'created_at', c.created_at,
										'updated_at', c.updated_at,
										'deleted_at', c.deleted_at
									)
								FROM tm_commodity c
								WHERE id = a.commodity_id
							),
							'price', a.price,
							'priceRupiahFormat', rupiah_format(a.price),
							'priceTier', (
								SELECT '{"sangat-rendah":{"title":"Sangat Rendah","color":"#208245"},"rendah":{"title":"Rendah","color":"#27AE65"},"sedang":{"title":"Sedang","color":"#FFFF00"},"tinggi":{"title":"Tinggi","color":"#C0392B"},"sangat-tinggi":{"title":"Sangat Tinggi","color":"#7C2019"}}'::JSON
							),
							'priceTierCode',
							(
								SELECT '["sangat-rendah","rendah","sedang","tinggi","sangat-tinggi"]'::JSON
							)
						)
					)
				FROM (
				    SELECT *
				    FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
					ORDER BY price desc
					OFFSET @page
					LIMIT @limit
				) a
			) "priceCity"
	`

	PriceDiffLevelHarga = `
		SELECT
			(
				SELECT get_level_harga_min(@provinceId, @commodityId, @selectedDate)
			) "priceMin",
			(
				SELECT get_level_harga_max(@provinceId, @commodityId, @selectedDate)
			) "priceMax",
			(
				SELECT get_level_harga_price_range(@provinceId, @commodityId, @selectedDate)
			) "priceDiff",
			(
				(
					(
						SELECT get_level_harga_max(@provinceId, @commodityId, @selectedDate)
					) -
					(
						SELECT get_level_harga_min(@provinceId, @commodityId, @selectedDate)
					)
				) / 5
			) "priceBarCategory"
	`

	PriceDiffLevelHargaAvgChild = `
		SELECT
			(
				SELECT get_level_harga_min_avg_child(@provinceId, @commodityId, @selectedDate)
			) "priceMin",
			(
				SELECT get_level_harga_max_avg_child(@provinceId, @commodityId, @selectedDate)
			) "priceMax",
			(
				SELECT get_level_harga_price_range_avg_child(@provinceId, @commodityId, @selectedDate)
			) "priceDiff",
			(
				(
					(
						SELECT get_level_harga_max_avg_child(@provinceId, @commodityId, @selectedDate)
					) -
					(
						SELECT get_level_harga_min_avg_child(@provinceId, @commodityId, @selectedDate)
					)
				) / 5
			) "priceBarCategory"
	`

	PriceGetLatestDateAvail = `
		SELECT
			CONCAT(
				CONCAT(
					TO_CHAR(last_update, 'YYYY'),
					CONCAT('-', REPLACE(TO_CHAR(last_update, 'MM'), ' ','') )
				),
				CONCAT('-', TO_CHAR(last_update, 'DD'))
			)
		FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
		WHERE id = @cityId
	`

	PriceGetLatestDateAvailAvgChild = `
		SELECT
			CONCAT(
				CONCAT(
					TO_CHAR(last_update, 'YYYY'),
					CONCAT('-', REPLACE(TO_CHAR(last_update, 'MM'), ' ','') )
				),
				CONCAT('-', TO_CHAR(last_update, 'DD'))
			)
		FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
		WHERE id = @cityId
	`

	PriceGetLatestDateProvinceAvail = `
		SELECT
			CONCAT(
				CONCAT(
					TO_CHAR(last_update, 'YYYY'),
					CONCAT('-', REPLACE(TO_CHAR(last_update, 'MM'), ' ','') )
				),
				CONCAT('-', TO_CHAR(last_update, 'DD'))
			)
		FROM get_level_harga_province(@commodityId, @selectedDate)
		WHERE province_id = @provinceId
	`

	PriceGetLatestDateProvinceAvailAvgChild = `
		SELECT
			CONCAT(
				CONCAT(
					TO_CHAR(last_update, 'YYYY'),
					CONCAT('-', REPLACE(TO_CHAR(last_update, 'MM'), ' ','') )
				),
				CONCAT('-', TO_CHAR(last_update, 'DD'))
			)
-- 		FROM get_level_harga_province(@commodityId, @selectedDate)
-- 		WHERE province_id = @provinceId
		FROM price_province_avg_child(@commodityId, @provinceId, @selectedDate)
	`

	PriceGetLatestDateNationalAvail = `
		SELECT
			CONCAT(
				CONCAT(
					TO_CHAR(last_update, 'YYYY'),
					CONCAT('-', REPLACE(TO_CHAR(last_update, 'MM'), ' ','') )
				),
				CONCAT('-', TO_CHAR(last_update, 'DD'))
			)
		FROM get_level_harga_national(@commodityId, @selectedDate)
		WHERE national_id = @nationalId
	`

	PriceGetLatestDateNationalAvailAvgChild = `
		SELECT
			CONCAT(
				CONCAT(
					TO_CHAR(last_update, 'YYYY'),
					CONCAT('-', REPLACE(TO_CHAR(last_update, 'MM'), ' ','') )
				),
				CONCAT('-', TO_CHAR(last_update, 'DD'))
			)
		FROM price_national_avg_child(@commodityId, @nationalId, @selectedDate)
	`

	PriceExists = `
		SELECT EXISTS(
			SELECT *
-- 			FROM get_level_harga(@provinceId, @commodityId, @selectedDate) a
			FROM price_city(@commodityId, @cityId, @startDate, @endDate)
		)
	`

	HistoryQuery = `SELECT * FROM tx_file_upload_history WHERE module_type = @module and status = 1 and file_name ilike @search ORDER BY id desc
					OFFSET @page
					LIMIT @limit`

	HistoryCountQuery = `SELECT count(1) FROM tx_file_upload_history WHERE module_type = @module and file_name ilike @search`
)
