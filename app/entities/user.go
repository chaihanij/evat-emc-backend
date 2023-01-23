package entities

import (
	"gitlab.com/chaihanij/evat/app/types"
)

type User struct {
	ID            string
	UID           string
	Email         string
	FirstName     string
	LastName      string
	Address       string
	Academy       string
	Branch        string
	Tel           string
	UserType      string
	Role          string
	Image         types.File
	Documents     types.Files
	Password      string
	Year          string
	TeamUID       string
	IsEmailVerify bool
	ActivateCode  string
	IsActive      bool
	CreatedAt     types.DateTime
	UpdatedAt     types.DateTime
}

type UserCreate struct {
	Email         string
	FirstName     *string
	LastName      *string
	Address       *string
	Academy       *string
	Branch        *string
	Tel           *string
	UserType      *types.UserType
	Role          *types.UserRole
	Password      *string
	Year          *string
	TeamUID       *string
	IsEmailVerify *bool
	ActivateCode  *string
	IsActive      *bool
}

type UserMinimalCreate struct {
	Email    string
	Password string
}

type UserPartialUpdate struct {
	UID           *string
	Email         *string
	FirstName     *string
	LastName      *string
	Address       *string
	Academy       *string
	Branch        *string
	Tel           *string
	UserType      *types.UserType
	Role          *types.UserRole
	Image         *types.File
	Documents     *types.Files
	Password      *string
	Year          *string
	TeamUID       *string
	IsEmailVerify *bool
	ActivateCode  *string
	IsActive      *bool
}

type UserFilter struct {
	UID   *string
	Email *string
	Year  *string

	IsEmailVerify *bool
	IsActive      *bool
	ActivateCode  *string
	Page          *int64
	PageSize      *int64
}
