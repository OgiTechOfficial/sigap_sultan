package models

import "github.com/golang-jwt/jwt/v4"

type LoginResponse struct {
	Id
	Name           string            `json:"name" db:"name"`
	Position       string            `json:"position" db:"position"`
	PositionId     int               `json:"position_id" db:"position_id"`
	AccessibleMenu []*AccessibleMenu `json:"accessible_menu" db:"-"`
	Token          string            `json:"token" db:"-"`
	jwt.RegisteredClaims
}

type ProfileResponse struct {
	Id
	Username        string  `json:"username" db:"username"`
	Email           string  `json:"email" db:"email"`
	NamaDepan       string  `json:"nama_depan" db:"nama_depan"`
	NamaBelakang    *string `json:"nama_belakang" db:"nama_belakang"`
	Jabatan         *string `json:"jabatan" db:"jabatan"`
	JabatanId       *int    `json:"jabatan_id" db:"jabatan_id"`
	Organisasi      *string `json:"organisasi" db:"organisasi"`
	BidangUnitKerja *string `json:"bidang_unit_kerja" db:"bidang_unit_kerja"`
}

type ForgotResponse struct {
	Link string `json:"link"`
}

type AccessibleMenu struct {
	Id
	Name     string `json:"name" db:"name"`
	Position string `json:"position" db:"position"`
}

type LoginRequestParams struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ForgotRequestParams struct {
	Email string `json:"email" validate:"required"`
	IsCms int    `json:"is_cms"`
}

type ResetRequestParams struct {
	Code               string `json:"code" validate:"required"`
	NewPassword        string `json:"newPassword" validate:"required"`
	NewPasswordConfirm string `json:"newPasswordConfirm" validate:"required"`
}

type ForgotTokenParam struct {
	Code string `json:"code" validate:"required"`
}

type UpdateProfileParam struct {
	Username        string  `json:"username" db:"username" validate:"required"`
	Email           string  `json:"email" db:"email" validate:"required"`
	NamaDepan       string  `json:"nama_depan" db:"nama_depan" validate:"required"`
	NamaBelakang    *string `json:"nama_belakang" db:"nama_belakang"`
	Jabatan         *string `json:"jabatan" db:"jabatan" validate:"required"`
	JabatanId       *int    `json:"jabatan_id" db:"jabatan_id" validate:"required"`
	Organisasi      *string `json:"organisasi" db:"organisasi"`
	BidangUnitKerja *string `json:"bidang_unit_kerja" db:"bidang_unit_kerja"`
}

type ChangePasswordParam struct {
	OldPassword        string `json:"oldPassword" validate:"required"`
	NewPassword        string `json:"newPassword" validate:"required"`
	NewPasswordConfirm string `json:"newPasswordConfirm" validate:"required"`
}
