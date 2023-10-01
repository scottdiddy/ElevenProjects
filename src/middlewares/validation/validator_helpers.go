package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationMiddleware struct {
	validator *validator.Validate
}

func NewValidatorMiddleware() *ValidationMiddleware {
	return &ValidationMiddleware{
		validator: validator.New(),
	}
}
func (mw *ValidationMiddleware) ValidateStruct(s any) (map[string]string, error, bool) {
	var validationErrorMap = make(map[string]string)
	if err := mw.validator.Struct(s); err != nil {
		if validation_Errors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validation_Errors {
				validationErrorMap[e.Field()] = e.Error()
			}
			return validationErrorMap, errors.New("invalid request Data"), true
		} else {
			return validationErrorMap, err, false
		}
	}
	return validationErrorMap, nil, false

}
