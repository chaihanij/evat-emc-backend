package models

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/logger"
)

func Test_UpdateMember(t *testing.T) {
	logger.Init()
	member := &entities.Member{
		FirstName: "วชิรวัฒน์",
		LastName:  "ไชยฮะนิจ",
		Address:   "292/510 ถนนรัชดาภิเษก แขวงลาดยาว เขตจตุจักร กรุงเทพมหานคร 10900",
		Email:     "chaihanij@gmail.com",
		Tel:       "0866304634",
		Academy:   "มหาวิทยาลัยเทคโนโลยีสุรนารี",
		Year:      "2023",
		Image:     "c11eecdd-f879-4608-b47b-6babe6a57179",
		Documents: []string{"d0e1b5d6-b0d0-4b11-b2bb-c4668344d965", "d08afe43-f546-4591-afc5-e5b2e7a3b59a"},
	}

	actual := UpdateMember(member)
	assert.NotNil(t, actual)
	if assert.NotNil(t, actual) {
		fmt.Println("value", actual)
		fmt.Println("reflect", reflect.TypeOf(actual))
	}
}
