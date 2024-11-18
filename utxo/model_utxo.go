package utxo

import "time"

// Utxo is an object representing a BitCoin unspent transaction
type Utxo struct {
	Model
	UtxoPointer
	ID           string    `as:"id"`
	XpubID       string    `as:"x_pub_id"`
	Satoshis     uint64    `as:"satoshis"`
	ScriptPubKey string    `as:"script_pub_key"`
	Type         string    `as:"type"`
	DraftID      string    `as:"draft_id"`
	ReservedAt   time.Time `as:"reserved_at"`
	SpendingTxID string    `as:"spending_tx_id"`
}

type UtxoPointer struct {
	TransactionID string  `as:"transaction_id"`
	OutputIndex   uint32  `as:"output_index"`
}

type Model struct {
	CreatedAt     time.Time  `as:"created_at"`
	UpdatedAt     time.Time  `as:"updated_at"`
	DeletedAt     time.Time  `as:"deleted_at"`
	EncryptionKey string     `as:"encryption_key"`
	RawXpubKey    string     `as:"raw_x_pub_key"`
}
