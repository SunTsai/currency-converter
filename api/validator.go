package api

import (
	"github.com/go-playground/validator/v10"

	"suntsai/currency-converter/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	currency, ok := fieldLevel.Field().Interface().(string)
	if !ok {
		return false
	}

	return util.IsSuportedCurrency(currency)
}
