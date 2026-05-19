package queries

const (
	//ReportNeracaByCommodity = `
	//	SELECT
	//		name title,
	//		(
	//			SELECT
	//				JSON_AGG(
	//					JSON_BUILD_OBJECT(
	//						TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
	//							SELECT kebutuhan
	//							FROM stock_akhir_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
	//							LIMIT 1
	//						)
	//					)
	//				)
	//			FROM (
	//				SELECT *
	//				FROM (
	//					SELECT
	//						TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
	//					FROM tx_commodity_stock
	//					WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
	//					GROUP BY DATE_TRUNC('month', last_update)
	//					ORDER BY last_update
	//				)
	//			) c
	//		) kebutuhan,
	//		(
	//			SELECT
	//				JSON_AGG(
	//					JSON_BUILD_OBJECT(
	//						TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
	//							SELECT ketersediaan
	//							FROM stock_akhir_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
	//							LIMIT 1
	//						)
	//					)
	//				)
	//			FROM (
	//				SELECT *
	//				FROM (
	//					SELECT
	//						TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
	//					FROM tx_commodity_stock
	//					WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
	//					GROUP BY DATE_TRUNC('month', last_update)
	//					ORDER BY last_update
	//				)
	//			) c
	//		) ketersediaan,
	//		(
	//			SELECT
	//				JSON_AGG(
	//					JSON_BUILD_OBJECT(
	//						TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
	//							SELECT stock
	//							FROM stock_akhir_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
	//							LIMIT 1
	//						)
	//					)
	//				)
	//			FROM (
	//				SELECT *
	//				FROM (
	//					SELECT
	//						TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
	//					FROM tx_commodity_stock
	//					WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
	//					GROUP BY DATE_TRUNC('month', last_update)
	//					ORDER BY last_update
	//				)
	//			) c
	//		) neraca
	//	FROM(
	//		SELECT
	//			tc.id,
	//			tc.name,
	//			"sequence"
	//		FROM tx_commodity_stock tcp
	//		JOIN tm_city tc ON tc.id = tcp.city_id
	//		WHERE
	//			tcp.commodity_id = @commodityId AND
	//			province_id = @provinceId AND
	//			(last_update BETWEEN @startDate AND @endDate)
	//		GROUP BY tc.id, name, "sequence"
	//		ORDER BY "sequence"
	//	) a
	//	OFFSET @page
	//	LIMIT @limit
	//`

	ReportNeracaByCommodity = `
		SELECT
			name title,
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'ketersediaan', (
								JSON_BUILD_OBJECT(
									TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
										SELECT ketersediaan::INTEGER
										FROM stock_akhir_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
										LIMIT 1
									)
								)
							),
							'kebutuhan', (
								JSON_BUILD_OBJECT(
									TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
										SELECT kebutuhan::INTEGER
										FROM stock_akhir_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
										LIMIT 1
									)
								)
							),
							'neraca', (
								JSON_BUILD_OBJECT(
									TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
										SELECT neraca::INTEGER
										FROM stock_akhir_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
										LIMIT 1
									)
								)
							)
						)
					)
				FROM (
					SELECT *
					FROM (
						SELECT
							TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
						FROM tx_commodity_stock
						WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
						GROUP BY DATE_TRUNC('month', last_update)
						ORDER BY last_update
					) stockData
				) c
			) stocks
		FROM(
			SELECT
				tc.id,
				tc.name,
				"sequence" 
			FROM tm_city tc
			WHERE
				province_id = @provinceId
			ORDER BY "sequence"
		) a
		OFFSET @page
		LIMIT @limit
	`

	ReportNeracaByCommodityDownload = `
	SELECT
		name title,
		(
			SELECT
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'ketersediaan', (
							JSON_BUILD_OBJECT(
								TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
									SELECT ketersediaan::INTEGER
									FROM stock_akhir_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
									LIMIT 1
								)
							)
						),
						'kebutuhan', (
							JSON_BUILD_OBJECT(
								TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
									SELECT kebutuhan::INTEGER
									FROM stock_akhir_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
									LIMIT 1
								)
							)
						),
						'neraca', (
							JSON_BUILD_OBJECT(
								TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
									SELECT neraca::INTEGER
									FROM stock_akhir_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
									LIMIT 1
								)
							)
						)
					)
				)
			FROM (
				SELECT *
				FROM (
					SELECT
						TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
					FROM tx_commodity_stock
					WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
					GROUP BY DATE_TRUNC('month', last_update)
					ORDER BY last_update
				) stockData
			) c
		) stocks
	FROM(
		SELECT
			tc.id,
			tc.name,
			"sequence" 
		FROM tm_city tc
		WHERE
			province_id = @provinceId
		ORDER BY "sequence"
	) a
`

	//ReportNeracaByCommodityAvgChild = `
	//	SELECT
	//		name title,
	//		(
	//			SELECT
	//				JSON_AGG(
	//					JSON_BUILD_OBJECT(
	//						TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
	//							kebutuhan
	//							FROM stock_city_avg_child(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
	//							LIMIT 1
	//						)
	//					)
	//				)
	//			FROM (
	//				SELECT *
	//				FROM (
	//					SELECT
	//						TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
	//					FROM tx_commodity_stock
	//					WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
	//					GROUP BY DATE_TRUNC('month', last_update)
	//					ORDER BY last_update
	//				) priceData
	//			) c
	//		) kebutuhan,
	//		(
	//			SELECT
	//				JSON_AGG(
	//					JSON_BUILD_OBJECT(
	//						TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
	//							ketersediaan
	//							FROM stock_city_avg_child(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
	//							LIMIT 1
	//						)
	//					)
	//				)
	//			FROM (
	//				SELECT *
	//				FROM (
	//					SELECT
	//						TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
	//					FROM tx_commodity_stock
	//					WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
	//					GROUP BY DATE_TRUNC('month', last_update)
	//					ORDER BY last_update
	//				) priceData
	//			) c
	//		) ketersediaan,
	//		(
	//			SELECT
	//				JSON_AGG(
	//					JSON_BUILD_OBJECT(
	//						TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
	//							neraca
	//							FROM stock_city_avg_child(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
	//							LIMIT 1
	//						)
	//					)
	//				)
	//			FROM (
	//				SELECT *
	//				FROM (
	//					SELECT
	//						TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
	//					FROM tx_commodity_stock
	//					WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
	//					GROUP BY DATE_TRUNC('month', last_update)
	//					ORDER BY last_update
	//				) priceData
	//			) c
	//		) neraca
	//	FROM(
	//		SELECT
	//			tc.id,
	//			tc.name,
	//			"sequence"
	//		FROM tm_city tc
	//		WHERE
	//			province_id = @provinceId
	//		ORDER BY "sequence"
	//	) a
	//	OFFSET @page
	//	LIMIT @limit
	//`

	//ReportNeracaByCommodityAvgChild = `
	//	SELECT
	//		name title,
	//		(
	//			SELECT
	//				JSON_AGG(
	//					JSON_BUILD_OBJECT(
	//						'ketersediaan', (
	//							JSON_BUILD_OBJECT(
	//								TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
	//									SELECT ketersediaan::INTEGER
	//									FROM stock_city_avg_child(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
	//									LIMIT 1
	//								)
	//							)
	//						),
	//						'kebutuhan', (
	//							JSON_BUILD_OBJECT(
	//								TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
	//									SELECT kebutuhan::INTEGER
	//									FROM stock_city_avg_child(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
	//									LIMIT 1
	//								)
	//							)
	//						),
	//						'neraca', (
	//							JSON_BUILD_OBJECT(
	//								TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
	//									SELECT neraca::INTEGER
	//									FROM stock_city_avg_child(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
	//									LIMIT 1
	//								)
	//							)
	//						)
	//					)
	//				)
	//			FROM (
	//				SELECT *
	//				FROM (
	//					SELECT
	//						TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
	//					FROM tx_commodity_stock
	//					WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
	//					GROUP BY DATE_TRUNC('month', last_update)
	//					ORDER BY last_update
	//				) stockData
	//			) c
	//		) stocks
	//	FROM(
	//		SELECT
	//			tc.id,
	//			tc.name,
	//			"sequence"
	//		FROM tm_city tc
	//		WHERE
	//			province_id = @provinceId
	//		ORDER BY "sequence"
	//	) a
	//	OFFSET @page
	//	LIMIT @limit
	//`

	ReportNeracaByCommodityAvgChild = `
		WITH date_series AS (
    SELECT TO_CHAR(datum, 'YYYY-MM') AS last_update
    FROM GENERATE_SERIES(
        @startDate, 
        @endDate, 
        '1 month'::INTERVAL
    ) datum
)
SELECT
    name AS title,
    (
        SELECT JSON_BUILD_OBJECT(
            'ketersediaan', (
                SELECT JSON_AGG(
                    JSON_BUILD_OBJECT(
                        TO_CHAR(CONCAT(ds.last_update, '-01 00:00:00')::TIMESTAMP, 'MMYYYY'),
                        COALESCE((
                            SELECT ketersediaan::INTEGER
                            FROM stock_akhir_city(@commodityId, a.id, CONCAT(ds.last_update, '-01'), end_of_month(CONCAT(ds.last_update, '-01')::DATE)::TEXT)
                            LIMIT 1
                        ), 0) -- Default ke 0 jika NULL
                    )
                    ORDER BY ds.last_update -- Sort secara ASC
                )
                FROM date_series ds
            ),
            'kebutuhan', (
                SELECT JSON_AGG(
                    JSON_BUILD_OBJECT(
                        TO_CHAR(CONCAT(ds.last_update, '-01 00:00:00')::TIMESTAMP, 'MMYYYY'),
                        COALESCE((
                            SELECT kebutuhan::INTEGER
                            FROM stock_akhir_city(@commodityId, a.id, CONCAT(ds.last_update, '-01'), end_of_month(CONCAT(ds.last_update, '-01')::DATE)::TEXT)
                            LIMIT 1
                        ), 0)
                    )
                    ORDER BY ds.last_update
                )
                FROM date_series ds
            ),
            'neraca', (
                SELECT JSON_AGG(
                    JSON_BUILD_OBJECT(
                        TO_CHAR(CONCAT(ds.last_update, '-01 00:00:00')::TIMESTAMP, 'MMYYYY'),
                        COALESCE((
                            SELECT stock::INTEGER
                            FROM stock_akhir_city(@commodityId, a.id, CONCAT(ds.last_update, '-01'), end_of_month(CONCAT(ds.last_update, '-01')::DATE)::TEXT)
                            LIMIT 1
                        ), 0)
                    )
                    ORDER BY ds.last_update
                )
                FROM date_series ds
            )
        )
    ) AS stocks
FROM (
    SELECT
        tc.id,
        tc.name,
        "sequence" 
    FROM tm_city tc
    WHERE province_id = @provinceId
    ORDER BY "sequence"
) a
		OFFSET @page
		LIMIT @limit
	`

	ReportNeracaByCommodityAvgChildDwonload = `
	WITH date_series AS (
    SELECT TO_CHAR(datum, 'YYYY-MM') AS last_update
    FROM GENERATE_SERIES(
        @startDate, 
        @endDate, 
        '1 month'::INTERVAL
    ) datum
)
SELECT
    name AS title,
    (
        SELECT JSON_BUILD_OBJECT(
            'ketersediaan', (
                SELECT JSON_AGG(
                    JSON_BUILD_OBJECT(
                        TO_CHAR(CONCAT(ds.last_update, '-01 00:00:00')::TIMESTAMP, 'MMYYYY'),
                        COALESCE((
                            SELECT ketersediaan::INTEGER
                            FROM stock_akhir_city(@commodityId, a.id, CONCAT(ds.last_update, '-01'), end_of_month(CONCAT(ds.last_update, '-01')::DATE)::TEXT)
                            LIMIT 1
                        ), 0) -- Default ke 0 jika NULL
                    )
                    ORDER BY ds.last_update -- Sort secara ASC
                )
                FROM date_series ds
            ),
            'kebutuhan', (
                SELECT JSON_AGG(
                    JSON_BUILD_OBJECT(
                        TO_CHAR(CONCAT(ds.last_update, '-01 00:00:00')::TIMESTAMP, 'MMYYYY'),
                        COALESCE((
                            SELECT kebutuhan::INTEGER
                            FROM stock_akhir_city(@commodityId, a.id, CONCAT(ds.last_update, '-01'), end_of_month(CONCAT(ds.last_update, '-01')::DATE)::TEXT)
                            LIMIT 1
                        ), 0)
                    )
                    ORDER BY ds.last_update
                )
                FROM date_series ds
            ),
            'neraca', (
                SELECT JSON_AGG(
                    JSON_BUILD_OBJECT(
                        TO_CHAR(CONCAT(ds.last_update, '-01 00:00:00')::TIMESTAMP, 'MMYYYY'),
                        COALESCE((
                            SELECT stock::INTEGER
                            FROM stock_akhir_city(@commodityId, a.id, CONCAT(ds.last_update, '-01'), end_of_month(CONCAT(ds.last_update, '-01')::DATE)::TEXT)
                            LIMIT 1
                        ), 0)
                    )
                    ORDER BY ds.last_update
                )
                FROM date_series ds
            )
        )
    ) AS stocks
FROM (
    SELECT
        tc.id,
        tc.name,
        "sequence" 
    FROM tm_city tc
    WHERE province_id = @provinceId
    ORDER BY "sequence"
) a
`

	ReportNeracaByCommodityCount = `
		SELECT
			COUNT(1)
		FROM(
			SELECT
				tc.id,
				tc.name,
				"sequence" 
			FROM tx_commodity_stock tcp
			JOIN tm_city tc ON tc.id = tcp.city_id
			WHERE
				tcp.commodity_id = @commodityId AND
				province_id = @provinceId AND
				(last_update BETWEEN @startDate AND @endDate)
			GROUP BY tc.id, name, "sequence"
			ORDER BY "sequence"
		) a
	`

	ReportNeracaByCommodityCountAvgChild = `
		SELECT
			COUNT(1)
		FROM(
			SELECT
				tc.id,
				tc.name,
				"sequence" 
			FROM tx_commodity_stock tcp
			JOIN tm_city tc ON tc.id = tcp.city_id
			WHERE
				tcp.commodity_id = @commodityId AND
				province_id = @provinceId AND
				(last_update BETWEEN @startDate AND @endDate)
			GROUP BY tc.id, name, "sequence"
			ORDER BY "sequence"
		) a
	`

	ReportNeracaByCity = `
		WITH bulan AS (
    SELECT TO_CHAR(DATE_TRUNC('month', dd::DATE), 'MMYYYY') AS bulan
    FROM generate_series(
        @startDate, 
        @endDate, 
        '1 month'::INTERVAL
    ) dd
),
data_stok AS (
    SELECT 
        tc.id AS commodity_id,
        TO_CHAR(DATE_TRUNC('month', tcs.last_update::DATE), 'MMYYYY') AS bulan,
        (
            CASE 
                WHEN tc.parent_id IS NULL THEN (
                    SELECT ketersediaan::INTEGER
                    FROM stock_akhir_city(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                ) 
                ELSE (
                    SELECT ketersediaan::INTEGER
                    FROM stock_city_avg_child(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                )
            END
        ) AS ketersediaan,
        (
            CASE 
                WHEN tc.parent_id IS NULL THEN (
                    SELECT kebutuhan::INTEGER
                    FROM stock_akhir_city(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                ) 
                ELSE (
                    SELECT kebutuhan::INTEGER
                    FROM stock_city_avg_child(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                )
            END
        ) AS kebutuhan,
        (
            CASE 
                WHEN tc.parent_id IS NULL THEN (
                    SELECT stock::INTEGER
                    FROM stock_akhir_city(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                ) 
                ELSE (
                    SELECT neraca::INTEGER
                    FROM stock_city_avg_child(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                )
            END
        ) AS neraca
    FROM tx_commodity_stock tcs
    JOIN tm_commodity tc ON tc.id = tcs.commodity_id
    WHERE tcs.last_update BETWEEN @startDate AND @endDate
    GROUP BY tc.id, bulan, tc.parent_id, tcs.last_update
)
SELECT
    tc.name AS title,
    JSON_BUILD_OBJECT(
        'ketersediaan', (
            SELECT JSON_AGG(
                JSON_BUILD_OBJECT(bulan.bulan, COALESCE(data_stok.ketersediaan, 0))
            )
            FROM bulan
            LEFT JOIN data_stok ON bulan.bulan = data_stok.bulan AND data_stok.commodity_id = tc.id
        ),
        'kebutuhan', (
            SELECT JSON_AGG(
                JSON_BUILD_OBJECT(bulan.bulan, COALESCE(data_stok.kebutuhan, 0))
            )
            FROM bulan
            LEFT JOIN data_stok ON bulan.bulan = data_stok.bulan AND data_stok.commodity_id = tc.id
        ),
        'neraca', (
            SELECT JSON_AGG(
                JSON_BUILD_OBJECT(bulan.bulan, COALESCE(data_stok.neraca, 0))
            )
            FROM bulan
            LEFT JOIN data_stok ON bulan.bulan = data_stok.bulan AND data_stok.commodity_id = tc.id
        )
    ) AS stocks
FROM tm_commodity tc
WHERE EXISTS (
    SELECT 1 FROM tx_commodity_stock tcs WHERE tcs.commodity_id = tc.id AND tcs.last_update BETWEEN @startDate AND @endDate
)
ORDER BY tc.name
		OFFSET @page
		LIMIT @limit
	`

	ReportNeracaByCityDownload = `
		WITH bulan AS (
    SELECT TO_CHAR(DATE_TRUNC('month', dd::DATE), 'MMYYYY') AS bulan
    FROM generate_series(
        @startDate, 
        @endDate, 
        '1 month'::INTERVAL
    ) dd
),
data_stok AS (
    SELECT 
        tc.id AS commodity_id,
        TO_CHAR(DATE_TRUNC('month', tcs.last_update::DATE), 'MMYYYY') AS bulan,
        (
            CASE 
                WHEN tc.parent_id IS NULL THEN (
                    SELECT ketersediaan::INTEGER
                    FROM stock_akhir_city(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                ) 
                ELSE (
                    SELECT ketersediaan::INTEGER
                    FROM stock_city_avg_child(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                )
            END
        ) AS ketersediaan,
        (
            CASE 
                WHEN tc.parent_id IS NULL THEN (
                    SELECT kebutuhan::INTEGER
                    FROM stock_akhir_city(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                ) 
                ELSE (
                    SELECT kebutuhan::INTEGER
                    FROM stock_city_avg_child(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                )
            END
        ) AS kebutuhan,
        (
            CASE 
                WHEN tc.parent_id IS NULL THEN (
                    SELECT stock::INTEGER
                    FROM stock_akhir_city(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                ) 
                ELSE (
                    SELECT neraca::INTEGER
                    FROM stock_city_avg_child(tc.id, @cityId, TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM')::TEXT, end_of_month(TO_DATE(TO_CHAR(tcs.last_update, 'YYYY-MM'), 'YYYY-MM'))::TEXT)
                    LIMIT 1
                )
            END
        ) AS neraca
    FROM tx_commodity_stock tcs
    JOIN tm_commodity tc ON tc.id = tcs.commodity_id
    WHERE tcs.last_update BETWEEN @startDate AND @endDate
    GROUP BY tc.id, bulan, tc.parent_id, tcs.last_update
)
SELECT
    tc.name AS title,
    JSON_BUILD_OBJECT(
        'ketersediaan', (
            SELECT JSON_AGG(
                JSON_BUILD_OBJECT(bulan.bulan, COALESCE(data_stok.ketersediaan, 0))
            )
            FROM bulan
            LEFT JOIN data_stok ON bulan.bulan = data_stok.bulan AND data_stok.commodity_id = tc.id
        ),
        'kebutuhan', (
            SELECT JSON_AGG(
                JSON_BUILD_OBJECT(bulan.bulan, COALESCE(data_stok.kebutuhan, 0))
            )
            FROM bulan
            LEFT JOIN data_stok ON bulan.bulan = data_stok.bulan AND data_stok.commodity_id = tc.id
        ),
        'neraca', (
            SELECT JSON_AGG(
                JSON_BUILD_OBJECT(bulan.bulan, COALESCE(data_stok.neraca, 0))
            )
            FROM bulan
            LEFT JOIN data_stok ON bulan.bulan = data_stok.bulan AND data_stok.commodity_id = tc.id
        )
    ) AS stocks
FROM tm_commodity tc
WHERE EXISTS (
    SELECT 1 FROM tx_commodity_stock tcs WHERE tcs.commodity_id = tc.id AND tcs.last_update BETWEEN @startDate AND @endDate
)
ORDER BY tc.name;
	`

	ReportNeracaByCityCount = `
		SELECT
			COUNT(1)
		FROM (
			SELECT
				tc.id,
				tc.name 
			FROM tx_commodity_stock tcp
			JOIN tm_commodity tc ON tc.id = tcp.commodity_id
			WHERE
				tcp.city_id = @cityId AND
				(last_update BETWEEN @startDate AND @endDate)
			GROUP BY tc.id, name
			ORDER BY name
		) a
	`
)
