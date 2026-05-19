package queries

const (
	ReportPriceByCommodity = `
		SELECT
			name title,
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
								SELECT rupiah_format(price)
								FROM price_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
								LIMIT 1
							)
						)
					)
				FROM (
					SELECT *
					FROM (
						SELECT
							TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
						FROM tx_commodity_price
						WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
						GROUP BY DATE_TRUNC('month', last_update)
						ORDER BY last_update
					)
				) c
			) prices
		FROM(
			SELECT
				tc.id,
				tc.name,
				"sequence" 
			FROM tx_commodity_price tcp
			JOIN tm_city tc ON tc.id = tcp.city_id
			WHERE
				tcp.commodity_id = @commodityId AND
				province_id = @provinceId AND
				(last_update BETWEEN @startDate AND @endDate)
			GROUP BY tc.id, name, "sequence"
			ORDER BY "sequence"
		) a
		OFFSET @page
		LIMIT @limit
	`

	ReportPriceByCommodityDownload = `
		SELECT
			name title,
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
								SELECT rupiah_format(price)
								FROM price_city(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
								LIMIT 1
							)
						)
					)
				FROM (
					SELECT *
					FROM (
						SELECT
							TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
						FROM tx_commodity_price
						WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
						GROUP BY DATE_TRUNC('month', last_update)
						ORDER BY last_update
					)
				) c
			) prices
		FROM(
			SELECT
				tc.id,
				tc.name,
				"sequence" 
			FROM tx_commodity_price tcp
			JOIN tm_city tc ON tc.id = tcp.city_id
			WHERE
				tcp.commodity_id = @commodityId AND
				province_id = @provinceId AND
				(last_update BETWEEN @startDate AND @endDate)
			GROUP BY tc.id, name, "sequence"
			ORDER BY "sequence"
		) a
	`

	ReportPriceByCommodityAvgChild = `
		SELECT
			name title,
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
								SELECT rupiah_format(price)
								FROM price_city_avg_child(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
								LIMIT 1
							)
						)
					)
				FROM (
					SELECT *
					FROM (
						SELECT
							TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
						FROM tx_commodity_price
						WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
						GROUP BY DATE_TRUNC('month', last_update)
						ORDER BY last_update
					) priceData
				) c
			) prices
		FROM(
			SELECT
				tc.id,
				tc.name,
				"sequence" 
			FROM tm_city tc
			--	FROM tx_commodity_price tcp
			--	JOIN tm_city tc ON tc.id = tcp.city_id
			WHERE
			--		tcp.commodity_id IN (
			--			SELECT id
			--			FROM tm_commodity
			--			WHERE parent_id = 1
			--		) AND
				province_id = @provinceId
			--		AND
			--		(last_update BETWEEN '2024-07-01 00:00:00' AND '2024-10-31 23:59:59')
			--	GROUP BY tc.id, name, "sequence"
			ORDER BY "sequence"
-- 			SELECT
-- 				tc.id,
-- 				tc.name,
-- 				"sequence" 
-- 			FROM tx_commodity_price tcp
-- 			JOIN tm_city tc ON tc.id = tcp.city_id
-- 			WHERE
-- 				tcp.commodity_id IN (
-- 				    SELECT id
-- 				    FROM tm_commodity
-- 				    WHERE parent_id = @commodityId
-- 				) AND
-- 				province_id = @provinceId AND
-- 				(last_update BETWEEN @startDate AND @endDate)
-- 			GROUP BY tc.id, name, "sequence"
-- 			ORDER BY "sequence"
		) a
		OFFSET @page
		LIMIT @limit
	`

	ReportPriceByCommodityAvgChildDownload = `
		SELECT
			name title,
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
								SELECT rupiah_format(price)
								FROM price_city_avg_child(@commodityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
								LIMIT 1
							)
						)
					)
				FROM (
					SELECT *
					FROM (
						SELECT
							TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
						FROM tx_commodity_price
						WHERE city_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
						GROUP BY DATE_TRUNC('month', last_update)
						ORDER BY last_update
					) priceData
				) c
			) prices
		FROM(
			SELECT
				tc.id,
				tc.name,
				"sequence" 
			FROM tm_city tc
			--	FROM tx_commodity_price tcp
			--	JOIN tm_city tc ON tc.id = tcp.city_id
			WHERE
			--		tcp.commodity_id IN (
			--			SELECT id
			--			FROM tm_commodity
			--			WHERE parent_id = 1
			--		) AND
				province_id = @provinceId
			--		AND
			--		(last_update BETWEEN '2024-07-01 00:00:00' AND '2024-10-31 23:59:59')
			--	GROUP BY tc.id, name, "sequence"
			ORDER BY "sequence"
-- 			SELECT
-- 				tc.id,
-- 				tc.name,
-- 				"sequence" 
-- 			FROM tx_commodity_price tcp
-- 			JOIN tm_city tc ON tc.id = tcp.city_id
-- 			WHERE
-- 				tcp.commodity_id IN (
-- 				    SELECT id
-- 				    FROM tm_commodity
-- 				    WHERE parent_id = @commodityId
-- 				) AND
-- 				province_id = @provinceId AND
-- 				(last_update BETWEEN @startDate AND @endDate)
-- 			GROUP BY tc.id, name, "sequence"
-- 			ORDER BY "sequence"
		) a
	`

	ReportPriceByCommodityCount = `
		SELECT
			COUNT(1)
		FROM(
			SELECT
				tc.id,
				tc.name,
				"sequence" 
			FROM tx_commodity_price tcp
			JOIN tm_city tc ON tc.id = tcp.city_id
			WHERE
				tcp.commodity_id = @commodityId AND
				province_id = @provinceId AND
				(last_update BETWEEN @startDate AND @endDate)
			GROUP BY tc.id, name, "sequence"
			ORDER BY "sequence"
		) a
	`

	ReportPriceByCommodityCountAvgChild = `
		SELECT
			COUNT(1)
		FROM(
			SELECT
				tc.id,
				tc.name,
				"sequence" 
			FROM tm_city tc 
			WHERE
				province_id = @provinceId
			GROUP BY tc.id, name, "sequence"
			ORDER BY "sequence"
		) a
	`

	ReportPriceByCity = `
		SELECT
			name title,
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
								CASE
								    WHEN a.parent_id IS NULL
									THEN (
									    SELECT rupiah_format(price)
										FROM get_harga_by_city_id(@cityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
										LIMIT 1
									) 
								    ELSE (
								    	SELECT rupiah_format(price)
										FROM get_harga_by_city_id_avg_child(@cityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
										LIMIT 1
								    )
								END
							)
						)
					)
				FROM (
					SELECT *
					FROM (
						SELECT
							TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
						FROM tx_commodity_price
						WHERE commodity_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
						GROUP BY DATE_TRUNC('month', last_update)
						ORDER BY last_update
					) priceData
				) c
			) prices
		FROM(
			SELECT
				tc.id,
				tc.name,
				tc.parent_id
			FROM tx_commodity_price tcp
			JOIN tm_commodity tc ON tc.id = tcp.commodity_id
			WHERE
				tcp.city_id = @cityId AND
				(last_update BETWEEN @startDate AND @endDate)
			GROUP BY tc.id, name
			ORDER BY name
		) a
		OFFSET @page
		LIMIT @limit
	`

	ReportPriceByCityDownload = `
		SELECT
			name title,
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							TO_CHAR(CONCAT(last_update, '-01 00:00:00')::timestamp, 'MMYYYY'), (
								CASE
								    WHEN a.parent_id IS NULL
									THEN (
									    SELECT rupiah_format(price)
										FROM get_harga_by_city_id(@cityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
										LIMIT 1
									) 
								    ELSE (
								    	SELECT rupiah_format(price)
										FROM get_harga_by_city_id_avg_child(@cityId, a.id, CONCAT(last_update, '-01'), end_of_month(CONCAT(last_update, '-01')::DATE)::TEXT)
										LIMIT 1
								    )
								END
							)
						)
					)
				FROM (
					SELECT *
					FROM (
						SELECT
							TO_CHAR(DATE_TRUNC('month', last_update), 'YYYY-mm') AS last_update
						FROM tx_commodity_price
						WHERE commodity_id = a.id AND (last_update BETWEEN @startDate AND @endDate)
						GROUP BY DATE_TRUNC('month', last_update)
						ORDER BY last_update
					) priceData
				) c
			) prices
		FROM(
			SELECT
				tc.id,
				tc.name,
				tc.parent_id
			FROM tx_commodity_price tcp
			JOIN tm_commodity tc ON tc.id = tcp.commodity_id
			WHERE
				tcp.city_id = @cityId AND
				(last_update BETWEEN @startDate AND @endDate)
			GROUP BY tc.id, name
			ORDER BY name
		) a
	`

	ReportPriceByCityCount = `
		SELECT
			COUNT(1)
		FROM (
			SELECT
				tc.id,
				tc.name 
			FROM tx_commodity_price tcp
			JOIN tm_commodity tc ON tc.id = tcp.commodity_id
			WHERE
				tcp.city_id = @cityId AND
				(last_update BETWEEN @startDate AND @endDate)
			GROUP BY tc.id, name
			ORDER BY name
		) a
	`
)
