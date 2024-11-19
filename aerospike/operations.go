package aerospike_db

import (
	"errors"
	"reflect"

	"github.com/aerospike/aerospike-client-go/v7"
	"github.com/google/uuid"
)

// TODO : Use Bin objects instead of BinMap in Put operations. Put methods require passing a map for bin values. This allocates an array of bins on each call, iterates on the map, and creates Bin objects. Avoid BinMap allocation by using PutBins to manually pass the bins, allocate []Bin, and iterate over the BinMap. This creates two allocations and an O(n) algorithm.
// refer : https://aerospike.com/developer/client/best_practices pt. 3

func (adb *AerospikeDB) CreateNewRecord(setName AerospikeSetName, data interface{}) (*aerospike.Key, error) {

	
	binMap, err := StructToBins(data)
	if err != nil {
		return nil, errors.New("failed to convert struct to bins" + err.Error())
	}

	id := uuid.New().String()
	key, err := aerospike.NewKey(adb.config.Namespace, string(setName), id)
	if err != nil {
		return nil, errors.New("failed to create Aerospike key")
	}

	err = adb.client.Put(adb.createPolicy, key, binMap)
	if err != nil {
		return nil, errors.New("failed to insert record")
	}

	return key, nil
}

func (adb *AerospikeDB) GetRecord(key *aerospike.Key, result interface{}) error {

	record, err := adb.client.Get(adb.readPolicy, key)
	if err != nil {
		return errors.New("failed to get record")
	}

	if err := BinsToStruct(record, result); err != nil {
		return errors.New("failed to convert bins to struct")
	}

	return nil
}

func (adb *AerospikeDB) UpdateRecord(data interface{}, key *aerospike.Key) error {

	binMap, err := StructToBins(data)
	if err != nil {
		return errors.New("failed to convert struct to bins" + err.Error())
	}

	err = adb.client.Put(adb.updatePolicy, key, binMap)
	if err != nil {
		return errors.New("failed to update record in Aerospike: " + err.Error())
	}

	return nil
}

func (adb *AerospikeDB) DeleteRecord(key *aerospike.Key) error {
	_, err := adb.client.Delete(adb.deletePolicy, key)
	if err != nil {
		return errors.New("failed to delete record with key")
	}
	return nil
}

func (adb *AerospikeDB) GetRecords(setName AerospikeSetName, conditions map[string]string, result interface{}) ([]interface{}, error) {

	statement := aerospike.NewStatement(adb.config.Namespace, string(setName))

	// note we can only pass one condition now there is an workaroud for multiple condition with policies which we havent yet implemented
	for binName, binValue := range conditions {
		filter := aerospike.NewEqualFilter(binName, binValue)
		statement.SetFilter(filter)
	}

	recordset, err := adb.client.Query(adb.queryPolicy, statement)
	if err != nil {
		return nil, err
	}
	defer recordset.Close()

	var resultList []interface{}

	for record := range recordset.Results() {

		if record.Err != nil {
			continue
		}
		newResult := reflect.New(reflect.TypeOf(result).Elem()).Interface()
		if err := BinsToStruct(record.Record, newResult); err != nil {
			continue
		}
		resultList = append(resultList, newResult)
	}
	return resultList, nil
}
