package bpx_api_sdk_go

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/constants"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/iface"
	"github.com/imroc/req/v3"
)

type BackpackExchange struct {
	url        string
	apiKey     string
	apiSecret  string
	privateKey *ed25519.PrivateKey
	client     *req.Client
}

func NewBackpack(url, apiKey, apiSecret string, isDebug bool) (*BackpackExchange, error) {
	bpx := &BackpackExchange{
		url:       url,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}

	if err := bpx.check(); err != nil {
		return nil, err
	}

	if err := bpx.decodeApiSecret(); err != nil {
		return nil, err
	}
	bpx.client = req.C().SetCommonHeader("Content-Type", "application/json; charset=utf-8")
	if isDebug {
		bpx.client.DevMode()
	}
	return bpx, nil
}

func (bpx *BackpackExchange) check() error {
	if bpx.url == "" {
		return constants.ErrURLIsNull
	}

	if bpx.apiKey == "" {
		return constants.ErrApiKeyIsNull
	}

	if bpx.apiSecret == "" {
		return constants.ErrApiSecretIsNull
	}
	return nil
}

func (bpx *BackpackExchange) decodeApiSecret() error {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(bpx.apiSecret)
	if err != nil {
		return fmt.Errorf("decode base64 error: %s", err)
	}

	if len(privateKeyBytes) != ed25519.SeedSize {
		fmt.Println()
		return fmt.Errorf("unexpected private key length: %d", len(privateKeyBytes))
	}
	privateKey := ed25519.NewKeyFromSeed(privateKeyBytes)
	bpx.privateKey = &privateKey
	return nil
}

func (bpx *BackpackExchange) GetURL() string {
	return bpx.url
}

func (bpx *BackpackExchange) SetURL(url string) error {
	if url == "" {
		return constants.ErrURLIsNull
	}
	bpx.url = url
	return nil
}

func (bpx *BackpackExchange) GetApiKey() string {
	return bpx.apiKey
}

func (bpx *BackpackExchange) SetApiKey(apiKey string) error {
	if apiKey == "" {
		return constants.ErrApiKeyIsNull
	}
	bpx.apiKey = apiKey
	return nil
}

func (bpx *BackpackExchange) GetApiSecret() string {
	return bpx.apiSecret
}

func (bpx *BackpackExchange) SetApiSecret(apiSecret string) error {
	bpx.apiSecret = apiSecret
	return bpx.decodeApiSecret()
}

func (bpx *BackpackExchange) Markets() iface.IMarket {
	return NewMarkets(bpx.url, bpx.client)
}

func (bpx *BackpackExchange) System() iface.ISystem {
	return NewSystem(bpx.url, bpx.client)
}

func (bpx *BackpackExchange) Trades() iface.ITrades {
	return NewTrade(bpx.url, bpx.client)
}

func (bpx *BackpackExchange) Capital() iface.ICapital {
	return NewCapital(bpx.url, bpx.apiKey, bpx.privateKey, bpx.client)
}

func (bpx *BackpackExchange) History() iface.IHistory {
	return NewHistory(bpx.url, bpx.apiKey, bpx.privateKey, bpx.client)
}

func (bpx *BackpackExchange) Order() iface.IOrder {
	return NewOrder(bpx.url, bpx.apiKey, bpx.privateKey, bpx.client)
}
