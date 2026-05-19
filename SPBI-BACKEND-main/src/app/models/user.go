package models

import "time"

type User struct {
	Id
	ClientId *string `json:"clientId"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	AuditRail
}

type TmUserForgotToken struct {
	Id
	UserId    *int       `json:"user_id"`
	Token     *string    `json:"token"`
	ExpiredAt *time.Time `json:"expired_at"`
	AuditRail
}
