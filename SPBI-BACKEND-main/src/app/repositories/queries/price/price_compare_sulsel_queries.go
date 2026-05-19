package price

const (
	PriceGetCompareBySulsel = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
							'higher', (
								SELECT * FROM price_higher_than_province(a.id, @provinceId, @selectedDate) 
							),
							'same', (
								SELECT * FROM price_same_as_province(a.id, @provinceId, @selectedDate)
							),
							'lower', (
								SELECT * FROM price_lower_than_province(a.id, @provinceId, @selectedDate)
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
									WHERE last_update BETWEEN CONCAT(selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									SELECT * FROM city_object(a.id)
								),
								'price', (a.price),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
								'tier', (
									SELECT * FROM price_compare_province(@commodityId, a.id, a.province_id, selectedDate)
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @selectedDate selectedDate
						FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
					) a
				) "priceLevel",
				(
					SELECT '{"higher":{"title":"Lebih Tinggi","color":"#FF6711"},"same":{"title":"Sama","color":"#32D583"},"lower":{"title":"Lebih Rendah","color":"#05603A"}}'::JSON
				) "priceTier",
			    (
					SELECT '["higher","same","lower"]'::JSON
				) "priceTierCode"
		`

	PriceGetCompareBySulselAvgChild = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
							'higher', (
								SELECT * FROM price_higher_than_province_avg_child(@commodityId, @provinceId, @selectedDate) 
							),
							'same', (
								SELECT * FROM price_same_as_province_avg_child(@commodityId, @provinceId, @selectedDate)
							),
							'lower', (
								SELECT * FROM price_lower_than_province_avg_child(@commodityId, @provinceId, @selectedDate)
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
							'clientId', 1,
							'province', (
								SELECT * FROM province_object(c.province_id)
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							)
						)
					FROM price_province_avg_child(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', null,
								'client_id', 1,
								'city', (
									SELECT * FROM city_object(a.id)
								),
								'price', (a.price),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
								'tier', (
									SELECT * FROM price_compare_province_avg_child(@commodityId, a.id, @provinceId, selectedDate)
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @selectedDate selectedDate
						FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
					) a
				) "priceLevel",
				(
					SELECT '{"higher":{"title":"Lebih Tinggi","color":"#FF6711"},"same":{"title":"Sama","color":"#32D583"},"lower":{"title":"Lebih Rendah","color":"#05603A"}}'::JSON
				) "priceTier",
			    (
					SELECT '["higher","same","lower"]'::JSON
				) "priceTierCode"
		`

	PriceGetCompareBySulselList = `
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
									WHERE last_update BETWEEN CONCAT(selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									SELECT * FROM city_object(a.id)
								),
								'price', (a.price),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
								'priceDiff', (
									a.price -
									(
										SELECT * FROM price_province_latest(a.province_id, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(a.price)
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @selectedDate selectedDate
						FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
					) a
				) "priceCity"
		`

	PriceGetCompareBySulselListAvgChild = `
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
									WHERE last_update BETWEEN CONCAT(selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									SELECT * FROM city_object(a.id)
								),
								'price', (a.price),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
								'priceDiff', (
									a.price -
									(
										SELECT * FROM price_province_latest_avg_child(@provinceId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(a.price)
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @selectedDate selectedDate
						FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
					) a
				) "priceCity"
		`

	PriceGetCompareBySulselListWithOrderAsc = `
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
									WHERE last_update BETWEEN CONCAT(selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									SELECT * FROM city_object(a.id)
								),
								'price', (a.price),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
								'priceDiff', (
									a.price -
									(
										SELECT * FROM price_province_latest(a.province_id, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									    a.price-
										(
											SELECT * FROM price_province_latest(a.province_id, @commodityId)
										)
									)
								)
							)
						)
					FROM (
						SELECT
							*
						FROM (
							SELECT
								*,
								@selectedDate selectedDate
							FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
						) b
						ORDER BY (
							SELECT (
								(
									SELECT price
									FROM get_level_harga(@provinceId, @commodityId, @selectedDate) c
									WHERE c.id = b.id
								) -
								(
									SELECT * FROM price_province_latest(province_id, @commodityId)
								)
							)
						)
					) a
				) "priceCity"
		`

	PriceGetCompareBySulselListWithOrderAscAvgChild = `
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
							)
						)
					FROM price_province_avg_child(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', null,
								'client_id', 1,
								'city', (
									SELECT * FROM city_object(a.id)
								),
								'price', (a.price),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
								'priceDiff', (
									a.price -
									(
										SELECT * FROM price_province_latest_avg_child(@provinceId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									    a.price-
										(
											SELECT * FROM price_province_latest_avg_child(@provinceId, @commodityId)
										)
									)
								)
							)
						)
					FROM (
						SELECT
							*
						FROM (
							SELECT
								*,
								@selectedDate selectedDate
							FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
						) b
						ORDER BY (
							SELECT (
								(
									SELECT price
									FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate) c
									WHERE c.id = b.id
								) -
								(
									SELECT * FROM price_province_latest_avg_child(@provinceId, @commodityId)
								)
							)
						)
					) a
				) "priceCity"
		`

	PriceGetCompareBySulselListWithOrderDesc = `
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
									WHERE last_update BETWEEN CONCAT(selectedDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(selectedDate, ' 23:59:59')::TIMESTAMPTZ
								),
								'client_id', 1,
								'city', (
									SELECT * FROM city_object(a.id)
								),
								'price', a.price,
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
								'priceDiff', (
									a.price -
									(
										SELECT * FROM price_province_latest(a.province_id, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									    a.price -
										(
											SELECT * FROM price_province_latest(a.province_id, @commodityId)
										)
									)
								)
							)
						)
					FROM (
						SELECT
							*
						FROM (
							SELECT
								*,
								@selectedDate selectedDate
							FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
						) b
						ORDER BY (
							SELECT (
								(
									SELECT price
									FROM get_level_harga(@provinceId, @commodityId, @selectedDate) c
									WHERE c.id = b.id
								) -
								(
									SELECT * FROM price_province_latest(province_id, @commodityId)
								)
							)
						) DESC
					) a
				) "priceCity"
		`

	PriceGetCompareBySulselListWithOrderDescAvgChild = `
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
							)
						)
					FROM price_province_avg_child(@commodityId, @provinceId) c
					WHERE idx = 1
				) "provincePrice",
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', null,
								'client_id', 1,
								'city', (
									SELECT * FROM city_object(a.id)
								),
								'price', (a.price),
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
								'priceDiff', (
									a.price -
									(
										SELECT * FROM price_province_latest_avg_child(@provinceId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
									    a.price-
										(
											SELECT * FROM price_province_latest_avg_child(@provinceId, @commodityId)
										)
									)
								)
							)
						)
					FROM (
						SELECT
							*
						FROM (
							SELECT
								*,
								@selectedDate selectedDate
							FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
						) b
						ORDER BY (
							SELECT (
								(
									SELECT price
									FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate) c
									WHERE c.id = b.id
								) -
								(
									SELECT * FROM price_province_latest_avg_child(@provinceId, @commodityId)
								)
							)
						) DESC
					) a
				) "priceCity"
		`

	PriceGetCompareBySulselCityHistory = `
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
										'price', a.price,
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
											(
												SELECT
													price
												FROM price_province(commodityId, provinceId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_province(commodityId, provinceId)
													WHERE idx = 1
												)
											)
										)
									)
								)
						)
					FROM (
						SELECT
						    *,
						    @commodityId::integer commodityId,
						    @provinceId::integer provinceId
						FROM price_city(@commodityId, @cityId, @startDate, @endDate)
						ORDER BY idx DESC
					) a
				) "priceDiff"
		`

	PriceGetCompareBySulselCityHistoryProvince = `
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
						WHERE a.id = @provinceId
				) city,
			    (
					SELECT
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
										'price', a.price,
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
											(
												SELECT
													price
												FROM price_province(commodityId, provinceId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_province(commodityId, provinceId)
													WHERE idx = 1
												)
											)
										)
									)
								)
						)
					FROM (
						SELECT
						    *,
						    @commodityId::integer commodityId,
						    @provinceId::integer provinceId
						FROM price_province(@commodityId, @provinceId, @startDate, @endDate)
						ORDER BY idx DESC
						LIMIT 1
					) a
				) "priceDiff"
		`

	PriceGetCompareBySulselCityHistoryAvgChild = `
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
										'price', a.price,
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
											(
												SELECT
													price
												FROM price_province_avg_child(commodityId, provinceId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_province_avg_child(commodityId, provinceId)
													WHERE idx = 1
												)
											)
										)
									)
								)
						)
					FROM (
						SELECT
						    *,
						    @commodityId::integer commodityId,
						    @provinceId::integer provinceId
						FROM price_city_avg_child(@commodityId, @cityId, @startDate, @endDate)
						ORDER BY idx DESC
					) a
				) "priceDiff"
		`

	PriceGetCompareBySulselCityHistoryProvinceAvgChild = `
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
						WHERE a.id = @provinceId
				) city,
			    (
					SELECT
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
										'price', a.price,
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
											(
												SELECT
													price
												FROM price_province_avg_child(commodityId, provinceId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_province_avg_child(commodityId, provinceId)
													WHERE idx = 1
												)
											)
										)
									)
								)
						)
					FROM (
						SELECT
						    *,
						    @commodityId::integer commodityId,
						    @provinceId::integer provinceId
						FROM price_province_avg_child(@commodityId, @provinceId, @startDate, @endDate)
					) a
				) "priceDiff"
		`

	PriceGetCompareBySulselCommodityHistoryHigher = `
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
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'city', (
										    SELECT * FROM city_object(a.id)
										),
										'price', a.price,
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
											(
												SELECT
													price
												FROM price_province(commodityId, provinceId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_province(commodityId, provinceId)
													WHERE idx = 1
												)
											)
										)
									)
								)
						)
					FROM (
						SELECT
							*,
							@commodityId::integer commodityId,
							@provinceId::integer provinceId,
							(
								price -
								(
									SELECT * FROM price_province_latest(@provinceId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga(@provinceId, @commodityId, @selectedDate) a
					) a
					WHERE priceDiff > 0
				) priceDiff
		`

	PriceGetCompareBySulselCommodityHistoryHigherAvgChild = `
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
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'city', (
										    SELECT * FROM city_object(a.id)
										),
										'price', a.price,
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
											(
												SELECT
													price
												FROM price_province_avg_child(commodityId, provinceId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_province_avg_child(commodityId, provinceId)
													WHERE idx = 1
												)
											)
										)
									)
								)
						)
					FROM (
						SELECT
							*,
							@commodityId::integer commodityId,
							@provinceId::integer provinceId,
							(
								price -
								(
									SELECT * FROM price_province_latest_avg_child(@provinceId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate) a
					) a
					WHERE priceDiff > 0
				) priceDiff
		`

	PriceGetCompareBySulselCommodityHistorySame = `
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
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'city', (
										    SELECT * FROM city_object(a.id)
										),
										'price', a.price,
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
											(
												SELECT
													price
												FROM price_province(commodityId, provinceId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_province(commodityId, provinceId)
													WHERE idx = 1
												)
											)
										)
									)
								)
						)
					FROM (
						SELECT
							*,
							@commodityId::integer commodityId,
							@provinceId::integer provinceId,
							(
								price -
								(
									SELECT * FROM price_province_latest(@provinceId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga(@provinceId, @commodityId, @selectedDate) a
					) a
					WHERE priceDiff = 0
				) priceDiff
		`

	PriceGetCompareBySulselCommodityHistorySameAvgChild = `
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
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'city', (
										    SELECT * FROM city_object(a.id)
										),
										'price', a.price,
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
											(
												SELECT
													price
												FROM price_province_avg_child(commodityId, provinceId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_province_avg_child(commodityId, provinceId)
													WHERE idx = 1
												)
											)
										)
									)
								)
						)
					FROM (
						SELECT
							*,
							@commodityId::integer commodityId,
							@provinceId::integer provinceId,
							(
								price -
								(
									SELECT * FROM price_province_latest_avg_child(@provinceId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate) a
					) a
					WHERE priceDiff = 0
				) priceDiff
		`

	PriceGetCompareBySulselCommodityHistoryLower = `
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
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'city', (
										    SELECT * FROM city_object(a.id)
										),
										'price', a.price,
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
											(
												SELECT
													price
												FROM price_province(commodityId, provinceId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_province(commodityId, provinceId)
													WHERE idx = 1
												)
											)
										)
									)
								)
						)
					FROM (
						SELECT
							*,
							@commodityId::integer commodityId,
							@provinceId::integer provinceId,
							(
								price -
								(
									SELECT * FROM price_province_latest(@provinceId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga(@provinceId, @commodityId, @selectedDate) a
					) a
					WHERE priceDiff < 0
				) priceDiff
		`

	PriceGetCompareBySulselCommodityHistoryLowerAvgChild = `
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
						(
							SELECT 
								JSON_AGG(
									JSON_BUILD_OBJECT(
										'city', (
										    SELECT * FROM city_object(a.id)
										),
										'price', a.price,
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
											(
												SELECT
													price
												FROM price_province_avg_child(commodityId, provinceId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_province_avg_child(commodityId, provinceId)
													WHERE idx = 1
												)
											)
										)
									)
								)
						)
					FROM (
						SELECT
							*,
							@commodityId::integer commodityId,
							@provinceId::integer provinceId,
							(
								price -
								(
									SELECT * FROM price_province_latest_avg_child(@provinceId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate) a
					) a
					WHERE priceDiff < 0
				) priceDiff
		`
)
