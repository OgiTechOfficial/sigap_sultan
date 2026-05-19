package models

type TmNationalResponse struct {
	Id
	Name   string          `json:"name" db:"name"`
	Assets *AssetsResponse `json:"assets" db:"assets"`
	AuditRail
}
