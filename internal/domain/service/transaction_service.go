package service

import (
	"github.com/modern-apis-architecture/banklo-authorizer/internal/api"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
	ext  *ExternalAuthorization
}

func NewTransactionService(repo repository.TransactionRepository, ext *ExternalAuthorization) *TransactionService {
	return &TransactionService{repo: repo, ext: ext}
}

func (ts *TransactionService) Confirmation(t *api.RequestTransaction) (*domain.TransactionId, error) {
	return nil, nil
}

func (ts *TransactionService) Reversal(t *api.RequestTransaction) (*domain.TransactionId, error) {
	return nil, nil
}

func (ts *TransactionService) Cancellation(t *api.RequestCancellation) (*domain.TransactionId, error) {
	return nil, nil
}
