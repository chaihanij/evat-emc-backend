package entities

import (
	"time"

	"gitlab.com/chaihanij/evat/app/types"
)

type User struct {
	ID            string
	Username      string
	UID           string
	Email         string
	FirstName     string
	LastName      string
	Address       string
	Tel           string
	Role          types.UserRole
	Password      string
	Year          string
	TeamUID       string
	IsEmailVerify bool
	ActivateCode  string
	AccessToken   string
	IsActive      bool
	LastLogin     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserCreate struct {
	Username      string
	Email         string
	FirstName     *string
	LastName      *string
	Address       *string
	Tel           *string
	Role          *types.UserRole
	Password      *string
	Year          *string
	TeamUID       *string
	IsEmailVerify *bool
	ActivateCode  *string
	AccessToken   *string
	IsActive      *bool
}

type UserMinimalCreate struct {
	Email    string
	Password string
}

type UserPartialUpdate struct {
	UID           *string
	Username      *string
	Email         *string
	FirstName     *string
	LastName      *string
	Address       *string
	Tel           *string
	Role          *types.UserRole
	Password      *string
	Year          *string
	TeamUID       *string
	IsEmailVerify *bool
	ActivateCode  *string
	AccessToken   *string
	IsActive      *bool
	LastLogin     *time.Time
}

type UserFilter struct {
	UID   *string
	Email *string
	Year  *string

	IsEmailVerify *bool
	IsActive      *bool
	ActivateCode  *string
	AccessToken   *string
	Page          *int64
	PageSize      *int64
}

type ResetPassword struct {
	UID             string
	Email           string
	Otp             string
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
}
