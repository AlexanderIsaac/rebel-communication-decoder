package service

import (
	"app/internal/core/port/outbound"
	errorMessage "app/internal/error"
	"errors"
	"strings"
)

/* Service for handling satellite information */
type SatelliteService struct {
	satelliteRepository outbound.SatellitePort
}

func NewSatelliteService(repo outbound.SatellitePort) *SatelliteService {
	return &SatelliteService{
		satelliteRepository: repo,
	}
}

/*
SaveLastMessage saves a new message for a specified satellite if it exists.
*/
func (s *SatelliteService) SaveReceivedMessage(name string, distance float64, message []string) (bool, error) {
	// Retrieve all satellites from the repository
	satellites, err := s.satelliteRepository.GetAllSatellites()

	if err != nil {
		return false, errors.New(errorMessage.SatelliteNotFoundMessage)
	}

	found := false
	// Check if the satellite with the given name exists.
	for _, satellite := range satellites {
		if strings.EqualFold(satellite.Name, name) {
			found = true
		}
	}

	// If the satellite is not found, return false and an error.
	if !found {
		return false, errors.New(errorMessage.SatelliteNotFoundMessage)
	}

	// Save the message for the found satellite.

	return s.satelliteRepository.SaveReceivedMessage(name, distance, message)
}
