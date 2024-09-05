package firestore

import (
	errorMessage "app/internal/error"
	"context"
	"encoding/base64"
	"errors"
	"log/slog"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreClient interface {
	Collection(collection string) *firestore.CollectionRef
}

type Client struct {
	client *firestore.Client
	ctx    context.Context
}

type Port interface {
	GetAll(collection string) ([]*firestore.DocumentSnapshot, error)
	Save(collection string, doc string, data map[string]interface{}) (bool, error)
	GetAllWithTime(collection string, minutesAgo int) ([]*firestore.DocumentSnapshot, error)
}

func NewClient() (*Client, error) {
	ctx := context.Background()
	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
	slog.Info(project)
	firestoreDB := os.Getenv("GOOGLE_CLOUD_FIRESTORE_DB")
	slog.Info(firestoreDB)
	json, err := base64.StdEncoding.DecodeString(os.Getenv("GOOGLE_CLOUD_CREDENTIALS"))
	if err != nil {
		return nil, err
	}
	// Initialize Firestore client with credentials
	client, err := firestore.NewClientWithDatabase(ctx, project, firestoreDB, option.WithCredentialsJSON(json))
	if err != nil {
		return nil, err
	}

	store := &Client{}

	store.client = client
	store.ctx = ctx

	return store, nil
}

func (f *Client) GetAll(collection string) ([]*firestore.DocumentSnapshot, error) {

	documents, err := f.client.Collection(collection).Documents(f.ctx).GetAll()

	if err != nil {
		return nil, errors.New(errorMessage.FirestoreRetrieving)
	}

	return documents, nil
}

func (f *Client) Save(collection string, doc string, data map[string]interface{}) (bool, error) {

	_, err := f.client.Collection(collection).Doc(doc).Set(f.ctx, data)

	if err != nil {
		return false, errors.New(errorMessage.FirestoreSaving)
	}

	return true, nil
}

func (f *Client) GetAllWithTime(collection string, minutesAgo int) ([]*firestore.DocumentSnapshot, error) {

	calculateMinutesAgo := time.Now().Add(-time.Duration(minutesAgo) * time.Minute)

	coll := f.client.Collection(collection)
	query := coll.Where("timestamp", ">=", calculateMinutesAgo)

	documents, err := query.Documents(f.ctx).GetAll()

	if err != nil {
		return nil, errors.New(errorMessage.FirestoreRetrieving)
	}

	return documents, nil
}
