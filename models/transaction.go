package models

type Transaction struct {
	ID              int     `json:"transaction_id"`
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
	EventDate       string  `json:"event_date"`
}
