package repository

import (
	"github.com/yogikaushik/bank-api/models"

	"database/sql"
)

type TransactionRepository interface {
	Create(transaction models.Transaction) (models.Transaction, error)
}

type transactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{DB: db}
}

func (r *transactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
	query := "INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Exec(query, transaction.AccountID, transaction.OperationTypeID, transaction.Amount, transaction.EventDate)
	if err != nil {
		return models.Transaction{}, err
	}
	transactionID, err := result.LastInsertId()
	if err != nil {
		return models.Transaction{}, err
	}

	return models.Transaction{
		ID:              int(transactionID),
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          transaction.Amount,
		EventDate:       transaction.EventDate,
	}, nil
}
