package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_utils "gitlab.com/chaihanij/evat/app/utils"
)

type mockStruct struct {
	data string
}

func Test_NewSuccessResponse(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		fakeData := mockStruct{"test"}

		newSuccessResponse := _utils.NewSuccessResponse(fakeData)
		assert.Equal(t, newSuccessResponse.StatusCode, _utils.StatusCode1000)
		assert.Equal(t, newSuccessResponse.StatusMessage, _utils.StatusMessageSuccess)
		assert.Equal(t, newSuccessResponse.Data, fakeData)

	})
}

func Test_NewErrorResponse(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockErrorMessage := "Error here!!"

		newSuccessResponse := _utils.NewErrorResponse(mockErrorMessage)
		assert.Equal(t, newSuccessResponse.StatusCode, _utils.StatusCode1002)
		assert.Equal(t, newSuccessResponse.StatusMessage, _utils.StatusMassageFail)
		assert.Equal(t, newSuccessResponse.Message, mockErrorMessage)
	})
}
