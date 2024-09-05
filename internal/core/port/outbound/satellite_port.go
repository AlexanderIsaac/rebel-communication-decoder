package outbound

import (
	"app/internal/adapter/outbound/repository/model"
)

type SatellitePort interface {
	GetAllSatellites() []model.Satellite
	SaveReceivedMessage(name string, distance float64, message []string) []model.LastMessageReceived
	GetLastMessagesReceived() []model.LastMessageReceived
}
