package queries

const (
	TmProvinceGet = `
			SELECT * 
			FROM tm_province
			ORDER BY @orderBy @ascending
			OFFSET @page
			LIMIT @limit
		`

	TmProvinceGetCount = `
			SELECT COUNT(1)
			FROM tm_province
		`

	TmProvinceGetByName = `
			SELECT * 
			FROM tm_province
			WHERE LOWER(name) LIKE @name
-- 			ORDER BY sequence
			OFFSET @page
			LIMIT @limit
		`

	TmProvinceGetByNameCount = `
			SELECT COUNT(1)
			FROM tm_province
			WHERE LOWER(name) LIKE @name
		`

	TmProvinceGetById = `
			SELECT * 
			FROM tm_province
			WHERE id = @id
		`

	TmProvinceGetByIdCount = `
			SELECT COUNT(1)
			FROM tm_province
			WHERE id = @id
		`

	TmProvinceInsertWhenNotExists = `
			INSERT INTO tm_province(
				name,
				created_at
			)
			SELECT 
				@name,
				@createdAt
			FROM tm_province
			WHERE NOT EXISTS (
				SELECT name
				FROM tm_province
				WHERE LOWER(name) = @name
			)
		`
)
