package responses

import (
	"net/http"
)

type ErrorResponse map[string]interface{}

func newErrorResponse(statusCode int) ErrorResponse {
	response := make(ErrorResponse)
	response["code"] = statusCode
	response["success"] = false

	return response
}

func ErrorBasicResponse(statusCode int, message *string) ErrorResponse {
	response := newErrorResponse(statusCode)
	if message != nil {
		response["message"] = *message
	} else {
		response["message"] = http.StatusText(statusCode)
	}
	return response
}

/*func ErrorResponseWithData() *ErrorResponse {
}*/

/*func ErrorResponseWithErrors(statusCode int, errors []*validation.Error) ErrorResponse {
	errs := make(map[string]string)
	response := newErrorResponse(statusCode)
	for _, err := range errors {
		errs[err.Key] = strings.TrimSpace(err.Message)
	}
	response["errors"] = errs
	return response
}
*/
