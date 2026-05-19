package price

const (
	PriceGetYty = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
							'higher', (
								SELECT
								    COUNT(1)
								FROM price_inflation(
									a.id,
									@provinceId,
									start_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								)
								WHERE inflation > 0
							),
							'same', (
								SELECT
								    COUNT(1)
								FROM price_inflation(
									a.id,
									@provinceId,
									start_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								)
								WHERE inflation = 0
							),
							'lower', (
								SELECT
								    COUNT(1)
								FROM price_inflation(
									a.id,
									@provinceId,
									start_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								)
								WHERE inflation < 0
							)
						)
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) summary,
				(
					SELECT JSON_BUILD_OBJECT(
						'id', a.id,
						'name', a.name
					) 
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
						JSON_BUILD_OBJECT(
							'id', commodity_price_province_id,
							'clientId', client_id,
							'province', (
								SELECT * FROM province_object(c.province_id)
							),
							'price', c.price,
							'priceRupiahFormat', (
								SELECT rupiah_format(c.price)
							),
							'tier', (
								SELECT
									CASE 
										WHEN (
											SELECT price_province_inflation(
												@commodityId,
												c.province_id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )
										) > 0
										THEN 'higher'
										WHEN (
											SELECT price_province_inflation(
												@commodityId,
												c.province_id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )
										) < 0
										THEN 'lower'
										WHEN (
											SELECT price_province_inflation(
												@commodityId,
												c.province_id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )
										) = 0
										THEN 'same'
									END
							)
						)
					FROM price_province(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', (
									SELECT commodity_price_city_id FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									JSON_BUILD_OBJECT(
										'id', id,
										'name', name,
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
											WHERE id = a.assets_relation_id
										)
									)
								),
								'price', (
									SELECT price FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT
												price
											FROM price_city(@commodityId, a.id)
											WHERE idx = 1
										)
									)
								),
								'tier', (
									SELECT
										CASE 
											WHEN (
												SELECT price_city_inflation(
													@commodityId,
													a.id,
													start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
											    )
											) > 0
											THEN 'higher'
											WHEN (
											    SELECT price_city_inflation(
													@commodityId,
													a.id,
													start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
											    )
											) < 0
											THEN 'lower'
											WHEN (
											    SELECT price_city_inflation(
													@commodityId,
													a.id,
													start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
											    )
											) = 0
											THEN 'same'
										END
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								)
							)
						)
					FROM (
						SELECT
							*,
							@selectedDate "selectedDate"
						FROM tm_city
						WHERE province_id = 73
						ORDER BY sequence
					) a
				) "priceLevel",
				(
					SELECT '{"higher":{"title":"Lebih Tinggi","color":"#FF6711"},"same":{"title":"Sama","color":"#32D583"},"lower":{"title":"Lebih Rendah","color":"#05603A"}}'::JSON
				) "priceTier",
				(
					SELECT '["higher","same","lower"]'::JSON
				) "priceTierCode";
		`

	PriceGetYtyAvgChild = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
							'higher', (
								SELECT
								    COUNT(1)
								FROM price_inflation_avg_child(
									a.id,
									@provinceId,
									start_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								)
								WHERE inflation > 0
							),
							'same', (
								SELECT
								    COUNT(1)
								FROM price_inflation_avg_child(
									a.id,
									@provinceId,
									start_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								)
								WHERE inflation = 0
							),
							'lower', (
								SELECT
								    COUNT(1)
								FROM price_inflation_avg_child(
									a.id,
									@provinceId,
									start_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year('2024-08-05')::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								)
								WHERE inflation < 0
							)
						)
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) summary,
				(
					SELECT JSON_BUILD_OBJECT(
						'id', a.id,
						'name', a.name
					) 
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
						JSON_BUILD_OBJECT(
							'id', null,
							'clientId', 1,
							'province', (
								SELECT * FROM province_object(c.province_id)
							),
							'price', c.price,
							'priceRupiahFormat', (
								SELECT rupiah_format(c.price)
							),
							'tier', (
								SELECT
									CASE 
										WHEN (
											SELECT price_province_inflation_avg_child(
												@commodityId,
												c.province_id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )
										) > 0
										THEN 'higher'
										WHEN (
											SELECT price_province_inflation_avg_child(
												@commodityId,
												c.province_id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )
										) < 0
										THEN 'lower'
										WHEN (
											SELECT price_province_inflation_avg_child(
												@commodityId,
												c.province_id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )
										) = 0
										THEN 'same'
									END
							)
						)
					FROM price_province_avg_child(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', (
									SELECT commodity_price_city_id FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									JSON_BUILD_OBJECT(
										'id', id,
										'name', name,
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
											WHERE id = a.assets_relation_id
										)
									)
								),
								'price', (
									SELECT price FROM price_city_avg_child(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT
												price
											FROM price_city_avg_child(@commodityId, a.id)
											WHERE idx = 1
										)
									)
								),
								'tier', (
									SELECT
										CASE 
											WHEN (
												SELECT price_city_inflation_avg_child(
													@commodityId,
													a.id,
													start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
											    )
											) > 0
											THEN 'higher'
											WHEN (
											    SELECT price_city_inflation_avg_child(
													@commodityId,
													a.id,
													start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
											    )
											) < 0
											THEN 'lower'
											WHEN (
											    SELECT price_city_inflation_avg_child(
													@commodityId,
													a.id,
													start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
											    )
											) = 0
											THEN 'same'
										END
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal_avg_child(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal_avg_child(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								)
							)
						)
					FROM (
						SELECT
							*,
							@selectedDate "selectedDate"
						FROM tm_city
						WHERE province_id = 73
						ORDER BY sequence
					) a
				) "priceLevel",
				(
					SELECT '{"higher":{"title":"Lebih Tinggi","color":"#FF6711"},"same":{"title":"Sama","color":"#32D583"},"lower":{"title":"Lebih Rendah","color":"#05603A"}}'::JSON
				) "priceTier",
				(
					SELECT '["higher","same","lower"]'::JSON
				) "priceTierCode";
		`

	PriceGetYtyList = `
			SELECT
				(
					SELECT JSON_BUILD_OBJECT(
						'id', a.id,
						'name', a.name
					) 
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
						JSON_BUILD_OBJECT(
							'id', commodity_price_province_id,
							'clientId', client_id,
							'province', (
								SELECT * FROM province_object(c.province_id)
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							),
							'priceDiff', (
								SELECT price_province_inflation(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT price_province_inflation(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							)::double precision
						)
					FROM price_province(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', (
									SELECT commodity_price_city_id FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									JSON_BUILD_OBJECT(
										'id', id,
										'name', name,
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
											WHERE id = a.assets_relation_id
										)
									)
								),
								'price', (
									SELECT price FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT
												price
											FROM price_city(@commodityId, a.id)
											WHERE idx = 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @selectedDate "selectedDate"
						FROM tm_city
						WHERE province_id = @provinceId
						ORDER BY sequence
					) a
				) "priceCity"
		`

	PriceGetYtyListWithOrder = `
			SELECT
				(
					SELECT JSON_BUILD_OBJECT(
						'id', a.id,
						'name', a.name
					) 
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
						JSON_BUILD_OBJECT(
							'id', commodity_price_province_id,
							'clientId', client_id,
							'province', (
								SELECT * FROM province_object(c.province_id)
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							),
							'priceDiff', (
								SELECT price_province_inflation(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT price_province_inflation(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							)::double precision
						)
					FROM price_province(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', (
									SELECT commodity_price_city_id FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									JSON_BUILD_OBJECT(
										'id', id,
										'name', name,
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
											WHERE id = a.assets_relation_id
										)
									)
								),
								'price', (
									SELECT price FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT
												price
											FROM price_city(@commodityId, a.id)
											WHERE idx = 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @selectedDate "selectedDate"
						FROM tm_city a
						WHERE province_id = @provinceId
						ORDER BY (
							SELECT price_city_inflation(
								@commodityId,
								a.id,
								start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
								end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
						   )
						)
					) a
				) "priceCity"
		`

	PriceGetYtyListWithOrderDesc = `
			SELECT
				(
					SELECT JSON_BUILD_OBJECT(
						'id', a.id,
						'name', a.name
					) 
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
						JSON_BUILD_OBJECT(
							'id', commodity_price_province_id,
							'clientId', client_id,
							'province', (
								SELECT * FROM province_object(c.province_id)
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							),
							'priceDiff', (
								SELECT price_province_inflation(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT price_province_inflation(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							)::double precision
						)
					FROM price_province(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', (
									SELECT commodity_price_city_id FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									JSON_BUILD_OBJECT(
										'id', id,
										'name', name,
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
											WHERE id = a.assets_relation_id
										)
									)
								),
								'price', (
									SELECT price FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT
												price
											FROM price_city(@commodityId, a.id)
											WHERE idx = 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @selectedDate "selectedDate"
						FROM tm_city a
						WHERE province_id = @provinceId
						ORDER BY (
							SELECT price_city_inflation(
								@commodityId,
								a.id,
								start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
								end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
						   )
						) DESC
					) a
				) "priceCity"
		`

	PriceGetYtyListAvgChild = `
			SELECT
				(
					SELECT JSON_BUILD_OBJECT(
						'id', a.id,
						'name', a.name
					) 
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
						JSON_BUILD_OBJECT(
							'id', null,
							'clientId', 1,
							'province', (
								SELECT * FROM province_object(c.province_id)
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							),
							'priceDiff', (
								SELECT price_province_inflation_avg_child(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation_avg_child(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT price_province_inflation_avg_child(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							)::double precision
						)
					FROM price_province_avg_child(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', (
									SELECT commodity_price_city_id FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									JSON_BUILD_OBJECT(
										'id', id,
										'name', name,
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
											WHERE id = a.assets_relation_id
										)
									)
								),
								'price', (
									SELECT price FROM price_city_avg_child(@commodityId, a.id, @selectedDate, @selectedDate)
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT
												price
											FROM price_city_avg_child(@commodityId, a.id)
											WHERE idx = 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal_avg_child(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal_avg_child(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation_avg_child(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @selectedDate "selectedDate"
						FROM tm_city
						WHERE province_id = @provinceId
						ORDER BY sequence
					) a
				) "priceCity"
		`

	PriceGetYtyListWithOrderAvgChild = `
			SELECT
				(
					SELECT JSON_BUILD_OBJECT(
						'id', a.id,
						'name', a.name
					) 
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
						JSON_BUILD_OBJECT(
							'id', null,
							'clientId', 1,
							'province', (
								SELECT * FROM province_object(c.province_id)
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							),
							'priceDiff', (
								SELECT price_province_inflation_avg_child(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation_avg_child(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT price_province_inflation_avg_child(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							)::double precision
						)
					FROM price_province_avg_child(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', (
									SELECT commodity_price_city_id FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									JSON_BUILD_OBJECT(
										'id', id,
										'name', name,
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
											WHERE id = a.assets_relation_id
										)
									)
								),
								'price', (
									SELECT price FROM price_city_avg_child(@commodityId, a.id, @selectedDate, @selectedDate)
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT
												price
											FROM price_city_avg_child(@commodityId, a.id)
											WHERE idx = 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal_avg_child(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal_avg_child(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation_avg_child(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @selectedDate "selectedDate"
						FROM tm_city
						WHERE province_id = @provinceId
						--ORDER BY sequence
					    ORDER BY (
							SELECT price_city_inflation_avg_child(
								@commodityId,
								a.id,
								start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
								end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
						   )
						)
					) a
				) "priceCity"
		`

	PriceGetYtyListWithOrderDescAvgChild = `
			SELECT
				(
					SELECT JSON_BUILD_OBJECT(
						'id', a.id,
						'name', a.name
					) 
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
						JSON_BUILD_OBJECT(
							'id', null,
							'clientId', 1,
							'province', (
								SELECT * FROM province_object(c.province_id)
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							),
							'priceDiff', (
								SELECT price_province_inflation_avg_child(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation_avg_child(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT price_province_inflation_avg_child(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
							   )
							)::double precision
						)
					FROM price_province_avg_child(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', (
									SELECT commodity_price_city_id FROM price_city(@commodityId, a.id)
									WHERE last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									JSON_BUILD_OBJECT(
										'id', id,
										'name', name,
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
											WHERE id = a.assets_relation_id
										)
									)
								),
								'price', (
									SELECT price FROM price_city_avg_child(@commodityId, a.id, @selectedDate, @selectedDate)
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT
												price
											FROM price_city_avg_child(@commodityId, a.id)
											WHERE idx = 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal_avg_child(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal_avg_child(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation_avg_child(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
								   )
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @selectedDate "selectedDate"
						FROM tm_city
						WHERE province_id = @provinceId
						--ORDER BY sequence
					    ORDER BY (
							SELECT price_city_inflation_avg_child(
								@commodityId,
								a.id,
								start_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
								end_of_month(TO_DATE(minus_1_year(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								start_of_month(TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								TO_DATE(NOW()::TEXT, 'YYYY-MM-DD')::TEXT
						   )
						) DESC
					) a
				) "priceCity"
		`

	PriceGetYtyCityHistory = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
						   'id', a.id,
						   'parentId', a.parent_id,
						   'name', a.name,
						   'assets', (
							   SELECT
								   JSON_BUILD_OBJECT(
									   'id', id,
									   'assetsType', assets_type,
									   'assetsLocation', assets_location,
									   'assetsLocationType', assets_location_type,
									   'assetsMediaType', assets_media_type,
									   'assetsExt', assets_ext,
									   'assetsName', assets_name
								   )
							   FROM assets
							   WHERE id = a.assets_relation_id
						   ),
						   'createdAt', a.created_at,
						   'updatedAt', a.updated_at,
						   'deletedAt', a.deleted_at
					   )
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
						JSON_BUILD_OBJECT(
							'id', a.id,
							'name', a.name,
							'assets', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', id,
										'assetsType', assets_type,
										'assetsLocation', assets_location,
										'assetsLocationType', assets_location_type,
										'assetsMediaType', assets_media_type,
										'assetsExt', assets_ext,
										'assetsName', assets_name
									)
								FROM assets
								WHERE id = a.assets_relation_id
							),
							'createdAt', a.created_at,
							'updatedAt', a.updated_at,
							'deletedAt', a.deleted_at
						)
						FROM tm_city a
						WHERE a.id = @cityId
				) city,
				(
					SELECT
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
										'inflation', (
											SELECT
												price_city_inflation(
													@commodityId,
													@cityId,
													start_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													start_of_month(TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT
												)
										),
										'inflationRupiahFormat', (
											SELECT
												rupiah_format(
													price_city_inflation_nominal(
														@commodityId,
														@cityId,
														start_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														end_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														start_of_month(TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT
													)::INTEGER
												)
										)
									)
								)
						)
					FROM (
						SELECT generate_series(@startDate, @endDate, '1 day'::interval) "last_update"
					)
				) inflations
		`

	PriceGetYtyCityHistoryAvgChild = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
						   'id', a.id,
						   'parentId', a.parent_id,
						   'name', a.name,
						   'assets', (
							   SELECT
								   JSON_BUILD_OBJECT(
									   'id', id,
									   'assetsType', assets_type,
									   'assetsLocation', assets_location,
									   'assetsLocationType', assets_location_type,
									   'assetsMediaType', assets_media_type,
									   'assetsExt', assets_ext,
									   'assetsName', assets_name
								   )
							   FROM assets
							   WHERE id = a.assets_relation_id
						   ),
						   'createdAt', a.created_at,
						   'updatedAt', a.updated_at,
						   'deletedAt', a.deleted_at
					   )
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
						JSON_BUILD_OBJECT(
							'id', a.id,
							'name', a.name,
							'assets', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', id,
										'assetsType', assets_type,
										'assetsLocation', assets_location,
										'assetsLocationType', assets_location_type,
										'assetsMediaType', assets_media_type,
										'assetsExt', assets_ext,
										'assetsName', assets_name
									)
								FROM assets
								WHERE id = a.assets_relation_id
							),
							'createdAt', a.created_at,
							'updatedAt', a.updated_at,
							'deletedAt', a.deleted_at
						)
						FROM tm_city a
						WHERE a.id = @cityId
				) city,
				(
					SELECT
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
										'inflation', (
											SELECT
												price_city_inflation_avg_child(
													@commodityId,
													@cityId,
													start_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													start_of_month(TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT
												)
										),
										'inflationRupiahFormat', (
											SELECT
												rupiah_format(
													price_city_inflation_nominal_avg_child(
														@commodityId,
														@cityId,
														start_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														end_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														start_of_month(TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT
													)::INTEGER
												)
										)
									)
								)
						)
					FROM (
						SELECT generate_series(@startDate, @endDate, '1 day'::interval) "last_update"
					)
				) inflations
		`

	PriceGetYtyCityHistoryProvince = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
						   'id', a.id,
						   'parentId', a.parent_id,
						   'name', a.name,
						   'assets', (
							   SELECT
								   JSON_BUILD_OBJECT(
									   'id', id,
									   'assetsType', assets_type,
									   'assetsLocation', assets_location,
									   'assetsLocationType', assets_location_type,
									   'assetsMediaType', assets_media_type,
									   'assetsExt', assets_ext,
									   'assetsName', assets_name
								   )
							   FROM assets
							   WHERE id = a.assets_relation_id
						   ),
						   'createdAt', a.created_at,
						   'updatedAt', a.updated_at,
						   'deletedAt', a.deleted_at
					   )
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
			    	SELECT * FROM province_object(@provinceId)
				) province,
				(
					SELECT
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
										'inflation', (
											SELECT
												price_province_inflation(
													@commodityId,
													@provinceId,
													start_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													start_of_month(TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT
												)
										),
										'inflationRupiahFormat', (
											SELECT
												rupiah_format(
													price_province_inflation_nominal(
														@commodityId,
														@provinceId,
														start_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														end_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														start_of_month(TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT
													)::INTEGER
												)
										)
									)
								)
						)
					FROM (
						SELECT generate_series(@startDate, @endDate, '1 day'::interval) "last_update"
					)
				) inflations
		`

	PriceGetYtyCityHistoryProvinceAvgChild = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
						   'id', a.id,
						   'parentId', a.parent_id,
						   'name', a.name,
						   'assets', (
							   SELECT
								   JSON_BUILD_OBJECT(
									   'id', id,
									   'assetsType', assets_type,
									   'assetsLocation', assets_location,
									   'assetsLocationType', assets_location_type,
									   'assetsMediaType', assets_media_type,
									   'assetsExt', assets_ext,
									   'assetsName', assets_name
								   )
							   FROM assets
							   WHERE id = a.assets_relation_id
						   ),
						   'createdAt', a.created_at,
						   'updatedAt', a.updated_at,
						   'deletedAt', a.deleted_at
					   )
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
			    	SELECT * FROM province_object(@provinceId)
				) province,
				(
					SELECT
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
										'inflation', (
											SELECT
												price_province_inflation_avg_child(
													@commodityId,
													@provinceId,
													start_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													start_of_month(TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT
												)
										),
										'inflationRupiahFormat', (
											SELECT
												rupiah_format(
													price_province_inflation_nominal_avg_child(
														@commodityId,
														@provinceId,
														start_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														end_of_month(TO_DATE(minus_1_year(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														start_of_month(TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														TO_DATE(last_update::TEXT, 'YYYY-MM-DD')::TEXT
													)::INTEGER
												)
										)
									)
								)
						)
					FROM (
						SELECT generate_series(@startDate, @endDate, '1 day'::interval) "last_update"
					)
				) inflations
		`

	PriceGetYtyCommodityHistoryHigher = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
						   'id', a.id,
						   'parentId', a.parent_id,
						   'name', a.name,
						   'assets', (
							   SELECT
								   JSON_BUILD_OBJECT(
									   'id', id,
									   'assetsType', assets_type,
									   'assetsLocation', assets_location,
									   'assetsLocationType', assets_location_type,
									   'assetsMediaType', assets_media_type,
									   'assetsExt', assets_ext,
									   'assetsName', assets_name
								   )
							   FROM assets
							   WHERE id = a.assets_relation_id
						   ),
						   'createdAt', a.created_at,
						   'updatedAt', a.updated_at,
						   'deletedAt', a.deleted_at
					   )
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
					    JSON_AGG(
							JSON_BUILD_OBJECT(
								'city', (
									SELECT
										JSON_BUILD_OBJECT(
											'id', a.id,
											'name', a.name,
											'assets', (
												SELECT
													JSON_BUILD_OBJECT(
														'id', id,
														'assetsType', assets_type,
														'assetsLocation', assets_location,
														'assetsLocationType', assets_location_type,
														'assetsMediaType', assets_media_type,
														'assetsExt', assets_ext,
														'assetsName', assets_name
													)
												FROM assets
												WHERE id = a.assets_relation_id
											),
											'createdAt', a.created_at,
											'updatedAt', a.updated_at,
											'deletedAt', a.deleted_at
										)
								),
								'inflation', (
									SELECT
										price_city_inflation(
											@commodityId,
											a.id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
										)
								),
								'inflationRupiahFormat', (
									SELECT
										rupiah_format(
											price_city_inflation_nominal(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
											)::INTEGER
										)
								)
							)
						)
					FROM (
					    SELECT
					        id,
					        name,
					        assets_relation_id,
					        a.created_at,
					        a.updated_at,
					        a.deleted_at,
					        (
					            SELECT
									price_city_inflation(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
									)
					        ) inflation
					    FROM tm_city a
						WHERE a.province_id = @provinceId
						ORDER BY sequence
					) a
					WHERE inflation > 0
				) cityInflations
		`

	PriceGetYtyCommodityHistoryHigherAvgChild = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
						   'id', a.id,
						   'parentId', a.parent_id,
						   'name', a.name,
						   'assets', (
							   SELECT
								   JSON_BUILD_OBJECT(
									   'id', id,
									   'assetsType', assets_type,
									   'assetsLocation', assets_location,
									   'assetsLocationType', assets_location_type,
									   'assetsMediaType', assets_media_type,
									   'assetsExt', assets_ext,
									   'assetsName', assets_name
								   )
							   FROM assets
							   WHERE id = a.assets_relation_id
						   ),
						   'createdAt', a.created_at,
						   'updatedAt', a.updated_at,
						   'deletedAt', a.deleted_at
					   )
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
					    JSON_AGG(
							JSON_BUILD_OBJECT(
								'city', (
									SELECT
										JSON_BUILD_OBJECT(
											'id', a.id,
											'name', a.name,
											'assets', (
												SELECT
													JSON_BUILD_OBJECT(
														'id', id,
														'assetsType', assets_type,
														'assetsLocation', assets_location,
														'assetsLocationType', assets_location_type,
														'assetsMediaType', assets_media_type,
														'assetsExt', assets_ext,
														'assetsName', assets_name
													)
												FROM assets
												WHERE id = a.assets_relation_id
											),
											'createdAt', a.created_at,
											'updatedAt', a.updated_at,
											'deletedAt', a.deleted_at
										)
								),
								'inflation', (
									SELECT
										price_city_inflation_avg_child(
											@commodityId,
											a.id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
										)
								),
								'inflationRupiahFormat', (
									SELECT
										rupiah_format(
											price_city_inflation_nominal_avg_child(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
											)::INTEGER
										)
								)
							)
						)
					FROM (
					    SELECT
					        id,
					        name,
					        assets_relation_id,
					        a.created_at,
					        a.updated_at,
					        a.deleted_at,
					        (
					            SELECT
									price_city_inflation_avg_child(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
									)
					        ) inflation
					    FROM tm_city a
						WHERE a.province_id = @provinceId
						ORDER BY sequence
					) a
					WHERE inflation > 0
				) cityInflations
		`

	PriceGetYtyCommodityHistorySame = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
						   'id', a.id,
						   'parentId', a.parent_id,
						   'name', a.name,
						   'assets', (
							   SELECT
								   JSON_BUILD_OBJECT(
									   'id', id,
									   'assetsType', assets_type,
									   'assetsLocation', assets_location,
									   'assetsLocationType', assets_location_type,
									   'assetsMediaType', assets_media_type,
									   'assetsExt', assets_ext,
									   'assetsName', assets_name
								   )
							   FROM assets
							   WHERE id = a.assets_relation_id
						   ),
						   'createdAt', a.created_at,
						   'updatedAt', a.updated_at,
						   'deletedAt', a.deleted_at
					   )
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
					    JSON_AGG(
							JSON_BUILD_OBJECT(
								'city', (
									SELECT
										JSON_BUILD_OBJECT(
											'id', a.id,
											'name', a.name,
											'assets', (
												SELECT
													JSON_BUILD_OBJECT(
														'id', id,
														'assetsType', assets_type,
														'assetsLocation', assets_location,
														'assetsLocationType', assets_location_type,
														'assetsMediaType', assets_media_type,
														'assetsExt', assets_ext,
														'assetsName', assets_name
													)
												FROM assets
												WHERE id = a.assets_relation_id
											),
											'createdAt', a.created_at,
											'updatedAt', a.updated_at,
											'deletedAt', a.deleted_at
										)
								),
								'inflation', (
									SELECT
										price_city_inflation(
											@commodityId,
											a.id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
										)
								),
								'inflationRupiahFormat', (
									SELECT
										rupiah_format(
											price_city_inflation_nominal(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
											)::INTEGER
										)
								)
							)
						)
					FROM (
					    SELECT
					        id,
					        name,
					        assets_relation_id,
					        a.created_at,
					        a.updated_at,
					        a.deleted_at,
					        (
					            SELECT
									price_city_inflation(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
									)
					        ) inflation
					    FROM tm_city a
						WHERE a.province_id = @provinceId
						ORDER BY sequence
					) a
					WHERE inflation = 0
				) cityInflations
		`

	PriceGetYtyCommodityHistorySameAvgChild = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
						   'id', a.id,
						   'parentId', a.parent_id,
						   'name', a.name,
						   'assets', (
							   SELECT
								   JSON_BUILD_OBJECT(
									   'id', id,
									   'assetsType', assets_type,
									   'assetsLocation', assets_location,
									   'assetsLocationType', assets_location_type,
									   'assetsMediaType', assets_media_type,
									   'assetsExt', assets_ext,
									   'assetsName', assets_name
								   )
							   FROM assets
							   WHERE id = a.assets_relation_id
						   ),
						   'createdAt', a.created_at,
						   'updatedAt', a.updated_at,
						   'deletedAt', a.deleted_at
					   )
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
					    JSON_AGG(
							JSON_BUILD_OBJECT(
								'city', (
									SELECT
										JSON_BUILD_OBJECT(
											'id', a.id,
											'name', a.name,
											'assets', (
												SELECT
													JSON_BUILD_OBJECT(
														'id', id,
														'assetsType', assets_type,
														'assetsLocation', assets_location,
														'assetsLocationType', assets_location_type,
														'assetsMediaType', assets_media_type,
														'assetsExt', assets_ext,
														'assetsName', assets_name
													)
												FROM assets
												WHERE id = a.assets_relation_id
											),
											'createdAt', a.created_at,
											'updatedAt', a.updated_at,
											'deletedAt', a.deleted_at
										)
								),
								'inflation', (
									SELECT
										price_city_inflation_avg_child(
											@commodityId,
											a.id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
										)
								),
								'inflationRupiahFormat', (
									SELECT
										rupiah_format(
											price_city_inflation_nominal_avg_child(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
											)::INTEGER
										)
								)
							)
						)
					FROM (
					    SELECT
					        id,
					        name,
					        assets_relation_id,
					        a.created_at,
					        a.updated_at,
					        a.deleted_at,
					        (
					            SELECT
									price_city_inflation_avg_child(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
									)
					        ) inflation
					    FROM tm_city a
						WHERE a.province_id = @provinceId
						ORDER BY sequence
					) a
					WHERE inflation = 0
				) cityInflations
		`

	PriceGetYtyCommodityHistoryLower = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
						   'id', a.id,
						   'parentId', a.parent_id,
						   'name', a.name,
						   'assets', (
							   SELECT
								   JSON_BUILD_OBJECT(
									   'id', id,
									   'assetsType', assets_type,
									   'assetsLocation', assets_location,
									   'assetsLocationType', assets_location_type,
									   'assetsMediaType', assets_media_type,
									   'assetsExt', assets_ext,
									   'assetsName', assets_name
								   )
							   FROM assets
							   WHERE id = a.assets_relation_id
						   ),
						   'createdAt', a.created_at,
						   'updatedAt', a.updated_at,
						   'deletedAt', a.deleted_at
					   )
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
					    JSON_AGG(
							JSON_BUILD_OBJECT(
								'city', (
									SELECT
										JSON_BUILD_OBJECT(
											'id', a.id,
											'name', a.name,
											'assets', (
												SELECT
													JSON_BUILD_OBJECT(
														'id', id,
														'assetsType', assets_type,
														'assetsLocation', assets_location,
														'assetsLocationType', assets_location_type,
														'assetsMediaType', assets_media_type,
														'assetsExt', assets_ext,
														'assetsName', assets_name
													)
												FROM assets
												WHERE id = a.assets_relation_id
											),
											'createdAt', a.created_at,
											'updatedAt', a.updated_at,
											'deletedAt', a.deleted_at
										)
								),
								'inflation', (
									SELECT
										price_city_inflation(
											@commodityId,
											a.id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
										)
								),
								'inflationRupiahFormat', (
									SELECT
										rupiah_format(
											price_city_inflation_nominal(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
											)::INTEGER
										)
								)
							)
						)
					FROM (
					    SELECT
					        id,
					        name,
					        assets_relation_id,
					        a.created_at,
					        a.updated_at,
					        a.deleted_at,
					        (
					            SELECT
									price_city_inflation(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
									)
					        ) inflation
					    FROM tm_city a
						WHERE a.province_id = @provinceId
						ORDER BY sequence
					) a
					WHERE inflation < 0
				) cityInflations
		`

	PriceGetYtyCommodityHistoryLowerAvgChild = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
						   'id', a.id,
						   'parentId', a.parent_id,
						   'name', a.name,
						   'assets', (
							   SELECT
								   JSON_BUILD_OBJECT(
									   'id', id,
									   'assetsType', assets_type,
									   'assetsLocation', assets_location,
									   'assetsLocationType', assets_location_type,
									   'assetsMediaType', assets_media_type,
									   'assetsExt', assets_ext,
									   'assetsName', assets_name
								   )
							   FROM assets
							   WHERE id = a.assets_relation_id
						   ),
						   'createdAt', a.created_at,
						   'updatedAt', a.updated_at,
						   'deletedAt', a.deleted_at
					   )
					FROM tm_commodity a
					WHERE a.id = @commodityId
				) commodity,
				(
					SELECT
					    JSON_AGG(
							JSON_BUILD_OBJECT(
								'city', (
									SELECT
										JSON_BUILD_OBJECT(
											'id', a.id,
											'name', a.name,
											'assets', (
												SELECT
													JSON_BUILD_OBJECT(
														'id', id,
														'assetsType', assets_type,
														'assetsLocation', assets_location,
														'assetsLocationType', assets_location_type,
														'assetsMediaType', assets_media_type,
														'assetsExt', assets_ext,
														'assetsName', assets_name
													)
												FROM assets
												WHERE id = a.assets_relation_id
											),
											'createdAt', a.created_at,
											'updatedAt', a.updated_at,
											'deletedAt', a.deleted_at
										)
								),
								'inflation', (
									SELECT
										price_city_inflation_avg_child(
											@commodityId,
											a.id,
											start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
										)
								),
								'inflationRupiahFormat', (
									SELECT
										rupiah_format(
											price_city_inflation_nominal_avg_child(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
											)::INTEGER
										)
								)
							)
						)
					FROM (
					    SELECT
					        id,
					        name,
					        assets_relation_id,
					        a.created_at,
					        a.updated_at,
					        a.deleted_at,
					        (
					            SELECT
									price_city_inflation_avg_child(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_year(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate::TEXT, 'YYYY-MM-DD')::TEXT
									)
					        ) inflation
					    FROM tm_city a
						WHERE a.province_id = @provinceId
						ORDER BY sequence
					) a
					WHERE inflation < 0
				) cityInflations
		`
)
