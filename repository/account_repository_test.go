package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yogikaushik/bank-api/models"
)

func TestCreateAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := models.Account{DocumentNumber: "12345678900"}

	mock.ExpectExec("INSERT INTO accounts").
		WithArgs(account.DocumentNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))

	createdAccount, err := repo.Create(account)
	assert.NoError(t, err)
	assert.Equal(t, 1, createdAccount.ID)
	assert.Equal(t, account.DocumentNumber, createdAccount.DocumentNumber)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := models.Account{ID: 1, DocumentNumber: "12345678900"}

	rows := sqlmock.NewRows([]string{"account_id", "document_number"}).
		AddRow(account.ID, account.DocumentNumber)

	mock.ExpectQuery("SELECT account_id, document_number FROM accounts WHERE account_id = ?").
		WithArgs(account.ID).
		WillReturnRows(rows)

	foundAccount, err := repo.FindByID(uint(account.ID))
	assert.NoError(t, err)
	assert.Equal(t, account.ID, foundAccount.ID)
	assert.Equal(t, account.DocumentNumber, foundAccount.DocumentNumber)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	mock.ExpectQuery("SELECT account_id, document_number FROM accounts WHERE account_id = ?").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.FindByID(1)
	assert.Error(t, err)
	assert.Equal(t, "account not found", err.Error())

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
