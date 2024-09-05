package service

import (
	"app/internal/adapter/outbound/repository/model"
	errorMessage "app/internal/error"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveReceivedMessageSuccess(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	service := NewSatelliteService(mockRepo)

	// Define test data
	satellites := []model.Satellite{
		{Name: "Kenobi", Position: model.Position{X: -500, Y: -200}},
		{Name: "Skywalker", Position: model.Position{X: 100, Y: -100}},
		{Name: "Sato", Position: model.Position{X: 500, Y: 100}},
	}
	message := []string{"", "este", "es", "un", "mensaje"}

	name := "Kenobi"
	distance := 100.0

	newMessages := []model.LastMessageReceived{{Name: name, Distance: distance, Message: message}}

	// Set up mock expectations
	mockRepo.On("GetAllSatellites").Return(satellites).Once()
	mockRepo.On("SaveReceivedMessage", name, distance, message).Return(newMessages).Once()

	// Test the method
	success, err := service.SaveReceivedMessage(name, distance, message)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, success)
	mockRepo.AssertExpectations(t)
}

func TestSaveReceivedMessage_SatelliteNotFound(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	service := NewSatelliteService(mockRepo)

	// Define test data
	satellites := []model.Satellite{}
	name := "Kenobi"
	distance := 100.0
	message := []string{""}

	// Set up mock expectations
	mockRepo.On("GetAllSatellites").Return(satellites).Once()

	// Test the method
	success, err := service.SaveReceivedMessage(name, distance, message)

	// Assertions
	assert.Error(t, err)
	assert.False(t, success)
	assert.Equal(t, errorMessage.SatelliteNotFoundMessage, err.Error())
	mockRepo.AssertExpectations(t)
}
