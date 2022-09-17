package service

import "github.com/modern-apis-architecture/banklo-authorizer/internal/domain"

type ExternalAuthorization struct {
	extAuth ClientExternalAuthorization
}

func NewExternalAuthorization(extAuth ClientExternalAuthorization) *ExternalAuthorization {
	return &ExternalAuthorization{extAuth: extAuth}
}

func (ea *ExternalAuthorization) Authorize(t *domain.Transaction) (*domain.ExternalTransactionId, error) {
	nt, err := ea.extAuth.Authorize(t)
	if err != nil {
		return nil, err
	}
	return &domain.ExternalTransactionId{
		Id: nt.Id,
	}, err
}

func (ea *ExternalAuthorization) Cancellation(t *domain.Transaction) (*domain.ExternalTransactionId, error) {
	nt, err := ea.extAuth.Cancellation(t)
	if err != nil {
		return nil, err
	}
	return &domain.ExternalTransactionId{Id: nt.Id}, err

}

func (ea *ExternalAuthorization) Reversal(t *domain.Transaction) (*domain.ExternalTransactionId, error) {
	nt, err := ea.extAuth.Reversal(t)
	if err != nil {
		return nil, err
	}
	return &domain.ExternalTransactionId{Id: nt.Id}, err
}
