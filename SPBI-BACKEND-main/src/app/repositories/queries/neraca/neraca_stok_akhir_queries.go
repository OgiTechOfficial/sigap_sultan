package neraca

const (
	NeraacStokAkhirSummary = `
		SELECT
			JSON_BUILD_OBJECT(
				'defisit', (
					SELECT neraca_defisit_cr(a.id, @provinceId, @startDate, @endDate) 
				),
				'rentan', (
					SELECT neraca_rentan_cr(a.id, @provinceId, @startDate, @endDate)
				),
				'waspada', (
					SELECT neraca_waspada_cr(a.id, @provinceId, @startDate, @endDate)
				),
				'aman', (
					SELECT neraca_aman_cr(a.id, @provinceId, @startDate, @endDate)
				)
			)
		FROM tm_commodity a
		WHERE a.id = @commodityId
	`

	NeracaStokAkhirMap = `
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
							'stock', COALESCE(a.stock, ROUND(a.stock::NUMERIC, 0), 0),
							'tier', (
							    CASE
									WHEN a.kebutuhan = 0
									THEN 'rentan'
								ELSE
									CASE
										WHEN (a.stock / a.kebutuhan)*100 >= 81 THEN 'aman'
										WHEN
											((a.stock / a.kebutuhan)*100 <= 80 AND
-- 											(a.stock / a.kebutuhan)*100 >= 47)
											(
												CASE
													WHEN (a.stock / a.kebutuhan)*100 >= 46.0 AND (a.stock / a.kebutuhan)*100 <= 47.0
													THEN ROUND(a.stock::NUMERIC, 2)
													ELSE ROUND(CAST((a.stock / kebutuhan)*100 AS NUMERIC), 2)
												END
											) >= 47::double precision)
										    THEN 'waspada'
										WHEN ((a.stock / a.kebutuhan)*100 >= 0 and (a.stock / a.kebutuhan)*100 <= 46) THEN 'rentan'
									    WHEN ((a.stock / a.kebutuhan)*100 < 0) THEN 'defisit'
										ELSE null
-- 										ELSE 'defisit'
									END
							    END
							)
						)
					)
				FROM get_level_stock_cr(@provinceId, @commodityId, @startDate, @endDate) a
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			)"stockTier",
			(
				SELECT '["defisit","rentan","waspada","aman"]'::JSON
			) "stockTierCode"
	`

	NeracaStokAkhirListByCommodity = `
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
									WHEN ((c.stock / c.kebutuhan)*100 < 0) THEN 'defisit'
									ELSE null
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
-- 							'tier', (
-- 							CASE
-- 								WHEN (a.stock / a.kebutuhan)*100 >= 81 THEN 'aman'
-- 								WHEN ((a.stock / a.kebutuhan)*100 <= 80 and (a.stock / a.kebutuhan)*100 >= 47) THEN 'waspada'
-- 								WHEN ((a.stock / a.kebutuhan)*100 >= 0 and (a.stock / a.kebutuhan)*100 <= 46) THEN 'rentan'
-- 								ELSE 'defisit'
-- 								END )
							'tier', (
							    CASE
									WHEN a.kebutuhan = 0
									THEN 'rentan'
								ELSE
									CASE
										WHEN (a.stock / a.kebutuhan)*100 >= 81 THEN 'aman'
										WHEN
											((a.stock / a.kebutuhan)*100 <= 80 AND
-- 											(a.stock / a.kebutuhan)*100 >= 47)
											(
												CASE
													WHEN (a.stock / a.kebutuhan)*100 >= 46.0 AND (a.stock / a.kebutuhan)*100 <= 47.0
													THEN ROUND(a.stock::NUMERIC, 2)
													ELSE ROUND(CAST((a.stock / kebutuhan)*100 AS NUMERIC), 2)
												END
											) >= 47::double precision)
										    THEN 'waspada'
										WHEN ((a.stock / a.kebutuhan)*100 >= 0 and (a.stock / a.kebutuhan)*100 <= 46) THEN 'rentan'
										WHEN ((a.stock / a.kebutuhan)*100 < 0) THEN 'defisit'
										ELSE null
									END
							    END
							)
						)
					)
				FROM (
				    SELECT
						*,
						(b.stock / b.kebutuhan)*100 "aaa"
					FROM get_level_stock_cr(@provinceId, @commodityId, @startDate, @endDate) b
					OFFSET @page
					LIMIT @limit
				) a
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			)"stockTier"
	`

	NeracaStokAkhirListByCommodityAsc = `
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
									WHEN ((c.stock / c.kebutuhan)*100 < 0) THEN 'defisit'
									ELSE null
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
									WHEN
										(a.stock / a.kebutuhan)*100 <= 80 and
										(
											CASE
												WHEN kebutuhan = 0
												THEN 100
												ELSE
													CASE
														WHEN (a.stock / kebutuhan)*100 >= 46.0 AND (a.stock / kebutuhan)*100 <= 47.0
														THEN ROUND(a.stock::NUMERIC, 2)
														ELSE ROUND(CAST((a.stock / kebutuhan)*100 AS NUMERIC), 2)
													END
											END
										)>= 47::double precision
									THEN 'waspada'
									WHEN ((a.stock / a.kebutuhan)*100 >= 0 and (a.stock / a.kebutuhan)*100 <= 46) THEN 'rentan'
									WHEN ((a.stock / a.kebutuhan)*100 < 0) THEN 'defisit'
									ELSE NULL
								END
							)
						)
					)
				FROM (
				    SELECT
						*,
						(b.stock / b.kebutuhan)*100 "aaa"
					FROM get_level_stock_cr(@provinceId, @commodityId, @startDate, @endDate) b
					ORDER BY b.stock
					OFFSET @page
					LIMIT @limit
				) a
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			)"stockTier"
	`

	NeracaStokAkhirListByCommodityDesc = `
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
									WHEN ((c.stock / c.kebutuhan)*100 < 0) THEN 'defisit'
									ELSE null
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
									WHEN
										(a.stock / a.kebutuhan)*100 <= 80 and
										(
											CASE
												WHEN kebutuhan = 0
												THEN 100
												ELSE
													CASE
														WHEN (a.stock / kebutuhan)*100 >= 46.0 AND (a.stock / kebutuhan)*100 <= 47.0
														THEN ROUND(a.stock::NUMERIC, 2)
														ELSE ROUND(CAST((a.stock / kebutuhan)*100 AS NUMERIC), 2)
													END
											END
										)>= 47::double precision
									THEN 'waspada'
									WHEN ((a.stock / a.kebutuhan)*100 >= 0 and (a.stock / a.kebutuhan)*100 <= 46) THEN 'rentan'
									WHEN ((a.stock / a.kebutuhan)*100 < 0) THEN 'defisit'
									ELSE NULL
								END
							)
						)
					)
				FROM (
				    SELECT *
					FROM(
						SELECT
							*,
							(b.stock / b.kebutuhan)*100 "aaa"
						FROM get_level_stock_cr(@provinceId, @commodityId, @startDate, @endDate) b
						WHERE b.stock IS NOT NULL
						ORDER BY b.stock DESC
					) a2
					UNION ALL
					SELECT
						*,
						(b.stock / b.kebutuhan)*100 "aaa"
					FROM get_level_stock_cr(@provinceId, @commodityId, @startDate, @endDate) b
					WHERE b.stock IS NULL
					OFFSET @page
					LIMIT @limit
				) a
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			)"stockTier"
	`

	NeracaStokAkhirCityHistory = `
		SELECT
			'ton' "unit",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
			    SELECT * FROM city_object(@cityId)
			) "city",
      (
          SELECT
            JSON_AGG(
              JSON_BUILD_OBJECT(
                'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
                'stock', stock,
                'stockFormat', thousand_format(stock::numeric)
              )
            )
          FROM(
              SELECT *
              FROM stock_akhir_city(@commodityId, @cityId, @startDate, @endDate) a
              ORDER BY last_update
          ) a
      ) "stock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaStokAkhirCommodityHistoryDefisit = `
		SELECT
			'ton' "unit",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT
				    JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_id,
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', a.tier
						)
					)
				FROM get_level_stock_cr(@provinceId, @commodityId, @startDate, @endDate) a
				WHERE tier = 'defisit'
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaStokAkhirCommodityHistoryRentan = `
		SELECT
			'ton' "unit",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT
				    JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_id,
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', a.tier
						)
					)
				FROM get_level_stock_cr(@provinceId, @commodityId, @startDate, @endDate) a
				WHERE tier = 'rentan'
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaStokAkhirCommodityHistoryWaspada = `
		SELECT
			'ton' "unit",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT
				    JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_id,
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', a.tier
						)
					)
				FROM get_level_stock_cr(@provinceId, @commodityId, @startDate, @endDate) a
				WHERE tier = 'waspada'
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaStokAkhirCommodityHistoryAman = `
		SELECT
			'ton' "unit",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT
				    JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_id,
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', a.tier
						)
					)
				FROM get_level_stock_cr(@provinceId, @commodityId, @startDate, @endDate) a
				WHERE tier = 'aman'
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaGetLatestDateAvail = `
		SELECT
			CONCAT(
				CONCAT(
					TO_CHAR(last_update, 'YYYY'),
					CONCAT('-', REPLACE(TO_CHAR(last_update, 'MM'), ' ','') )
				),
				CONCAT('-', TO_CHAR(last_update, 'DD'))
			)
		FROM get_level_stock(@provinceId, @commodityId, @selectedDate)
		WHERE id = @cityId
	`

	NeracaGetLatestDateAvailProvince = `
		SELECT
			CONCAT(
				CONCAT(
					TO_CHAR(last_update, 'YYYY'),
					CONCAT('-', REPLACE(TO_CHAR(last_update, 'MM'), ' ','') )
				),
				CONCAT('-', TO_CHAR(last_update, 'DD'))
			)
		FROM get_level_stock_province_cr(@commodityId, @startDate, @endDate)
		WHERE province_id = @cityId
	`

	NeracaStokAkhirByCommodityCityHistory = `
		SELECT
			'ton' "unit",
			'percentage' "unitDiff",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT * FROM city_object(@cityId)
			) "city",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'period', CONCAT(REPLACE(TO_CHAR(last_update, 'Month'), ' ',''), CONCAT(' ', TO_CHAR(last_update, 'YYYY'))),
							'neraca', ROUND(stock::NUMERIC, 0),
							'ketersediaan', ROUND(ketersediaan::NUMERIC, 0),
							'produksi', (
						    	SELECT ROUND((70/10*ketersediaan/10)::numeric, 0)
							),
							'kebutuhan', ROUND(kebutuhan::NUMERIC, 0),
							'tier', tier
						)
					)
				FROM (
				    SELECT
				        id,
						name,
						assets_relation_id,
						province_id,
						max(commodity_stock_id) commodity_stock_id,
						client_id,
						commodity_id,
						stock,
						tier,
						kebutuhan,
						ketersediaan,
						last_update
					FROM stock_akhir_city(@commodityId, @cityId)
					GROUP BY
						id,
						name,
						assets_relation_id,
						province_id,
						client_id,
						commodity_id,
						stock,
						tier,
						kebutuhan,
						ketersediaan,
						last_update
					ORDER BY last_update DESC
				) b
			) "stock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaStokAkhirByCommodityProvinceHistory = `
		SELECT
			'ton' "unit",
			'percentage' "unitDiff",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT * FROM city_object(@cityId)
			) "city",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'period', CONCAT(REPLACE(TO_CHAR(last_update, 'Month'), ' ',''), CONCAT(' ', TO_CHAR(last_update, 'YYYY'))),
							'neraca', ROUND(stock::NUMERIC, 0),
							'ketersediaan', ROUND(ketersediaan::NUMERIC, 0),
							'produksi', (
						    	SELECT ROUND((70/10*ketersediaan/10)::numeric, 0)
							),
							'kebutuhan', ROUND(kebutuhan::NUMERIC, 0),
							'tier', tier
						)
					)
				FROM (
				    SELECT
				        id,
						name,
						assets_relation_id,
						max(commodity_stock_id) commodity_stock_id,
						client_id,
						commodity_id,
						stock,
						tier,
						kebutuhan,
						ketersediaan,
						last_update
					FROM stock_akhir_province(@commodityId, @cityId)
					GROUP BY
						id,
						name,
						assets_relation_id,
						client_id,
						commodity_id,
						stock,
						tier,
						kebutuhan,
						ketersediaan,
						last_update
					ORDER BY last_update DESC
				) b
			) "stock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaStokAkhirByCommodityCityHistoryChart = `
		SELECT
			'ton' "unit",
			'percentage' "unitDiff",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT * FROM city_object(@cityId)
			) "city",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'period', CONCAT(REPLACE(TO_CHAR(last_update, 'Month'), ' ',''), CONCAT(' ', TO_CHAR(last_update, 'YYYY'))),
							'neraca', stock,
							'ketersediaan', ketersediaan,
							'produksi', (
						    	SELECT 70/10*ketersediaan::float8
							),
							'kebutuhan', kebutuhan,
							'tier', tier
						)
					)
				FROM(
				    SELECT *
					FROM stock_akhir_city(@commodityId, @cityId, @startDate, @endDate)
					ORDER BY last_update
				) a
			) "stock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaStokAkhirByCommodityProvinceHistoryChart = `
		SELECT
			'ton' "unit",
			'percentage' "unitDiff",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT * FROM city_object(@cityId)
			) "city",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'period', CONCAT(REPLACE(TO_CHAR(last_update, 'Month'), ' ',''), CONCAT(' ', TO_CHAR(last_update, 'YYYY'))),
							'neraca', stock,
							'ketersediaan', ketersediaan,
							'produksi', (
						    	SELECT 70/10*ketersediaan::float8
							),
							'kebutuhan', kebutuhan,
							'tier', tier
						)
					)
				FROM(
				    SELECT *
					FROM stock_akhir_province(@commodityId, @cityId)
					WHERE last_update BETWEEN @startDate::TIMESTAMPTZ AND @endDate::TIMESTAMPTZ
					ORDER BY last_update
				) a
			) "stock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaStokAkhirCityCommodity = `
		SELECT
			'ton' "unit",
			'percentage' "unitDiff",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
		    (
		    	SELECT * FROM city_object(@cityId)
			) "city",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_id,
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', a.tier
						)
					)
				FROM (
				    SELECT
						*,
						@startDate "startDate",
						@endDate "endDate"
					FROM v_stock_city a
					WHERE
						a.commodity_id = @commodityId AND
						a.id = @cityId
				) a
				WHERE
					a.last_update <= CONCAT(a.startDate, ' 23:59:59')::TIMESTAMPTZ AND
					a.last_update >= CONCAT(a.endDate, ' 23:59:59')::TIMESTAMPTZ
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaCompareWithPriceCommodityHistory = `
		SELECT
			'ton' "unit",
			'percentage' "unitDiff",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_id,
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', a.tier
						)
					)
				FROM stock_akhir_city(@commodityId, @startDate, @endDate) a
				WHERE
				    a.id = @cityId
-- 				(
-- 				    SELECT
-- 						*,
-- 						@startDate "startDate",
-- 						@endDate "endDate"
-- 					FROM v_stock_city a
-- 					WHERE
-- 						a.commodity_id = @commodityId AND
-- 						a.id = @cityId
-- 				) a
-- 				WHERE
-- 					a.last_update <= CONCAT(a.startDate, ' 23:59:59')::TIMESTAMPTZ AND
-- 					a.last_update >= CONCAT(a.endDate, ' 23:59:59')::TIMESTAMPTZ
			) "city",
			(
		        SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
							'stock', ROUND(stock::NUMERIC, 0),
							'stockFormat', thousand_format(ROUND(stock::NUMERIC, 0)),
							'price', (
								SELECT ROUND(AVG(price), 0) 
								FROM price_city(@commodityId, @cityId, TO_CHAR(last_update, 'YYYY-MM-DD'), end_of_month(TO_CHAR(last_update, 'YYYY-MM-DD'))::TEXT)
							),
							'priceRupiahFormat', (
							    SELECT thousand_format(ROUND(AVG(price), 0)::numeric) 
								FROM price_city(@commodityId, @cityId, TO_CHAR(last_update, 'YYYY-MM-DD'), end_of_month(TO_CHAR(last_update, 'YYYY-MM-DD'))::TEXT)
							)
						)
					)
		        FROM (
		            SELECT *
		        	FROM stock_akhir_city(@commodityId, @cityId, @startDate, @endDate)
		        	ORDER BY last_update
		        ) a
		    ) "stock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaCompareWithPriceCommodityHistoryProvince = `
		SELECT
			'ton' "unit",
			'percentage' "unitDiff",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_id,
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', a.tier
						)
					)
				FROM stock_akhir_city(@commodityId, @startDate, @endDate) a
				WHERE
				    a.id = @cityId
-- 				(
-- 				    SELECT
-- 						*,
-- 						@startDate "startDate",
-- 						@endDate "endDate"
-- 					FROM v_stock_city a
-- 					WHERE
-- 						a.commodity_id = @commodityId AND
-- 						a.id = @cityId
-- 				) a
-- 				WHERE
-- 					a.last_update <= CONCAT(a.startDate, ' 23:59:59')::TIMESTAMPTZ AND
-- 					a.last_update >= CONCAT(a.endDate, ' 23:59:59')::TIMESTAMPTZ
			) "city",
			(
		        SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
							'stock', ROUND(stock::NUMERIC, 0),
							'stockFormat', thousand_format(ROUND(stock::NUMERIC, 0)),
							'price', (
								SELECT ROUND(AVG(price), 0) 
								FROM price_province_cr(@commodityId, @cityId, TO_CHAR(last_update, 'YYYY-MM-DD'))
							),
							'priceRupiahFormat', (
							    SELECT thousand_format(ROUND(AVG(price), 0)::numeric) 
								FROM price_province_cr(@commodityId, @cityId, TO_CHAR(last_update, 'YYYY-MM-DD'))
							)
						)
					)
		        FROM (
		            SELECT *
		        	FROM stock_akhir_province(@commodityId, @cityId, @startDate, @endDate)
		        	ORDER BY last_update
		        ) a
		    ) "stock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaCompareWithPriceCommodityHistoryAvgChild = `
		SELECT
			'ton' "unit",
			'percentage' "unitDiff",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_id,
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', a.tier
						)
					)
				FROM stock_akhir_city(@commodityId, @startDate, @endDate) a
				WHERE
				    a.id = @cityId
-- 				(
-- 				    SELECT
-- 						*,
-- 						@startDate "startDate",
-- 						@endDate "endDate"
-- 					FROM v_stock_city a
-- 					WHERE
-- 						a.commodity_id = @commodityId AND
-- 						a.id = @cityId
-- 				) a
-- 				WHERE
-- 					a.last_update <= CONCAT(a.startDate, ' 23:59:59')::TIMESTAMPTZ AND
-- 					a.last_update >= CONCAT(a.endDate, ' 23:59:59')::TIMESTAMPTZ
			) "city",
			(
		        SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
							'stock', ROUND(stock::NUMERIC, 0),
							'stockFormat', thousand_format(ROUND(stock::NUMERIC, 0)),
							'price', (
								SELECT ROUND(AVG(price), 0) 
								FROM price_city_avg_child(@commodityId, @cityId, TO_CHAR(last_update, 'YYYY-MM-DD'), end_of_month(TO_CHAR(last_update, 'YYYY-MM-DD'))::TEXT)
							),
							'priceRupiahFormat', (
							    SELECT thousand_format(ROUND(AVG(price), 0)::numeric) 
								FROM price_city_avg_child(@commodityId, @cityId, TO_CHAR(last_update, 'YYYY-MM-DD'), end_of_month(TO_CHAR(last_update, 'YYYY-MM-DD'))::TEXT)
							)
						)
					)
		        FROM (
		            SELECT *
		        	FROM stock_akhir_city(@commodityId, @cityId, @startDate, @endDate)
		        	ORDER BY last_update
		        ) a
		    ) "stock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaCompareWithPriceCommodityHistoryProvinceAvgChild = `
		SELECT
			'ton' "unit",
			'percentage' "unitDiff",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				)
				FROM tm_commodity tc
				WHERE tc.id = @commodityId
			) "commodity",
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', a.commodity_id,
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
							'stock', ROUND(a.stock::NUMERIC, 0),
							'tier', a.tier
						)
					)
				FROM stock_akhir_city(@commodityId, @startDate, @endDate) a
				WHERE
				    a.id = @cityId
-- 				(
-- 				    SELECT
-- 						*,
-- 						@startDate "startDate",
-- 						@endDate "endDate"
-- 					FROM v_stock_city a
-- 					WHERE
-- 						a.commodity_id = @commodityId AND
-- 						a.id = @cityId
-- 				) a
-- 				WHERE
-- 					a.last_update <= CONCAT(a.startDate, ' 23:59:59')::TIMESTAMPTZ AND
-- 					a.last_update >= CONCAT(a.endDate, ' 23:59:59')::TIMESTAMPTZ
			) "city",
			(
		        SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
							'stock', ROUND(stock::NUMERIC, 0),
							'stockFormat', thousand_format(ROUND(stock::NUMERIC, 0)),
							'price', (
								SELECT ROUND(AVG(price), 0) 
								FROM price_province_avg_child(@commodityId, @cityId, TO_CHAR(last_update, 'YYYY-MM-DD'))
							),
							'priceRupiahFormat', (
							    SELECT thousand_format(ROUND(AVG(price), 0)::numeric) 
								FROM price_province_avg_child(@commodityId, @cityId, TO_CHAR(last_update, 'YYYY-MM-DD'))
							)
						)
					)
		        FROM (
		            SELECT *
		        	FROM stock_akhir_province(@commodityId, @cityId, @startDate, @endDate)
		        	ORDER BY last_update
		        ) a
		    ) "stock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`
)
