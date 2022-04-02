package router

import (
	"drones/app/http/requests"
	"drones/app/http/responses"
	"drones/pkg/validation"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

func (r *Router) inject(constructor requests.RequestConstructor) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := constructor()
		jsonTags := map[string]string{}
		val := reflect.Indirect(reflect.ValueOf(request))
		for i := 0; i < val.Type().NumField(); i++ {
			jsonTags[val.Type().Field(i).Name] = val.Type().Field(i).Tag.Get("json")
		}

		if err := c.BindJSON(request); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorBasicResponse(http.StatusBadRequest, nil))
			c.Abort()
			return
		}

		validate := validation.GetValidator()
		err := validate.Validator.Struct(request)
		if err != nil {
			errors := map[string]string{}

			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())

				errors[jsonTags[err.Field()]] = validation.GetRuleMessage(err.Tag(), map[string]string{err.ActualTag(): err.Param()})
			}

			c.JSON(422, gin.H{
				"success": false,
				"code":    422,
				"errors":  errors,
			})

			c.Abort()
			return
		}

		c.Set("request", request)
		c.Next()
	}
}
