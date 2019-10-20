package entity

// Account definition struct
type Account struct {
	AccountID       int     `json:"account_id"`
	CreditLimit     float64 `json:"available_credit_limit" valid:"optional"`
	WithdrawlLimit  float64 `json:"available_withdrawal_limit" valid:"optional"`
	CreditAvaliable float64 `json:"-"`
}
