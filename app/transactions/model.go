package transactions

import "tms/app/types"

type Transaction struct {
	Id            string `json:"id"`
	AccountId     string `json:"account_id"`
	Debit         string `json:"debit"`
	Credit        string `json:"credit"`
	TransactionId string `json:"transaction_id"`
}

type TransactionCreate struct {
	TransactionType types.TransactionType `json:"transaction_type"`
	Amount          float64               `json:"amount"`
	AccountId       string                `json:"account_id"`
	TransactionId   string                `json:"transaction_id"`
}
