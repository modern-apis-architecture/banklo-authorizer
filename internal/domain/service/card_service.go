package service

import (
	"encoding/json"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/config"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain"
	"io"
	"io/ioutil"
	"net/http"
)

type CardService struct {
	hcli *http.Client
	cfg  *config.CardService
}

func (cs *CardService) CardById(cid string) (*domain.Card, error) {
	req, err := http.NewRequest(http.MethodGet, cs.cfg.Url+"/cards/"+cid, nil)
	if err != nil {
		return nil, err
	}
	res, err := cs.hcli.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}
	c := &domain.Card{}
	jsonErr := json.Unmarshal(body, c)
	if jsonErr != nil {
		return nil, err
	}
	return c, nil
}
