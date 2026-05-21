package queries

const (
	TmCommodityList = `
			SELECT a.id, a.client_id, a.parent_id, NULL::varchar as code, a.name, a.created_at, a.updated_at, a.deleted_at, a.assets_relation_id, a.sequence, a.unit_id, a.unit_id_neraca, b.name as unit_name 
			FROM tm_commodity a
			LEFT JOIN settings b ON b.id = a.unit_id 
			ORDER BY sequence
		`

	TmCommodityListNeraca = `
			SELECT a.id, a.client_id, a.parent_id, NULL::varchar as code, a.name, a.created_at, a.updated_at, a.deleted_at, a.assets_relation_id, a.sequence, a.unit_id, a.unit_id_neraca, b.name as unit_name 
			FROM tm_commodity a
			LEFT JOIN settings b ON b.id = a.unit_id_neraca 
			ORDER BY sequence
		`

	TmCommodityParentListOrderBySequence = `
			SELECT a.id, a.client_id, a.parent_id, NULL::varchar as code, a.name, a.created_at, a.updated_at, a.deleted_at, a.assets_relation_id, a.sequence 
			FROM tm_commodity a
			WHERE a.parent_id IS NULL
			ORDER BY sequence
		`

	TmCommodityGetById = `
			SELECT a.id, a.client_id, a.parent_id, NULL::varchar as code, a.name, a.created_at, a.updated_at, a.deleted_at, a.assets_relation_id, a.sequence, a.unit_id, a.unit_id_neraca, b.name as unit_name 
			FROM tm_commodity a
			LEFT JOIN settings b ON b.id = a.unit_id 
			WHERE a.id = @id
		`

	TmCommodityGetByParentId = `
			SELECT a.id, a.client_id, a.parent_id, NULL::varchar as code, a.name, a.created_at, a.updated_at, a.deleted_at, a.assets_relation_id, a.sequence 
			FROM tm_commodity a
			WHERE a.parent_id = @id
		`

	TmCommodityGetByName = `
			SELECT *
			FROM (
			    SELECT a.id, a.client_id, a.parent_id, NULL::varchar as code, a.name, a.created_at, a.updated_at, a.deleted_at, a.assets_relation_id, a.sequence, a.unit_id, a.unit_id_neraca, b.name as unit_name 
				FROM tm_commodity a
				LEFT JOIN settings b ON b.id = a.unit_id 
				WHERE LOWER(a.name) LIKE @name 
			) B
		`

	TmCommodityGetByEqualName = `
			SELECT *
			FROM (
			    SELECT a.id, a.client_id, a.parent_id, NULL::varchar as code, a.name, a.created_at, a.updated_at, a.deleted_at, a.assets_relation_id, a.sequence, a.unit_id, a.unit_id_neraca, b.name as unit_name 
				FROM tm_commodity a
				LEFT JOIN settings b ON b.id = a.unit_id 
				WHERE LOWER(a.name) = @name 
			) B
		`

	TmCommodityInsertParentWhenNotExists = `
			INSERT INTO tm_commodity(
				client_id,
				name,
				created_at,
				assets_relation_id
			)
			SELECT
				1,
				CAST(@name AS VARCHAR),
				@createdAt,
				CAST(@assetRelationId AS NUMERIC)
			FROM tm_commodity
			WHERE NOT EXISTS (
				SELECT name
				FROM tm_commodity
				WHERE LOWER(name) = @name
			)
			LIMIT 1
			RETURNING id
		`

	TmCommodityInsertChildWhenNotExists = `
			INSERT INTO tm_commodity(
				client_id,
				parent_id,
				name,
				created_at,
				assets_relation_id
			)
			SELECT
				1,
				CAST(@parentId AS NUMERIC),
				CAST(@name AS VARCHAR),
				@createdAt,
				CAST(@assetRelationId AS NUMERIC)
			FROM tm_commodity
			WHERE NOT EXISTS (
				SELECT name
				FROM tm_commodity
				WHERE LOWER(name) = @name
			)
			LIMIT 1
			RETURNING id
		`

	TmCommodityIsParent = `
		SELECT EXISTS(
			SELECT
				id
			FROM tm_commodity
			WHERE id = @commodityId AND parent_id IS NULL
		)
	`
)
