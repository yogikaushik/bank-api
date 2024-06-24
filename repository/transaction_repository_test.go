package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yogikaushik/bank-api/models"
)

func TestCreateTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewTransactionRepository(db)

	transaction := models.Transaction{
		AccountID:       1,
		OperationTypeID: 4,
		Amount:          123.45,
		EventDate:       time.Now().Format(time.DateTime),
	}

	mock.ExpectExec("INSERT INTO transactions").
		WithArgs(transaction.AccountID, transaction.OperationTypeID, transaction.Amount, transaction.EventDate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	createdTransaction, err := repo.Create(transaction)
	assert.NoError(t, err)
	assert.Equal(t, 1, createdTransaction.ID)
	assert.Equal(t, transaction.AccountID, createdTransaction.AccountID)
	assert.Equal(t, transaction.OperationTypeID, createdTransaction.OperationTypeID)
	assert.Equal(t, transaction.Amount, createdTransaction.Amount)
	tm, _ := time.Parse(time.DateTime, transaction.EventDate)
	createdTm, _ := time.Parse(time.DateTime, createdTransaction.EventDate)
	assert.WithinDuration(t, tm, createdTm, time.Second)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
