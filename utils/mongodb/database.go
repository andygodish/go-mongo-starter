package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	db *mongo.Database
}

func (md *mongoDatabase) Client() ClientHelper {
	client := md.db.Client()
	return &mongoClient{client: client}
}

func (md *mongoDatabase) CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error {
	return md.db.CreateCollection(ctx, name, opts...)
}
