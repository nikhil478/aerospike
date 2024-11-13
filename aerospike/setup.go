package aerospike_db

import (
	"fmt"

	"github.com/aerospike/aerospike-client-go/v7"
)

// AerospikeSetName represents the name of a set in Aerospike
type AerospikeSetName string

const (
	Notif    AerospikeSetName = "notif"
	Activity AerospikeSetName = "activity"
)

// AerospikeConfig holds the Aerospike cluster configuration
type AerospikeConfig struct {
	Address   string `as:"address"`
	Port      int    `as:"port"`
	Namespace string `as:"namespace"`
}

// DefaultConfig provides default Aerospike connection settings
// TODO: get this config from env variable
var DefaultConfig = AerospikeConfig{
	Address:   "127.0.0.1", // Default address
	Port:      3000,        // Default port
	Namespace: "dev",       // Default namespace
}

type AerospikeDB struct {
	createPolicy *aerospike.WritePolicy
	queryPolicy  *aerospike.QueryPolicy
	updatePolicy *aerospike.WritePolicy
	savePolicy   *aerospike.WritePolicy
	readPolicy   *aerospike.BasePolicy
	deletePolicy *aerospike.WritePolicy
	scanPolicy   *aerospike.ScanPolicy
	client       *aerospike.Client
	config       *AerospikeConfig
}

// NewAerospikeClient initializes and returns an Aerospike client
func NewAerospikeClient(config *AerospikeConfig) (*AerospikeDB, error) {
	client, err := aerospike.NewClient(config.Address, config.Port)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Aerospike at %s:%d: %v", config.Address, config.Port, err)
	}

	return &AerospikeDB{
		client:       client,
		createPolicy: getCreatePolicy(),
		queryPolicy:  getQueryPolicy(),
		updatePolicy: getUpdatePolicy(),
		savePolicy:   getSavePolicy(),
		readPolicy:   getReadPolicy(),
		deletePolicy: getDeletePolicy(),
		scanPolicy:   getScanPolicy(),
		config:       config,
	}, nil
}

func (db *AerospikeDB) Close() {
	if db.client != nil {
		db.client.Close()
	}
}