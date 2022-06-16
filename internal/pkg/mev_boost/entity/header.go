package entity

import (
	"time"
)

type RelayPayload struct {
	ID               uint64    `db:"id" json:"id"`
	SlotNumber       uint64    `db:"slot" json:"slot_number"`
	BlockHash        string    `db:"block_hash" json:"block_hash"`
	BlockNumber      uint64    `db:"block_number" json:"block_number"`
	FeeRecipient     string    `db:"fee_recipient" json:"fee_recipient"`
	TransactionsRoot string    `db:"transactions_root" json:"transactions_root"`
	Pubkey           string    `db:"pubkey" json:"pubkey"`
	Signature        string    `db:"signature" json:"signature"`
	RelayAddr        string    `db:"relay_adr" json:"relay_adr"`
	RelayTimestamp   time.Time `db:"relay_timestamp" json:"relay_timestamp"`
	Created          time.Time `db:"created_at" json:"created_at"`
}

type RelayPayloadPaginated struct {
	Total        uint64         `json:"total"`
	PerPage      uint64         `json:"per_page"`
	CurrentPage  uint64         `json:"current_page"`
	LastPage     uint64         `json:"last_page"`
	FirstPageURL string         `json:"first_page_url"`
	LastPageUrl  string         `json:"last_page_url"`
	NextPageURL  string         `json:"next_page_url"`
	PrevPageURL  string         `json:"prev_page_url"`
	From         uint64         `json:"from"`
	To           uint64         `json:"to"`
	Data         []RelayPayload `json:"data"`
}
