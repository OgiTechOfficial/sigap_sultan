package neraca

const (
	NeracaStokAkhirByCityMap = `
		SELECT
			'ton' "unit",
			(
			    SELECT * FROM city_object(@cityId)
			) "city",
			(
			    SELECT
					JSON_BUILD_OBJECT(
						'defisit', (
							SELECT neraca_by_city_defisit(@cityId, @startDate, @endDate) 
						),
						'rentan', (
							SELECT neraca_by_city_rentan(@cityId, @startDate, @endDate)
						),
						'waspada', (
							SELECT neraca_by_city_waspada(@cityId, @startDate, @endDate)
						),
						'aman', (
							SELECT neraca_by_city_aman(@cityId, @startDate, @endDate)
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
						'tier', (
						    CASE
						        WHEN c.kebutuhan = 0
								THEN 'aman'
							ELSE
								CASE
									WHEN (c.stock / c.kebutuhan) >= 81 THEN 'aman'
									WHEN
										 (c.stock / c.kebutuhan) <= 80 and
										 (
												CASE
													WHEN kebutuhan = 0
													THEN 100
													ELSE
														CASE
															WHEN (c.stock / kebutuhan)*100 >= 46.0 AND (c.stock / kebutuhan)*100 <= 47.0
															THEN ROUND(c.stock::NUMERIC, 2)
															ELSE ROUND(CAST((c.stock / kebutuhan)*100 AS NUMERIC), 2)
														END
												END
										)>= 47::double precision
									THEN 'waspada'
									WHEN (c.stock / c.kebutuhan) >= 0 and (c.stock / c.kebutuhan) <= 46 THEN 'rentan'
									ELSE 'defisit'
									END 
							END
						)
					)
				FROM get_level_stock_province_cr(@commodityId, @startDate, @endDate) c
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
								FROM stock_akhir_city(tc.id, @cityId, @startDate, @endDate)
								LIMIT 1
							),
							'tier', (
								SELECT tier
								FROM stock_akhir_city(tc.id, @cityId, @startDate, @endDate)
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
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			)"stockTier",
			(
				SELECT '["defisit","rentan","waspada","aman"]'::JSON
			) "stockTierCode"
	`

	NeracaStokAkhirByCityMapProvince = `
		SELECT
			'ton' "unit",
			(
			    SELECT * FROM city_object(@cityId)
			) "city",
			(
			    SELECT
					JSON_BUILD_OBJECT(
						'defisit', (
							SELECT neraca_by_city_defisit_province(@cityId, @startDate, @endDate) 
						),
						'rentan', (
							SELECT neraca_by_city_rentan_province(@cityId, @startDate, @endDate)
						),
						'waspada', (
							SELECT neraca_by_city_waspada_province(@cityId, @startDate, @endDate)
						),
						'aman', (
							SELECT neraca_by_city_aman_province(@cityId, @startDate, @endDate)
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
						'tier', (
						    CASE
						        WHEN c.kebutuhan = 0
								THEN 'aman'
							ELSE
								CASE
									WHEN (c.stock / c.kebutuhan) >= 81 THEN 'aman'
									WHEN
										 (c.stock / c.kebutuhan) <= 80 and
										 (
												CASE
													WHEN kebutuhan = 0
													THEN 100
													ELSE
														CASE
															WHEN (c.stock / kebutuhan)*100 >= 46.0 AND (c.stock / kebutuhan)*100 <= 47.0
															THEN ROUND(c.stock::NUMERIC, 2)
															ELSE ROUND(CAST((c.stock / kebutuhan)*100 AS NUMERIC), 2)
														END
												END
										)>= 47::double precision
									THEN 'waspada'
									WHEN (c.stock / c.kebutuhan) >= 0 and (c.stock / c.kebutuhan) <= 46 THEN 'rentan'
									ELSE 'defisit'
									END 
							END
						)
					)
				FROM get_level_stock_province_cr(@commodityId, @startDate, @endDate) c
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
								FROM stock_akhir_province(tc.id, @cityId, @startDate, @endDate)
								LIMIT 1
							),
							'tier', (
								SELECT tier
								FROM stock_akhir_province(tc.id, @cityId, @startDate, @endDate)
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
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			)"stockTier",
			(
				SELECT '["defisit","rentan","waspada","aman"]'::JSON
			) "stockTierCode"
	`

	NeracaStokAkhirListByCity = `
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
				FROM get_level_stock_province_cr(@commodityId, @startDate, @endDate) c
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

	NeracaStokAkhirListByCityAsc = `
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
				FROM get_level_stock_province_cr(@commodityId, @startDate, @endDate) c
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

	NeracaStokAkhirListByCityDesc = `
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
				FROM get_level_stock_province_cr(@commodityId, @startDate, @endDate) c
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
