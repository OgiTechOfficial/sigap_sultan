package neraca

const (
	//NeracaKebutuhanByCityMap = `
	//	SELECT
	//		'ton' "unit",
	//		(
	//		    SELECT * FROM city_object(@cityId)
	//		) "city",
	//		(
	//		    SELECT
	//				JSON_BUILD_OBJECT(
	//					'defisit', (
	//						SELECT neraca_defisit_cr(a.id, @provinceId, @startDate, @endDate)
	//					),
	//					'rentan', (
	//						SELECT neraca_rentan_cr(a.id, @provinceId, @startDate, @endDate)
	//					),
	//					'waspada', (
	//						SELECT neraca_waspada_cr(a.id, @provinceId, @startDate, @endDate)
	//					),
	//					'aman', (
	//						SELECT neraca_aman_cr(a.id, @provinceId, @startDate, @endDate)
	//					)
	//				)
	//			FROM tm_commodity a
	//			WHERE a.id = 1
	//		) "summary",
	//		(
	//			SELECT
	//				JSON_BUILD_OBJECT(
	//					'id', province_id,
	//					'clientId', client_id,
	//					'province', (
	//						SELECT * FROM province_object(@provinceId)
	//					),
	//					'stock', c.stock,
	//					'tier', (
	//						CASE
	//							WHEN (c.stock / c.kebutuhan) >= 81 THEN 'aman'
	//							WHEN (c.stock / c.kebutuhan) <= 80 and (c.stock / c.kebutuhan) >= 47 THEN 'waspada'
	//							WHEN (c.stock / c.kebutuhan) >= 0 and (c.stock / c.kebutuhan) <= 46 THEN 'rentan'
	//							ELSE 'defisit'
	//							END
	//					)
	//				)
	//			FROM get_level_kebutuhan_province_cr(@commodityId, @startDate, @endDate) c
	//			WHERE province_id = @provinceId
	//		) "provinceStock",
	//		(
	//			SELECT
	//				JSON_AGG(
	//					JSON_BUILD_OBJECT(
	//						'id', NULL,
	//						'commodity', JSON_BUILD_OBJECT(
	//							'id', tc.id,
	//							'name', tc.name,
	//							'assets', (
	//								SELECT
	//									JSON_BUILD_OBJECT(
	//										'id', id,
	//										'assetsType', assets_type,
	//										'assetsLocation', assets_location,
	//										'assetsLocationType', assets_location_type,
	//										'assetsMediaType', assets_media_type,
	//										'assetsExt', assets_ext,
	//										'assetsName', assets_name
	//									)
	//								FROM assets
	//								WHERE id = tc.assets_relation_id
	//							)
	//						),
	//						'stock', (
	//							SELECT ROUND(stock::NUMERIC, 0)
	//							FROM get_level_kebutuhan_cr(@provinceId, @commodityId, @startDate, @endDate) a
	//							LIMIT 1
	//						),
	//						'tier', (
	//							SELECT tier
	//							FROM stock_akhir_city(tc.id, @cityId)
	//							LIMIT 1
	//						)
	//					)
	//				)
	//			FROM (
	//				SELECT *
	//				FROM tm_commodity tc
	//				WHERE tc.parent_id IS NULL
	//				ORDER BY "sequence"
	//			) tc
	//		) "commodityStock",
	//		(
	//			SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
	//		)"stockTier",
	//		(
	//			SELECT '["defisit","rentan","waspada","aman"]'::JSON
	//		) "stockTierCode"
	//`
	//
	//NeracaKebutuhanByCityMapProvince = `
	//	SELECT
	//		'ton' "unit",
	//		(
	//		    SELECT * FROM city_object(@cityId)
	//		) "city",
	//		(
	//		    SELECT
	//				JSON_BUILD_OBJECT(
	//					'defisit', (
	//						SELECT neraca_defisit_cr(a.id, @provinceId, @startDate, @endDate)
	//					),
	//					'rentan', (
	//						SELECT neraca_rentan_cr(a.id, @provinceId, @startDate, @endDate)
	//					),
	//					'waspada', (
	//						SELECT neraca_waspada_cr(a.id, @provinceId, @startDate, @endDate)
	//					),
	//					'aman', (
	//						SELECT neraca_aman_cr(a.id, @provinceId, @startDate, @endDate)
	//					)
	//				)
	//			FROM tm_commodity a
	//			WHERE a.id = 1
	//		) "summary",
	//		(
	//			SELECT
	//				JSON_BUILD_OBJECT(
	//					'id', province_id,
	//					'clientId', client_id,
	//					'province', (
	//						SELECT * FROM province_object(@provinceId)
	//					),
	//					'stock', c.stock,
	//					'tier', (
	//						CASE
	//							WHEN (c.stock / c.stock) >= 81 THEN 'aman'
	//							WHEN (c.stock / c.stock) <= 80 and (c.stock / c.stock) >= 47 THEN 'waspada'
	//							WHEN (c.stock / c.stock) >= 0 and (c.stock / c.stock) <= 46 THEN 'rentan'
	//							ELSE 'defisit'
	//							END
	//					)
	//				)
	//			FROM get_level_kebutuhan_province_cr(@commodityId, @startDate, @endDate) c
	//			WHERE province_id = @provinceId
	//		) "provinceStock",
	//		(
	//			SELECT
	//				JSON_AGG(
	//					JSON_BUILD_OBJECT(
	//						'id', NULL,
	//						'commodity', JSON_BUILD_OBJECT(
	//							'id', tc.id,
	//							'name', tc.name,
	//							'assets', (
	//								SELECT
	//									JSON_BUILD_OBJECT(
	//										'id', id,
	//										'assetsType', assets_type,
	//										'assetsLocation', assets_location,
	//										'assetsLocationType', assets_location_type,
	//										'assetsMediaType', assets_media_type,
	//										'assetsExt', assets_ext,
	//										'assetsName', assets_name
	//									)
	//								FROM assets
	//								WHERE id = tc.assets_relation_id
	//							)
	//						),
	//						'stock', (
	//							SELECT ROUND(stock::NUMERIC, 0)
	//							FROM get_level_kebutuhan_province_cr(tc.id, @startDate, @endDate)
	//							WHERE province_id = @cityId
	//							LIMIT 1
	//						),
	//						'tier', (
	//							SELECT tier
	//							FROM get_level_kebutuhan_province_cr(tc.id, @startDate, @endDate)
	//							WHERE province_id = @cityId
	//							LIMIT 1
	//						)
	//					)
	//				)
	//			FROM (
	//				SELECT *
	//				FROM tm_commodity tc
	//				WHERE tc.parent_id IS NULL
	//				ORDER BY "sequence"
	//			) tc
	//		) "commodityStock",
	//		(
	//			SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
	//		)"stockTier",
	//		(
	//			SELECT '["defisit","rentan","waspada","aman"]'::JSON
	//		) "stockTierCode"
	//`

	NeracaKebutuhanByCityMap = `
		SELECT
			'ton' "unit",
			(
			    SELECT * FROM city_object(@cityId)
			) "city",
			(
			    SELECT
					JSON_BUILD_OBJECT(
						'meningkat', (
							SELECT kebutuhan_by_city_meningkat(@cityId, @startDate, @endDate) 
						),
						'menurun', (
							SELECT kebutuhan_by_city_menurun(@cityId, @startDate, @endDate)
						),
						'stabil', (
							SELECT kebutuhan_by_city_stabil(@cityId, @startDate, @endDate)
						)
					)
			) "summary",
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', id,
						'clientId', client_id,
						'province', (
							SELECT * FROM province_object(@provinceId)
						),
						'stock', c.stock,
						'tier', c.tier
					)
				FROM get_level_kebutuhan_city_province(@commodityId, @provinceId, @startDate, @endDate) c
			) "provinceStock",
			(
				SELECT 
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', NULL,
							'commodity', JSON_BUILD_OBJECT(
								'id', tc.id,
								'name', tc.name,
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
									WHERE id = tc.assets_relation_id
								)
							),
							'stock', (
								SELECT ROUND(stock::NUMERIC, 0)
								FROM get_level_kebutuhan_city(tc.id, @cityId, @startDate, @endDate)
								LIMIT 1
							),
							'tier', (
								SELECT tier
								FROM get_level_kebutuhan_city(tc.id, @cityId, @startDate, @endDate)
								LIMIT 1
							)
						)
					)
				FROM (
					SELECT *
					FROM tm_commodity tc
					WHERE tc.parent_id IS NULL
					ORDER BY "sequence"
				) tc
			) "commodityStock",
			(
				SELECT '{"menurun":{"title":"Menurun","color":"#FF6711"},"stabil":{"title":"Stabil","color":"#32D583"},"meningkat":{"title":"Meningkat","color":"#05603A"}}'::JSON
			) "stockTier",
			(
				SELECT '["menurun","stabil","meningkat"]'::JSON
			) "stockTierCode"
	`

	NeracaKebutuhanByCityMapProvince = `
		SELECT
			'ton' "unit",
			(
			    SELECT * FROM city_object(@cityId)
			) "city",
			(
			    SELECT
					JSON_BUILD_OBJECT(
						'meningkat', (
							SELECT kebutuhan_by_city_meningkat_province(@cityId, @startDate, @endDate) 
						),
						'menurun', (
							SELECT kebutuhan_by_city_menurun_province(@cityId, @startDate, @endDate)
						),
						'stabil', (
							SELECT kebutuhan_by_city_stabil_province(@cityId, @startDate, @endDate)
						)
					)
			) "summary",
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', province_id,
						'clientId', client_id,
						'province', (
							SELECT * FROM province_object(@provinceId)
						),
						'stock', c.stock,
						'tier', c.tier
					)
				FROM get_level_kebutuhan_province_cr(@commodityId, @startDate, @endDate) c
				WHERE province_id = @provinceId
			) "provinceStock",
			(
				SELECT 
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', NULL,
							'commodity', JSON_BUILD_OBJECT(
								'id', tc.id,
								'name', tc.name,
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
									WHERE id = tc.assets_relation_id
								)
							),
							'stock', (
								SELECT ROUND(stock::NUMERIC, 0)
								FROM get_level_kebutuhan_city_province(tc.id, @provinceId, @startDate, @endDate)
								LIMIT 1
							),
							'tier', (
								SELECT tier
								FROM get_level_kebutuhan_city_province(tc.id, @provinceId, @startDate, @endDate)
								LIMIT 1
							)
						)
					)
				FROM (
					SELECT *
					FROM tm_commodity tc
					WHERE tc.parent_id IS NULL
					ORDER BY "sequence"
				) tc
			) "commodityStock",
			(
				SELECT '{"menurun":{"title":"Menurun","color":"#FF6711"},"stabil":{"title":"Stabil","color":"#32D583"},"meningkat":{"title":"Meningkat","color":"#05603A"}}'::JSON
			) "stockTier",
			(
				SELECT '["menurun","stabil","meningkat"]'::JSON
			) "stockTierCode"
	`

	NeracaKebutuhanListByCity = `
		SELECT
			'ton' "unit",
		    (
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
		    ) "commodity",
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', province_id,
						'clientId', client_id,
						'province', (
							SELECT * FROM province_object(@provinceId)
						),
						'stock', c.stock,
						'tier', (
							CASE
								WHEN (c.stock / c.kebutuhan) >= 81 THEN 'aman'
								WHEN (c.stock / c.kebutuhan) <= 80 and (c.stock / c.kebutuhan) >= 47 THEN 'waspada'
								WHEN (c.stock / c.kebutuhan) >= 0 and (c.stock / c.kebutuhan) <= 46 THEN 'rentan'
								ELSE 'defisit'
								END 
						)
					)
				FROM get_level_kebutuhan_province_cr(@commodityId, @startDate, @endDate) c
				WHERE province_id = @provinceId
			) "provinceStock",
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', (
							CASE
								WHEN (a.stock / a.kebutuhan)*100 >= 81 THEN 'aman'
								WHEN ((a.stock / a.kebutuhan)*100 <= 80 and (a.stock / a.kebutuhan)*100 >= 47) THEN 'waspada'
								WHEN ((a.stock / a.kebutuhan)*100 >= 0 and (a.stock / a.kebutuhan)*100 <= 46) THEN 'rentan'
								ELSE 'defisit'
								END )
						)
					)
				FROM (
				    SELECT
						*,
						(b.stock / b.kebutuhan)*100 "aaa"
					FROM get_level_stock(@provinceId, @commodityId, @selectedDate) b
					OFFSET @page
					LIMIT @limit
				) a
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			)"stockTier"
	`

	NeracaKebutuhanListByCityAsc = `
		SELECT
			'ton' "unit",
		    (
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
		    ) "commodity",
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', province_id,
						'clientId', client_id,
						'province', (
							SELECT * FROM province_object(@provinceId)
						),
						'stock', c.stock,
						'tier', (
							CASE
								WHEN (c.stock / c.kebutuhan) >= 81 THEN 'aman'
								WHEN (c.stock / c.kebutuhan) <= 80 and (c.stock / c.kebutuhan) >= 47 THEN 'waspada'
								WHEN (c.stock / c.kebutuhan) >= 0 and (c.stock / c.kebutuhan) <= 46 THEN 'rentan'
								ELSE 'defisit'
								END 
						)
					)
				FROM get_level_kebutuhan_province_cr(@commodityId, @startDate, @endDate) c
				WHERE province_id = @provinceId
			) "provinceStock",
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', (
							CASE
								WHEN (a.stock / a.kebutuhan)*100 >= 81 THEN 'aman'
								WHEN ((a.stock / a.kebutuhan)*100 <= 80 and (a.stock / a.kebutuhan)*100 >= 47) THEN 'waspada'
								WHEN ((a.stock / a.kebutuhan)*100 >= 0 and (a.stock / a.kebutuhan)*100 <= 46) THEN 'rentan'
								ELSE 'defisit'
								END )
						)
					)
				FROM (
				    SELECT
						*,
						(b.stock / b.kebutuhan)*100 "aaa"
					FROM get_level_stock(@provinceId, @commodityId, @selectedDate) b
					ORDER BY aaa
					OFFSET @page
					LIMIT @limit
				) a
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			)"stockTier"
	`

	NeracaKebutuhanListByCityDesc = `
		SELECT
			'ton' "unit",
		    (
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
		    ) "commodity",
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', province_id,
						'clientId', client_id,
						'province', (
							SELECT * FROM province_object(@provinceId)
						),
						'stock', c.stock,
						'tier', (
							CASE
								WHEN (c.stock / c.kebutuhan) >= 81 THEN 'aman'
								WHEN (c.stock / c.kebutuhan) <= 80 and (c.stock / c.kebutuhan) >= 47 THEN 'waspada'
								WHEN (c.stock / c.kebutuhan) >= 0 and (c.stock / c.kebutuhan) <= 46 THEN 'rentan'
								ELSE 'defisit'
								END 
						)
					)
				FROM get_level_kebutuhan_province_cr(@commodityId, @startDate, @endDate) c
				WHERE province_id = @provinceId
			) "provinceStock",
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', (
							CASE
								WHEN (a.stock / a.kebutuhan)*100 >= 81 THEN 'aman'
								WHEN ((a.stock / a.kebutuhan)*100 <= 80 and (a.stock / a.kebutuhan)*100 >= 47) THEN 'waspada'
								WHEN ((a.stock / a.kebutuhan)*100 >= 0 and (a.stock / a.kebutuhan)*100 <= 46) THEN 'rentan'
								ELSE 'defisit'
								END )
						)
					)
				FROM (
				    SELECT
						*,
						(b.stock / b.kebutuhan)*100 "aaa"
					FROM get_level_stock(@provinceId, @commodityId, @selectedDate) b
					ORDER BY aaa DESC
					OFFSET @page
					LIMIT @limit
				) a
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			)"stockTier"
	`
)
