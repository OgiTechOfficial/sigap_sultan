package neraca

const (
	NeracaKetersediaanMap = `	
		SELECT
			'ton' "unit",
			(
				SELECT JSON_BUILD_OBJECT(
					'id',tc.id,
					'name',tc.name
				) FROM tm_commodity tc WHERE tc.id = 1
			) "commodity",
		    (
		        SELECT 
					JSON_BUILD_OBJECT (
						'meningkat', (
							SELECT COUNT(1)
							FROM (
								SELECT
								CASE
									WHEN a.latest > a.latest2 THEN 'meningkat'
									WHEN a.latest < a.latest2 THEN 'menurun'
									when a.latest = latest2 then 'stabil'
            						else ''
								END AS hasil
								FROM get_ketersediaan_diff_cr(@provinceId, @commodityId, @startDate, @endDate) a
							) d
							WHERE hasil='meningkat'
						),
						'menurun', (
							SELECT COUNT(1)
							FROM (
							SELECT
								CASE
									WHEN a.latest > a.latest2 THEN 'meningkat'
									WHEN a.latest < a.latest2 THEN 'menurun'
									when a.latest = latest2 then 'stabil'
            						else ''
								END AS hasil
							FROM get_ketersediaan_diff_cr(@provinceId, @commodityId, @startDate, @endDate) a) d
							WHERE hasil='menurun'
						),
						'stabil', (
							SELECT COUNT(1)
							FROM (
								SELECT
									CASE
										WHEN a.latest > a.latest2 THEN 'meningkat'
										WHEN a.latest < a.latest2 THEN 'menurun'
										when a.latest = latest2 then 'stabil'
            						else ''
									END AS hasil
								FROM get_ketersediaan_diff_cr(@provinceId, @commodityId, @startDate, @endDate) a
							) d
							WHERE hasil='stabil'
						)
					)
		    ) "summary",
			(
				SELECT
					JSON_BUILD_OBJECT(
						'id', province_id,
						'clientId', client_id,
						'province', (
							SELECT * FROM province_object(73)
						),
						'stock', c.stock,
						'stockDiff', ROUND(c.stockdiff::NUMERIC, 2),
						'tier', c.tier
					)
				FROM get_level_ketersediaan_province_cr(@commodityId, @startDate, @endDate) c
				WHERE province_id = 73
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
							'stockDiff', ROUND(a.stockDiff::NUMERIC, 2),
							'tier', a.tier
						)
					)
				FROM get_level_ketersediaan_cr(@provinceId, @commodityId, @startDate, @endDate) a
			) "cityStock",
			(
				SELECT '{"menurun":{"title":"Menurun","color":"#FF6711"},"stabil":{"title":"Stabil","color":"#32D583"},"meningkat":{"title":"Meningkat","color":"#05603A"}}'::JSON
			) "stockTier",
			(
				SELECT '["menurun","stabil","meningkat"]'::JSON
			) "stockTierCode"
	`

	NeracaKetersediaanListByCommodity = `
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
	-- 								WHEN (c.stock / c.kebutuhan) <= 80 and (c.stock / c.kebutuhan) >= 47 THEN 'waspada'
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
							'stockDiff', ROUND(a.stockDiff::NUMERIC, 2),
							'tier', a.tier
						)
					)
				FROM get_level_ketersediaan_cr(@provinceId, @commodityId, @startDate, @endDate) a
				OFFSET @page
				LIMIT @limit
			) "cityStock",
			(
				SELECT '{"menurun":{"title":"Menurun","color":"#FF6711"},"stabil":{"title":"Stabil","color":"#32D583"},"meningkat":{"title":"Meningkat","color":"#05603A"}}'::JSON
			) "stockTier",
			(
				SELECT '["menurun","stabil","meningkat"]'::JSON
			) "stockTierCode"
	`

	NeracaKetersediaanListByCommodityAsc = `
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
								WHEN kebutuhan = 0
									THEN 'aman'
								ELSE
									CASE
										WHEN (c.stock / c.kebutuhan) >= 81 THEN 'aman'
		-- 								WHEN (c.stock / c.kebutuhan) <= 80 and (c.stock / c.kebutuhan) >= 47 THEN 'waspada'
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
							'stockDiff', ROUND(a.stockDiff::NUMERIC, 2),
							'tier', a.tier
						)
					)
				FROM(
				    SELECT *
				    FROM get_level_ketersediaan_cr(@provinceId, @commodityId, @startDate, @endDate) a
				    ORDER BY a.stock
					OFFSET @page
					LIMIT @limit
				) a
			) "cityStock",
			(
				SELECT '{"menurun":{"title":"Menurun","color":"#FF6711"},"stabil":{"title":"Stabil","color":"#32D583"},"meningkat":{"title":"Meningkat","color":"#05603A"}}'::JSON
			) "stockTier",
			(
				SELECT '["menurun","stabil","meningkat"]'::JSON
			) "stockTierCode"
	`

	NeracaKetersediaanListByCommodityDesc = `
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
								WHEN kebutuhan = 0
									THEN 'aman'
								ELSE
									CASE
										WHEN (c.stock / c.kebutuhan) >= 81 THEN 'aman'
		-- 								WHEN (c.stock / c.kebutuhan) <= 80 and (c.stock / c.kebutuhan) >= 47 THEN 'waspada'
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
							'stockDiff', ROUND(a.stockDiff::NUMERIC, 2),
							'tier', a.tier
						)
					)
				FROM(
					SELECT *
					FROM(
					    SELECT *
						FROM get_level_ketersediaan_cr(@provinceId, @commodityId, @startDate, @endDate) a
					    WHERE a.stock IS NOT NULL
						ORDER BY a.stock DESC
					) a2
					UNION ALL
					SELECT *
					FROM get_level_ketersediaan_cr(@provinceId, @commodityId, @startDate, @endDate) a
					WHERE a.stock IS NULL
					OFFSET @page
					LIMIT @limit
				) a
			) "cityStock",
			(
				SELECT '{"menurun":{"title":"Menurun","color":"#FF6711"},"stabil":{"title":"Stabil","color":"#32D583"},"meningkat":{"title":"Meningkat","color":"#05603A"}}'::JSON
			) "stockTier",
			(
				SELECT '["menurun","stabil","meningkat"]'::JSON
			) "stockTierCode"
	`

	NeracaKetersediaanCommodityHistoryMenurun = `
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
							'ketersediaan', a.stock,
							'ketersediaanDiffPercentage', ROUND(a.stockDiff::NUMERIC, 2),
							'tier', a.tier
						)
					)
				FROM get_level_ketersediaan_cr(@provinceId, @commodityId, @startDate, @endDate) a
				WHERE a.tier = 'menurun'
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaKetersediaanCommodityHistoryStabil = `
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
							'ketersediaan', ROUND(a.stock::NUMERIC, 0),
							'ketersediaanDiffPercentage', ROUND(a.stockDiff::NUMERIC, 2),
							'tier', a.tier
						)
					)
				FROM get_level_ketersediaan_cr(@provinceId, @commodityId, @startDate, @endDate) a
				WHERE a.tier = 'stabil'
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaKetersediaanCommodityHistoryMeningkat = `
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
							'ketersediaan', ROUND(a.stock::NUMERIC, 0),
							'ketersediaanDiffPercentage', ROUND(a.stockDiff::NUMERIC, 2),
							'tier', a.tier
						)
					)
				FROM get_level_ketersediaan_cr(@provinceId, @commodityId, @startDate, @endDate) a
				WHERE a.tier = 'meningkat'
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaKetersediaanCommodityCityHistory = `
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
-- 							'id', a.commodity_id,
-- 							'client_id', a.client_id,
-- 							'city', (
-- 								SELECT
-- 									JSON_BUILD_OBJECT(
-- 										'id', b.id,
-- 										'name', b.name,
-- 										'assets', (
-- 											SELECT
-- 												JSON_BUILD_OBJECT(
-- 													'id', id,
-- 													'assetsType', assets_type,
-- 													'assetsLocation', assets_location,
-- 													'assetsLocationType', assets_location_type,
-- 													'assetsMediaType', assets_media_type,
-- 													'assetsExt', assets_ext,
-- 													'assetsName', assets_name
-- 												)
-- 											FROM assets
-- 											WHERE id = b.assets_relation_id
-- 										),
-- 										'created_at', b.created_at,
-- 										'updated_at', b.updated_at,
-- 										'deleted_at', b.deleted_at
-- 									)
-- 								FROM tm_city b
-- 								WHERE id = a.id
-- 							),
-- 							'ketersediaan', a.ketersediaan,
							'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
							'stock', ROUND(a.ketersediaan::numeric, 0),
							'stockFormat', ''
-- 							'tier', a.tier
						)
					)
-- 				FROM get_level_ketersediaan(@provinceId, @commodityId, @startDate, @endDate) a
-- 				FROM v_stock_city a
				FROM(
				    SELECT *
				    FROM stock_akhir_city(@commodityId, @startDate, @endDate) 
					WHERE
						id = @cityId
					ORDER BY last_update
				) a
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaKebutuhanCommodityHistoryMenurun = `
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
							'kebutuhan', ROUND(a.stock::NUMERIC, 0),
							'kebutuhanDiffPercentage', ROUND(a.stockDiff::NUMERIC, 2),
							'tier', a.tier
						)
					)
				FROM get_level_kebutuhan_cr(@provinceId, @commodityId, @startDate, @endDate) a
				WHERE a.tier = 'menurun'
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaKebutuhanCommodityHistoryStabil = `
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
							'kebutuhan', ROUND(a.stock::NUMERIC, 0),
							'kebutuhanDiffPercentage', ROUND(a.stockDiff::NUMERIC, 2),
							'tier', a.tier
						)
					)
				FROM get_level_kebutuhan_cr(@provinceId, @commodityId, @startDate, @endDate) a
				WHERE a.tier = 'stabil'
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaKebutuhanCommodityHistoryMeningkat = `
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
							'kebutuhan', ROUND(a.stock::NUMERIC, 0),
							'kebutuhanDiffPercentage', ROUND(a.stockDiff::NUMERIC, 2),
							'tier', a.tier
						)
					)
				FROM get_level_kebutuhan_cr(@provinceId, @commodityId, @startDate, @endDate) a
				WHERE a.tier = 'meningkat'
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`

	NeracaKebutuhanCommodityCityHistory = `
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
-- 							'id', a.commodity_id,
-- 							'client_id', a.client_id,
-- 							'city', (
-- 								SELECT
-- 									JSON_BUILD_OBJECT(
-- 										'id', b.id,
-- 										'name', b.name,
-- 										'assets', (
-- 											SELECT
-- 												JSON_BUILD_OBJECT(
-- 													'id', id,
-- 													'assetsType', assets_type,
-- 													'assetsLocation', assets_location,
-- 													'assetsLocationType', assets_location_type,
-- 													'assetsMediaType', assets_media_type,
-- 													'assetsExt', assets_ext,
-- 													'assetsName', assets_name
-- 												)
-- 											FROM assets
-- 											WHERE id = b.assets_relation_id
-- 										),
-- 										'created_at', b.created_at,
-- 										'updated_at', b.updated_at,
-- 										'deleted_at', b.deleted_at
-- 									)
-- 								FROM tm_city b
-- 								WHERE id = a.id
-- 							),
-- 							'kebutuhan', a.kebutuhan,
-- 							'tier', a.tier
							'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
							'stock', ROUND(a.kebutuhan::numeric, 0),
							'stockFormat', ''
						)
					)
-- 				FROM get_level_kebutuhan(@provinceId, @commodityId, @startDate, @endDate) a
-- 				FROM v_stock_city a
				FROM(
				    SELECT *
				    FROM stock_akhir_city(@commodityId, @startDate, @endDate) 
					WHERE
						id = @cityId
					ORDER BY last_update
				) a
			) "cityStock",
			(
				SELECT '{"defisit":{"title":"Defisit","start":null,"end":0,"color":"#B11016","unit":"percentage","backgroundColor":"FEF3F2"},"rentan":{"title":"Rentan","start":0,"end":46,"color":"#FF6711","unit":"percentage","backgroundColor":"FEEFC6"},"waspada":{"title":"Waspada","start":47,"end":80,"color":"#E4B701","unit":"percentage","backgroundColor":"FFFAEB"},"aman":{"title":"Aman","start":81,"end":100,"color":"#05603A","unit":"percentage","backgroundColor":"D1FADF"}}'::JSON
			) "stockTier"
	`
)
