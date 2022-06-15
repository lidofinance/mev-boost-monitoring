package dto

import (
	"fmt"
	"strconv"
	"time"

	"github.com/lidofinance/mev-boost-monitoring/internal/pkg/mev_boost/entity"
)

type RelayPayload struct {
	Message struct {
		Value  string `json:"value"`
		Header struct {
			GasUsed          string `json:"gas_used"`
			GasLimit         string `json:"gas_limit"`
			Timestamp        string `json:"timestamp"`
			BlockHash        string `json:"block_hash"`
			ExtraData        string `json:"extra_data"`
			LogsBloom        string `json:"logs_bloom"`
			StateRoot        string `json:"state_root"`
			ParentHash       string `json:"parent_hash"`
			PrevRandao       string `json:"prev_randao"`
			BlockNumber      string `json:"block_number"`
			FeeRecipient     string `json:"fee_recipient"`
			ReceiptsRoot     string `json:"receipts_root"`
			BaseFeePerGas    string `json:"base_fee_per_gas"`
			TransactionsRoot string `json:"transactions_root"`
		} `json:"header"`
		Pubkey string `json:"pubkey"`
	} `json:"message"`
	Signature string `json:"signature"`
}

type CustomRelayPayload struct {
	SlotNumber       uint64 `json:"slot" binding:"required"`
	BlockHash        string `json:"block_hash" binding:"required"`
	BlockNumber      uint64 `json:"block_number" binding:"required"`
	FeeRecipient     string `json:"fee_recipient" binding:"required"`
	TransactionsRoot string `json:"transactions_root" binding:"required"`
	Pubkey           string `json:"pubkey" binding:"required"`
	Signature        string `json:"signature" binding:"required"`
	RelayAddr        string `json:"relay_adr" binding:"required"`
	RelayTimestamp   string `json:"timestamp"`
}

func (p *CustomRelayPayload) Validate() (*entity.RelayPayload, error) {
	if p.BlockHash == "" {
		return nil, fmt.Errorf(`field: block_hash is invalid`)
	}

	if p.FeeRecipient == "" {
		return nil, fmt.Errorf(`field: fee_recipient is invalid`)
	}

	if p.TransactionsRoot == "" {
		return nil, fmt.Errorf(`field: transactions_root is invalid`)
	}

	if p.Pubkey == "" {
		return nil, fmt.Errorf(`field: pubkey is invalid`)
	}

	if p.Signature == "" {
		return nil, fmt.Errorf(`field: signature is invalid`)
	}

	if p.RelayAddr == "" {
		return nil, fmt.Errorf(`field: relay_adr is invalid`)
	}

	i, err := strconv.ParseInt(p.RelayTimestamp, 10, 64)
	if err != nil {
		return nil, fmt.Errorf(`field: timestamp is invalid`)
	}
	tm := time.Unix(i, 0)

	return &entity.RelayPayload{
		SlotNumber:       p.SlotNumber,
		BlockHash:        p.BlockHash,
		BlockNumber:      p.BlockNumber,
		FeeRecipient:     p.FeeRecipient,
		TransactionsRoot: p.TransactionsRoot,
		Pubkey:           p.Pubkey,
		Signature:        p.Signature,
		RelayAddr:        p.RelayAddr,
		RelayTimestamp:   tm,
	}, nil
}
