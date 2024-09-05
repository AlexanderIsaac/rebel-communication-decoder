package repository

import (
	"testing"

	"app/internal/adapter/outbound/repository/model"

	"github.com/stretchr/testify/assert"
)

func TestNewSatelliteRepository(t *testing.T) {
	repo := NewSatelliteRepository()

	satellites := repo.GetAllSatellites()
	assert.Len(t, satellites, 3)
	assert.Equal(t, "Kenobi", satellites[0].Name)
	assert.Equal(t, model.Position{X: -500, Y: -200}, satellites[0].Position)
	assert.Equal(t, "Skywalker", satellites[1].Name)
	assert.Equal(t, model.Position{X: 100, Y: -100}, satellites[1].Position)
	assert.Equal(t, "Sato", satellites[2].Name)
	assert.Equal(t, model.Position{X: 500, Y: 100}, satellites[2].Position)
}

func TestSaveSatellite(t *testing.T) {
	repo := NewSatelliteRepository()

	repo.SaveSatellite("Tatooine", model.Position{X: 200, Y: 300})
	satellites := repo.GetAllSatellites()

	assert.Len(t, satellites, 4)
	assert.Equal(t, "Tatooine", satellites[3].Name)
	assert.Equal(t, model.Position{X: 200, Y: 300}, satellites[3].Position)
}

func TestSaveReceivedMessage(t *testing.T) {
	repo := NewSatelliteRepository()

	repo.SaveReceivedMessage("Kenobi", 100.0, []string{"", "este", "es", "un", "mensaje"})
	repo.SaveReceivedMessage("Skywalker", 200.0, []string{"", "este", "es", "", "mensaje"})

	messages := repo.GetLastMessagesReceived()

	assert.Len(t, messages, 2)
	assert.Equal(t, "Kenobi", messages[0].Name)
	assert.Equal(t, 100.0, messages[0].Distance)
	assert.Equal(t, []string{"", "este", "es", "un", "mensaje"}, messages[0].Message)
	assert.Equal(t, "Skywalker", messages[1].Name)
	assert.Equal(t, 200.0, messages[1].Distance)
	assert.Equal(t, []string{"", "este", "es", "", "mensaje"}, messages[1].Message)
}

func TestOverwriteReceivedMessage(t *testing.T) {
	repo := NewSatelliteRepository()

	repo.SaveReceivedMessage("Kenobi", 100.0, []string{"", "", "", "", "mensaje"})
	repo.SaveReceivedMessage("Kenobi", 150.0, []string{"", "este", "es", "un", "mensaje"})

	messages := repo.GetLastMessagesReceived()

	assert.Len(t, messages, 1)
	assert.Equal(t, "Kenobi", messages[0].Name)
	assert.Equal(t, 150.0, messages[0].Distance)
	assert.Equal(t, []string{"", "este", "es", "un", "mensaje"}, messages[0].Message)
}
