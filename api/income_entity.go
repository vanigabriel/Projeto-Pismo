package api

// Account defines income data type
type Account struct {
	CreditLimit    Amount `json:"available_credit_limit" valid:"optional"`
	WithdrawlLimit Amount `json:"available_withdrawal_limit" valid:"optional"`
}

// Amount defines income data type
type Amount struct {
	Value float64 `json:"amount"`
}
