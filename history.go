package bpx_api_sdk_go

import (
	"crypto/ed25519"
	"fmt"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/constants"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/types"
	"github.com/imroc/req/v3"
	"strings"
)

type history struct {
	url        string
	apiKey     string
	privateKey *ed25519.PrivateKey
	client     *req.Client
}

func NewHistory(url string, apiKey string, privateKey *ed25519.PrivateKey, client *req.Client) *history {
	return &history{url: url, apiKey: apiKey, privateKey: privateKey, client: client}
}

func (h *history) GetOrderHistory(params *types.OrderHistoryParams) ([]types.OrderHistoryResult, error) {
	url := fmt.Sprintf("%s/wapi/v1/history/orders?limit=%d&offset=%d", h.url, params.Limit, params.Offset)
	queryMap := make(map[string]interface{})
	if params == nil {
		return nil, fmt.Errorf("order history params is null")
	}

	params.Symbol = strings.ToUpper(params.Symbol)
	if params.Symbol != "" {
		url += fmt.Sprintf("&symbol=%s", params.Symbol)
		queryMap["symbol"] = params.Symbol
	}
	queryMap["limit"] = params.Limit
	queryMap["offset"] = params.Offset

	signRes := Sign(h.privateKey, constants.InstructionOrderHistoryQueryAll, queryMap, nil)
	headers := SetHeaders(h.apiKey, signRes)

	response, err := h.client.R().SetHeaders(headers).Get(url)
	if err != nil {
		return nil, err
	}

	var ohr []types.OrderHistoryResult
	result, err := dealResponse(&ohr, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.OrderHistoryResult), nil
}

func (h *history) GetFills(params *types.FillHistoryParams) ([]types.FillHistoryResult, error) {
	url := fmt.Sprintf("%s/wapi/v1/history/fills?limit=%d&offset=%d", h.url, params.Limit, params.Offset)

	queryMap := make(map[string]interface{})
	if params == nil {
		return nil, fmt.Errorf("order history params is null")
	}

	params.Symbol = strings.ToUpper(params.Symbol)
	if params.Symbol != "" {
		url += fmt.Sprintf("&symbol=%s", params.Symbol)
		queryMap["symbol"] = params.Symbol
	}

	if params.OrderId != "" {
		url += fmt.Sprintf("&orderId=%s", params.OrderId)
		queryMap["orderId"] = params.OrderId
	}

	queryMap["limit"] = params.Limit
	queryMap["offset"] = params.Offset
	signRes := Sign(h.privateKey, constants.InstructionFillHistoryQueryAll, queryMap, nil)
	headers := SetHeaders(h.apiKey, signRes)

	response, err := h.client.R().SetHeaders(headers).Get(url)
	if err != nil {
		return nil, err
	}

	var fhr []types.FillHistoryResult
	result, err := dealResponse(&fhr, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.FillHistoryResult), nil
}
