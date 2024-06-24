package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yogikaushik/bank-api/errors"
	"github.com/yogikaushik/bank-api/models"
)

// MockAccountRepository is a mock implementation of the AccountRepository interface
type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) Create(account models.Account) (models.Account, error) {
	args := m.Called(account)
	return args.Get(0).(models.Account), args.Error(1)
}

func (m *MockAccountRepository) FindByID(id uint) (models.Account, error) {
	args := m.Called(id)
	return args.Get(0).(models.Account), args.Error(1)
}

func TestCreateAccount(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	service := NewAccountService(mockRepo)

	account := models.Account{ID: 1, DocumentNumber: "12345678900"}
	mockRepo.On("Create", mock.Anything).Return(account, nil)

	createdAccount, err := service.CreateAccount("12345678900")
	assert.NoError(t, err)
	assert.Equal(t, account.ID, createdAccount.ID)
	assert.Equal(t, account.DocumentNumber, createdAccount.DocumentNumber)

	mockRepo.AssertExpectations(t)
}

func TestGetAccount(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	service := NewAccountService(mockRepo)

	account := models.Account{ID: 1, DocumentNumber: "12345678900"}
	mockRepo.On("FindByID", uint(1)).Return(account, nil)

	foundAccount, err := service.GetAccount(1)
	assert.NoError(t, err)
	assert.Equal(t, account.ID, foundAccount.ID)
	assert.Equal(t, account.DocumentNumber, foundAccount.DocumentNumber)

	mockRepo.AssertExpectations(t)
}

func TestGetAccount_NotFound(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	service := NewAccountService(mockRepo)

	mockRepo.On("FindByID", uint(1)).Return(models.Account{}, errors.ErrAccountNotFound)

	_, err := service.GetAccount(1)
	assert.Error(t, err)
	assert.Equal(t, errors.ErrAccountNotFound, err)

	mockRepo.AssertExpectations(t)
}
