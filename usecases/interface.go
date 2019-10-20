package usecases

import "github.com/vanigabriel/Projeto-Pismo/entity"

// Accounts definition interface
type Accounts interface {
	AddAccount(acc *entity.Account) (*entity.Account, error)
	UpdateAccount(accountID int, creditLimit float64, creditWithdraw float64) (*entity.Account, error)
	GetAccounts() ([]entity.Account, error)
}

// Transactions definition interface
type Transactions interface {
	InsertTransaction(transaction *entity.Transaction) error
}

// Repository interface of accounts and transactions
type Repository interface {
	Accounts
	Transactions
}
