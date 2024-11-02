package main

import (
	"errors"
	"log"
	"reflect"
	"strings"
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
func StructToBins(data interface{}) (aerospike.BinMap, error) {
	bins := aerospike.BinMap{}
	v := reflect.ValueOf(data)

	// If pointer, get the underlying element
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("input must be a struct or pointer to struct")
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Get the aerospike tag
		tag := field.Tag.Get("as")
		if tag == "" {
			continue // Skip fields without aerospike tags
		}

		// Parse tag options
		tagParts := strings.Split(tag, ",")
		binName := tagParts[0]
		omitempty := len(tagParts) > 1 && tagParts[1] == "omitempty"

		// Handle pointer fields
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				if !omitempty {
					return nil, errors.New("nil value for required field: " + field.Name)
				}
				continue
			}
			value = value.Elem()
		}

		// Skip zero values if omitempty is set
		if omitempty && value.IsZero() {
			continue
		}

		// Convert the value to an appropriate type for Aerospike
		switch value.Kind() {
		case reflect.String, reflect.Int, reflect.Int64, reflect.Float64, reflect.Bool:
			bins[binName] = value.Interface()
		case reflect.Struct:
			if value.Type() == reflect.TypeOf(time.Time{}) {
				bins[binName] = value.Interface()
			} else {
				// Handle nested structs if needed
				nestedBins, err := StructToBins(value.Interface())
				if err != nil {
					return nil, err
				}
				for k, v := range nestedBins {
					bins[binName+"_"+k] = v
				}
			}
		default:
			// Handle other types or return error for unsupported types
			return nil, errors.New("unsupported field type: " + field.Name)
		}
	}

	return bins, nil
}
