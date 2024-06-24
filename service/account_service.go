package service

import (
	"github.com/yogikaushik/bank-api/errors"
	"github.com/yogikaushik/bank-api/models"
	"github.com/yogikaushik/bank-api/repository"
)

type AccountService interface {
	CreateAccount(documentNumber string) (models.Account, error)
	GetAccount(id uint) (models.Account, error)
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

func (s *accountService) CreateAccount(documentNumber string) (models.Account, error) {
	account := models.Account{DocumentNumber: documentNumber}
	return s.repo.Create(account)
}

func (s *accountService) GetAccount(id uint) (models.Account, error) {
	account, err := s.repo.FindByID(id)
	if err != nil {
		return models.Account{}, errors.ErrAccountNotFound
	}
	return account, nil
}
