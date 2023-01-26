package utils

import (
	"net/http"
	"time"

	"gitlab.com/chaihanij/evat/app/errors"

	"github.com/gin-gonic/gin"
)

const (
	StatusCode1000 = 1000 //Success
	StatusCode1001 = 1001 //ParameterError
	StatusCode1002 = 1002 //UnprocessableEntity
	StatusCode1003 = 1003 //InternalError
	StatusCode1004 = 1004 //RecordNotFoundError
	StatusCode1005 = 1005 //Unauthorized
	StatusCode1006 = 1006 //Forbidden
	StatusCode1007 = 1007 //MongoDuplicateKeyError
	StatusCode1008 = 1008 //HttpClientError
	StatusCode1009 = 1009 //PaymentRegistrationError
	StatusCode5000 = 5000 //Default
)

const (
	StatusMassageFail    string = "fail"
	StatusMessageSuccess string = "success"
)

type BaseSuccessResponse struct {
	StatusCode    int         `json:"statusCode" example:"1000"`
	StatusMessage string      `json:"statusMessage" example:"Success"`
	Timestamp     time.Time   `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          interface{} `json:"data,omitempty"`
	MetaData      interface{} `json:"metaData,omitempty"`
}

type BaseRedirectSuccess struct {
	CardId string
	Status string
	Desc   string
	Ref    string
}

func JSONSuccessResponse(c *gin.Context, data interface{}) {
	r := new(BaseSuccessResponse)
	r.StatusCode = StatusCode1000
	r.StatusMessage = StatusMessageSuccess
	r.Timestamp = time.Now()
	r.Data = data
	c.AbortWithStatusJSON(http.StatusOK, *r)
}

func JSONSuccessCodeWithMetaDataResponse(c *gin.Context, data interface{}, metaData interface{}) {
	r := new(BaseSuccessResponse)
	r.StatusCode = StatusCode1000
	r.StatusMessage = StatusMessageSuccess
	r.Timestamp = time.Now()
	r.Data = data
	r.MetaData = metaData
	c.AbortWithStatusJSON(http.StatusOK, *r)
}

func NewSuccessResponse(data interface{}) BaseSuccessResponse {
	r := BaseSuccessResponse{
		StatusCode:    StatusCode1000,
		StatusMessage: StatusMessageSuccess,
		Timestamp:     time.Now(),
		Data:          data,
	}
	return r
}

type ErrorResponse struct {
	StatusCode    int       `json:"statusCode" example:"1001"`
	StatusMessage string    `json:"statusMessage" example:"fail"`
	Message       string    `json:"message" example:"error message will be show here"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}

func NewErrorResponse(message string) ErrorResponse {
	r := ErrorResponse{
		StatusCode:    StatusCode1002,
		StatusMessage: StatusMassageFail,
		Message:       message,
		Timestamp:     time.Now(),
	}
	return r
}

type ErrorResponseData struct {
	StatusCode    int         `json:"statusCode" example:"1001"`
	StatusMessage string      `json:"statusMessage" example:"fail"`
	Message       string      `json:"message" example:"error message will be show here"`
	Timestamp     time.Time   `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          interface{} `json:"data"`
}

func JSONErrorResponse(c *gin.Context, err error) {
	httpStatusCode := http.StatusInternalServerError
	statusCode := StatusCode5000
	message := ""

	switch err.(type) {
	case errors.ParameterError:
		httpStatusCode = http.StatusBadRequest
		statusCode = StatusCode1001
		message = err.Error()
	case errors.UnprocessableEntity:
		httpStatusCode = http.StatusBadRequest
		statusCode = StatusCode1002
		message = err.Error()
	case errors.InternalError:
		httpStatusCode = http.StatusInternalServerError
		statusCode = StatusCode1003
		message = err.Error()
	case errors.RecordNotFoundError:
		httpStatusCode = http.StatusNotFound
		statusCode = StatusCode1004
		message = err.Error()
		err = nil
	case errors.Unauthorized:
		httpStatusCode = http.StatusUnauthorized
		statusCode = StatusCode1005
		message = err.Error()
	case errors.Forbidden:
		httpStatusCode = http.StatusForbidden
		statusCode = StatusCode1006
		message = err.Error()
	case errors.DuplicateKeyError:
		httpStatusCode = http.StatusInternalServerError
		statusCode = StatusCode1007
		message = err.Error()
	case errors.HttpClientError:
		httpStatusCode = http.StatusInternalServerError
		statusCode = StatusCode1008
		message = err.Error()
	default:
		message = err.Error()
	}

	errorResponse := ErrorResponse{
		StatusCode:    statusCode,
		StatusMessage: StatusMassageFail,
		Message:       message,
		Timestamp:     time.Now(),
	}
	c.AbortWithStatusJSON(httpStatusCode, errorResponse)
}

// JSONErrorWithDataResponse response error json
func JSONErrorWithDataResponse(c *gin.Context, err error, data interface{}) {
	httpStatusCode := http.StatusInternalServerError
	statusCode := StatusCode5000
	message := ""

	switch err.(type) {
	case errors.ParameterError:
		httpStatusCode = http.StatusBadRequest
		statusCode = StatusCode1001
		message = err.Error()
	case errors.UnprocessableEntity:
		httpStatusCode = http.StatusBadRequest
		statusCode = StatusCode1002
		message = err.Error()
	case errors.InternalError:
		httpStatusCode = http.StatusInternalServerError
		statusCode = StatusCode1003
		message = err.Error()
	case errors.RecordNotFoundError:
		httpStatusCode = http.StatusNotFound
		statusCode = StatusCode1004
		message = err.Error()
		err = nil
	case errors.Unauthorized:
		httpStatusCode = http.StatusUnauthorized
		statusCode = StatusCode1005
		message = err.Error()
	case errors.Forbidden:
		httpStatusCode = http.StatusForbidden
		statusCode = StatusCode1006
		message = err.Error()
	case errors.HttpClientError:
		httpStatusCode = http.StatusInternalServerError
		statusCode = StatusCode1008
		message = err.Error()
	default:
		message = err.Error()
	}

	errorResponse := ErrorResponseData{
		StatusCode:    statusCode,
		StatusMessage: StatusMassageFail,
		Message:       message,
		Timestamp:     time.Now(),
		Data:          data,
	}

	c.AbortWithStatusJSON(httpStatusCode, errorResponse)
}
