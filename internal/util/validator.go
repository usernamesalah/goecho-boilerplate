package util

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate

	// map for custom error message
	customMsgForTag = map[string]string{}
)

// GetValidator Initiatilize validator in singleton way
func GetValidator() *validator.Validate {

	if validate == nil {
		validate = validator.New()
		validate.RegisterValidation("alphanumunder", ValidateAlphaNumericAndUnderscore)
		validate.RegisterValidation("slug", ValidateSlug)
	}

	return validate
}

func getErrMessage(verr validator.FieldError) (msg string) {
	customMsgFormat := customMsgForTag[verr.Tag()]
	if customMsgFormat == "" {
		msg = verr.Field() + " is " + verr.Tag()
	} else {
		msg = fmt.Sprintf(customMsgFormat, verr.Field())
	}
	return msg
}

func Validation(val interface{}) error {
	validate := GetValidator()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	if err := validate.Struct(val); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, verr := range err.(validator.ValidationErrors) {
			err = errors.New(getErrMessage(verr))
			break
		}

		if err != nil {
			return ErrBadRequest(err, err.Error())
		}
	}
	return nil
}

const (
	alphaNumericWithUnderscoreRegexString = `^[a-zA-Z0-9]+(_[a-zA-Z0-9]+)*[a-zA-Z0-9]+$`
	slugRegexString                       = `^[a-zA-Z]+([_a-zA-Z0-9.]+)*[a-zA-Z0-9]+$`
)

// ValidateAlphaNumericAndUnderscore implements validator.Func
func ValidateAlphaNumericAndUnderscore(fl validator.FieldLevel) bool {
	return regexp.MustCompile(alphaNumericWithUnderscoreRegexString).MatchString(fl.Field().String())
}

// ValidateAlphaNumericAndUnderscore implements validator.Func
func ValidateSlug(fl validator.FieldLevel) bool {
	return regexp.MustCompile(slugRegexString).MatchString(fl.Field().String())
}
