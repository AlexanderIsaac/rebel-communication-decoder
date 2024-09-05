package repository

import (
	"app/internal/adapter/outbound/firestore"
	"app/internal/adapter/outbound/repository/model"
	"time"
)

type SatelliteRepository struct {
	satellites              []model.Satellite
	store                   firestore.Port
	satellitesCollection    string
	satellitesMsgCollection string
}

func NewSatelliteRepository(store firestore.Port) *SatelliteRepository {
	repo := &SatelliteRepository{}
	repo.store = store
	repo.satellitesCollection = "satellites"
	repo.satellitesMsgCollection = "satellites_messages"
	return repo
}

func (repo *SatelliteRepository) GetAllSatellites() ([]model.Satellite, error) {
	var satellites []model.Satellite
	docs, err := repo.store.GetAll(repo.satellitesCollection)
	if err != nil {
		return satellites, err
	}

	// Iterate through documents and decode into Satellite model
	for _, doc := range docs {
		var satellite model.Satellite
		if err := doc.DataTo(&satellite); err != nil {
			continue
		}
		satellites = append(satellites, satellite)
	}
	return satellites, nil
}

func (repo *SatelliteRepository) SaveSatellite(name string, position model.Position) []model.Satellite {
	repo.satellites = append(repo.satellites, model.Satellite{Name: name, Position: position})
	return repo.satellites
}

func (repo *SatelliteRepository) SaveReceivedMessage(name string, distance float64, message []string) (bool, error) {

	var structData = make(map[string]interface{})
	structData["name"] = name
	structData["distance"] = distance
	structData["message"] = message
	structData["timestamp"] = time.Now()

	return repo.store.Save(repo.satellitesMsgCollection, name, structData)
}

func (repo *SatelliteRepository) GetLastMessagesReceived() ([]model.LastMessageReceived, error) {
	var lastMessagesReceived []model.LastMessageReceived
	docs, err := repo.store.GetAllWithTime(repo.satellitesMsgCollection, 5)

	if err != nil {
		return lastMessagesReceived, err
	}

	// Iterate through documents and decode into Satellite model
	for _, doc := range docs {
		var message model.LastMessageReceived
		if err := doc.DataTo(&message); err != nil {
			continue
		}
		lastMessagesReceived = append(lastMessagesReceived, message)
	}

	return lastMessagesReceived, nil
}
