package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yogikaushik/bank-api/models"
)

// MockAccountService is a mock implementation of the AccountService interface
type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) CreateAccount(documentNumber string) (models.Account, error) {
	args := m.Called(documentNumber)
	return args.Get(0).(models.Account), args.Error(1)
}

func (m *MockAccountService) GetAccount(id uint) (models.Account, error) {
	args := m.Called(id)
	return args.Get(0).(models.Account), args.Error(1)
}

func TestCreateAccount(t *testing.T) {
	mockService := new(MockAccountService)
	controller := NewAccountController(mockService)

	account := models.Account{ID: 1, DocumentNumber: "12345678900"}
	mockService.On("CreateAccount", "12345678900").Return(account, nil)

	input := AccountInput{DocumentNumber: "12345678900"}
	body, _ := json.Marshal(input)
	req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateAccount)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var responseAccount models.Account
	err = json.NewDecoder(rr.Body).Decode(&responseAccount)
	assert.NoError(t, err)
	assert.Equal(t, account.ID, responseAccount.ID)
	assert.Equal(t, account.DocumentNumber, responseAccount.DocumentNumber)
	mockService.AssertExpectations(t)
}

func TestGetAccount(t *testing.T) {
	mockService := new(MockAccountService)
	controller := NewAccountController(mockService)

	account := models.Account{ID: 1, DocumentNumber: "12345678900"}
	mockService.On("GetAccount", uint(1)).Return(account, nil)

	req, err := http.NewRequest("GET", "/accounts/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/accounts/{id}", controller.GetAccount).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var responseAccount models.Account
	err = json.NewDecoder(rr.Body).Decode(&responseAccount)
	assert.NoError(t, err)
	assert.Equal(t, account.ID, responseAccount.ID)
	assert.Equal(t, account.DocumentNumber, responseAccount.DocumentNumber)
	mockService.AssertExpectations(t)
}
