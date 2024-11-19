package utxo

// Utxo is an object representing a BitCoin unspent transaction
type Utxo struct {
	Model
	UtxoPointer
	ID           string    `as:"id"`
	XpubID       string    `as:"x_pub_id"`
	Satoshis     int64     `as:"satoshis"` // as aerospike doesnt support uin64 i m converting this too int64 for now https://github.com/aerospike/aerospike-client-go/issues/62
	ScriptPubKey string    `as:"script_pub_key"`
	Type         string    `as:"type"`
	DraftID      string    `as:"draft_id"`
	ReservedAt   int `as:"reserved_at"`
	SpendingTxID string    `as:"spending_tx_id"`
}

type UtxoPointer struct {
	TransactionID string `as:"transaction_id"`
	OutputIndex   uint32 `as:"output_index"`
}

type Model struct {
	CreatedAt     int `as:"created_at"`
	UpdatedAt     int `as:"updated_at"`
	DeletedAt     int `as:"deleted_at"`
	EncryptionKey string    `as:"encryption_key"`
	RawXpubKey    string    `as:"raw_x_pub_key"`
}
