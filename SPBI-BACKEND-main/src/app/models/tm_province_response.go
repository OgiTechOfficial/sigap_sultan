package models

type TmProvinceResponse struct {
	Id
	Name   string          `json:"name" db:"name"`
	Assets *AssetsResponse `json:"assets" db:"assets"`
	AuditRail
}
