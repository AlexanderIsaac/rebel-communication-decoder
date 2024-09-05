package service

import (
	"app/internal/adapter/outbound/repository/model"

	"github.com/stretchr/testify/mock"
)

// Mock implementation of the SatellitePort interface.
type MockSatellitePort struct {
	mock.Mock
}

// GetAllSatellites mocks the GetAllSatellites method
func (m *MockSatellitePort) GetAllSatellites() []model.Satellite {
	args := m.Called()
	return args.Get(0).([]model.Satellite)
}

// GetLastMessagesReceived mocks the GetLastMessagesReceived method
func (m *MockSatellitePort) GetLastMessagesReceived() []model.LastMessageReceived {
	args := m.Called()
	return args.Get(0).([]model.LastMessageReceived)
}

// SaveReceivedMessage mocks the SaveReceivedMessage method
func (m *MockSatellitePort) SaveReceivedMessage(name string, distance float64, message []string) []model.LastMessageReceived {
	args := m.Called(name, distance, message)
	return args.Get(0).([]model.LastMessageReceived)
}
