package inbound

type SatellitePort interface {
	SaveReceivedMessage(name string, distance float64, message []string) bool
}
