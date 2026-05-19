package models

type TmPosition struct {
	Id
	ClientId int32  `json:"clientId" db:"client_id"`
	Name     string `json:"name"`
	AuditRail
}

type PrivilegesRequest struct {
	Position   *string           `json:"position"`
	Privileges []MenuPermissions `json:"privileges"`
}

type MenuPermissions struct {
	MenuId      int32              `json:"menuId"`
	Permissions *PermissionRequest `json:"permissions"`
}

type PermissionRequest struct {
	Read   int8 `json:"read" db:"read"`
	Create int8 `json:"create" db:"create"`
	Update int8 `json:"update" db:"update"`
	Delete int8 `json:"delete" db:"delete"`
}
