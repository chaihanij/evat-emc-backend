package entities

import (
	"time"
)

type User struct {
	ID            string
	UID           string
	Username      string
	Email         string
	FirstName     string
	LastName      string
	Address       string
	Tel           string
	Occupation    string
	Academy       string
	Role          string
	Password      string
	Year          string
	TeamUUID      string
	Team          interface{}
	IsEmailVerify bool
	ActivateCode  string
	AccessToken   string
	IsActive      bool
	LastLogin     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     string
	UpdatedBy     string
}

type UserMinimalCreate struct {
	Email    string
	Password string
}

type UserPartialUpdate struct {
	ID            *string
	UID           string
	Username      *string
	Email         *string
	FirstName     *string
	LastName      *string
	Address       *string
	Tel           *string
	Occupation    *string
	Academy       *string
	Role          *string
	Password      *string
	Year          *string
	TeamUUID      *string
	IsEmailVerify *bool
	ActivateCode  *string
	AccessToken   *string
	IsActive      *bool
	LastLogin     *time.Time
	UpdatedBy     *string
}

type UserFilter struct {
	UID           *string
	Email         *string
	Year          *string
	IsEmailVerify *bool
	IsActive      *bool
	ActivateCode  *string
	AccessToken   *string

	Page     *int64
	PageSize *int64
}

type ResetPassword struct {
	UID             string
	TeamUUID        string
	ActivateCode    string
	Email           string
	Otp             string
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
}
