package bpx_api_sdk_go

import (
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/constants"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/types"
	"github.com/imroc/req/v3"
	"strings"
)

type order struct {
	url        string
	apiKey     string
	privateKey *ed25519.PrivateKey
	client     *req.Client
}

func NewOrder(url string, apiKey string, privateKey *ed25519.PrivateKey, client *req.Client) *order {
	return &order{url: url, apiKey: apiKey, privateKey: privateKey, client: client}
}

func (o *order) GetOrder(params *types.OrderParams) (*types.OrderResult, error) {
	if params.Symbol == "" {
		return nil, constants.ErrSymbolIsNull
	}

	if params.OrderId == "" && params.ClientId <= 0 {
		return nil, fmt.Errorf("must specify either `client_id` or `order_id`")
	}
	params.Symbol = strings.ToUpper(params.Symbol)
	queryMap := make(map[string]interface{})
	url := fmt.Sprintf("%s/api/v1/order?symbol=%s", o.url, params.Symbol)
	if params.ClientId > 0 {
		url += fmt.Sprintf("&clientId=%d", params.ClientId)
		queryMap["clientId"] = params.ClientId
	}

	if params.OrderId != "" {
		url += fmt.Sprintf("&orderId=%s", params.OrderId)
		queryMap["orderId"] = params.OrderId
	}

	queryMap["symbol"] = params.Symbol
	signRes := Sign(o.privateKey, constants.InstructionOrderQuery, queryMap, nil)
	headers := SetHeaders(o.apiKey, signRes)

	response, err := o.client.R().SetHeaders(headers).Get(url)
	if err != nil {
		return nil, err
	}
	result, err := dealResponse(new(types.OrderResult), response, err)
	if err != nil {
		return nil, err
	}

	fmtRes := result.(types.OrderResult)
	return &fmtRes, nil
}

func (o *order) ExecuteOrder(params *types.ExecuteOrderParams) (*types.ExecuteOrderResult, error) {
	if params.OrderType != constants.OrderMarket && params.OrderType != constants.OrderLimit {
		return nil, fmt.Errorf(
			"orderType %s MUST BE %s or %s",
			params.OrderType,
			constants.OrderMarket,
			constants.OrderLimit,
		)
	}

	if params.Side != constants.SideBid && params.Side != constants.SideAsk {
		return nil, fmt.Errorf(
			"order side %s MUST BE %s or %s",
			params.Side,
			constants.SideBid,
			constants.SideAsk,
		)
	}

	if params.Symbol == "" {
		return nil, constants.ErrSymbolIsNull
	}
	params.Symbol = strings.ToUpper(params.Symbol)

	bodyMap := make(map[string]interface{})
	bodyMap["orderType"] = params.OrderType
	bodyMap["side"] = params.Side
	bodyMap["symbol"] = params.Symbol

	// optional
	if params.ClientId > 0 {
		bodyMap["clientId"] = params.ClientId
	}
	bodyMap["postOnly"] = true

	if params.Price != "" {
		bodyMap["price"] = params.Price
	}

	if params.Quantity != "" {
		bodyMap["quantity"] = params.Quantity
	}

	if params.QuoteQuantity != "" {
		bodyMap["quoteQuantity"] = params.QuoteQuantity
	}

	if params.SelfTradePrevention != "" {
		bodyMap["selfTradePrevention"] = params.SelfTradePrevention
	}

	if params.TimeInForce != "" {
		bodyMap["timeInForce"] = params.TimeInForce
	}

	if params.TriggerPrice != "" {
		bodyMap["triggerPrice"] = params.TriggerPrice
	}

	signRes := Sign(o.privateKey, constants.InstructionOrderExecute, nil, bodyMap)
	headers := SetHeaders(o.apiKey, signRes)

	_body, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, fmt.Errorf("marshal body params error: %s", err)
	}

	response, err := o.client.R().SetHeaders(headers).SetBody(_body).Post(fmt.Sprintf("%s/api/v1/order", o.url))
	if err != nil {
		return nil, err
	}
	result, err := dealResponse(new(types.ExecuteOrderResult), response, err)
	if err != nil {
		return nil, err
	}

	fmtRes := result.(types.ExecuteOrderResult)
	return &fmtRes, nil
}

func (o *order) CancelOrder(params *types.OrderParams) (*types.OrderResult, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol is null")
	}
	params.Symbol = strings.ToUpper(params.Symbol)

	bodyMap := make(map[string]interface{})
	bodyMap["symbol"] = params.Symbol

	if params.ClientId > 0 {
		bodyMap["clientId"] = params.ClientId
	}

	if params.OrderId != "" {
		bodyMap["orderId"] = params.OrderId
	}

	signRes := Sign(o.privateKey, constants.InstructionOrderCancel, nil, bodyMap)
	headers := SetHeaders(o.apiKey, signRes)

	_body, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, fmt.Errorf("marshal body params error: %s", err)
	}

	response, err := o.client.R().SetHeaders(headers).SetBody(_body).Delete(fmt.Sprintf("%s/api/v1/order", o.url))
	if err != nil {
		return nil, err
	}
	result, err := dealResponse(new(types.OrderResult), response, err)
	if err != nil {
		return nil, err
	}
	fmtRes := result.(types.OrderResult)
	return &fmtRes, nil
}

func (o *order) GetOrders(symbol string) ([]types.OrderResult, error) {
	url := fmt.Sprintf("%s/api/v1/orders", o.url)
	symbol = strings.ToUpper(symbol)
	queryMap := make(map[string]interface{})
	if symbol != "" {
		url += fmt.Sprintf("?symbol=%s", symbol)
		queryMap["symbol"] = symbol
	}

	signRes := Sign(o.privateKey, constants.InstructionOrderQueryAll, queryMap, nil)
	headers := SetHeaders(o.apiKey, signRes)
	response, err := o.client.R().SetHeaders(headers).Get(url)
	if err != nil {
		return nil, err
	}

	var orders []types.OrderResult
	result, err := dealResponse(&orders, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.OrderResult), nil
}

func (o *order) CancelOrders(symbol string) ([]types.OrderResult, error) {
	if symbol == "" {
		return nil, constants.ErrSymbolIsNull
	}
	symbol = strings.ToUpper(symbol)

	bodyMap := make(map[string]interface{})
	bodyMap["symbol"] = symbol

	signRes := Sign(o.privateKey, constants.InstructionOrderCancelAll, nil, bodyMap)
	headers := SetHeaders(o.apiKey, signRes)

	_body, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, fmt.Errorf("marshal body params error: %s", err)
	}

	response, err := o.client.R().SetHeaders(headers).SetBody(_body).Delete(fmt.Sprintf("%s/api/v1/orders", o.url))
	if err != nil {
		return nil, err
	}
	var orders []types.OrderResult
	result, err := dealResponse(&orders, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.OrderResult), nil
}
