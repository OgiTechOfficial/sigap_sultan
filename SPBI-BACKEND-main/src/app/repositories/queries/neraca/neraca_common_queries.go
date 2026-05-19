package neraca

const (
	NeracaExist = `
		SELECT
			JSON_OBJECT_AGG(
				TO_CHAR(DATE_TRUNC('day', b.date_range):: date, 'YYYYMMDD'), (
					SELECT EXISTS(
						SELECT
							*
						FROM tx_commodity_stock a
						WHERE
							commodity_id = @commodityId AND
							last_update BETWEEN CONCAT(TO_CHAR(DATE_TRUNC('day', b.date_range):: date, 'YYYY-MM-01'), ' 00:00:00')::TIMESTAMPTZ AND CONCAT(TO_CHAR(DATE_TRUNC('day', b.date_range):: date, 'YYYY-MM-DD'), ' 23:59:59')::TIMESTAMPTZ
					)
				)
			) price_exists_date
		FROM(
		    SELECT
				generate_series(
					start_date::DATE,
					end_date::DATE,
					'1 month'
				) date_range
		    FROM (
		        SELECT
		            @startDate "start_date",
            		@endDate "end_date"
		    ) dateRange
		) b
	`

	NeracaLatestDateExist = `
		SELECT
			TO_CHAR(DATE_TRUNC('day', a.last_update):: date, 'YYYY-MM-01')
		FROM tx_commodity_stock a
		WHERE
			commodity_id = @commodityId
		ORDER BY last_update DESC
		LIMIT 1
	`
)
