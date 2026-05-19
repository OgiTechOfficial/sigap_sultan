package price

const (
	PriceExist = `
		SELECT
			JSON_OBJECT_AGG(
				TO_CHAR(DATE_TRUNC('day', b.date_range):: date, 'YYYYMMDD'), (
					SELECT EXISTS(
						SELECT
							*
						FROM tx_commodity_price a
						WHERE
							commodity_id = @commodityId AND
							last_update BETWEEN CONCAT(TO_CHAR(DATE_TRUNC('day', b.date_range):: date, 'YYYY-MM-DD'), ' 00:00:00')::TIMESTAMPTZ AND CONCAT(TO_CHAR(DATE_TRUNC('day', b.date_range):: date, 'YYYY-MM-DD'), ' 23:59:59')::TIMESTAMPTZ
					)
				)
			) price_exists_date
		FROM(
		    SELECT 
		    	generate_series(
		    		CONCAT(start_date, ' 00:00:00')::TIMESTAMPTZ,
		    		CONCAT(end_date,' 23:59:59')::TIMESTAMPTZ,
		    		'1 day'::INTERVAL
		    	) date_range
		    FROM (
		        SELECT
		            @startDate "start_date",
            		@endDate "end_date"
		    ) dateRange
		) b
	`

	PriceExistParent = `
		SELECT
			JSON_OBJECT_AGG(
				TO_CHAR(DATE_TRUNC('day', b.date_range):: date, 'YYYYMMDD'), (
					SELECT EXISTS(
						SELECT
							*
						FROM tx_commodity_price a
						WHERE
							commodity_id IN (
							    SELECT id
							    FROM tm_commodity
							    WHERE parent_id = @commodityId
							) AND
							last_update BETWEEN CONCAT(TO_CHAR(DATE_TRUNC('day', b.date_range):: date, 'YYYY-MM-DD'), ' 00:00:00')::TIMESTAMPTZ AND CONCAT(TO_CHAR(DATE_TRUNC('day', b.date_range):: date, 'YYYY-MM-DD'), ' 23:59:59')::TIMESTAMPTZ
					)
				)
			) price_exists_date
		FROM(
		    SELECT 
		    	generate_series(
		    		CONCAT(start_date, ' 00:00:00')::TIMESTAMPTZ,
		    		CONCAT(end_date,' 23:59:59')::TIMESTAMPTZ,
		    		'1 day'::INTERVAL
		    	) date_range
		    FROM (
		        SELECT
		            @startDate "start_date",
            		@endDate "end_date"
		    ) dateRange
		) b
	`

	PriceLatestDateExist = `
		SELECT
			TO_CHAR(DATE_TRUNC('day', a.last_update):: date, 'YYYY-MM-DD')
		FROM tx_commodity_price a
		WHERE
			commodity_id = @commodityId
		ORDER BY last_update DESC
		LIMIT 1
	`

	PriceLatestDateExistWithoutCommodity = `
		SELECT
			TO_CHAR(DATE_TRUNC('day', a.last_update):: date, 'YYYY-MM-DD')
		FROM tx_commodity_price a
		ORDER BY last_update DESC
		LIMIT 1
	`

	PriceLatestDateExistChild = `
		SELECT
			TO_CHAR(DATE_TRUNC('day', a.last_update):: date, 'YYYY-MM-DD')
		FROM tx_commodity_price a
		WHERE
			commodity_id IN (
			    SELECT id
			    FROM tm_commodity
			    WHERE parent_id = @commodityId
			) OR
			commodity_id = @commodityId
		ORDER BY last_update DESC
		LIMIT 1
	`
)
