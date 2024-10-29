package main

import (
    "log"
    "time"

    "github.com/aerospike/aerospike-client-go/v7"
)

func main() {
    // ***
    // Setup
    // ***

    address := "127.0.0.1"      // Aerospike Cloud cluster address
    port := 3000                // Aerospike Cloud cluster port
    namespace := "test"         // Cluster namespace
    set := "foo"                // Set name within namespace

    // Create the client and connect to the database
    client, err := aerospike.NewClient(address, port)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // ***
    // Write a record
    // ***

    // Create a WritePolicy to set the TotalTimeout for writes
    // default 1000 ms
    writePolicy := aerospike.NewWritePolicy(0, 0)
    writePolicy.TotalTimeout = 5000 * time.Millisecond

	key := CreateNewRecords(namespace, set, writePolicy, client)

	

    // ***
    // Read back the record we just wrote
    // ***

    // Create a Policy to set the TotalTimeout for reads
    // default 1000 ms
    readPolicy := aerospike.NewPolicy()
    readPolicy.TotalTimeout = 5000 * time.Millisecond

	GetRecords(key, readPolicy, client)

    
}


func CreateNewRecords(namespace, set string, writePolicy *aerospike.WritePolicy, client *aerospike.Client) (key *aerospike.Key) {

	// Create the record key
    // A tuple consisting of namespace, set name, and user defined key
    key, err := aerospike.NewKey(namespace, set, "bar")
    if err != nil {
        log.Fatal(err)
    }

    // Create a bin to store data within the new record
    bin := aerospike.NewBin("myBin", "Hello World!")

    //Write the record to your database
    err = client.PutBins(writePolicy, key, bin)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Succesfully wrote record")
	return key
}


func GetRecords(key *aerospike.Key, readPolicy *aerospike.BasePolicy, client *aerospike.Client) {

	// Create the record key
    // A tuple consisting of namespace, set name, and user defined key
    // Read the record
    record, err := client.Get(readPolicy, key)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Record: %s", record.Bins)
}