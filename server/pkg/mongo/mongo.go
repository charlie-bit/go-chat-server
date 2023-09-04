package mongo

import (
	"chat_socket/server/config"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	db *mongo.Client
}

// NewMongo Initialize MongoDB connection.
func NewMongo() (*Mongo, error) {
	var (
		uri          string
		mongodbHosts string
	)
	for i, v := range config.Cfg.Mongo.Address {
		if i == len(config.Cfg.Mongo.Address)-1 {
			mongodbHosts += v
		} else {
			mongodbHosts += v + ","
		}
	}
	if config.Cfg.Mongo.Password != "" && config.Cfg.Mongo.Username != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s/%s?maxPoolSize=%d&authSource=admin",
			config.Cfg.Mongo.Username, config.Cfg.Mongo.Password, mongodbHosts,
			config.Cfg.Mongo.Database, config.Cfg.Mongo.MaxPoolSize)
	} else {
		uri = fmt.Sprintf("mongodb://%s/%s/?maxPoolSize=%d&authSource=admin",
			mongodbHosts, config.Cfg.Mongo.Database,
			config.Cfg.Mongo.MaxPoolSize)
	}
	var mongoClient *mongo.Client
	var err error = nil
	for i := 0; i <= 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err == nil {
			return &Mongo{db: mongoClient}, nil
		}
		if cmdErr, ok := err.(mongo.CommandError); ok {
			if cmdErr.Code == 13 || cmdErr.Code == 18 {
				return nil, err
			} else {
				fmt.Printf("Failed to connect to MongoDB: %s\n", err)
			}
		}
	}
	return nil, err
}

func (m *Mongo) GetDatabase() *mongo.Database {
	return m.db.Database(config.Cfg.Mongo.Database)
}
