package controller

import (
	"encoding/json"
	"net/http"

	"github.com/yogikaushik/bank-api/service"
)

// TransactionController handles transaction-related requests
type TransactionController struct {
	Service service.TransactionService
}

// NewTransactionController creates a new TransactionController
func NewTransactionController(s service.TransactionService) *TransactionController {
	return &TransactionController{Service: s}
}

type TransactionInput struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Create a new transaction with the given details
// @Tags transactions
// @Accept  json
// @Produce  json
//
//	@Param transaction body TransactionInput true "Transaction details"
//
// @Success 200 {object} models.Transaction
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /transactions [post]
func (c *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var input TransactionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transaction, err := c.Service.CreateTransaction(input.AccountID, input.OperationTypeID, input.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}
