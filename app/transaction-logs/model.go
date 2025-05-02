package transactionlogs

import "tms/app/types"

type TransactionLog struct {
	Id              string                `json:"id"`
	TransactionType types.TransactionType `json:"transaction_type"`
	Amount          float64               `json:"amount"`
	AccountId       string                `json:"account_id"`
	Status          types.Status          `json:"status"`
}

type TransactionLogCreate struct {
	TransactionType types.TransactionType `json:"transaction_type"`
	Amount          float64               `json:"amount"`
	AccountId       string                `json:"account_id"`
}
