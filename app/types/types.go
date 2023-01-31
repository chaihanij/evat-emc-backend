package types

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

type TeamType string

const (
	TeamTypePopulation TeamType = "POPULATION"
	UserRoleStudent    TeamType = "STUDENT"
)

func isValidTeamType(fl validator.FieldLevel) bool {
	status := TeamType(fl.Field().String())
	switch status {
	case TeamTypePopulation, UserRoleStudent:
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

type MemberType string

const (
	MemberTypeMember MemberType = "MEMBER"
	MemberTypeMentor MemberType = "MENTOR"
)

func isValidMemberType(fl validator.FieldLevel) bool {
	status := MemberType(fl.Field().String())
	switch status {
	case MemberTypeMember, MemberTypeMentor:
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
	case "teamType":
		return "invalid type"
	}
	return fe.Error() // default error
}

func HandleValidateError(err error, req interface{}) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			fieldName := fe.Field()
			field, ok := reflect.TypeOf(req).Elem().FieldByName(fieldName)
			if ok {
				fieldName, ok := field.Tag.Lookup("json")
				if ok {
					msg := fmt.Sprintf("%s %s", fieldName, MsgForTag(fe))
					return errors.New(msg)
				}
			} else {
				msg := fmt.Sprintf("%s %s", fieldName, MsgForTag(fe))
				return errors.New(msg)
			}
		}
	}
	return nil
}

func init() {
	_ = Validate.RegisterValidation("teamType", isValidTeamType)
	_ = Validate.RegisterValidation("userRole", isValidUserRole)
	_ = Validate.RegisterValidation("memberType", isValidMemberType)
	_ = Validate.RegisterValidation("passwordComplexity", isValidatePasswordComplexity)
	Validate.RegisterTagNameFunc(GetTagName)
}
