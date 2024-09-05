package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"app/internal/adapter/outbound/repository/model"
)

func TestSaveSatellite(t *testing.T) {
	mockStore := new(MockFirestorePort)
	repo := NewSatelliteRepository(mockStore)

	// Call method
	position := model.Position{X: 1, Y: 2}
	satellites := repo.SaveSatellite("Sat1", position)

	// Assertions
	assert.Len(t, satellites, 1)
	assert.Equal(t, "Sat1", satellites[0].Name)
	assert.Equal(t, position, satellites[0].Position)
}

func TestSaveReceivedMessage(t *testing.T) {
	mockStore := new(MockFirestorePort)
	repo := NewSatelliteRepository(mockStore)

	// Prepare mock data
	messageData := map[string]interface{}{
		"name":      "Sat1",
		"distance":  100.0,
		"message":   []string{"Hello", "World"},
		"timestamp": time.Now(),
	}
	mockStore.On("Save", "satellites_messages", "Sat1", messageData).Return(true, nil)

	// Call method
	_, err := repo.SaveReceivedMessage("Sat1", 100.0, []string{"Hello", "World"})

	// Assertions
	// assert.True(t, success)
	assert.NoError(t, err)
	mockStore.AssertExpectations(t)
}
