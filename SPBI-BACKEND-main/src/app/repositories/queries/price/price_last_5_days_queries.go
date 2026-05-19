package price

const (
	PriceLast5DaysPriceByCityId = `
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
				) city,
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', b.id,
								'parentId', b.parent_id,
								'name', b.name,
								'createdAt', b.created_at,
								'updatedAt', b.updated_at,
								'deletedAt', b.deleted_at,
								'period', (
									SELECT JSON_BUILD_OBJECT(
									   'startDate', startDate,
									   'endDate', endDate
									)
								),
								'currentPrice', (
									SELECT price
									FROM get_harga_by_city_id_avg_child(a.id, b.id, @startDate, @endDate)
									ORDER BY last_update DESC
                                	LIMIT 1
								),
								'currentPriceRupiahFormat', (
									SELECT rupiah_format(
										(
											SELECT price
											FROM get_harga_by_city_id_avg_child(a.id, b.id, @startDate, @endDate)
											ORDER BY last_update DESC
                                			LIMIT 1
									   )
								   )
								),
								'priceDiffLast5Days', (
									SELECT (
									   (
										   SELECT price
										   FROM get_harga_by_city_id_avg_child(a.id, b.id, @startDate, @endDate)
										   ORDER BY last_update DESC
										   LIMIT 1
									   ) -
									   (
										   SELECT price
										   FROM get_harga_by_city_id_avg_child(a.id, b.id, @startDate, @endDate)
										   ORDER BY last_update ASC
										   LIMIT 1
									   )
									)
								),
								'priceDiffLast5DaysFormat', (
									SELECT rupiah_format(
										(
											SELECT (
												  (
													  SELECT price
													  FROM get_harga_by_city_id_avg_child(a.id, b.id, @startDate, @endDate)
													  ORDER BY last_update DESC
													  LIMIT 1
												  ) -
												  (
													  SELECT price
													  FROM get_harga_by_city_id_avg_child(a.id, b.id, @startDate, @endDate)
													  ORDER BY last_update ASC
													  LIMIT 1
												  )
											)
									   )
								   )
								),
								'price', (
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
										FROM get_harga_by_city_id_avg_child(a.id, b.id, @startDate, @endDate)
										ORDER BY last_update
									) price
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @startDate startDate,
						    @endDate endDate
						FROM
							tm_commodity
						WHERE parent_id IS NULL 
						ORDER BY sequence
					) b
				) commodities
			FROM tm_city a
			WHERE a.id = @cityId
		`

	PriceLast5DaysPriceByCityIdCount = `
			SELECT (
				SELECT
					COUNT(1)
				FROM
					tm_commodity b
				WHERE parent_id IS NULL 
			)
			FROM tm_city a
			WHERE a.id = @cityId
		`

	PriceLast5DaysPriceByCommodityId = `
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
				) commodity,
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', b.id,
								'name', b.name,
								'createdAt', b.created_at,
								'updatedAt', b.updated_at,
								'deletedAt', b.deleted_at,
								'period', (
									SELECT JSON_BUILD_OBJECT(
									   'startDate', startDate,
									   'endDate', endDate
									)
								),
								'currentPrice', (
									SELECT price
									FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
									ORDER BY last_update DESC
                                	LIMIT 1
								),
								'currentPriceRupiahFormat', (
									SELECT rupiah_format(
										(
											SELECT price
											FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
											ORDER BY last_update DESC
                                			LIMIT 1
									   )
								   )
								),
								'priceDiffLast5Days', (
									SELECT (
									   (
										   SELECT price
										   FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
										   ORDER BY last_update DESC
										   LIMIT 1
									   ) -
									   (
										   SELECT price
										   FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
										   ORDER BY last_update ASC
										   LIMIT 1
									   )
									)
								),
								'priceDiffLast5DaysFormat', (
									SELECT rupiah_format(
										(
											SELECT (
												  (
													  SELECT price
													  FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
													  ORDER BY last_update DESC
													  LIMIT 1
												  ) -
												  (
													  SELECT price
													  FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
													  ORDER BY last_update ASC
													  LIMIT 1
												  )
											)
									   )
								   )
								),
								'price', (
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
										FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
										ORDER BY last_update
									) price
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @startDate startDate,
						    @endDate endDate
						FROM
							tm_city
						WHERE province_id = 73
						ORDER BY name
					) b
				) cities
			FROM tm_commodity a
			WHERE a.id = @commodityId
		`

	PriceLast5DaysPriceByCommodityIdNew = `
			SELECT
				JSON_BUILD_OBJECT(
					'id', a.id,
					'parentId', a.parent_id,
					'name', a.name,
					'assets', null,
					'createdAt', a.created_at,
					'updatedAt', a.updated_at,
					'deletedAt', a.deleted_at
				) commodity,
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', b.id,
								'name', b.name,
								'createdAt', b.created_at,
								'updatedAt', b.updated_at,
								'deletedAt', b.deleted_at,
								'period', (
									SELECT JSON_BUILD_OBJECT(
									   'startDate', startDate,
									   'endDate', endDate
									)
								),
								'currentPrice', (
									SELECT price
									FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
									ORDER BY last_update DESC
                                	LIMIT 1
								),
								'currentPriceRupiahFormat', (
									SELECT rupiah_format(
										(
											SELECT price
											FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
											ORDER BY last_update DESC
                                			LIMIT 1
									   )
								   )
								),
								'priceDiffLast5Days', (
									SELECT (
									   (
										   SELECT price
										   FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
										   ORDER BY last_update DESC
										   LIMIT 1
									   ) -
									   (
										   SELECT price
										   FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
										   ORDER BY last_update ASC
										   LIMIT 1
									   )
									)
								),
								'priceDiffLast5DaysFormat', (
									SELECT rupiah_format(
										(
											SELECT (
												  (
													  SELECT price
													  FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
													  ORDER BY last_update DESC
													  LIMIT 1
												  ) -
												  (
													  SELECT price
													  FROM get_harga_by_city_id(b.id, a.id, @startDate, @endDate)
													  ORDER BY last_update ASC
													  LIMIT 1
												  )
											)
									   )
								   )
								),
								'price', (
									SELECT 
										CASE
											WHEN(
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
													FROM get_harga_by_city_id(b.id, a.id, b.startDate, b.endDate)
													ORDER BY last_update
												) as last5dayPrice
											) IS NOT NULL
											THEN (
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
													FROM get_harga_by_city_id(b.id, a.id, b.startDate, b.endDate)
													ORDER BY last_update
												) as last5dayPrice
											)
											ELSE(
												(
													SELECT
														JSON_AGG(
															JSON_BUILD_OBJECT(
																'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
																'price', (
																	SELECT AVG(price)::INTEGER
																	FROM tx_commodity_price tcp 
																	WHERE
																		city_id = b.id AND
																		commodity_id IN (
																			SELECT id FROM tm_commodity tc WHERE tc.parent_id = @commodityId
																		) AND
																		last_update = c.last_update
																),
																'priceRupiahFormat', (
																	SELECT rupiah_format(
																		(
																			SELECT AVG(price)::INTEGER
																			FROM tx_commodity_price tcp 
																			WHERE
																				city_id = b.id AND
																				commodity_id IN (
																					SELECT id FROM tm_commodity tc WHERE tc.parent_id = @commodityId
																				) AND
																				last_update = c.last_update
																		)
																	)
																)
															)
														)
													FROM (
														SELECT
															city_id,
															last_update
														FROM tx_commodity_price a
														WHERE
															city_id = b.id AND
															commodity_id IN (
																SELECT id FROM tm_commodity tc WHERE tc.parent_id = @commodityId
															) AND
															last_update BETWEEN CONCAT(b.startDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(b.endDate, ' 23:59:59')::TIMESTAMPTZ
														GROUP BY
															city_id,
															last_update
													) c
												)
											)
										END
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @startDate startDate,
						    @endDate endDate
						FROM
							tm_city
						WHERE province_id = 73
						ORDER BY sequence
					) b
				) cities
			FROM tm_commodity a
			WHERE a.id = @commodityId
		`

	PriceLast5DaysPriceByCommodityIdNewAvgChild = `
			SELECT
				JSON_BUILD_OBJECT(
					'id', a.id,
					'parentId', a.parent_id,
					'name', a.name,
					'assets', null,
					'createdAt', a.created_at,
					'updatedAt', a.updated_at,
					'deletedAt', a.deleted_at
				) commodity,
				(
					SELECT
						JSON_AGG(
							JSON_BUILD_OBJECT(
								'id', b.id,
								'name', b.name,
								'createdAt', b.created_at,
								'updatedAt', b.updated_at,
								'deletedAt', b.deleted_at,
								'period', (
									SELECT JSON_BUILD_OBJECT(
									   'startDate', startDate,
									   'endDate', endDate
									)
								),
								'currentPrice', (
									SELECT price
									FROM get_harga_by_city_id_avg_child(b.id, a.id, @startDate, @endDate)
									ORDER BY last_update DESC
                                	LIMIT 1
								),
								'currentPriceRupiahFormat', (
									SELECT rupiah_format(
										(
											SELECT price
											FROM get_harga_by_city_id_avg_child(b.id, a.id, @startDate, @endDate)
											ORDER BY last_update DESC
                                			LIMIT 1
									   )
								   )
								),
								'priceDiffLast5Days', (
									SELECT (
									   (
										   SELECT price
										   FROM get_harga_by_city_id_avg_child(b.id, a.id, @startDate, @endDate)
										   ORDER BY last_update DESC
										   LIMIT 1
									   ) -
									   (
										   SELECT price
										   FROM get_harga_by_city_id_avg_child(b.id, a.id, @startDate, @endDate)
										   ORDER BY last_update ASC
										   LIMIT 1
									   )
									)
								),
								'priceDiffLast5DaysFormat', (
									SELECT rupiah_format(
										(
											SELECT (
												  (
													  SELECT price
													  FROM get_harga_by_city_id_avg_child(b.id, a.id, @startDate, @endDate)
													  ORDER BY last_update DESC
													  LIMIT 1
												  ) -
												  (
													  SELECT price
													  FROM get_harga_by_city_id_avg_child(b.id, a.id, @startDate, @endDate)
													  ORDER BY last_update ASC
													  LIMIT 1
												  )
											)
									   )
								   )
								),
								'price', (
									SELECT 
										CASE
											WHEN(
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
													FROM get_harga_by_city_id_avg_child(b.id, a.id, b.startDate, b.endDate)
													ORDER BY last_update
												) as last5dayPrice
											) IS NOT NULL
											THEN (
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
													FROM get_harga_by_city_id_avg_child(b.id, a.id, b.startDate, b.endDate)
													ORDER BY last_update
												) as last5dayPrice
											)
											ELSE(
												(
													SELECT
														JSON_AGG(
															JSON_BUILD_OBJECT(
																'date', CONCAT(CONCAT(TO_CHAR(last_update, 'DD'), CONCAT('/', REPLACE(TO_CHAR(last_update, 'MM'), ' ',''))), CONCAT('/', TO_CHAR(last_update, 'YY'))),
																'price', (
																	SELECT AVG(price)::INTEGER
																	FROM tx_commodity_price tcp 
																	WHERE
																		city_id = b.id AND
																		commodity_id IN (
																			SELECT id FROM tm_commodity tc WHERE tc.parent_id = @commodityId
																		) AND
																		last_update = c.last_update
																),
																'priceRupiahFormat', (
																	SELECT rupiah_format(
																		(
																			SELECT AVG(price)::INTEGER
																			FROM tx_commodity_price tcp 
																			WHERE
																				city_id = b.id AND
																				commodity_id IN (
																					SELECT id FROM tm_commodity tc WHERE tc.parent_id = @commodityId
																				) AND
																				last_update = c.last_update
																		)
																	)
																)
															)
														)
													FROM (
														SELECT
															city_id,
															last_update
														FROM tx_commodity_price a
														WHERE
															city_id = b.id AND
															commodity_id IN (
																SELECT id FROM tm_commodity tc WHERE tc.parent_id = @commodityId
															) AND
															last_update BETWEEN CONCAT(b.startDate, ' 00:00:00')::TIMESTAMPTZ AND CONCAT(b.endDate, ' 23:59:59')::TIMESTAMPTZ
														GROUP BY
															city_id,
															last_update
													) c
												)
											)
										END
								)
							)
						)
					FROM (
						SELECT
						    *,
						    @startDate startDate,
						    @endDate endDate
						FROM
							tm_city
						WHERE province_id = 73
						ORDER BY sequence
					) b
				) cities
			FROM tm_commodity a
			WHERE a.id = @commodityId
		`

	PriceLast5DaysPriceByCommodityIdCount = `
			SELECT (
				SELECT
					COUNT(1)
				FROM
					tm_city b
				WHERE province_id = 73
			)
			FROM tm_commodity a
			WHERE a.id = @commodityId
		`
)
