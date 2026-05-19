package queries

const (
	TmPositionList = `
		SELECT
			a.id,
			a.client_id,
			a.name,
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'menu', c.name,
					'permissions', JSON_BUILD_OBJECT(
						'read', b.read_status,
						'create', b.create_status,
						'update', b.update_status,
						'delete', b.delete_status
					)
				)
			) privileges,
			a.created_at,
			a.updated_at,
			a.deleted_at
		FROM tm_position a
		JOIN map_jabatan_menu b ON a.id = b.position_id
		JOIN tm_menu c on b.menu_id = c.id
		WHERE a.deleted_at is null
		GROUP BY a.id, a.name, a.client_id, a.id, a.created_at, a.updated_at, a.deleted_at
		ORDER BY name
		OFFSET @page
		LIMIT @limit
	`

	TmPositionCount = `
		SELECT COUNT(1)
		FROM tm_position a
		JOIN map_jabatan_menu b ON a.id = b.position_id
		JOIN tm_menu c on b.menu_id = c.id
		WHERE a.deleted_at is null
		GROUP BY a.id, a.name, a.client_id, a.id, a.created_at, a.updated_at, a.deleted_at
	`

	TmPositionGetById = `
		SELECT
			a.id,
			a.client_id,
			a.name,
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'menu', c.name,
					'permissions', JSON_BUILD_OBJECT(
						'read', b.read_status,
						'create', b.create_status,
						'update', b.update_status,
						'delete', b.delete_status
					)
				)
			) privileges,
			a.created_at,
			a.updated_at,
			a.deleted_at
		FROM tm_position a
		JOIN map_jabatan_menu b ON a.id = b.position_id
		JOIN tm_menu c on b.menu_id = c.id
		WHERE a.deleted_at is null and a.id = @id
		GROUP BY a.id, a.name, a.client_id, a.id, a.created_at, a.updated_at, a.deleted_at
	`

	TmPositionGetByName = `
		SELECT
			a.id,
			a.client_id,
			a.name,
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'menu', c.name,
					'permissions', JSON_BUILD_OBJECT(
						'read', b.read_status,
						'create', b.create_status,
						'update', b.update_status,
						'delete', b.delete_status
					)
				)
			) privileges,
			a.created_at,
			a.updated_at,
			a.deleted_at
		FROM tm_position a
		JOIN map_jabatan_menu b ON a.id = b.position_id
		JOIN tm_menu c on b.menu_id = c.id
		WHERE a.deleted_at is null and LOWER(a.name) LIKE @name
		GROUP BY a.id, a.name, a.client_id, a.id, a.created_at, a.updated_at, a.deleted_at
	`

	TmPositionGetByNameCount = `
		SELECT COUNT(1)
		FROM tm_position a
		JOIN map_jabatan_menu b ON a.id = b.position_id
		JOIN tm_menu c on b.menu_id = c.id
		WHERE a.deleted_at is null and LOWER(a.name) LIKE @name
		GROUP BY a.id, a.name, a.client_id, a.id, a.created_at, a.updated_at, a.deleted_at
	`

	TmPositionInsert = `
		INSERT INTO tm_position(
			client_id,
			name,
			created_at
		) VALUES(
			$1,
			$2,
			$3
		)
		RETURNING id
	`

	//TmPositionInsert = `
	//	INSERT INTO tm_position(
	//		client_id,name,created_at
	//	) VALUES
	//`

	TmPositionUpdate = `
		UPDATE tm_position
		SET name = @name
		WHERE id = @id
	`

	TmPositionDelete = `
		DELETE FROM tm_position
		WHERE id = @id
	`

	TmPositionSoftDelete = `
		UPDATE tm_position
		SET deleted_at = now()
		WHERE id = @id
	`
)
