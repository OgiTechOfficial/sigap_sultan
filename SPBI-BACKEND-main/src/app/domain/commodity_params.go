package domain

type TmCommodityParam struct {
	ClientId         int32   `json:"clientId" db:"client_id"`
	ParentId         *int32  `json:"parentId" db:"parent_id"`
	Code             *string `json:"code" db:"code"`
	Name             string  `json:"name" db:"name"`
	AssetsRelationId *int32  `json:"assetsRelationId" db:"assets_relation_id"`
}

type TmCommodityRequestParam struct {
	Code       *string `json:"code" db:"code"`
	Name       string  `json:"name" db:"name"`
	ModuleType string  `json:"moduleType" db:"moduleType"`
}
