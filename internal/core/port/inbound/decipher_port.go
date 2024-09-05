package inbound

import (
	"app/internal/core/model"
)

type DecipherPort interface {
	GetLocation(distances []model.Distance) model.Position
	GetMessage(messages [][]string) string
	GetSplitLocation() model.Position
	GetSplitMessage() string
}
