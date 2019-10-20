package usecases

import (
	"errors"
	"sort"
	"time"

	"github.com/vanigabriel/Projeto-Pismo/entity"
)

// Statements defines a set of transactions
type Statements struct {
	Transactions []entity.Transaction
	Wallet
}

// SortedTransactions is used internaly in the procedure of correcting balance
type SortedTransactions struct {
	trx []*entity.Transaction
}

// InitStatements return new Statements
func InitStatements() *Statements {
	return &Statements{}
}

// InsertTransaction insert a transaction in Statements
// Return error if anything not right occours
func (statement *Statements) InsertTransaction(transaction *entity.Transaction) error {
	// If not payment, amount must be negative
	if transaction.OperationID != 4 && transaction.Amount > 0 {
		return errors.New("Operation type cannot accept negative amount")
	}

	// If payment, amount must be positive
	if transaction.OperationID != 4 && transaction.Amount > 0 {
		return errors.New("Operation type cannot accept positive amount")
	}

	statement.Lock()
	defer statement.Unlock()

	_, err := statement.findAccountID(transaction.AccountID)
	if err != nil {
		return errors.New("Account not found")
	}
	transaction.Balance = transaction.Amount

	err = statement.correctLimit(transaction)
	if err != nil {
		return err
	}

	if transaction.EventDate == "" {
		currentDate := time.Now()
		transaction.EventDate = currentDate.Format("2006-01-02")
	}

	transaction.TransactionID = len(statement.Transactions) + 1

	// If payment, correct balance of olders transactions
	if transaction.OperationID == 4 {
		statement.correctBalance(transaction)
	}

	statement.Transactions = append(statement.Transactions, *transaction)

	return nil
}

func (statement *Statements) correctLimit(transaction *entity.Transaction) error {
	account, err := statement.findAccountID(transaction.AccountID)
	if err != nil {
		return errors.New("Account not found")
	}

	// Check if exists avaliable credit from past payments
	if account.CreditAvaliable > 0 {
		if transaction.Balance+account.CreditAvaliable >= 0 {
			account.CreditAvaliable += transaction.Balance
			transaction.Balance = 0
		} else {
			transaction.Balance += account.CreditAvaliable
			account.CreditAvaliable = 0
		}
	}

	// Check if is necessary to get from limit
	switch transaction.OperationID {
	case 1, 2:
		if account.CreditLimit+transaction.Balance < 0 {
			return errors.New("Insuficient creditLimit")
		}
		account.CreditLimit += transaction.Balance
	case 3:
		if account.WithdrawlLimit+transaction.Balance < 0 {
			return errors.New("Insuficient withdrawLimit")
		}
		account.WithdrawlLimit += transaction.Balance
	}

	return nil
}

// correctBalance occours when a payment is made
// goes to olders transactions and correct their balance
func (statement *Statements) correctBalance(transaction *entity.Transaction) error {
	trx, err := statement.getAccountsTransactions(transaction)
	if err != nil {
		return err
	}

	account, err := statement.findAccountID(transaction.AccountID)
	if err != nil {
		return errors.New("Account not found")
	}

	balance := transaction.Amount
	sorted := sortTransactions(trx)

	for i := 0; i < len(sorted.trx) && balance > 0; i++ {

		if ((*sorted.trx[i]).Balance + balance) >= 0 {
			// Return limit
			switch (*sorted.trx[i]).OperationID {
			case 1, 2:
				account.CreditLimit -= (*sorted.trx[i]).Balance
			case 3:
				account.WithdrawlLimit -= (*sorted.trx[i]).Balance
			}

			balance = balance + (*sorted.trx[i]).Balance
			(*sorted.trx[i]).Balance = 0
			continue
		}

		// Return limits
		switch (*sorted.trx[i]).OperationID {
		case 1, 2:
			account.CreditLimit -= ((*sorted.trx[i]).Balance + balance)
		case 3:
			account.WithdrawlLimit -= ((*sorted.trx[i]).Balance + balance)
		}

		(*sorted.trx[i]).Balance += balance
		balance = 0
	}
	transaction.Balance = balance

	if transaction.Balance > 0 {
		account.CreditAvaliable += transaction.Balance
	}
	return nil
}

// getAccountsTransactions return all open (with balance) transactions of account
func (statement *Statements) getAccountsTransactions(transaction *entity.Transaction) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction

	for i := 0; i < len(statement.Transactions); i++ {
		if (statement.Transactions[i].AccountID == transaction.AccountID) && (statement.Transactions[i].Balance < 0) {
			transactions = append(transactions, &statement.Transactions[i])
		}
	}
	return transactions, nil
}

func sortTransactions(trx []*entity.Transaction) *SortedTransactions {
	var t SortedTransactions
	t.trx = trx

	sort.Sort(t)

	return &t
}

// Len is used to sort
func (t SortedTransactions) Len() int {
	return len(t.trx)
}

// Less is used to sort
func (t SortedTransactions) Less(i, j int) bool {
	if (*t.trx[i]).OperationID > (*t.trx[j]).OperationID {
		return true
	} else if (*t.trx[i]).OperationID == (*t.trx[j]).OperationID {
		if (*t.trx[i]).EventDate < (*t.trx[j]).EventDate {
			return true
		}
		return false
	}
	return false
}

func (t SortedTransactions) Swap(i, j int) {
	t.trx[i], t.trx[j] = t.trx[j], t.trx[i]
}
