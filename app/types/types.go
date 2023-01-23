package types

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type UserType string

const (
	UserTypeAdviser   UserType = "ADVISER"
	UserTypeMember    UserType = "MEMBER"
	UserTypeCommittee UserType = "COMMMITEE"
)

func isValidUserType(fl validator.FieldLevel) bool {
	status := UserType(fl.Field().String())
	switch status {
	case UserTypeAdviser, UserTypeMember, UserTypeCommittee:
		return true
	default:
		return false
	}
}

type UserRole string

const (
	UserRoleSuperAdmin UserRole = "SUPER_ADMIN"
	UserRoleAdmin      UserRole = "ADMIN"
	UserRoleUSER       UserRole = "USER"
	UserRoleCommittee  UserRole = "COMMMITEE"
)

func isValidUserRole(fl validator.FieldLevel) bool {
	status := UserRole(fl.Field().String())
	switch status {
	case UserRoleSuperAdmin, UserRoleAdmin, UserRoleUSER, UserRoleCommittee:
		return true
	default:
		return false
	}
}

func isValidatePasswordComplexity(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	// Must have at least one lower case
	isValid := regexp.
		MustCompile(`[a-z]`).
		MatchString(password)
	if !isValid {
		return false
	}

	// Must have at least one upper case
	isValid = regexp.
		MustCompile(`[A-Z]`).
		MatchString(password)
	if !isValid {
		return false
	}

	// Must have at least one number
	isValid = regexp.
		MustCompile(`\d`).
		MatchString(password)
	if !isValid {
		return false
	}

	// Must have at least one symbol
	isValid = regexp.
		MustCompile(`[-+_!@#$%^&*.,?]`).
		MatchString(password)
	if !isValid {
		return false
	}

	return true
}

var Validate = validator.New()

func init() {
	_ = Validate.RegisterValidation("userType", isValidUserType)
	_ = Validate.RegisterValidation("userRole", isValidUserRole)
	_ = Validate.RegisterValidation("passwordComplexity", isValidatePasswordComplexity)
}

func MsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "passwordComplexity":
		return "Invalid email"
	}

	return fe.Error() // default error
}
