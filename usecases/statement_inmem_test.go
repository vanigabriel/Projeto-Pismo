package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanigabriel/Projeto-Pismo/entity"
)

func TestInsertTransaction(t *testing.T) {
	statements := InitStatements()

	transaction := entity.Transaction{
		AccountID:   2,
		OperationID: 1,
		Amount:      -200,
	}

	err := statements.InsertTransaction(&transaction)
	assert.NotNil(t, err)

	account := entity.Account{
		AccountID:      1,
		CreditLimit:    100,
		WithdrawlLimit: 100,
	}

	insertedacc, err := statements.AddAccount(&account)

	transaction.AccountID = insertedacc.AccountID

	// Add a transaction without limit
	err = statements.InsertTransaction(&transaction)
	assert.NotNil(t, err)

	// Add a payment
	payment := entity.Transaction{
		AccountID:   1,
		OperationID: 4,
		Amount:      300,
	}
	err = statements.InsertTransaction(&payment)
	assert.Nil(t, err)

	acc, _ := statements.findAccountID(1)
	assert.Equal(t, acc.CreditAvaliable, 300.)

	// Add a valid transaction with CreditAvaliable
	transaction = entity.Transaction{
		AccountID:   1,
		OperationID: 1,
		Amount:      -200,
	}
	err = statements.InsertTransaction(&transaction)
	assert.Nil(t, err)

	acc, _ = statements.findAccountID(1)
	assert.Equal(t, acc.CreditAvaliable, 100.)

	/*
		statements.UpdateAccount(1, 4900, 4900)

		transaction = entity.Transaction{
			AccountID:   1,
			OperationID: 1,
			Amount:      -200,
		}
		transaction2 := entity.Transaction{
			AccountID:   1,
			OperationID: 2,
			Amount:      -200,
		}
		transaction3 := entity.Transaction{
			AccountID:   1,
			OperationID: 3,
			Amount:      -200,
		}
		transaction4 := entity.Transaction{
			AccountID:   1,
			OperationID: 3,
			Amount:      -200,
			EventDate:   "2019-06-10",
		}
		transaction5 := entity.Transaction{
			AccountID:   1,
			OperationID: 3,
			Amount:      -200,
			EventDate:   "2019-06-11",
		}
		transaction6 := entity.Transaction{
			AccountID:   1,
			OperationID: 3,
			Amount:      -200,
			EventDate:   "2019-07-10",
		}

		log.Println(statements.Accounts)

		statements.InsertTransaction(&transaction3)
		statements.InsertTransaction(&transaction2)
		statements.InsertTransaction(&transaction)
		statements.InsertTransaction(&transaction4)
		statements.InsertTransaction(&transaction5)
		statements.InsertTransaction(&transaction6)

		for i := 0; i < len(statements.Transactions); i++ {
			log.Println(statements.Transactions[i])
		}

		log.Println("before payments")
		log.Println(statements.Accounts)

		payment := entity.Transaction{
			AccountID:   1,
			OperationID: 4,
			Amount:      300,
		}
		payment2 := entity.Transaction{
			AccountID:   1,
			OperationID: 4,
			Amount:      5000,
		}
		statements.InsertTransaction(&payment)
		statements.InsertTransaction(&payment2)

		for i := 0; i < len(statements.Transactions); i++ {
			log.Println(statements.Transactions[i])
		}

		log.Println(statements.Accounts)

		transaction = entity.Transaction{
			AccountID:   1,
			OperationID: 1,
			Amount:      -500,
		}
		statements.InsertTransaction(&transaction)
		for i := 0; i < len(statements.Transactions); i++ {
			log.Println(statements.Transactions[i])
		}

		log.Println(statements.Accounts)*/

}

func TestSortTransactions(t *testing.T) {
	statements := InitStatements()

	transaction := entity.Transaction{
		AccountID:   1,
		OperationID: 1,
		Amount:      -200,
	}
	transaction2 := entity.Transaction{
		AccountID:   1,
		OperationID: 2,
		Amount:      -200,
	}
	transaction3 := entity.Transaction{
		AccountID:   1,
		OperationID: 3,
		Amount:      -200,
	}
	transaction4 := entity.Transaction{
		AccountID:   1,
		OperationID: 3,
		Amount:      -200,
		EventDate:   "2019-06-10",
	}
	transaction5 := entity.Transaction{
		AccountID:   1,
		OperationID: 3,
		Amount:      -200,
		EventDate:   "2019-06-11",
	}
	transaction6 := entity.Transaction{
		AccountID:   1,
		OperationID: 3,
		Amount:      -200,
		EventDate:   "2019-07-10",
	}

	account := entity.Account{
		CreditLimit:    5000.12,
		WithdrawlLimit: 5000.12,
	}

	statements.AddAccount(&account)
	statements.InsertTransaction(&transaction3)
	statements.InsertTransaction(&transaction2)
	statements.InsertTransaction(&transaction)
	statements.InsertTransaction(&transaction4)
	statements.InsertTransaction(&transaction5)
	statements.InsertTransaction(&transaction6)

	statements.correctBalance(&transaction3)
	tr, _ := statements.getAccountsTransactions(&transaction)
	trx := sortTransactions(tr)
	assert.Equal(t, trx.trx[0].OperationID, 3)
	assert.Equal(t, trx.trx[0].EventDate, "2019-06-10")
	assert.Equal(t, trx.trx[1].OperationID, 3)
	assert.Equal(t, trx.trx[1].EventDate, "2019-06-11")
	assert.Equal(t, trx.trx[2].OperationID, 3)
	assert.Equal(t, trx.trx[2].EventDate, "2019-07-10")
	assert.Equal(t, trx.trx[3].OperationID, 3)
	assert.Equal(t, trx.trx[4].OperationID, 2)
	assert.Equal(t, trx.trx[5].OperationID, 1)
}
