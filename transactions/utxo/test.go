package utxo

import (
	"fmt"
	"time"
	"github.com/nikhil478/aerospike/aerospike"
)

type Utxo struct {
	TxID        string
	OutputIndex int
}

func StressTest(db *aerospike_db.AerospikeDB) {
	// Outer loop to simulate stress at increasing iterations
	// for j := 1; j <= 100000; j *= 10 {
		start := time.Now()

		// Initialize a Utxo instance to be inserted
		utxo := Utxo{
			TxID:        "transactionID",  // Example TxID
			OutputIndex: 1,
		}

		// Inner loop to simulate repeated record creation (10 times per outer loop iteration)
		// for i := 0; i < j; i++ {
			_, err := db.CreateNewRecord(aerospike_db.Activity, &utxo)
			if err != nil {
				fmt.Printf("error while inserting new record: %s\n", err.Error())
			}
		// }

		// Calculate elapsed time
		elapsed := time.Since(start)
		fmt.Printf("Stress test iteration %d completed. Time taken: %v\n", 1, elapsed)
	// }
}
