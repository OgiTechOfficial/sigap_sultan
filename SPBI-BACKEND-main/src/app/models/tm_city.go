package models

type TmCity struct {
	Id
	ProvinceId       int32  `json:"provinceId" db:"province_id"`
	Name             string `json:"name" db:"name"`
	Sequence         int32  `json:"sequence" db:"sequence"`
	AssetsRelationId int32  `json:"assets_relation_id" db:"assets_relation_id"`
	AuditRail
}
