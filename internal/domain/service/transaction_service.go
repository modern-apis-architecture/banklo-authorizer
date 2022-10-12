package service

import (
	"context"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/api"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
	ext  *ExternalAuthorization
	cs   *CardService
}

func NewTransactionService(repo repository.TransactionRepository, ext *ExternalAuthorization, cs *CardService) *TransactionService {
	return &TransactionService{repo: repo, ext: ext, cs: cs}
}

func (ts *TransactionService) Confirmation(ctx context.Context, t *api.RequestTransaction) (*domain.TransactionId, error) {
	card, err := ts.cs.CardById(ctx,t.TransactionData.CardId)
	if err != nil {
		return nil, err
	}
	dt := &domain.Transaction{
		AuthorizationCode: t.AuthorizationCode,
		AcquirerCode:      t.AcquirerCode,
		MerchantCode:      t.MerchantCode,
		CurrencyCode:      t.CurrencyCode,
		CountryCode:       t.CountryCode,
		ProductId:         t.ProductId,
		PosId:             t.PosId,
		WithPassword:      t.TransactionData.WithPassword,
		Type:              string(t.TransactionData.Operation),
		Amount:            t.TransactionData.Amount,
		CardId:            card.Id,
		Id:                t.TransactionData.TransactionId,
	}
	exTid, err := ts.ext.Authorize(ctx,dt)
	if err != nil {
		return nil, err
	}
	dt.ExternalTid = exTid.Id
	tid, err := ts.repo.Store(dt)
	if err != nil {
		return nil, err
	}
	return tid, nil
}

func (ts *TransactionService) Reversal(ctx context.Context,t *api.RequestTransaction) (*domain.TransactionId, error) {
	card, err := ts.cs.CardById(nil, t.TransactionData.CardId)
	if err != nil {
		return nil, err
	}
	dt := &domain.Transaction{
		AuthorizationCode: t.AuthorizationCode,
		AcquirerCode:      t.AcquirerCode,
		MerchantCode:      t.MerchantCode,
		CurrencyCode:      t.CurrencyCode,
		CountryCode:       t.CountryCode,
		ProductId:         t.ProductId,
		PosId:             t.PosId,
		WithPassword:      t.TransactionData.WithPassword,
		Type:              string(t.TransactionData.Operation),
		Amount:            t.TransactionData.Amount,
		CardId:            card.Id,
	}
	exTid, err := ts.ext.Reversal(ctx,dt)
	if err != nil {
		return nil, err
	}
	dt.ExternalTid = exTid.Id
	tid, err := ts.repo.Store(dt)
	if err != nil {
		return nil, err
	}
	return tid, nil
}

func (ts *TransactionService) Cancellation(ctx context.Context,t *api.RequestCancellation) (*domain.TransactionId, error) {
	card, err := ts.cs.CardById(nil, t.TransactionData.CardId)
	if err != nil {
		return nil, err
	}
	dt := &domain.Transaction{
		AuthorizationCode: t.AuthorizationCode,
		AcquirerCode:      t.AcquirerCode,
		MerchantCode:      t.MerchantCode,
		CurrencyCode:      t.CurrencyCode,
		CountryCode:       t.CountryCode,
		ProductId:         t.ProductId,
		PosId:             t.PosId,
		WithPassword:      t.TransactionData.WithPassword,
		Type:              string(t.TransactionData.Operation),
		Amount:            t.TransactionData.Amount,
		CardId:            card.Id,
	}
	exTid, err := ts.ext.Cancellation(ctx,dt)
	if err != nil {
		return nil, err
	}
	dt.ExternalTid = exTid.Id
	tid, err := ts.repo.Store(dt)
	if err != nil {
		return nil, err
	}
	return tid, nil
}
