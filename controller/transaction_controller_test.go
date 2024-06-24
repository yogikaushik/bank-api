package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yogikaushik/bank-api/models"
)

// MockTransactionService is a mock implementation of the TransactionService interface
type MockTransactionService struct {
	mock.Mock
}

func (m *MockTransactionService) CreateTransaction(accountID, operationTypeID int, amount float64) (models.Transaction, error) {
	args := m.Called(accountID, operationTypeID, amount)
	return args.Get(0).(models.Transaction), args.Error(1)
}

func TestCreateTransaction(t *testing.T) {
	mockService := new(MockTransactionService)
	controller := NewTransactionController(mockService)

	transaction := models.Transaction{ID: 1, AccountID: 1, OperationTypeID: 4, Amount: 123.45}
	mockService.On("CreateTransaction", 1, 4, 123.45).Return(transaction, nil)

	input := TransactionInput{AccountID: 1, OperationTypeID: 4, Amount: 123.45}
	body, _ := json.Marshal(input)
	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateTransaction)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var responseTransaction models.Transaction
	err = json.NewDecoder(rr.Body).Decode(&responseTransaction)
	assert.NoError(t, err)
	assert.Equal(t, transaction.ID, responseTransaction.ID)
	assert.Equal(t, transaction.AccountID, responseTransaction.AccountID)
	assert.Equal(t, transaction.OperationTypeID, responseTransaction.OperationTypeID)
	assert.Equal(t, transaction.Amount, responseTransaction.Amount)
	mockService.AssertExpectations(t)
}
