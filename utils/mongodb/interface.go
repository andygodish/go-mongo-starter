package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

/* This package implement an interface wrapper around the official mongo-driver that allows
 * the driver to be mocked for testing.
 * https://godoc.org/go.mongodb.org/mongo-driver/mongo
 *
 * DO NOT ADD any functions here that are NOT part of the official mongo-driver
 * General purpose mongo utility functions should be defined elsewhere in this package
 */

// ClientHelper describes the mongo.Client type
type ClientHelper interface {
	Database(string) DatabaseHelper
	Connect() error
}

// DatabaseHelper describes the mongo.Databse type
type DatabaseHelper interface {
	Client() ClientHelper
	// Collection(name string) CollectionHelper
	CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error
}
