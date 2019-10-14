package entity

// Account definition struct
type Account struct {
	AccountId      int
	CreditLimit    float64 `json:"available_credit_limit" valid:"required"`
	WithdrawlLimit float64 `json:"available_credit_limit" valid:"required"`
}
