package models

import "time"

type Id struct {
	Id *int32 `json:"id" db:"id"`
}

type AuditRail struct {
	CreatedAt *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

type TitleColor struct {
	Title string `json:"title"`
	Color string `json:"color"`
}

type NeracaTitleColor struct {
	Title           string `json:"title"`
	Start           int32  `json:"start"`
	End             int32  `json:"end"`
	Color           string `json:"color"`
	Unit            string `json:"unit"`
	BackgroundColor string `json:"backgroundColor"`
}

type NeracaKetersediaanColor struct {
	Title string `json:"title"`
	Color string `json:"color"`
}
