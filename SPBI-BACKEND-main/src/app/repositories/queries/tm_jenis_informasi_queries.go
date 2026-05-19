package queries

const (
	TmJenisInformasiList = `
		SELECT
			id,
			client_id,
			parent_id,
			name,
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', id,
							'parentId', b.parent_id,
							'name', name,
							'code', code
						)
					)
				FROM (
					SELECT
						id,
						parent_id,
						name,
						code
					FROM tm_jenis_informasi
					WHERE parent_id IS NOT NULL AND parent_id = a.id
					ORDER BY id
				) b
			) detail_jenis_informasi,
			created_at,
			updated_at,
			deleted_at
		FROM tm_jenis_informasi a 
		WHERE parent_id IS NULL 
		ORDER BY name
	`

	TmJenisInformasiGetByName = `
		SELECT
			id,
			client_id,
			name,
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', id,
							'parentId', parent_id,
							'name', name,
							'code', code
						)
					)
				FROM tm_jenis_informasi
				WHERE parent_id IS NOT NULL AND parent_id = a.id
			) detail_jenis_informasi,
			created_at,
			updated_at,
			deleted_at
		FROM (
			SELECT * FROM tm_jenis_informasi
			ORDER BY name
		) a
		WHERE LOWER(name) LIKE @name
	`

	TmJenisInformasiGetById = `
		SELECT
			id,
			client_id,
			name,
			(
				SELECT
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'id', id,
							'parentId', parent_id,
							'name', name,
							'code', code
						)
					)
				FROM tm_jenis_informasi
				WHERE parent_id IS NOT NULL AND parent_id = a.id
			) detail_jenis_informasi,
			created_at,
			updated_at,
			deleted_at
		FROM (
			SELECT * FROM tm_jenis_informasi
			ORDER BY name
		) a
		WHERE id = @id
	`
)
