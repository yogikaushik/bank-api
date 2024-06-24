package service

import (
	"time"

	"github.com/yogikaushik/bank-api/models"
	"github.com/yogikaushik/bank-api/repository"
)

type TransactionService interface {
	CreateTransaction(accountID, operationTypeID int, amount float64) (models.Transaction, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) CreateTransaction(accountID, operationTypeID int, amount float64) (models.Transaction, error) {
	transaction := models.Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		EventDate:       time.Now().Format(time.DateTime),
	}
	return s.repo.Create(transaction)
}
