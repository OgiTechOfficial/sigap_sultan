package models

type Assets struct {
	Id                 Id     `json:"id"`
	AssetsType         string `json:"assetsType" db:"assets_type"`
	AssetsLocation     string `json:"assetsLocation" db:"assets_location"`
	AssetsLocationType string `json:"assetsLocationType" db:"assets_location_type"`
	AssetsMediaType    string `json:"assetsMediaType" db:"assets_media_type"`
	AssetsExt          string `json:"assetsExt" db:"assets_ext"`
	AssetsRelationId   int    `json:"assetsRelationId" db:"assets_relation_id"`
	AuditRail
}
