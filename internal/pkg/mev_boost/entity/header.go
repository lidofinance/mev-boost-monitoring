package entity

import (
	"time"
)

type RelayPayload struct {
	ID               uint64    `db:"slot" json:"id"`
	SlotNumber       uint64    `db:"slot" json:"slot_number"`
	BlockHash        string    `db:"block_hash" json:"block_hash"`
	BlockNumber      uint64    `db:"block_number" json:"block_number"`
	FeeRecipient     string    `db:"fee_recipient" json:"fee_recipient"`
	TransactionsRoot string    `db:"transactions_root" json:"transactions_root"`
	Pubkey           string    `db:"pubkey" json:"pubkey"`
	Signature        string    `db:"signature" json:"signature"`
	RelayAddr        string    `db:"relay_adr" json:"relay_addr"`
	RelayTimestamp   time.Time `db:"timestamp" json:"relay_timestamp"`
}
