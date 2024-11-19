package utxo

import (
	"fmt"
	"reflect"
	"time"

	aerospike_db "github.com/nikhil478/aerospike/aerospike"
)

func CreateUtxos(db *aerospike_db.AerospikeDB) {
	for i := 1; i < 100; i++ {
		utxo := Utxo{
			UtxoPointer: &UtxoPointer{
				TransactionID: "5966d43afa98a3a1733ef37092439c98a3b964208c4babbd6d31c291a390611e",
				OutputIndex:   uint32(i),
			},
			Satoshis:     100,
			XpubID:       "xPubID",
			ScriptPubKey: "ScriptPubKey",
			Type:         "Type",
			DraftID:      "draftID3",
			Model: &Model{
				CreatedAt: int(time.Now().Unix()),
				UpdatedAt: int(time.Now().Unix()),
			},
		}
		key, err := db.CreateNewRecord(aerospike_db.NewUtxo, &utxo)
		if err != nil {
			fmt.Print("error while creating new utxo", err.Error())
		}
		fmt.Printf("key entered successfully %v", *key)
	}
}

func UpdateUtxos(db *aerospike_db.AerospikeDB) {
	t1 := time.Now()
	results, err := db.GetRecords(aerospike_db.NewUtxo, map[string]any{
		"draft_id": "draftID3",
		"x_pub_id": "xPubID",
	}, &Utxo{})
	if err != nil {
		fmt.Printf("there is an error while processing get records func", err.Error())
	}

	for record := range results {

		fmt.Printf("record value :", *record.Record)

		if record.Err != nil {
			continue
		}
		
		newResult := reflect.New(reflect.TypeOf(&Utxo{}).Elem()).Interface()
		if err := aerospike_db.BinsToStruct(record.Record, newResult); err != nil {
			continue
		}
		utxo , _ :=  newResult.(Utxo)
		utxo.ReservedAt = 1234
		db.UpdateRecord(utxo, record.Record.Key)
	}
	end := time.Now()
    elapsed := end.Sub(t1)
    fmt.Println("Execution time:", elapsed)
}
