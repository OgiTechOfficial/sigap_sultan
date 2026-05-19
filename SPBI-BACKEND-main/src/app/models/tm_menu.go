package models

type TmMenu struct {
	Id
	ClientId int32  `json:"clientId" db:"client_id"`
	Name     string `json:"name" db:"name"`
	Position int8   `json:"position" db:"position"`
	AuditRail
	Url string `json:"url" db:"url"`
}
