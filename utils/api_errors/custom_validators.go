package apiErrors

import (
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Tag     string
	Handler validator.Func
}

func RegisterCustomValidator(customValidators ...CustomValidator) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for _, cv := range customValidators {
			err := v.RegisterValidation(cv.Tag, cv.Handler)
			if err != nil {
				log.Fatal("error on registering validators")
			}
		}
		return
	}
	log.Fatal("error on registering validators")
}
