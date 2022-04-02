package router

import (
	"drones/app/http/responses"
	"drones/pkg/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (r *Router) inject(request interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.BindJSON(request); err != nil {
			responses.
				NewContextResponse(c).
				Error().
				Code(http.StatusBadRequest).
				Message("failed to process your request").
				Send()

			c.Abort()
			return
		}

		validate := validation.GetValidator(r.App.DB)
		err := validate.Validator.Struct(request)
		if err != nil {
			errors := map[string]string{}

			for _, err := range err.(validator.ValidationErrors) {
				// TODO: replace struct field with the received json tag name including nested structs.
				errors[err.StructField()] = validation.GetRuleMessage(err.Tag(), map[string]string{err.ActualTag(): err.Param()})
			}

			responses.
				NewContextResponse(c).
				Error().
				Code(http.StatusUnprocessableEntity).
				Message("invalid data").
				Data(map[string]interface{}{
					"errors": errors,
				}).
				Send()
			c.Abort()
			return
		}

		c.Set("request", request)
		c.Next()
	}
}
