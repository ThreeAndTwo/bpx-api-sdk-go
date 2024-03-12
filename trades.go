package bpx_api_sdk_go

import (
	"fmt"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/constants"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/types"
	"github.com/imroc/req/v3"
	"strings"
)

type trade struct {
	url    string
	client *req.Client
}

func NewTrade(url string, client *req.Client) *trade {
	return &trade{url: url, client: client}
}

func (t *trade) GetRecentTrades(symbol string, limit int64) ([]types.RecentTrade, error) {
	if symbol == "" {
		return nil, constants.ErrSymbolIsNull
	}

	url := fmt.Sprintf("%s/api/v1/trades?symbol=%s", t.url, strings.ToUpper(symbol))
	limit = formatLimit(limit)
	url += fmt.Sprintf("&limit=%d", limit)
	response, err := t.client.R().Get(url)
	if err != nil {
		return nil, err
	}

	var data []types.RecentTrade
	result, err := dealResponse(&data, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.RecentTrade), nil
}

func (t *trade) GetHistoricalTrades(symbol string, limit, offset int64) ([]types.TradeHistory, error) {
	if symbol == "" {
		return nil, constants.ErrSymbolIsNull
	}

	url := fmt.Sprintf("%s/api/v1/trades/history?symbol=%s", t.url, strings.ToUpper(symbol))
	limit = formatLimit(limit)
	url += fmt.Sprintf("&limit=%d", limit)
	offset = formatOffset(offset)
	url += fmt.Sprintf("&offset=%d", offset)

	response, err := t.client.R().Get(url)
	if err != nil {
		return nil, err
	}

	var data []types.TradeHistory
	result, err := dealResponse(&data, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.TradeHistory), nil
}
