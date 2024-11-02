package main

import (
	"log"
	"time"

	"github.com/aerospike/aerospike-client-go/v7"
)

const (
	address   = "127.0.0.1" // Aerospike Cloud cluster address
	port      = 3000        // Aerospike Cloud cluster port
	namespace = "test"      // Cluster namespace
	set       = "foo"       // Set name within namespace
)

// Create a WritePolicy to set the TotalTimeout of 5000ms for writes
func GetCreatePolicy() *aerospike.WritePolicy {
	writePolicy := aerospike.NewWritePolicy(0, 0)
	writePolicy.TotalTimeout = 5000 * time.Millisecond
	writePolicy.RecordExistsAction = aerospike.CREATE_ONLY
	return writePolicy
}

// Create a UpdatePolicy to set the TotalTimeout of 5000ms for writes
func GetUpdatePolicy() *aerospike.WritePolicy {
	writePolicy := aerospike.NewWritePolicy(0, 0)
	writePolicy.TotalTimeout = 5000 * time.Millisecond
	writePolicy.RecordExistsAction = aerospike.UPDATE_ONLY
	return writePolicy
}

// Create a UpdatePolicy to set the TotalTimeout of 5000ms for writes
func GetSavePolicy() *aerospike.WritePolicy {
	writePolicy := aerospike.NewWritePolicy(0, 0)
	writePolicy.TotalTimeout = 5000 * time.Millisecond
	return writePolicy
}

// Update or insert (upsert) the record. Merges new bin data if the record exists.
func GetReadPolicy() *aerospike.BasePolicy {
	readPolicy := aerospike.NewPolicy()
	readPolicy.TotalTimeout = 5000 * time.Millisecond
	return readPolicy
}

func GetDeletePolicy() *aerospike.WritePolicy {
	deletePolicy := aerospike.NewWritePolicy(0, 0)
	deletePolicy.DurableDelete = true
	return deletePolicy
}

func main() {

	client, err := aerospike.NewClient(address, port)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	key := CreateNewRecords(namespace, set, client)

	GetRecords(key, client)

	_, err = client.Delete(GetDeletePolicy(), key)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateNewRecords(namespace, set string, client *aerospike.Client) (key *aerospike.Key) {

	// Create the record key
	// A tuple consisting of namespace, set name, and user defined key
	key, err := aerospike.NewKey(namespace, set, "bar")
	if err != nil {
		log.Fatal(err)
	}

	type CreateUser struct {
		Email string
		Name  string
		Age   int
	}

	user := map[string]interface{}{
		"email": "nikhilmatta@gmail.com",
		"nikhil":  "Nikhil Matta",
		"age":   21,
	}

    binMap := aerospike.BinMap(user)

	// bin, _ := StructToBins(user)

	// Create a bin to store data within the new record
	// client.PutBins(GetCreatePolicy(), key, aerospike.NewBin("myBin", "Hello World!"))

	//Write the record to your database
	err = client.Put(GetCreatePolicy(), key, binMap)


	if err != nil {
		log.Fatal(err)
	}
	log.Println("Succesfully wrote record")
	return key
}

func GetRecords(key *aerospike.Key, client *aerospike.Client) {

	// Create the record key
	// A tuple consisting of namespace, set name, and user defined key
	// Read the record
	record, err := client.Get(GetReadPolicy(), key)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Record: %s", record.Bins)
}

// structToBins converts any struct to Aerospike BinMap using reflection



