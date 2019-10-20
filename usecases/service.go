package usecases

import (
	"github.com/asaskevich/govalidator"
	"github.com/vanigabriel/Projeto-Pismo/entity"
)

// Service defines a service
type Service struct {
	repo Repository
}

// NewService gets a Respository type and returns a Service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// AddAccount receive an account and try to add in wallet
// If Account already exists, returns error
func (s *Service) AddAccount(acc *entity.Account) (*entity.Account, error) {
	_, err := govalidator.ValidateStruct(acc)
	if err != nil {
		return nil, err
	}

	return s.repo.AddAccount(acc)
}

// UpdateAccount receive AccountID and values creditLimit and CreditWithdraw to add into Account
// return updated Account or nil and error if not possible
func (s *Service) UpdateAccount(accountID int, creditLimit float64, creditWithdraw float64) (*entity.Account, error) {
	return s.repo.UpdateAccount(accountID, creditLimit, creditWithdraw)
}

// GetAccounts returns a array of entity.Account
// If doesn't exists any Account, return error
func (s *Service) GetAccounts() ([]entity.Account, error) {
	return s.repo.GetAccounts()
}

// InsertTransaction insert a transaction in Statements
// Return error if anything not right occours
func (s *Service) InsertTransaction(transaction *entity.Transaction) error {
	if transaction.OperationID == 0 {
		transaction.OperationID = 4
	}
	_, err := govalidator.ValidateStruct(transaction)
	if err != nil {
		return err
	}

	return s.repo.InsertTransaction(transaction)
}
