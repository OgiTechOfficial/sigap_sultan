package models

type TmJenisInformasi struct {
	Id
	ClientId int32   `json:"clientId" db:"client_id"`
	ParentId *string `json:"parentId" db:"parent_id"`
	Name     string  `json:"name" db:"name"`
	AuditRail
}

type JenisInformasiRequest struct {
	Name     string  `json:"name" db:"name"`
	ParentId *string `json:"parentId" db:"parent_id"`
}
