package service

import (
	"github.com/modern-apis-architecture/banklo-authorizer/internal/api"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (ts *TransactionService) Confirmation(t *api.RequestTransaction) {

}

func (ts *TransactionService) Reversal(t *api.RequestTransaction) {

}

func (ts *TransactionService) Cancellation(t *api.RequestCancellation) {

}
