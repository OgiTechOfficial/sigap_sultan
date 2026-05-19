package models

type TmJenisInformasiResponse struct {
	Id
	ClientId             int32                         `json:"clientId" db:"client_id"`
	ParentId             *int32                        `json:"parentId" db:"parent_id"`
	Name                 string                        `json:"name" db:"name"`
	DetailJenisInformasi *DetailJenisInformasiResponse `json:"detailJenisInformasi" db:"detailJenisInformasi"`
	AuditRail
}

type DetailJenisInformasiResponse []struct {
	Id
	ParentId *int32 `json:"parentId" db:"parentId"`
	Name     string `json:"name" db:"name"`
	Code     string `json:"code" db:"code"`
}
