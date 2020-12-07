package mongostore

import (
	"log"

	"github.com/andygodish/go-mongo-starter/utils/mongodb"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

// Store represents your mongo store object, contains your database
type Store struct {
	db mongodb.DatabaseHelper
}

// Constructor constructs a new mongo store by first 1 - creating a new client
// 2 - connect to the client
func Constructor(connURI string) (*Store, error) {
	mongoConnString, err := connstring.Parse(connURI)
	if err != nil {
		log.Printf("Error parsing mongoDB url: %s", err)
	}

	client, err := mongodb.NewClient(options.Client().ApplyURI(connURI))
	if err != nil {
		log.Printf("Error creating mongoDB client: %s", err)
	}

	err = client.Connect()
	if err != nil {
		log.Printf("Unable to connect to mongoDB: %s", err)
		return nil, err
	}
	log.Printf("Connected to mongoDB: %s", mongoConnString.AppName)

	db := client.Database(mongoConnString.Database)

	return &Store{db: db}, nil
}

// NewStore is used to generate a db for testing purposes
func NewStore(db mongodb.DatabaseHelper) *Store {
	return &Store{db: db}
}

// RawDB returns the raw database
func (s Store) RawDB() mongodb.DatabaseHelper {
	return s.db
}
