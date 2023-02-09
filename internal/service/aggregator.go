package service

import (
	"io"
	"log-aggregator/internal/mongodb"
	"log-aggregator/internal/rest"
	"log-aggregator/internal/socket"
	"strings"
	"time"
)

const fiveSeconds = time.Second * 5

// TODO: get logs since some date
// TODO: which container to get
// TODO: error handling

func Run() {
	for true {
		logs := collectLogs()
		dump(logs)
		time.Sleep(fiveSeconds)
	}
}

func dump(logMsgs []string) {
	var conn *mongodb.Connection = mongodb.Connect()

	var bulkDateTime time.Time = time.Now()

	var logDocuments []interface{}
	for i := 0; i < len(logMsgs); i++ {
		logDocuments = append(logDocuments, mongodb.Log{ContainerId: "f709242fb73e91b4988d3faf8487ed1e0cfef7c4bdb456f91d86dc733158de3", Log: logMsgs[i], DateTime: bulkDateTime})
	}

	_, err := conn.InsertMany(logDocuments)

	if err != nil {
		panic(err)
	}
}

func collectLogs() []string {
	httpc := socket.HttpUnixSocket()

	response, err := rest.Get(httpc)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBodyAsBytes, err := io.ReadAll(response.Body)
	responseBodyAsString := string(responseBodyAsBytes)

	return strings.Split(responseBodyAsString, "\n")
}
