package repository

import "app/internal/adapter/outbound/repository/model"

type SatelliteRepository struct {
	satellites           []model.Satellite
	lastMessagesReceived []model.LastMessageReceived
}

func NewSatelliteRepository() *SatelliteRepository {
	repo := &SatelliteRepository{}
	repo.SaveSatellite("Kenobi", model.Position{X: -500, Y: -200})
	repo.SaveSatellite("Skywalker", model.Position{X: 100, Y: -100})
	repo.SaveSatellite("Sato", model.Position{X: 500, Y: 100})
	return repo
}

func (repo *SatelliteRepository) GetAllSatellites() []model.Satellite {
	return repo.satellites
}

func (repo *SatelliteRepository) SaveSatellite(name string, position model.Position) []model.Satellite {
	repo.satellites = append(repo.satellites, model.Satellite{Name: name, Position: position})
	return repo.satellites
}

func (repo *SatelliteRepository) SaveReceivedMessage(name string, distance float64, message []string) []model.LastMessageReceived {
	// Filter out any previous messages with the same name
	var filteredMessages []model.LastMessageReceived
	for _, msg := range repo.lastMessagesReceived {
		if msg.Name != name {
			filteredMessages = append(filteredMessages, msg)
		}
	}

	filteredMessages = append(filteredMessages, model.LastMessageReceived{Name: name, Distance: distance, Message: message})
	repo.lastMessagesReceived = filteredMessages
	return repo.lastMessagesReceived
}

func (repo *SatelliteRepository) GetLastMessagesReceived() []model.LastMessageReceived {
	return repo.lastMessagesReceived
}
