package queries

const (
	TmMenuList = `
		SELECT 
			id,
			client_id,
			name,
			position,
			created_at,
			updated_at,
			deleted_at,
			url
		FROM tm_menu
		ORDER BY position
		OFFSET @page
		LIMIT @limit
	`

	TmMenuListCount = `
		SELECT COUNT(1) 
		FROM tm_menu
	`

	TmMenuGetByName = `
		SELECT 
			id,
			client_id,
			name,
			position,
			created_at,
			updated_at,
			deleted_at,
			url
		FROM tm_menu
		WHERE LOWER(name) LIKE @name
		ORDER BY name
		OFFSET @page
		LIMIT @limit
	`

	TmMenuGetByNameCount = `
		SELECT COUNT(1) 
		FROM tm_menu
		WHERE LOWER(name) LIKE @name
	`

	TmMenuGetById = `
		SELECT 
			id,
			client_id,
			name,
			position,
			created_at,
			updated_at,
			deleted_at,
			url
		FROM tm_menu
		WHERE id = @id
	`

	TmMenuInsert = `
		INSERT INTO tm_menu(
			client_id,
			name,
			position,
			created_at
		) VALUES($1, $2, $3, $4)
		RETURNING id
	`

	TmUserForgotToken = `
		INSERT INTO tm_user_forgot_token(
			user_id,
			token,
			expired_at,
			created_at
		) VALUES($1, $2, $3, $4)
		RETURNING id
	`

	TmMenuUpdate = `
		UPDATE 
			tm_menu
		SET client_id = @clientId, name = @name, position = @position, updated_at = @updatedAt
		WHERE id = @id
	`

	TmMenuDeleteMaster = `
		DELETE FROM tm_menu WHERE id = @id
	`

	TmMenuDeleteChild = `
		DELETE FROM map_jabatan_menu WHERE menu_id IN (SELECT id FROM tm_menu WHERE id = @id)
	`
)
