package models

type TxFileUploadHistory struct {
	Id
	FileName   string  `json:"fileName" db:"file_name"`
	RowTotal   int     `json:"rowTotal" db:"row_total"`
	Status     int     `json:"status" db:"status"`
	ModuleType string  `json:"moduleType" db:"module_type"`
	Errors     *string `json:"errors" db:"errors"`
	AuditRail
}
