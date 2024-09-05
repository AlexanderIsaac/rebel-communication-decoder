package main

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Value string `validate:"required"`
}

func TestValidate_ValidData(t *testing.T) {
	validate := validator.New()
	cv := &CustomValidator{validator: validate}

	data := TestStruct{
		Value: "value",
	}

	err := cv.Validate(data)
	assert.NoError(t, err)
}

func TestValidate_MissingValue(t *testing.T) {
	validate := validator.New()
	cv := &CustomValidator{validator: validate}

	data := TestStruct{
		Value: "",
	}

	err := cv.Validate(data)
	assert.Error(t, err)
}

func TestCreateLogger(t *testing.T) {
	logger := createLogger()
	assert.NotNil(t, logger)
}

func TestConfigureEcho(t *testing.T) {
	logger := createLogger()
	e := configureEcho(logger)
	assert.NotNil(t, e)
	assert.IsType(t, &echo.Echo{}, e)
}

func TestSetupServiceAndRoutes(t *testing.T) {
	e := echo.New()
	logger := createLogger()
	setupServicesAndRoutes(e, logger)
	assert.NotNil(t, e)
}
