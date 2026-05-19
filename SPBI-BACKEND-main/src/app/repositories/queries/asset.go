package queries

const (
	AssetGetSearchByName = `
		SELECT * 
		FROM assets
		WHERE assets_name LIKE @name
		LIMIT 1
	`
)
