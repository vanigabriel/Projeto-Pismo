package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanigabriel/Projeto-Pismo/entity"
)

func TestAddAccount(t *testing.T) {

	wallet := InitWallet()

	account := entity.Account{
		AccountID:      1,
		CreditLimit:    123.12,
		WithdrawlLimit: 123.12,
	}

	// First add
	_, err := wallet.AddAccount(&account)
	assert.Nil(t, err)

	// Duplicity add
	_, err = wallet.AddAccount(&account)
	assert.NotNil(t, err)

	account2 := entity.Account{
		CreditLimit:    123.12,
		WithdrawlLimit: 123.12,
	}
	// Without AccountID
	inserted, err := wallet.AddAccount(&account2)
	assert.Nil(t, err)
	assert.NotEqual(t, inserted.AccountID, 0)
	assert.Equal(t, inserted.CreditLimit, 123.12)

}

func TestFindAccount(t *testing.T) {

	wallet := InitWallet()

	account := entity.Account{
		AccountID:      1,
		CreditLimit:    123.12,
		WithdrawlLimit: 123.12,
	}

	account2 := entity.Account{
		AccountID:      99,
		CreditLimit:    123.12,
		WithdrawlLimit: 123.12,
	}

	// Search in empty wallet
	_, err := wallet.findAccountID(1)
	assert.NotNil(t, err)

	// Add accounts
	wallet.AddAccount(&account)
	wallet.AddAccount(&account2)

	// Check if exists
	finded, err := wallet.findAccountID(1)
	assert.Nil(t, err)
	assert.EqualValues(t, finded.AccountID, 1)

	// Check if exists
	finded, err = wallet.findAccountID(99)
	assert.Nil(t, err)
	assert.EqualValues(t, finded.AccountID, 99)

	// Search a not inserted account
	_, err = wallet.findAccountID(2)
	assert.NotNil(t, err)

}

func TestUpdateAccount(t *testing.T) {
	wallet := InitWallet()

	account := entity.Account{
		AccountID:      1,
		CreditLimit:    100.,
		WithdrawlLimit: 200.,
	}

	// Update a not inserted account
	_, err := wallet.UpdateAccount(1, 200., 200.)
	assert.NotNil(t, err)

	// Add accounts
	wallet.AddAccount(&account)

	// Update account correctly
	updated, err := wallet.UpdateAccount(1, 200., 200.)
	assert.Nil(t, err)
	assert.EqualValues(t, updated.CreditLimit, 300)
	assert.EqualValues(t, updated.WithdrawlLimit, 400)

	finded, _ := wallet.findAccountID(1)
	assert.Equal(t, updated, finded)

	// Update account correctly
	updated, err = wallet.UpdateAccount(1, -200., -200.)
	assert.Nil(t, err)
	assert.EqualValues(t, updated.CreditLimit, 100)
	assert.EqualValues(t, updated.WithdrawlLimit, 200)

	finded, _ = wallet.findAccountID(1)
	assert.Equal(t, updated, finded)

	// Update creditlimit and withdrawlimit to negative
	_, err = wallet.UpdateAccount(1, -200., -200.)
	assert.NotNil(t, err)

}

func TestGetAccounts(t *testing.T) {
	wallet := InitWallet()

	account := entity.Account{
		AccountID:      1,
		CreditLimit:    123.12,
		WithdrawlLimit: 123.12,
	}

	account2 := entity.Account{
		AccountID:      99,
		CreditLimit:    123.12,
		WithdrawlLimit: 123.12,
	}

	var accounts []entity.Account

	// Get empty wallet
	_, err := wallet.GetAccounts()
	assert.NotNil(t, err)

	// Add accounts
	wallet.AddAccount(&account)
	wallet.AddAccount(&account2)

	// Get accounts correctly
	accounts, err = wallet.GetAccounts()
	assert.Nil(t, err)
	assert.NotEqual(t, accounts[0].AccountID, 0)
	assert.Equal(t, len(accounts), 2)

}
