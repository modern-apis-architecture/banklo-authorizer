package domain

type TransactionId struct {
	Id string `json:"id"`
}

func NewTransactionId(tid string) *TransactionId {
	return &TransactionId{Id: tid}
}
