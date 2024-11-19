package utxo

import (
	"fmt"
	"time"

	aerospike_db "github.com/nikhil478/aerospike/aerospike"
)

func CreateUtxos(db *aerospike_db.AerospikeDB) {
	for i := 1000; i < 10000; i++ {
		utxo := Utxo{
			UtxoPointer: UtxoPointer{
				TransactionID: "5966d43afa98a3a1733ef37092439c98a3b964208c4babbd6d31c291a390611e",
				OutputIndex:   uint32(i),
			},
			Satoshis:     100,
			XpubID:       "xPubID",
			ScriptPubKey: "ScriptPubKey",
			Type:         "Type",
			DraftID:      "draftID",
			Model: Model{
				CreatedAt: int(time.Now().Unix()),
				UpdatedAt: int(time.Now().Unix()),
			},
		}
		key, err := db.CreateNewRecord(aerospike_db.Utxo, &utxo)
		if err != nil {
			fmt.Print("error while creating new utxo", err.Error())
		}
		fmt.Printf("key entered successfully %v", *key)
	}
}

func UpdateUtxos(db *aerospike_db.AerospikeDB) {
	freeUtxos, err := db.GetRecords(aerospike_db.Utxo, map[string]any{
		"draft_id": "draftID",
	}, &Utxo{})
	if err != nil {
		fmt.Printf("there is an error while processing get records func", err.Error())
	}

	fmt.Printf("utxo list %v", freeUtxos...)
}