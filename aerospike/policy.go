package aerospike_db

import (
	"time"

	"github.com/aerospike/aerospike-client-go/v7"
)

func getCreatePolicy() *aerospike.WritePolicy {
	writePolicy := aerospike.NewWritePolicy(0, 0)
	writePolicy.TotalTimeout = 5000 * time.Millisecond
	writePolicy.RecordExistsAction = aerospike.CREATE_ONLY
	return writePolicy
}

func getQueryPolicy() *aerospike.QueryPolicy {
	queryPolicy := aerospike.NewQueryPolicy()
	queryPolicy.TotalTimeout = 5000 * time.Millisecond
	return queryPolicy
}


func getUpdatePolicy() *aerospike.WritePolicy {
	writePolicy := aerospike.NewWritePolicy(0, 0)
	writePolicy.TotalTimeout = 5000 * time.Millisecond
	writePolicy.RecordExistsAction = aerospike.UPDATE_ONLY
	return writePolicy
}


func getSavePolicy() *aerospike.WritePolicy {
	writePolicy := aerospike.NewWritePolicy(0, 0)
	writePolicy.TotalTimeout = 5000 * time.Millisecond
	return writePolicy
}


func getReadPolicy() *aerospike.BasePolicy {
	readPolicy := aerospike.NewPolicy()
	readPolicy.TotalTimeout = 5000 * time.Millisecond
	return readPolicy
}

func getDeletePolicy() *aerospike.WritePolicy {
	deletePolicy := aerospike.NewWritePolicy(0, 0)
	deletePolicy.DurableDelete = true
	return deletePolicy
}

func getScanPolicy() *aerospike.ScanPolicy {
	scanPolicy := aerospike.NewScanPolicy()
	scanPolicy.TotalTimeout = 5000 * time.Millisecond
	return scanPolicy
}