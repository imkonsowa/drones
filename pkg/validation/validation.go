package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"strings"
)

// Validation is a wrapper for validator.Validate to provide the app dependencies
type Validation struct {
	Validator *validator.Validate
}

var v *Validation

func GetValidator(db *gorm.DB) *Validation {
	if v == nil {
		v = getNewConfiguredValidator(db)
	}

	return v
}

func getNewConfiguredValidator(db *gorm.DB) *Validation {
	validate := validator.New()

	err := validate.RegisterValidation("uniqueDB", func(fl validator.FieldLevel) bool {
		params := strings.Split(fl.Param(), ".")
		if len(params) != 2 {
			panic("invalid role [uniqueDB] received params")
		}

		var count int64
		q := fmt.Sprintf("select count(*) from %s where %s = ?", params[0], params[1])
		err := db.Raw(q, fl.Field().String()).Count(&count).Error

		if err != nil {
			log.Printf("error while processing validation role [uniqueDB]; err: %v\n", err)
			return false
		}

		if count > 0 {
			return false
		}

		return true
	})

	if err != nil {
		// TODO: log register role reg error
	}

	return &Validation{
		Validator: validate,
	}
}
