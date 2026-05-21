package queries

const (
	TmCityGet = `
			SELECT id, province_id, name, created_at, updated_at, deleted_at, sequence, assets_relation_id 
			FROM tm_city
			ORDER BY name
			OFFSET @page
			LIMIT @limit
		`

	TmCityGetCount = `
			SELECT COUNT(1)
			FROM tm_city
		`

	TmCityGetByName = `
			SELECT id, province_id, name, created_at, updated_at, deleted_at, sequence, assets_relation_id 
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
			SELECT id, province_id, name, created_at, updated_at, deleted_at, sequence, assets_relation_id 
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
		SELECT id, province_id, name, created_at, updated_at, deleted_at, sequence, assets_relation_id 
		FROM tm_city
		WHERE id = @id
	`
)
