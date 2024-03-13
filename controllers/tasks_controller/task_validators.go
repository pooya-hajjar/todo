package tasksController

import "github.com/go-playground/validator/v10"

var StatusValidator validator.Func = func(fl validator.FieldLevel) bool {
	status := fl.Field().Interface().(int)
	return status == -2 || status == -1 || status == 0 || status == 1
}
