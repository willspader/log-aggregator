package mongodb

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const database = "log_aggregator_db"
const collectionPrefix = "log-"

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

func (conn Connection) InsertMany(items []interface{}) (*mongo.InsertManyResult, error) {
	return conn.getCollection().InsertMany(context.TODO(), items)
}

func (conn Connection) getCollection() *mongo.Collection {
	return conn.Client.Database(database).Collection(collectionPrefix + getCollectionDate())
}

func getCollectionDate() string {
	year, month, day := time.Now().Date()

	return strconv.Itoa(year) + getMonthDayTwoDigits(strconv.Itoa(int(month))) + getMonthDayTwoDigits(strconv.Itoa(day))
}

func getMonthDayTwoDigits(s string) string {
	if len(s) == 2 {
		return s
	}
	return "0" + s
}
