package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

type RequestValidationError struct {
	Message string `json:"message" default:"validation error"`
	Error 	string `json:"error"`
}

func newRequestValidationError(c echo.Context, code int, err error) error {
	rve := &RequestValidationError{Error: err.Error()}
	defaults.SetDefaults(rve)
	return c.JSON(code, rve)
}

// v is an argument that is a pointer to a value of the type that implements 
// the interface you want to validate with.
// Errors are handled in the routes.
func bindAndValidateRequestBody(c echo.Context, v interface{}) error {
	if err := c.Bind(v); err != nil {
		return err
	}
	if err := c.Validate(v); err != nil {
		return err
	}
	return nil
}