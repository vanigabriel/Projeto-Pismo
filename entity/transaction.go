package entity

// Transaction definition struct
type Transaction struct {
	TransactionId int
	AccountId     int     `json:"account_id" valid:"required"`
	OperationId   int     `json:"operation_type_id" valid:"required"`
	Amount        float64 `json:"amount" valid:"required"`
}
