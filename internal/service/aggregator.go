package service

import (
	"fmt"
	"io"
	"log-aggregator/internal/mongodb"
	"log-aggregator/internal/rest"
	"log-aggregator/internal/socket"
	"time"
)

const fiveSeconds = time.Second * 5

func Run() {
	for true {
		dump()
		time.Sleep(fiveSeconds)
	}
}

func dump() {
	var conn *mongodb.Connection = mongodb.Connect()

	var logMsg string = collectLogs()

	result, err := conn.InsertOne("log", mongodb.Log{ContainerId: "f709242fb73e91b4988d3faf8487ed1e0cfef7c4bdb456f91d86dc733158de3", Log: logMsg, DateTime: time.Now()})

	if err != nil {
		panic(err)
	}

	fmt.Printf("document inserted successfully. {_id:%s}\n", result.InsertedID)
}

func collectLogs() string {
	httpc := socket.HttpUnixSocket()

	response, err := rest.Get(httpc)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBodyAsBytes, err := io.ReadAll(response.Body)
	responseBodyAsString := string(responseBodyAsBytes)

	return responseBodyAsString
}
