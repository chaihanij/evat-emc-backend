package types

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

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

func GetTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

func MsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "field is required"
	case "email":
		return "invalid email"
	case "passwordComplexity":
		return "weak password"
	case "userRole":
		return "invalid role"
	}
	return fe.Error() // default error
}

func init() {
	_ = Validate.RegisterValidation("userRole", isValidUserRole)
	_ = Validate.RegisterValidation("passwordComplexity", isValidatePasswordComplexity)
	Validate.RegisterTagNameFunc(GetTagName)
}
