package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"regexp"
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

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("error while processing validation role [uniqueDB]; err: %v\n", err)
			return false
		}

		if count > 0 {
			return false
		}

		return true
	})

	err = validate.RegisterValidation("existsDB", func(fl validator.FieldLevel) bool {
		params := strings.Split(fl.Param(), ".")
		if len(params) != 2 {
			panic("invalid role [existsDB] received params")
		}

		var count int64
		q := fmt.Sprintf("select count(*) from %s where %s = ?", params[0], params[1])
		err := db.Raw(q, fl.Field().String()).Count(&count).Error

		if err != nil {
			log.Printf("error while processing validation role [existsDB]; err: %v\n", err)
			return false
		}

		if count > 0 {
			return true
		}

		return false
	})

	err = validate.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		regex := fl.Param()
		val := fl.Field().String()

		reg := regexp.MustCompile(regex)
		return reg.MatchString(val)
	})

	if err != nil {
		// TODO: log register role reg error
	}

	return &Validation{
		Validator: validate,
	}
}
