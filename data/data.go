package data

import (
	"encoding/json"
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

const (
	cmd  = "select * from channel"
	db   = "amg"
	addr = "http://host.docker.internal:8086"
)

const (
	errResultsLength       = "unsupported length of results: %v"
	errRowsLength          = "unsupported length of rows: %v"
	errCouldNotCloseClient = "could not close HTTP client"
	errJsonNumber          = "JSON as number failed, original error: {%v}"
	errQueryFailed         = "db query failed, original error: {%v}"
	errPassing             = "{%v} failed, original error:\n{%v}"
)

func createClient() (client.Client, error) {
	return client.NewHTTPClient(client.HTTPConfig{Addr: addr})
}

func RequestChannel() ([]float64, error) { return requestChannel(createClient) }

func PopulateChannel(numbers []float64) error { return populateChannel(numbers, createClient) }

func requestChannel(createDbClient func() (client.Client, error)) ([]float64, error) {
	dbClient, err := createDbClient()
	if err != nil {
		return nil, err
	}
	defer closeClient(dbClient)

	results, err := query(dbClient, cmd, db)
	if err != nil {
		return nil, err
	}
	return unwrapValues(results)
}

func populateChannel(numbers []float64, createDbClient func() (client.Client, error)) error {
	dbClient, err := createDbClient()
	if err != nil {
		return fmt.Errorf(errPassing, "CreateClient", err.Error())
	}
	defer closeClient(dbClient)

	config := client.BatchPointsConfig{Database: db}
	batchPoints, err := client.NewBatchPoints(config)
	if err != nil {
		return fmt.Errorf(errPassing, "NewBatchPoints", err.Error())
	}
	tags := map[string]string{}
	for _, number := range numbers {
		fields := map[string]interface{}{"value": number}
		now := time.Now()
		point, err := client.NewPoint("channel", tags, fields, now)
		if err != nil {
			return fmt.Errorf(errPassing, "NewPoint", err.Error())
		}
		batchPoints.AddPoint(point)
	}
	return dbClient.Write(batchPoints)
}

func unwrapValues(results []client.Result) ([]float64, error) {
	if len(results) != 1 {
		return nil, fmt.Errorf(errResultsLength, len(results))
	}
	result := results[0]
	rows := result.Series
	if len(rows) != 1 {
		return nil, fmt.Errorf(errRowsLength, len(rows))
	}
	return copyNumbers(rows[0].Values)
}

func copyNumbers(values [][]interface{}) ([]float64, error) {
	length := len(values)
	numbers := make([]float64, length)
	for i := 0; i < length; i++ {
		value := values[i][1]
		number, err := value.(json.Number).Float64()
		if err != nil {
			return nil, fmt.Errorf(errJsonNumber, err.Error())
		}
		numbers[i] = number
	}
	return numbers, nil
}

func query(c client.Client, cmd string, db string) ([]client.Result, error) {
	q := client.Query{Command: cmd, Database: db}
	response, err := c.Query(q)
	if err != nil {
		return nil, fmt.Errorf(errQueryFailed, err.Error())
	}
	return response.Results, nil
}

func closeClient(c client.Client) {
	err := c.Close()
	if err != nil {
		log.Panic(errCouldNotCloseClient)
	}
}
