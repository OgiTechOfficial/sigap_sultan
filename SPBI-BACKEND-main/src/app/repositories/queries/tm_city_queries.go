package queries

const (
	TmCityGet = `
			SELECT * 
			FROM tm_city
			ORDER BY @orderBy @ascending
			OFFSET @page
			LIMIT @limit
		`

	TmCityGetCount = `
			SELECT COUNT(1)
			FROM tm_city
		`

	TmCityGetByName = `
			SELECT * 
			FROM tm_city
			WHERE LOWER(name) LIKE @name
			ORDER BY sequence
			OFFSET @page
			LIMIT @limit
		`

	TmCityGetByNameCount = `
			SELECT COUNT(1)
			FROM tm_city
			WHERE LOWER(name) LIKE @name
		`

	TmCityGetByProvinceId = `
			SELECT * 
			FROM tm_city
			WHERE province_id = @provinceId
			ORDER BY sequence
			OFFSET @page
			LIMIT @limit
		`

	TmCityGetByProvinceIdCount = `
			SELECT COUNT(1)
			FROM tm_city
			WHERE province_id = @provinceId
		`

	TmCityGetById = `
		SELECT * 
		FROM tm_city
		WHERE id = @id
	`
)
