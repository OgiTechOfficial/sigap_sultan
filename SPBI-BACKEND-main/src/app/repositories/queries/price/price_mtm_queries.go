package price

const (
	PriceGetMtmSummary = `
		SELECT
			JSON_BUILD_OBJECT(
				'higher', (
					SELECT
						COUNT(1)
					FROM price_inflation(
						a.id,
						@provinceId,
						start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
						end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
						start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
						TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
					)
					WHERE inflation > 0
				),
				'same', (
					SELECT
						COUNT(1)
					FROM price_inflation(
						a.id,
						@provinceId,
						start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
						end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
						start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
						TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
					)
					WHERE inflation = 0
				),
				'lower', (
					SELECT
						COUNT(1)
					FROM price_inflation(
						a.id,
						@provinceId,
						start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
						end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
						start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
						TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
					)
					WHERE inflation < 0
				)
			)
		FROM tm_commodity a
		WHERE a.id = @commodityId
	`

	PriceGetProvincePrice = `
		SELECT
		    JSON_BUILD_OBJECT(
				'id', commodity_price_id,
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							   )
							) > 0
							THEN 'higher'
							WHEN (
								SELECT price_province_inflation(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							   )
							) < 0
							THEN 'lower'
							WHEN (
								SELECT price_province_inflation(
									@commodityId,
									c.province_id,
									start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							   )
							) = 0
							THEN 'same'
						END
				)
			)
		FROM get_level_harga_province(@commodityId, @selectedDate) c
		WHERE province_id = @provinceId
	`

	PriceGetPriceList = `
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
					'price', (a.price),
					'priceRupiahFormat', 
					(
						SELECT rupiah_format(a.price)
					),
					'tier', (
						SELECT
							CASE 
								WHEN (
									SELECT price_city_inflation(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
									)
								) > 0
								THEN 'higher'
								WHEN (
									SELECT price_city_inflation(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
									)
								) < 0
								THEN 'lower'
								WHEN (
									SELECT price_city_inflation(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
									)
								) = 0
								THEN 'same'
							END
					),
					'priceDiff', (
						SELECT price_city_inflation_nominal(
							@commodityId,
							a.id,
							start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
							end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
							start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
							TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
					   )
					),
					'priceDiffFormat', (
						SELECT rupiah_format(
						   (
							   SELECT price_city_inflation_nominal(
									@commodityId,
									a.id,
									start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
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
			FROM get_level_harga(@provinceId, @commodityId, @selectedDate) a
		) a
	`

	PriceGetMtm = `
			WITH
			ctePriceInflation AS(
				SELECT
					*
				FROM price_inflation(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
				)
			),
			ctePriceProvinceInflation AS (
				SELECT
					*
				FROM price_province_inflation(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
			   )
			)
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
							'higher', (
								SELECT
									COUNT(1)
								FROM ctePriceInflation
								WHERE inflation > 0
							),
							'same', (
								SELECT
									COUNT(1)
								FROM ctePriceInflation
								WHERE inflation = 0
							),
							'lower', (
								SELECT
									COUNT(1)
								FROM ctePriceInflation
								WHERE inflation < 0
							)
						)
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
 							'id', commodity_price_id,
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
 											SELECT * FROM ctePriceProvinceInflation
 										) > 0
 										THEN 'higher'
 										WHEN (
 											SELECT * FROM ctePriceProvinceInflation
 										) < 0
 										THEN 'lower'
 										WHEN (
 											SELECT * FROM ctePriceProvinceInflation
 										) = 0
 										THEN 'same'
 									END
 							)
 						)
 					FROM get_level_harga_province(@commodityId, @selectedDate) c
 					WHERE province_id = @provinceId
 				) "provincePrice",
 				(
 					SELECT
 						JSON_AGG(
 							JSON_BUILD_OBJECT(
 								'id', null,
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
 													'assetsName', assets_name
 													--'assetsUrl', CONCAT(
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
													--'assetsUrl', CONCAT((SELECT name FROM settings WHERE parent_id = (SELECT id FROM settings WHERE name = 'BASE_URL')), '/assets?assets_location=', CONCAT(assets_location, '/',assets_name))
 												)
 											FROM assets
 											WHERE id = a.assets_relation_id
 										)
 									)
 								),
 								'price', (a.price),
 								'priceRupiahFormat', 
 								(
 									SELECT rupiah_format(a.price)
 								),
 								'tier', (
 									SELECT
 										CASE 
 											WHEN (
 												SELECT price_city_inflation(
 													@commodityId,
 													a.id,
 													start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
 													end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
 													start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
 													TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
 											    )
 											) > 0
 											THEN 'higher'
 											WHEN (
 											    SELECT price_city_inflation(
 													@commodityId,
 													a.id,
 													start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
 													end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
 													start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
 													TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
 											    )
 											) < 0
 											THEN 'lower'
 											WHEN (
 											    SELECT price_city_inflation(
 													@commodityId,
 													a.id,
 													start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
 													end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
 													start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
 													TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
 											    )
 											) = 0
 											THEN 'same'
 										END
 								),
 								'priceDiff', (
 								    SELECT price_city_inflation_nominal(
 										@commodityId,
 										a.id,
 										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
 										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
 										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
 										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
 								   )
 								),
 								'priceDiffFormat', (
 									SELECT rupiah_format(
 									   (
 										   SELECT price_city_inflation_nominal(
 												@commodityId,
 												a.id,
 												start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
 												end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
 												start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
 												TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
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
 						FROM get_level_harga(@provinceId, @commodityId, @selectedDate) a
 					) a
 				) "priceLevel",
				(
					SELECT '{"higher":{"title":"Lebih Tinggi","color":"#FF6711"},"same":{"title":"Sama","color":"#32D583"},"lower":{"title":"Lebih Rendah","color":"#05603A"}}'::JSON
				) "priceTier",
				(
					SELECT '["higher","same","lower"]'::JSON
				) "priceTierCode";
		`

	PriceGetMtmAvgChild = `
			WITH
			ctePriceInflation AS(
				SELECT
					*
				FROM price_inflation_avg_child(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
				)
			),
			ctePriceProvinceInflation AS (
				SELECT
					*
				FROM price_province_inflation_avg_child(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
			   )
			)
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
							'higher', (
								SELECT
									COUNT(1)
								FROM ctePriceInflation
								WHERE inflation > 0
							),
							'same', (
								SELECT
									COUNT(1)
								FROM ctePriceInflation
								WHERE inflation = 0
							),
							'lower', (
								SELECT
									COUNT(1)
								FROM ctePriceInflation
								WHERE inflation < 0
							)
						)
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
 											SELECT * FROM ctePriceProvinceInflation
 										) > 0
 										THEN 'higher'
 										WHEN (
 											SELECT * FROM ctePriceProvinceInflation
 										) < 0
 										THEN 'lower'
 										WHEN (
 											SELECT * FROM ctePriceProvinceInflation
 										) = 0
 										THEN 'same'
 									END
 							)
 						)
					FROM price_province_avg_child(@commodityId, @provinceId, @selectedDate) c
 				) "provincePrice",
 				(
 					SELECT
 						JSON_AGG(
 							JSON_BUILD_OBJECT(
 								'id', null,
 								'client_id', 1,
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
														'assetsName', assets_name
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
 								'price', (a.price),
 								'priceRupiahFormat', 
 								(
 									SELECT rupiah_format(a.price)
 								),
 								'tier', (
 									SELECT
 										CASE 
 											WHEN (
 												SELECT price_city_inflation_avg_child(
 													@commodityId,
 													a.id,
 													start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
 													end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
 													start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
 													TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
 											    )
 											) > 0
 											THEN 'higher'
 											WHEN (
 											    SELECT price_city_inflation_avg_child(
 													@commodityId,
 													a.id,
 													start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
 													end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
 													start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
 													TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
 											    )
 											) < 0
 											THEN 'lower'
 											WHEN (
 											    SELECT price_city_inflation_avg_child(
 													@commodityId,
 													a.id,
 													start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
 													end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
 													start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
 													TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
 											    )
 											) = 0
 											THEN 'same'
 										END
 								),
 								'priceDiff', (
 								    SELECT price_city_inflation_nominal_avg_child(
 										@commodityId,
 										a.id,
 										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
 										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
 										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
 										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
 								   )
 								),
 								'priceDiffFormat', (
 									SELECT rupiah_format(
 									   (
 										   SELECT price_city_inflation_nominal_avg_child(
 												@commodityId,
 												a.id,
 												start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
 												end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
 												start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
 												TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
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
 						FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate) a
						--FROM price_city_avg_child(@commodityId, @cityId, @selectedDate, @selectedDate) a
 					) a
 				) "priceLevel",
				(
					SELECT '{"higher":{"title":"Lebih Tinggi","color":"#FF6711"},"same":{"title":"Sama","color":"#32D583"},"lower":{"title":"Lebih Rendah","color":"#05603A"}}'::JSON
				) "priceTier",
				(
					SELECT '["higher","same","lower"]'::JSON
				) "priceTierCode";
		`

	PriceGetMtmList = `
			WITH
			ctePriceInflation AS(
				SELECT
					*
				FROM price_inflation(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
				)
			),
			ctePriceProvinceInflation AS (
				SELECT
					*
				FROM price_province_inflation(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
			   )
			)
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT * FROM ctePriceProvinceInflation
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
													'assetsName', assets_name
													--'assetsUrl', CONCAT(
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
													--'assetsUrl', CONCAT((SELECT name FROM settings WHERE parent_id = (SELECT id FROM settings WHERE name = 'BASE_URL')), '/assets?assets_location=', CONCAT(assets_location, '/',assets_name))
												)
											FROM assets
											WHERE id = a.assets_relation_id
										)
									)
								),
								'price', (
									SELECT price
									FROM tx_commodity_price tcp
									WHERE
										commodity_id = @commodityId AND
										city_id = a.id AND
										last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
									ORDER BY id DESC
									LIMIT 1
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
										    SELECT price::INTEGER
											FROM tx_commodity_price tcp
											WHERE
												commodity_id = @commodityId AND
												city_id = a.id AND
												last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
											ORDER BY id DESC
											LIMIT 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
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

	PriceGetMtmListWithOrder = `
			WITH
			ctePriceInflation AS(
				SELECT
					*
				FROM price_inflation(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
				)
			),
			ctePriceProvinceInflation AS (
				SELECT
					*
				FROM price_province_inflation(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
			   )
			)
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT * FROM ctePriceProvinceInflation
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
									SELECT price
									FROM tx_commodity_price tcp
									WHERE
										commodity_id = @commodityId AND
										city_id = a.id AND
										last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
									ORDER BY id DESC
									LIMIT 1
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
										    SELECT price::INTEGER
											FROM tx_commodity_price tcp
											WHERE
												commodity_id = @commodityId AND
												city_id = a.id AND
												last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
											ORDER BY id DESC
											LIMIT 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
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
						--ORDER BY sequence
						ORDER BY (
							SELECT price_city_inflation_nominal(
								@commodityId,
								a.id,
								start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
								end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
								TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							)
						)
					) a
				) "priceCity"
		`

	PriceGetMtmListWithOrderDesc = `
			WITH
			ctePriceInflation AS(
				SELECT
					*
				FROM price_inflation(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
				)
			),
			ctePriceProvinceInflation AS (
				SELECT
					*
				FROM price_province_inflation(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
			   )
			)
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT * FROM ctePriceProvinceInflation
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
									SELECT price
									FROM tx_commodity_price tcp
									WHERE
										commodity_id = @commodityId AND
										city_id = a.id AND
										last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
									ORDER BY id DESC
									LIMIT 1
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
										    SELECT price::INTEGER
											FROM tx_commodity_price tcp
											WHERE
												commodity_id = @commodityId AND
												city_id = a.id AND
												last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
											ORDER BY id DESC
											LIMIT 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
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
						--ORDER BY sequence
						ORDER BY (
							SELECT price_city_inflation_nominal(
								@commodityId,
								a.id,
								start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
								end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
								TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							)
						) DESC
					) a
				) "priceCity"
		`

	PriceGetMtmListAvgChild = `
			WITH
			ctePriceInflation AS(
				SELECT
					*
				FROM price_inflation_avg_child(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
				)
			),
			ctePriceProvinceInflation AS (
				SELECT
					*
				FROM price_province_inflation_avg_child(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
			   )
			)
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation_avg_child(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT * FROM ctePriceProvinceInflation
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
									SELECT AVG(price)::INTEGER
									FROM tx_commodity_price tcp
									WHERE
										commodity_id IN (
											SELECT id
											FROM tm_commodity
											WHERE parent_id = @commodityId
										) AND
										city_id = a.id AND
										last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
									GROUP BY
										last_update
									ORDER BY last_update DESC
									LIMIT 1
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT AVG(price)::INTEGER
											FROM tx_commodity_price tcp
											WHERE
												commodity_id IN (
													SELECT id
													FROM tm_commodity
													WHERE parent_id = @commodityId
												) AND
												city_id = a.id AND
												last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
											GROUP BY
												last_update
											ORDER BY last_update DESC
											LIMIT 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal_avg_child(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal_avg_child(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation_avg_child(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
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

	PriceGetMtmListWithOrderAvgChild = `
			WITH
			ctePriceInflation AS(
				SELECT
					*
				FROM price_inflation_avg_child(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
				)
			),
			ctePriceProvinceInflation AS (
				SELECT
					*
				FROM price_province_inflation_avg_child(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
			   )
			)
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation_avg_child(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT * FROM ctePriceProvinceInflation
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
									SELECT AVG(price)::INTEGER
									FROM tx_commodity_price tcp
									WHERE
										commodity_id IN (
											SELECT id
											FROM tm_commodity
											WHERE parent_id = @commodityId
										) AND
										city_id = a.id AND
										last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
									GROUP BY
										last_update
									ORDER BY last_update DESC
									LIMIT 1
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT AVG(price)::INTEGER
											FROM tx_commodity_price tcp
											WHERE
												commodity_id IN (
													SELECT id
													FROM tm_commodity
													WHERE parent_id = @commodityId
												) AND
												city_id = a.id AND
												last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
											GROUP BY
												last_update
											ORDER BY last_update DESC
											LIMIT 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal_avg_child(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal_avg_child(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation_avg_child(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
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
						--ORDER BY sequence
						ORDER BY (
							SELECT price_city_inflation_nominal_avg_child(
								@commodityId,
								a.id,
								start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
								end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
								TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
						   )
						)
					) a
				) "priceCity"
		`

	PriceGetMtmListWithOrderDescAvgChild = `
			WITH
			ctePriceInflation AS(
				SELECT
					*
				FROM price_inflation_avg_child(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
				)
			),
			ctePriceProvinceInflation AS (
				SELECT
					*
				FROM price_province_inflation_avg_child(
					@commodityId,
					@provinceId,
					start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
					start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
					TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
			   )
			)
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
									TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
							   )
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
								   (
									   SELECT price_province_inflation_avg_child(
											@commodityId,
											c.province_id,
											start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
											TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
									   )::INTEGER
								   )
							   )
							),
							'priceDiffPercentage', (
							    SELECT * FROM ctePriceProvinceInflation
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
									SELECT AVG(price)::INTEGER
									FROM tx_commodity_price tcp
									WHERE
										commodity_id IN (
											SELECT id
											FROM tm_commodity
											WHERE parent_id = @commodityId
										) AND
										city_id = a.id AND
										last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
									GROUP BY
										last_update
									ORDER BY last_update DESC
									LIMIT 1
								),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(
										(
											SELECT AVG(price)::INTEGER
											FROM tx_commodity_price tcp
											WHERE
												commodity_id IN (
													SELECT id
													FROM tm_commodity
													WHERE parent_id = @commodityId
												) AND
												city_id = a.id AND
												last_update BETWEEN CONCAT(@selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(@selectedDate, ' 23:59:59')::TIMESTAMPTZ
											GROUP BY
												last_update
											ORDER BY last_update DESC
											LIMIT 1
										)
									)
								),
								'priceDiff', (
								    SELECT price_city_inflation_nominal_avg_child(
										@commodityId,
										a.id,
										start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
								   )
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									   (
										   SELECT price_city_inflation_nominal_avg_child(
												@commodityId,
												a.id,
												start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
												TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
										   )::INTEGER
									   )
								   )
								),
								'priceDiffPercentage', (
								    SELECT price_city_inflation_avg_child(
								    	@commodityId,
								    	a.id,
								    	start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
										TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
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
						--ORDER BY sequence
						ORDER BY (
							SELECT price_city_inflation_nominal_avg_child(
								@commodityId,
								a.id,
								start_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT,'YYYY-MM-DD')::TEXT)::TEXT,
								end_of_month(TO_DATE(minus_1_month(@selectedDate)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
								start_of_month(TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT)::TEXT,
								TO_DATE(@selectedDate, 'YYYY-MM-DD')::TEXT
						   )
						) DESC
					) a
				) "priceCity"
		`

	PriceGetMtmCityHistory = `
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
													start_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
														start_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														end_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
					) a
				) inflations
		`

	PriceGetMtmCityHistoryAvgChild = `
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
													start_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
														start_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														end_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
					) a
				) inflations
		`

	PriceGetMtmCityHistoryProvince = `
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
													start_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
														start_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														end_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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

	PriceGetMtmCityHistoryProvinceAvgChild = `
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
-- 				(
-- 					SELECT
-- 						JSON_BUILD_OBJECT(
-- 							'id', a.id,
-- 							'name', a.name,
-- 							'assets', (
-- 								SELECT
-- 									JSON_BUILD_OBJECT(
-- 										'id', id,
-- 										'assetsType', assets_type,
-- 										'assetsLocation', assets_location,
-- 										'assetsLocationType', assets_location_type,
-- 										'assetsMediaType', assets_media_type,
-- 										'assetsExt', assets_ext,
-- 										'assetsName', assets_name
-- 									)
-- 								FROM assets
-- 								WHERE id = a.assets_relation_id
-- 							),
-- 							'createdAt', a.created_at,
-- 							'updatedAt', a.updated_at,
-- 							'deletedAt', a.deleted_at
-- 						)
-- 						FROM tm_city a
-- 						WHERE a.id = @provinceId
-- 				) city,
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
													start_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
													end_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
														start_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
														end_of_month(TO_DATE(minus_1_month(last_update::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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

	PriceGetMtmCommodityHistoryHigher = `
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
										start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
											start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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

	PriceGetMtmCommodityHistoryHigherAvgChild = `
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
										start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
											start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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

	PriceGetMtmCommodityHistorySame = `
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
											start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
												start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
										start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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

	PriceGetMtmCommodityHistorySameAvgChild = `
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
											start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
												start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
												end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
										start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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

	PriceGetMtmCommodityHistoryLower = `
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
										start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
											start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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

	PriceGetMtmCommodityHistoryLowerAvgChild = `
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
										start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
										end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
											start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
											end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
									start_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
									end_of_month(TO_DATE(minus_1_month(@selectedDate::TEXT)::TEXT, 'YYYY-MM-DD')::TEXT)::TEXT,
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
