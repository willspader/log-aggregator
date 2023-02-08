package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const database = "log_aggregator_db"

type Connection struct {
	Client *mongo.Client
}

type Log struct {
	ContainerId string
	Log         string
	DateTime    time.Time
}

func Connect() *Connection {
	// Create a new client and connect to the server
	var uri string = "mongodb://mongoadmin:secret@localhost:27017/?maxPoolSize=20&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	return &Connection{Client: client}
}

func (conn Connection) InsertOne(collection string, log Log) (*mongo.InsertOneResult, error) {
	return conn.getCollection(collection).InsertOne(context.TODO(), log)
}

func (conn Connection) getCollection(name string) *mongo.Collection {
	return conn.Client.Database(database).Collection(name)
}
