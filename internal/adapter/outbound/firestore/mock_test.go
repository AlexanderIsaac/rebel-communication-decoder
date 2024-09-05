package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/mock"
)

// Mock for Firestore Client
type MockFirestoreClient struct {
	mock.Mock
}

func (m *MockFirestoreClient) Collection(collection string) *MockFirestoreCollectionRef {
	args := m.Called(collection)
	return args.Get(0).(*MockFirestoreCollectionRef)
}

// Mock for Firestore CollectionRef
type MockFirestoreCollectionRef struct {
	mock.Mock
}

func (m *MockFirestoreCollectionRef) Documents(ctx context.Context) *MockFirestoreDocumentIterator {
	args := m.Called(ctx)
	return args.Get(0).(*MockFirestoreDocumentIterator)
}

func (m *MockFirestoreCollectionRef) Doc(doc string) *MockFirestoreDocumentRef {
	args := m.Called(doc)
	return args.Get(0).(*MockFirestoreDocumentRef)
}

// Mock for Firestore DocumentRef
type MockFirestoreDocumentRef struct {
	mock.Mock
}

func (m *MockFirestoreDocumentRef) Set(ctx context.Context, data interface{}) (*firestore.WriteResult, error) {
	args := m.Called(ctx, data)
	return args.Get(0).(*firestore.WriteResult), args.Error(1)
}

// Mock for Firestore DocumentIterator
type MockFirestoreDocumentIterator struct {
	mock.Mock
}

func (m *MockFirestoreDocumentIterator) GetAll() ([]*firestore.DocumentSnapshot, error) {
	args := m.Called()
	return args.Get(0).([]*firestore.DocumentSnapshot), args.Error(1)
}
