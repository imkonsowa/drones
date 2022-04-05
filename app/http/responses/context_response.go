package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type contextResponse struct {
	context *gin.Context

	success bool
	code    int
	message string
	data    map[string]interface{}
}

func NewContextResponse(context *gin.Context) *contextResponse {
	return &contextResponse{context: context}
}

func (r *contextResponse) Code(code int) *contextResponse {
	r.code = code
	return r
}

func (r *contextResponse) Message(message string) *contextResponse {
	r.message = message
	return r
}

func (r *contextResponse) Error() *contextResponse {
	r.success = false
	return r
}

func (r *contextResponse) Success() *contextResponse {
	r.success = false
	if r.code == 0 {
		r.code = http.StatusOK
	}
	return r
}

func (r *contextResponse) Data(data map[string]interface{}) *contextResponse {
	r.data = data
	return r
}

func (r *contextResponse) Send() {
	if r.code == 0 {
		panic("invalid response code")
	}

	if len(r.message) == 0 {
		if r.code >= 200 && r.code < 300 {
			r.message = "operation done successfully"
		} else {
			r.message = "failed to process you request"
		}
	}

	payload := gin.H{
		"success": r.success,
		"code":    r.code,
		"message": r.message,
	}

	if r.data != nil {
		for key, val := range r.data {
			payload[key] = val
		}
	}

	r.context.JSON(r.code, payload)

	return
}
