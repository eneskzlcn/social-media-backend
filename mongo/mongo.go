package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/eneskzlcn/hackernews/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client *mongo.Client
	config config.DB
}

func NewDB(config config.DB) *DB{
	return &DB{config: config}
}
func(db * DB) Connect(config config.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(createDSN(config)))
	if err != nil {
		return errors.New("cannot connected to the database")
	}
	db.client = client
	return nil
}

// createDSN function creates required data source name(connection string) for mongo client.
func createDSN(db config.DB) string {
	return fmt.Sprintf("mongodb://%s:%d",db.Host, db.Port)
}
