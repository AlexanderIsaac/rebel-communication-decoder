package main

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestStruct struct {
	Value string `validate:"required"`
}

func TestValidateValidData(t *testing.T) {
	validate := validator.New()
	cv := &CustomValidator{validator: validate}

	data := TestStruct{
		Value: "value",
	}

	err := cv.Validate(data)
	assert.NoError(t, err)
}

func TestValidateMissingValue(t *testing.T) {
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

type MockFirestoreClient struct {
	mock.Mock
}

func (m *MockFirestoreClient) NewClient() error {
	args := m.Called()
	return args.Error(0)
}
func TestSetupServiceAndRoutes(t *testing.T) {
	e := echo.New()
	logger := createLogger()
	mockStore := new(MockFirestoreClient)

	mockStore.On("NewClient").Return(mockStore, nil).Once()
	setupServicesAndRoutes(e, logger)
	assert.NotNil(t, e)
}
