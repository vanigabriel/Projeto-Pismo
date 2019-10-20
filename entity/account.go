package entity

// Account definition struct
type Account struct {
	AccountID       int
	CreditLimit     float64 `json:"available_credit_limit" valid:"required"`
	WithdrawlLimit  float64 `json:"available_withdrawal_limit" valid:"required"`
	CreditAvaliable float64
}
