package price

const (
	PriceGetCompareByNational = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
							'higher', (
								SELECT * FROM price_higher_than_national(a.id, @provinceId, @nationalId, @selectedDate) 
							),
							'same', (
								SELECT * FROM price_same_as_national(a.id, @provinceId, @nationalId, @selectedDate) 
							),
							'lower', (
								SELECT * FROM price_lower_than_national(a.id, @provinceId, @nationalId, @selectedDate) 
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
							'id', commodity_price_national_id,
							'client_id', 1,
							'country', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', tn.id,
										'name', tn.name,
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
											WHERE id = tn.assets_relation_id
										),
										'createdAt', tn.created_at,
										'updatedAt', tn.updated_at,
										'deletedAt', tn.deleted_at
									)
								FROM tm_national tn 
								WHERE tn.id = 1
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							)
						)
					FROM price_national(@commodityId, 1) c
					WHERE idx = 1
				) "nationalPrice",
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
								SELECT * FROM price_province_compare_national(@commodityId, c.province_id, @nationalId, @selectedDate)
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
										SELECT * FROM price_national_latest(@nationalId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
										a.price -
										(
											SELECT * FROM price_national_latest(@nationalId, @commodityId)
										)
									)
								),
								'tier', (
									SELECT * FROM price_compare_national(@commodityId, a.id, a.province_id, @nationalId, selectedDate)
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

	PriceGetCompareByNationalAvgChild = `
			SELECT
				(
					SELECT
						JSON_BUILD_OBJECT(
							'higher', (
								SELECT * FROM price_higher_than_national_avg_child(a.id, @provinceId, @nationalId, @selectedDate) 
							),
							'same', (
								SELECT * FROM price_same_as_national_avg_child(a.id, @provinceId, @nationalId, @selectedDate) 
							),
							'lower', (
								SELECT * FROM price_lower_than_national_avg_child(a.id, @provinceId, @nationalId, @selectedDate) 
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
							'client_id', 1,
							'country', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', tn.id,
										'name', tn.name,
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
											WHERE id = tn.assets_relation_id
										),
										'createdAt', tn.created_at,
										'updatedAt', tn.updated_at,
										'deletedAt', tn.deleted_at
									)
								FROM tm_national tn 
								WHERE tn.id = 1
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							)
						)
					FROM price_national_avg_child(@commodityId, 1) c
					WHERE idx = 1
				) "nationalPrice",
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
								SELECT * FROM price_province_compare_national_avg_child(@commodityId, c.province_id, @nationalId, @selectedDate)
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
								'price', a.price,
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
							    'priceDiff', (
									a.price -
									(
										SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
										a.price -
										(
											SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
										)
									)
								),
								'tier', (
									SELECT * FROM price_compare_national_avg_child(@commodityId, a.id, @provinceId, @nationalId, selectedDate)
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

	PriceGetCompareByNationalList = `
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
							'id', commodity_price_national_id,
							'client_id', 1,
							'country', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', tn.id,
										'name', tn.name,
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
											WHERE id = tn.assets_relation_id
										),
										'createdAt', tn.created_at,
										'updatedAt', tn.updated_at,
										'deletedAt', tn.deleted_at
									)
								FROM tm_national tn 
								WHERE tn.id = @nationalId
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							)
						)
					FROM price_national(@commodityId, @nationalId) c
					WHERE idx = 1
				) "nationalPrice",
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
								c.price -
								(
									SELECT * FROM price_national_latest(@nationalId, @commodityId)
								)
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
									c.price -
									(
										SELECT * FROM price_national_latest(@nationalId, @commodityId)
									)
								)
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
										SELECT * FROM price_national_latest(@nationalId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
										a.price -
										(
											SELECT * FROM price_national_latest(@nationalId, @commodityId)
										)
									)
								)
							)
						)
					FROM (
						SELECT
							*,
							@selectedDate selectedDate
						FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
					) a
				) "cityPrice"
		`

	PriceGetCompareByNationalListAvgChild = `
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
							'client_id', 1,
							'country', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', tn.id,
										'name', tn.name,
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
											WHERE id = tn.assets_relation_id
										),
										'createdAt', tn.created_at,
										'updatedAt', tn.updated_at,
										'deletedAt', tn.deleted_at
									)
								FROM tm_national tn 
								WHERE tn.id = @nationalId
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							)
						)
					FROM price_national_avg_child(@commodityId, @nationalId) c
					WHERE idx = 1
				) "nationalPrice",
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
								c.price -
								(
									SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
								)
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
									c.price -
									(
										SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
									)
								)
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
								'price', a.price,
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
							    'priceDiff', (
									a.price -
									(
										SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
										a.price -
										(
											SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
										)
									)
								)
							)
						)
					FROM (
						SELECT
							*,
							@selectedDate selectedDate
						FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
					) a
				) "cityPrice"
		`

	PriceGetCompareByNationalListWithOrderAsc = `
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
							'id', commodity_price_national_id,
							'client_id', 1,
							'country', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', tn.id,
										'name', tn.name,
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
											WHERE id = tn.assets_relation_id
										),
										'createdAt', tn.created_at,
										'updatedAt', tn.updated_at,
										'deletedAt', tn.deleted_at
									)
								FROM tm_national tn 
								WHERE tn.id = @nationalId
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							)
						)
					FROM price_national(@commodityId, @nationalId) c
					WHERE idx = 1
				) "nationalPrice",
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
								c.price -
								(
									SELECT * FROM price_national_latest(@nationalId, @commodityId)
								)
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
									c.price -
									(
										SELECT * FROM price_national_latest(@nationalId, @commodityId)
									)
								)
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
										SELECT * FROM price_national_latest(@nationalId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
										a.price -
										(
											SELECT * FROM price_national_latest(@nationalId, @commodityId)
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
									SELECT * FROM price_national_latest(@nationalId, @commodityId)
								)
							)
						)
					) a
				) "cityPrice"
		`

	PriceGetCompareByNationalListWithOrderAscAvgChild = `
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
							'client_id', 1,
							'country', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', tn.id,
										'name', tn.name,
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
											WHERE id = tn.assets_relation_id
										),
										'createdAt', tn.created_at,
										'updatedAt', tn.updated_at,
										'deletedAt', tn.deleted_at
									)
								FROM tm_national tn 
								WHERE tn.id = @nationalId
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							)
						)
					FROM price_national_avg_child(@commodityId, @nationalId) c
					WHERE idx = 1
				) "nationalPrice",
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
								c.price -
								(
									SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
								)
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
									c.price -
									(
										SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
									)
								)
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
								'price', a.price,
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
							    'priceDiff', (
									a.price -
									(
										SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
										a.price -
										(
											SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
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
									SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
								)
							)
						)
					) a
				) "cityPrice"
		`

	PriceGetCompareByNationalListWithOrderDesc = `
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
							'id', commodity_price_national_id,
							'client_id', 1,
							'country', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', tn.id,
										'name', tn.name,
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
											WHERE id = tn.assets_relation_id
										),
										'createdAt', tn.created_at,
										'updatedAt', tn.updated_at,
										'deletedAt', tn.deleted_at
									)
								FROM tm_national tn 
								WHERE tn.id = @nationalId
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							)
						)
					FROM price_national(@commodityId, @nationalId) c
					WHERE idx = 1
				) "nationalPrice",
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
								c.price -
								(
									SELECT * FROM price_national_latest(@nationalId, @commodityId)
								)
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
									c.price -
									(
										SELECT * FROM price_national_latest(@nationalId, @commodityId)
									)
								)
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
										SELECT * FROM price_national_latest(@nationalId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
										a.price -
										(
											SELECT * FROM price_national_latest(@nationalId, @commodityId)
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
									SELECT * FROM price_national_latest(@nationalId, @commodityId)
								)
							)
						) DESC
					) a
				) "cityPrice"
		`

	PriceGetCompareByNationalListWithOrderDescAvgChild = `
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
							'client_id', 1,
							'country', (
								SELECT
									JSON_BUILD_OBJECT(
										'id', tn.id,
										'name', tn.name,
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
											WHERE id = tn.assets_relation_id
										),
										'createdAt', tn.created_at,
										'updatedAt', tn.updated_at,
										'deletedAt', tn.deleted_at
									)
								FROM tm_national tn 
								WHERE tn.id = @nationalId
							),
							'price', c.price,
							'priceRupiahFormat', (
							    SELECT rupiah_format(c.price)
							)
						)
					FROM price_national_avg_child(@commodityId, @nationalId) c
					WHERE idx = 1
				) "nationalPrice",
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
								c.price -
								(
									SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
								)
							),
							'priceDiffFormat', (
								SELECT rupiah_format(
									c.price -
									(
										SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
									)
								)
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
								'price', a.price,
								'priceRupiahFormat', 
								(
									SELECT rupiah_format(a.price)
								),
							    'priceDiff', (
									a.price -
									(
										SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
									)
								),
								'priceDiffFormat', (
									SELECT rupiah_format(
										a.price -
										(
											SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
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
									SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
								)
							)
						) DESC
					) a
				) "cityPrice"
		`

	PriceGetCompareByNationalCityHistory = `
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
												SELECT * FROM price_national_latest(@nationalId, @commodityId)
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT * FROM price_national_latest(@nationalId, @commodityId)
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

	PriceGetCompareByNationalCityHistoryAvgChild = `
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
												SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
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

	PriceGetCompareByNationalCityHistoryProvince = `
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
												FROM price_national(commodityId, nationalId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_national(commodityId, nationalId)
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
						    @nationalId::integer nationalId
						FROM price_national(@commodityId, @nationalId, @startDate, @endDate)
						ORDER BY idx DESC
						LIMIT 1
					) a
				) "priceDiff"
		`

	PriceGetCompareByNationalCityHistoryProvinceAvgChild = `
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
												FROM price_national_avg_child(commodityId, nationalId)
												WHERE idx = 1
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT
														price
													FROM price_national_avg_child(commodityId, nationalId)
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
						    @nationalId::integer nationalId
						FROM price_national_avg_child(@commodityId, @nationalId, @startDate, @endDate)
					) a
				) "priceDiff"
		`

	PriceGetCompareByNationalCommodityHistoryHigher = `
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
										'price', (a.price),
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
										    (
												SELECT * FROM price_national_latest(@nationalId, @commodityId)
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT * FROM price_national_latest(@nationalId, @commodityId)
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
									SELECT * FROM price_national_latest(@nationalId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
					) a
					WHERE priceDiff > 0
				) priceDiff
		`

	PriceGetCompareByNationalCommodityHistoryHigherAvgChild = `
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
										'price', (a.price),
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
										    (
												SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
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
									SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
					) a
					WHERE priceDiff > 0
				) priceDiff
		`

	PriceGetCompareByNationalCommodityHistorySame = `
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
										'price', (a.price),
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
										    (
												SELECT * FROM price_national_latest(@nationalId, @commodityId)
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT * FROM price_national_latest(@nationalId, @commodityId)
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
									SELECT * FROM price_national_latest(@nationalId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
					) a
					WHERE priceDiff = 0
				) priceDiff
		`

	PriceGetCompareByNationalCommodityHistorySameAvgChild = `
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
										'price', (a.price),
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
										    (
												SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
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
									SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
					) a
					WHERE priceDiff = 0
				) priceDiff
		`

	PriceGetCompareByNationalCommodityHistoryLower = `
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
										'price', (a.price),
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
										    (
												SELECT * FROM price_national_latest(@nationalId, @commodityId)
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT * FROM price_national_latest(@nationalId, @commodityId)
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
									SELECT * FROM price_national_latest(@nationalId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga(@provinceId, @commodityId, @selectedDate)
					) a
					WHERE priceDiff < 0
				) priceDiff
		`

	PriceGetCompareByNationalCommodityHistoryLowerAvgChild = `
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
										'price', (a.price),
										'priceRupiahFormat', (
											SELECT rupiah_format(a.price)
										),
										'priceDiff', (
											a.price -
										    (
												SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
											)
										),
										'priceDiffFormat', (
											SELECT rupiah_format(
												a.price -
												(
													SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
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
									SELECT * FROM price_national_latest_avg_child(@nationalId, @commodityId)
								)
							) priceDiff
						FROM get_level_harga_avg_child(@provinceId, @commodityId, @selectedDate)
					) a
					WHERE priceDiff < 0
				) priceDiff
		`
)
