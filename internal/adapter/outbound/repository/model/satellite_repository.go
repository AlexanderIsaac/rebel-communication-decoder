package model

import (
	"time"
)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Satellite struct {
	Name     string   `json:"name"`
	Position Position `json:"position"`
}

type LastMessageReceived struct {
	Name      string    `json:"name"`
	Distance  float64   `json:"distance"`
	Message   []string  `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
