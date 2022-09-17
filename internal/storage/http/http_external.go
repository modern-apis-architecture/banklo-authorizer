package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/config"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain"
	"io"
	"io/ioutil"
	"net/http"
)

var ErrorDeclined = errors.New("transaction declined")

type HttpExternalAuthorization struct {
	cli *http.Client
	cfg *config.ExternalAuthorization
}

func NewHttpExternalAuthorization(cli *http.Client, cfg *config.ExternalAuthorization) *HttpExternalAuthorization {
	return &HttpExternalAuthorization{cli: cli, cfg: cfg}
}

func (hea *HttpExternalAuthorization) Authorize(t *domain.Transaction) (*domain.ExternalTransactionResponse, error) {
	req, _ := hea.createRequest("/transactions", t)
	resp, err := hea.cli.Do(req)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
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

func (hea *HttpExternalAuthorization) Cancellation(t *domain.Transaction) (*domain.ExternalTransactionResponse, error) {
	req, _ := hea.createRequest("/transactions"+t.Id+"/cancellation", t)
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

func (hea *HttpExternalAuthorization) Reversal(t *domain.Transaction) (*domain.ExternalTransactionResponse, error) {
	req, _ := hea.createRequest("/transactions"+t.Id+"/reversal", t)
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

func (hea *HttpExternalAuthorization) createRequest(path string, t *domain.Transaction) (*http.Request, error) {
	body, _ := json.Marshal(t)
	return http.NewRequest(http.MethodPost, hea.cfg.Url+path, bytes.NewBuffer(body))
}
