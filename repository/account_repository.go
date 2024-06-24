package repository

import (
	"fmt"

	"github.com/yogikaushik/bank-api/models"

	"database/sql"
)

type AccountRepository interface {
	Create(account models.Account) (models.Account, error)
	FindByID(id uint) (models.Account, error)
}

type accountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{DB: db}
}

func (r *accountRepository) Create(account models.Account) (models.Account, error) {
	query := "INSERT INTO accounts (document_number) VALUES (?)"
	result, err := r.DB.Exec(query, account.DocumentNumber)
	if err != nil {
		return models.Account{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.Account{}, err
	}
	return models.Account{ID: int(id), DocumentNumber: account.DocumentNumber}, nil
}

func (r *accountRepository) FindByID(id uint) (models.Account, error) {

	query := "SELECT account_id, document_number FROM accounts WHERE account_id = ?"
	row := r.DB.QueryRow(query, id)

	var account models.Account
	err := row.Scan(&account.ID, &account.DocumentNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return account, fmt.Errorf("account not found")
		}
		return account, err
	}
	return account, nil

}
