package validators

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type JSONValidator struct {
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func (JSONValidator) descriptive(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}

func (j JSONValidator) Validate(c *gin.Context, s interface{}) []ValidationError {

	var verr validator.ValidationErrors
	var validateErrors []ValidationError

	c.ShouldBindJSON(&s)

	validate := validator.New()
	err := validate.Struct(s)

	if errors.As(err, &verr) {
		validateErrors = j.descriptive(verr)
		c.JSON(http.StatusBadRequest, gin.H{"error": validateErrors})
	}
	return validateErrors
}
