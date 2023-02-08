package main

import (
	"context"
	"fmt"
	"io"
	"log-aggregator/internal/mongodb"
	"net"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type log struct {
	ContainerId string
	Log         string
	DateTime    time.Time
}

func main() {
	var conn *mongodb.Connection = mongodb.Connect()

	fmt.Println("Init Unix HTTP client")

	httpc := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}

	var response *http.Response
	response, err := httpc.Get("http://localhost/v1.41/containers/f709242fb73e91b4988d3faf8487eda1e0cfef7c4bdb456f91d86dc733158de3/logs?stdout=true")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var coll *mongo.Collection = conn.Client.Database("log_aggregator_db").Collection("log")

	responseBodyAsBytes, err := io.ReadAll(response.Body)
	responseBodyAsString := string(responseBodyAsBytes)

	result, err := coll.InsertOne(context.TODO(), log{ContainerId: "f709242fb73e91b4988d3faf8487ed1e0cfef7c4bdb456f91d86dc733158de3", Log: responseBodyAsString, DateTime: time.Now()})

	if err != nil {
		panic(err)
	}

	fmt.Printf("document inserted successfully. {_id:%s}\n", result.InsertedID)
}
