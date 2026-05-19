package queries

const (
	PriceGetListByCommodityIdAndCityId = `
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
					SELECT (
					   (
						   SELECT price
						   FROM get_harga_by_city_id(@cityId, @commodityId, @startDate, @endDate)
						   ORDER BY last_update DESC
						   LIMIT 1
					   ) -
					   (
						   SELECT price
						   FROM get_harga_by_city_id(@cityId, @commodityId, @startDate, @endDate)
						   ORDER BY last_update ASC
						   LIMIT 1
					   )
				   )
				) priceDiff,
				(
					SELECT rupiah_format(
					   (
						   SELECT (
							  (
								  SELECT price
								  FROM get_harga_by_city_id(@cityId, @commodityId, @startDate, @endDate)
								  ORDER BY last_update DESC
								  LIMIT 1
							  ) -
							  (
								  SELECT price
								  FROM get_harga_by_city_id(@cityId, @commodityId, @startDate, @endDate)
								  ORDER BY last_update ASC
								  LIMIT 1
							  )
						  )
					   )
				   )
				) priceDiffFormat,
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
								'price', price,
								'priceRupiahFormat', rupiah_format(price)
							)
						)
					FROM (
						 SELECT last_update, price
						 FROM get_harga_by_city_id(@cityId, @commodityId, @startDate, @endDate)
						 ORDER BY last_update
					) priceData
				) price
		`

	PriceGetListByCommodityIdAndCityIdAvgChild = `
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
					SELECT (
					   (
						   SELECT price
						   FROM get_harga_by_city_id_avg_child(@cityId, @commodityId, @startDate, @endDate)
						   ORDER BY last_update DESC
						   LIMIT 1
					   ) -
					   (
						   SELECT price
						   FROM get_harga_by_city_id_avg_child(@cityId, @commodityId, @startDate, @endDate)
						   ORDER BY last_update ASC
						   LIMIT 1
					   )
				   )
				) priceDiff,
				(
					SELECT rupiah_format(
					   (
						   SELECT (
							  (
								  SELECT price
								  FROM get_harga_by_city_id_avg_child(@cityId, @commodityId, @startDate, @endDate)
								  ORDER BY last_update DESC
								  LIMIT 1
							  ) -
							  (
								  SELECT price
								  FROM get_harga_by_city_id_avg_child(@cityId, @commodityId, @startDate, @endDate)
								  ORDER BY last_update ASC
								  LIMIT 1
							  )
						  )
					   )
				   )
				) priceDiffFormat,
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
								'price', price,
								'priceRupiahFormat', rupiah_format(price)
							)
						)
					FROM (
						 SELECT last_update, price
						 FROM get_harga_by_city_id_avg_child(@cityId, @commodityId, @startDate, @endDate)
						 ORDER BY last_update
					) priceData
				) price
		`

	PriceGetListByCommodityIdAndCityIdCount = `
			SELECT
				COUNT(1)
			FROM (
				 SELECT last_update, price
				 FROM get_harga_by_city_id(@cityId, @commodityId, @startDate, @endDate)
				 ORDER BY last_update
			) a
		`

	PriceGetListByCommodityIdAndProvinceId = `
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
				) city,
				(
					SELECT (
					   (
						   SELECT price
						   FROM get_harga_province(@provinceId, @commodityId, @startDate, @endDate)
						   ORDER BY last_update DESC
						   LIMIT 1
					   ) -
					   (
						   SELECT price
						   FROM get_harga_province(@provinceId, @commodityId, @startDate, @endDate)
						   ORDER BY last_update ASC
						   LIMIT 1
					   )
				   )
				) priceDiff,
				(
					SELECT rupiah_format(
					   (
						   SELECT (
							  (
								  SELECT price
								  FROM get_harga_province(@provinceId, @commodityId, @startDate, @endDate)
								  ORDER BY last_update DESC
								  LIMIT 1
							  ) -
							  (
								  SELECT price
								  FROM get_harga_province(@provinceId, @commodityId, @startDate, @endDate)
								  ORDER BY last_update ASC
								  LIMIT 1
							  )
						  )
					   )
				   )
				) priceDiffFormat,
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
								'price', price,
								'priceRupiahFormat', rupiah_format(price)
							)
						)
					FROM (
						 SELECT last_update, price
						 FROM get_harga_province(@provinceId, @commodityId, @startDate, @endDate)
						 ORDER BY last_update
					) priceData
				) price
		`

	PriceGetListByCommodityIdAndProvinceIdAvgChild = `
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
				) city,
				(
					SELECT (
					   (
						   SELECT price
						   FROM get_harga_province_avg_child(@provinceId, @commodityId, @startDate, @endDate)
						   ORDER BY last_update DESC
						   LIMIT 1
					   ) -
					   (
						   SELECT price
						   FROM get_harga_province_avg_child(@provinceId, @commodityId, @startDate, @endDate)
						   ORDER BY last_update ASC
						   LIMIT 1
					   )
				   )
				) priceDiff,
				(
					SELECT rupiah_format(
					   (
						   SELECT (
							  (
								  SELECT price
								  FROM get_harga_province_avg_child(@provinceId, @commodityId, @startDate, @endDate)
								  ORDER BY last_update DESC
								  LIMIT 1
							  ) -
							  (
								  SELECT price
								  FROM get_harga_province_avg_child(@provinceId, @commodityId, @startDate, @endDate)
								  ORDER BY last_update ASC
								  LIMIT 1
							  )
						  )
					   )
				   )
				) priceDiffFormat,
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
								'price', price,
								'priceRupiahFormat', rupiah_format(price)
							)
						)
					FROM (
						 SELECT last_update, price
						 FROM get_harga_province_avg_child(@provinceId, @commodityId, @startDate, @endDate)
						 ORDER BY last_update
					) priceData
				) price
		`

	PriceGetListByCommodityIdAndProvinceIdCount = `
			SELECT
				COUNT(1)
			FROM (
				 SELECT last_update, price
				 FROM get_harga_province(@provinceId, @commodityId, @startDate, @endDate)
				 ORDER BY last_update
			) a
		`
)
