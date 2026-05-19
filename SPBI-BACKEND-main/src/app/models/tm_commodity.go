package models

type TmCommodity struct {
	Id
	ClientId int32   `json:"clientId" db:"client_id"`
	ParentId *int32  `json:"parentId" db:"parent_id"`
	Code     *string `json:"code" db:"code"`
	Name     string  `json:"name" db:"name"`
	AuditRail
	AssetsRelationId *int32  `json:"assetsRelationId" db:"assets_relation_id"`
	Sequence         *int32  `json:"sequence" db:"sequence"`
	UnitId           *int32  `json:"unitId" db:"unit_id"`
	UnitIdNeraca     *int32  `json:"unitIdNeraca" db:"unit_id_neraca"`
	UnitName         *string `json:"unit"`
}
