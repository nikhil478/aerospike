package main

import (
	"fmt"
	"log"

	aerospike_db "github.com/nikhil478/aerospike/aerospike"
	"github.com/nikhil478/aerospike/utxo"
)

const (
	address   = "127.0.0.1" // Aerospike Cloud cluster address
	port      = 3000        // Aerospike Cloud cluster port
	namespace = "test"      // Cluster namespace
	set       = "foo"       // Set name within namespace
)

func main() {

	config := aerospike_db.AerospikeConfig{
		Address: address,
		Port: port,
		Namespace: namespace,
	} 

	asDb, err := aerospike_db.NewAerospikeClient(&config)
	if err != nil {
		fmt.Printf("there is some issue while creating aerospike instance %v", asDb)
		log.Fatal(err)
	}
	defer asDb.Close()

	utxo.UpdateUtxos(asDb)

	// retrieveObj := aerospike_db.AerospikeConfig{}

	// asKey, err := asDb.CreateNewRecord(aerospike_db.Activity , config)
	// if err != nil {
	// 	fmt.Printf("error whil inserting new record %s", err.Error())
	// }
	
	// err = asDb.GetRecord(asKey, &retrieveObj)

	// if err != nil {
	// 	fmt.Printf("error while fetching record %s", err.Error())
	// }

	// interfaceType := aerospike_db.AerospikeConfig{}

	// arrObj, err := asDb.GetRecords("activity", map[string]string{}, &interfaceType)

	// if err != nil {
	// 	fmt.Printf("error while fetching records %s", err.Error())
	// }

	// fmt.Printf("fetched records succesfully %v", arrObj)

	// err= asDb.DeleteRecord(asKey)

	// if err != nil {
	// 	fmt.Printf("error while deleting records %s", err.Error())
	// }

	// fmt.Printf("deleted record succesfully ")

	
}