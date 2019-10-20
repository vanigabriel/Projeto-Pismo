package usecases

import (
	"errors"
	"log"
	"sync"

	"github.com/vanigabriel/Projeto-Pismo/entity"
)

// Wallet defines a set of accounts
type Wallet struct {
	sync.Mutex
	Accounts []entity.Account
}

// InitWallet returns a new Wallet
func InitWallet() *Wallet {
	return &Wallet{}
}

func (wallet *Wallet) findAccountID(accID int) (*entity.Account, error) {
	if accID == 0 {
		return nil, errors.New("accID zero")
	}

	for i := 0; i < len(wallet.Accounts); i++ {
		if wallet.Accounts[i].AccountID == accID {
			return &wallet.Accounts[i], nil
		}
	}

	return nil, errors.New("Not found")
}

// AddAccount receive an account and try to add in wallet
// If Account already exists, returns error
func (wallet *Wallet) AddAccount(acc *entity.Account) (*entity.Account, error) {
	wallet.Lock()
	defer wallet.Unlock()

	if acc.AccountID == 0 {
		acc.AccountID = len(wallet.Accounts) + 1
	}

	log.Println(acc)

	if _, err := wallet.findAccountID(acc.AccountID); err == nil {
		return nil, errors.New("Account already exists, not possible to overwrite")
	}

	wallet.Accounts = append(wallet.Accounts, *acc)

	inserted, err := wallet.findAccountID(acc.AccountID)
	if err != nil {
		log.Println(acc.AccountID, inserted.AccountID)
		return nil, errors.New("Error after insert, not found")
	}

	return inserted, nil
}

// GetAccounts returns a array of entity.Account
// If doesn't exists any Account, return error
func (wallet *Wallet) GetAccounts() ([]entity.Account, error) {
	wallet.Lock()
	defer wallet.Unlock()

	if len(wallet.Accounts) == 0 {
		return nil, errors.New("No account avaliable")
	}

	return wallet.Accounts, nil
}

// UpdateAccount receive AccountID and values creditLimit and CreditWithdraw to add into Account
// return updated Account or nil and error if not possible
func (wallet *Wallet) UpdateAccount(AccountID int, creditLimit float64, creditWithdraw float64) (*entity.Account, error) {
	wallet.Lock()
	defer wallet.Unlock()

	account, err := wallet.findAccountID(AccountID)
	if err != nil {
		return nil, errors.New("Account not found")
	}

	if account.CreditLimit+creditLimit < 0 || account.WithdrawlLimit+creditWithdraw < 0 {
		return nil, errors.New("New values of CreditLimit or WithdrawLimit cannot be negative")
	}

	account.CreditLimit += creditLimit
	account.WithdrawlLimit += creditWithdraw

	return account, nil
}
