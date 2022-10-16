package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/config"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

var ErrorDeclined = errors.New("transaction declined")

type HttpExternalAuthorization struct {
	cli *http.Client
	cfg *config.Config
}

func NewHttpExternalAuthorization(cli *http.Client, cfg *config.Config) *HttpExternalAuthorization {
	return &HttpExternalAuthorization{cli: cli, cfg: cfg}
}

func (hea *HttpExternalAuthorization) Authorize(ctx context.Context, t *domain.Transaction) (*domain.ExternalTransactionResponse, error) {
	et := &ExternalRequestTransaction{
		AcquirerCode:      t.AcquirerCode,
		AuthorizationCode: t.AuthorizationCode,
		CountryCode:       t.CountryCode,
		CurrencyCode:      t.CurrencyCode,
		MerchantCode:      t.MerchantCode,
		ExternalTransactionData: ExternalTransactionData{
			Amount:          t.Amount,
			CardId:          t.CardId,
			TransactionId:   t.Id,
			TransactionType: t.Type,
			WithPassword:    false,
		},
	}
	req, _ := hea.createRequest(ctx, "/transactions", et)
	resp, err := hea.cli.Do(req)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if err != nil {
		log.Errorf(" register transaction error %v", err)
		return nil, err
	}
	if resp.StatusCode != 201 {
		log.Errorf(" register transaction status code %v", resp.StatusCode)
		return nil, ErrorDeclined
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	etr := &domain.ExternalTransactionResponse{}
	jsonErr := json.Unmarshal(body, etr)
	if jsonErr != nil {
		log.Errorf(" error to decode register transaction  %v", jsonErr)
		return nil, err
	}
	return etr, nil
}

func (hea *HttpExternalAuthorization) Cancellation(ctx context.Context, t *domain.Transaction) (*domain.ExternalTransactionResponse, error) {
	req, _ := hea.createRequest(ctx, "/transactions"+t.Id+"/cancellation", t)
	resp, err := hea.cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, ErrorDeclined
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	etr := &domain.ExternalTransactionResponse{}
	jsonErr := json.Unmarshal(body, etr)
	if jsonErr != nil {
		return nil, err
	}
	return etr, nil
}

func (hea *HttpExternalAuthorization) Reversal(ctx context.Context, t *domain.Transaction) (*domain.ExternalTransactionResponse, error) {
	req, _ := hea.createRequest(ctx, "/transactions"+t.Id+"/reversal", t)
	resp, err := hea.cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, ErrorDeclined
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	etr := &domain.ExternalTransactionResponse{}
	jsonErr := json.Unmarshal(body, etr)
	if jsonErr != nil {
		return nil, err
	}
	return etr, nil
}

func (hea *HttpExternalAuthorization) createRequest(ctx context.Context, path string, t interface{}) (*http.Request, error) {
	body, _ := json.Marshal(t)
	log.Info(string(body))
	req, err := http.NewRequest(http.MethodPost, hea.cfg.ExternalAuthorization.Url+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ctx.Value("external-auth").(string))
	return req, nil
}
