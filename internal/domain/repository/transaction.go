package repository

import "github.com/modern-apis-architecture/banklo-authorizer/internal/domain"

type TransactionRepository interface {
	Store(transaction *domain.Transaction) (*domain.TransactionId, error)
}
