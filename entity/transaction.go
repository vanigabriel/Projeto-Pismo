package entity

// Transaction definition struct
type Transaction struct {
	TransactionID int
	AccountID     int     `json:"account_id" valid:"required"`
	OperationID   int     `json:"operation_type_id" valid:"required"`
	Amount        float64 `json:"amount" valid:"required"`
	Balance       float64
	EventDate     string
	DueDate       string
}
