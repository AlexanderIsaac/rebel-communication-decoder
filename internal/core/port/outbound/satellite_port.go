package outbound

import (
	"app/internal/adapter/outbound/repository/model"
)

type SatellitePort interface {
	GetAllSatellites() ([]model.Satellite, error)
	SaveReceivedMessage(name string, distance float64, message []string) (bool, error)
	GetLastMessagesReceived() ([]model.LastMessageReceived, error)
}
