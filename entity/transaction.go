package entity

// Transaction definition struct
type Transaction struct {
	TransactionID int     `json:"transaction_id"`
	AccountID     int     `json:"account_id" valid:"required"`
	OperationID   int     `json:"operation_type_id" valid:"required"`
	Amount        float64 `json:"amount" valid:"required"`
	Balance       float64 `json:"balance"`
	EventDate     string  `json:"event_date"`
	DueDate       string
}
