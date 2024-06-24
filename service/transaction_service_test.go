package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yogikaushik/bank-api/models"
)

// MockTransactionRepository is a mock implementation of the TransactionRepository interface
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(models.Transaction), args.Error(1)
}

func TestCreateTransaction(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)

	transaction := models.Transaction{
		AccountID:       1,
		OperationTypeID: 4,
		Amount:          123.45,
		EventDate:       time.Now().Format(time.RFC3339),
	}

	mockRepo.On("Create", mock.Anything).Return(transaction, nil)

	createdTransaction, err := service.CreateTransaction(1, 4, 123.45)
	assert.NoError(t, err)
	assert.Equal(t, transaction.AccountID, createdTransaction.AccountID)
	assert.Equal(t, transaction.OperationTypeID, createdTransaction.OperationTypeID)
	assert.Equal(t, transaction.Amount, createdTransaction.Amount)
	assert.Equal(t, transaction.EventDate, createdTransaction.EventDate)

	mockRepo.AssertExpectations(t)
}
