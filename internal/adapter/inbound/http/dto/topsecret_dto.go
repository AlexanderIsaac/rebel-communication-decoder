package dto

type SatelliteSplit struct {
	Distance float64  `json:"distance" validate:"required,numeric"`
	Message  []string `json:"message" validate:"required,dive"`
}

type Satellite struct {
	SatelliteSplit
	Name string `json:"name" validate:"required"`
}

type TopSecretDTO struct {
	Satellites []Satellite `json:"satellites" validate:"required,dive"`
}
