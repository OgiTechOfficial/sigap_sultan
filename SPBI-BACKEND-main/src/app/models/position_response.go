package models

type TmPositionResponse struct {
	Id
	ClientId   int32               `json:"clientId" db:"client_id"`
	Name       string              `json:"name" db:"name"`
	Privileges *PrivilegesResponse `json:"privileges" db:"privileges"`
	AuditRail
}

type PrivilegesResponse []struct {
	Menu        string             `json:"menu" db:"menu"`
	Permissions PermissionResponse `json:"permissions" db:"permissions"`
}

type PermissionResponse struct {
	Read   int8 `json:"read" db:"read"`
	Create int8 `json:"create" db:"create"`
	Update int8 `json:"update" db:"update"`
	Delete int8 `json:"delete" db:"delete"`
}
