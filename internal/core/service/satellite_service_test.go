package service

import (
	"app/internal/adapter/outbound/repository/model"
	errorMessage "app/internal/error"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveReceivedMessageSatelliteExists(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	service := NewSatelliteService(mockRepo)

	satellites := []model.Satellite{
		{Name: "Satellite1"},
	}

	mockRepo.On("GetAllSatellites").Return(satellites, nil)
	mockRepo.On("SaveReceivedMessage", "Satellite1", 100.0, []string{"Message"}).Return(true, nil)

	success, err := service.SaveReceivedMessage("Satellite1", 100.0, []string{"Message"})

	assert.True(t, success)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSaveReceivedMessageSatelliteNotFound(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	service := NewSatelliteService(mockRepo)

	satellites := []model.Satellite{
		{Name: "Satellite2"},
	}

	mockRepo.On("GetAllSatellites").Return(satellites, nil)

	success, err := service.SaveReceivedMessage("Satellite1", 100.0, []string{"Message"})

	assert.False(t, success)
	assert.EqualError(t, err, errorMessage.SatelliteNotFoundMessage)
	mockRepo.AssertExpectations(t)
}

func TestSaveReceivedMessageGetAllSatellitesError(t *testing.T) {
	mockRepo := new(MockSatellitePort)
	service := NewSatelliteService(mockRepo)

	mockRepo.On("GetAllSatellites").Return([]model.Satellite{}, errors.New("repository error"))

	success, err := service.SaveReceivedMessage("Satellite1", 100.0, []string{"Message"})

	assert.False(t, success)
	assert.EqualError(t, err, errorMessage.SatelliteNotFoundMessage)
	mockRepo.AssertExpectations(t)
}
