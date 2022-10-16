package service

import (
	"context"
	"encoding/json"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/config"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

type CardService struct {
	hcli *http.Client
	cfg  *config.Config
}

func (cs *CardService) CardById(ctx context.Context, cid string) (*domain.Card, error) {
	req, err := http.NewRequest(http.MethodGet, cs.cfg.CardService.Url+"/cards/"+cid, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ctx.Value("external-auth").(string))
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
	log.Infof("card data %s", string(body))
	c := &domain.Card{}
	jsonErr := json.Unmarshal(body, c)
	if jsonErr != nil {
		return nil, err
	}
	return c, nil
}

func NewCardService(hcli *http.Client, cfg *config.Config) *CardService {
	return &CardService{
		hcli: hcli,
		cfg:  cfg,
	}
}
