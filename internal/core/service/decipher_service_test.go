package service

import (
	"app/internal/adapter/outbound/repository/model"
	coremodel "app/internal/core/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocationSuccess(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	ds := NewDecipherService(mockRepo)

	// Define test data
	satellites := []model.Satellite{
		{Name: "Kenobi", Position: model.Position{X: -500, Y: -200}},
		{Name: "Skywalker", Position: model.Position{X: 100, Y: -100}},
		{Name: "Sato", Position: model.Position{X: 500, Y: 100}},
	}
	distances := []coremodel.Distance{
		{Name: "Kenobi", Distance: 100.0},
		{Name: "Skywalker", Distance: 115.5},
		{Name: "Sato", Distance: 142.7},
	}

	// // Mock expectations
	mockRepo.On("GetAllSatellites").Return(satellites, nil).Once()

	// Test successful location calculation
	expectedPosition := coremodel.Position{X: -487.29, Y: 1557.01}
	position, err := ds.GetLocation(distances)

	assert.NoError(t, err)
	assert.Equal(t, expectedPosition, position)

}

func TestGetLocationError(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	ds := NewDecipherService(mockRepo)

	distances := []coremodel.Distance{}

	// // Test location calculation failure due to missing satellites
	mockRepo.On("GetAllSatellites").Return([]model.Satellite{}, nil).Once()
	position, err := ds.GetLocation(distances)

	assert.Error(t, err)
	assert.Equal(t, coremodel.Position{}, position)
}

func TestGetMessageSuccess(t *testing.T) {
	ds := NewDecipherService(nil) // No dependency for this method

	// Define test data
	messages := [][]string{
		{"", "este", "es", "un", "mensaje"},
		{"este", "", "un", "mensaje"},
		{"", "", "es", "mensaje"},
	}

	// Test successful message decoding
	expectedMessage := "este es un mensaje"
	message, err := ds.GetMessage(messages)
	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, message)
}

func TestGetMessageError(t *testing.T) {
	ds := NewDecipherService(nil) // No dependency for this method

	// Test message decoding failure (no valid messages)
	messages := [][]string{}
	message, err := ds.GetMessage(messages)
	assert.Error(t, err)
	assert.Equal(t, "", message)
}

func TestGetSplitLocationSuccess(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	ds := NewDecipherService(mockRepo)

	// Define test data
	messages := []model.LastMessageReceived{
		{Name: "Kenobi", Distance: 100.0},
		{Name: "Skywalker", Distance: 115.5},
		{Name: "Sato", Distance: 142.7},
	}
	satellites := []model.Satellite{
		{Name: "Kenobi", Position: model.Position{X: -500, Y: -200}},
		{Name: "Skywalker", Position: model.Position{X: 100, Y: -100}},
		{Name: "Sato", Position: model.Position{X: 500, Y: 100}},
	}

	// Mock expectations
	mockRepo.On("GetLastMessagesReceived").Return(messages).Once()
	mockRepo.On("GetAllSatellites").Return(satellites).Once()

	// Test successful split location calculation
	expectedPosition := coremodel.Position{X: -487.29, Y: 1557.01}
	position, err := ds.GetSplitLocation()
	assert.NoError(t, err)
	assert.Equal(t, expectedPosition, position)
}

func TestGetSplitLocationError(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	ds := NewDecipherService(mockRepo)

	// Test split location calculation failure (not enough messages)
	mockRepo.On("GetLastMessagesReceived").Return([]model.LastMessageReceived{}).Once()
	position, err := ds.GetSplitLocation()
	assert.Error(t, err)
	assert.Equal(t, coremodel.Position{}, position)
}

func TestGetSplitMessageSuccess(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	ds := NewDecipherService(mockRepo)

	// Define test data
	messages := []model.LastMessageReceived{
		{Name: "Kenobi", Message: []string{"", "este", "es", "un", "mensaje"}},
		{Name: "Skywalker", Message: []string{"este", "", "un", "mensaje"}},
		{Name: "Sato", Message: []string{"", "", "es", "mensaje"}},
	}
	// Mock expectations
	mockRepo.On("GetLastMessagesReceived").Return(messages).Once()

	// Test successful split message decoding
	expectedMessage := "este es un mensaje"
	message, err := ds.GetSplitMessage()
	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, message)

}

func TestGetSplitMessageError(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	ds := NewDecipherService(mockRepo)

	// Test split message decoding failure (no valid messages)
	mockRepo.On("GetLastMessagesReceived").Return([]model.LastMessageReceived{}).Once()
	message, err := ds.GetSplitMessage()
	assert.Error(t, err)
	assert.Equal(t, "", message)
}
