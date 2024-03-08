package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

func ValidateRequest[T any](req T) error {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	if err := validate.Struct(req); err != nil {
		// Traducir el error
		errors := err.(validator.ValidationErrors)
		var errorMsgs []string

		for _, e := range errors {
			errorMsgs = append(errorMsgs, fmt.Sprintf("%s: %s", e.Field(), e.Translate(trans)))
		}

		return fmt.Errorf(strings.Join(errorMsgs, ", "))
	}

	return nil
}
