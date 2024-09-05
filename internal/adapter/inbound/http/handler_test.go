package http

import (
	coremodel "app/internal/core/model"

	"github.com/stretchr/testify/mock"
)

// Mock DecipherService
type MockDecipherService struct {
	mock.Mock
}

func (m *MockDecipherService) GetLocation(distances []coremodel.Distance) (coremodel.Position, error) {
	args := m.Called(distances)
	return args.Get(0).(coremodel.Position), args.Error(1)
}

func (m *MockDecipherService) GetMessage(messages [][]string) (string, error) {
	args := m.Called(messages)
	return args.String(0), args.Error(1)
}

func (m *MockDecipherService) GetSplitLocation() (coremodel.Position, error) {
	args := m.Called()
	return args.Get(0).(coremodel.Position), args.Error(1)
}

func (m *MockDecipherService) GetSplitMessage() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// Mock SatelliteService
type MockSatelliteService struct {
	mock.Mock
}

func (m *MockSatelliteService) SaveReceivedMessage(name string, distance float64, message []string) (bool, error) {
	args := m.Called(name, distance, message)
	return args.Bool(0), args.Error(1)
}
