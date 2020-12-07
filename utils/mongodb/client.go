package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	client *mongo.Client
}

//NewClient creates a new mongo client
func NewClient(opts ...*options.ClientOptions) (ClientHelper, error) {
	client, err := mongo.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	return &mongoClient{client: client}, nil
}

func (mc *mongoClient) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mc.client.Connect(ctx)
}

func (mc *mongoClient) Database(dbName string) DatabaseHelper {
	db := mc.client.Database(dbName)
	return &mongoDatabase{db: db}
}
