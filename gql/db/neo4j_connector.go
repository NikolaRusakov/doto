package db

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
	"os"
)

var (
	Driver  neo4j.Driver
	err     error
	Session neo4j.Session
	Query   string
)

func Run() (neo4j.Session, error) {
	var (
		//driver           neo4j.Driver
		session neo4j.Session
		//recordsProcessed interface{}
		err error
	)

	// Construct a new Driver
	if Driver, err = neo4j.NewDriver(
		os.Getenv("NEO_URL"),
		neo4j.BasicAuth(os.Getenv("NEO_NAME"),
			os.Getenv("NEO_PASS"),
			""),
		func(config *neo4j.Config) {
			config.Log = neo4j.ConsoleLogger(neo4j.ERROR)
		}); err != nil {
		return nil, err
	}
	defer Driver.Close()

	// Acquire a Session
	if session, err = Driver.Session(neo4j.AccessModeRead); err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer session.Close()
	/*
		// Execute the transaction function
		if recordsProcessed, err = session.ReadTransaction(executeQuery); err != nil {
			return err
		}

		fmt.Printf("\n%d records processed\n", recordsProcessed)
	*/
	return session, nil
}

func executeQuery(tx neo4j.Transaction) (interface{}, error) {
	var (
		counter int
		result  neo4j.Result
		err     error
	)

	// Execute the Query on the provided transaction
	if result, err = tx.Run(Query, nil); err != nil {
		return nil, err
	}

	// Print headers
	//processHeaders(result)

	// Loop through record stream until EOF or error
	for result.Next() {
		processRecord(result.Record())
		counter++
	}
	// Check if we encountered any error during record streaming
	if err = result.Err(); err != nil {
		return nil, err
	}

	// Return counter
	return counter, nil
}

func processRecord(record neo4j.Record) {
	for index, value := range record.Values() {
		if index > 0 {
			fmt.Print("\t")
		}
		fmt.Printf("%-10v", value)
	}
	fmt.Print("\n")
}
