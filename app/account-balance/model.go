package accountbalance

import "tms/app/types"

type AccountBalance struct {
	Id        string `json:"id"`
	AccountId string `json:"account_id"`
	Balance   string `json:"balance"`

	TransactionId string `json:"transaction_id"`
}

type AccountBalanceCreate struct {
	TransactionType types.TransactionType `json:"transaction_type"`
	Amount          float64               `json:"amount"`
	AccountId       string                `json:"account_id"`
	TransactionId   string                `json:"transaction_id"`
}
