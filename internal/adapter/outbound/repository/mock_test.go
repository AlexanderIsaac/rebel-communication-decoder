package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/mock"
)

// Mock Firestore Port
type MockFirestorePort struct {
	mock.Mock
}

func (m *MockFirestorePort) GetAll(collection string) ([]*firestore.DocumentSnapshot, error) {
	args := m.Called(collection)
	return args.Get(0).([]*firestore.DocumentSnapshot), args.Error(1)
}

func (m *MockFirestorePort) Save(collection string, doc string, data map[string]interface{}) (bool, error) {
	args := m.Called(collection, doc, data)
	return args.Bool(0), args.Error(1)
}

func (m *MockFirestorePort) GetAllWithTime(collection string, minutesAgo int) ([]*firestore.DocumentSnapshot, error) {
	args := m.Called(collection, minutesAgo)
	return args.Get(0).([]*firestore.DocumentSnapshot), args.Error(1)
}
