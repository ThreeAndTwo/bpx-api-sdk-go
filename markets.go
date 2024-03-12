package bpx_api_sdk_go

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/constants"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/types"
	"github.com/imroc/req/v3"
	"reflect"
	"strings"
)

type markets struct {
	url    string
	client *req.Client
}

func NewMarkets(url string, client *req.Client) *markets {
	return &markets{
		url:    url,
		client: client,
	}
}

func (m *markets) GetAssets() ([]types.Asset, error) {
	response, err := m.client.R().Get(fmt.Sprintf("%s/api/v1/assets", m.url))
	if err != nil {
		return nil, err
	}

	var _assets []types.Asset
	result, err := dealResponse(&_assets, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.Asset), nil
}

func (m *markets) GetMarkets() ([]types.Markets, error) {
	response, err := m.client.R().Get(fmt.Sprintf("%s/api/v1/markets", m.url))
	if err != nil {
		return nil, err
	}

	var _markets []types.Markets
	result, err := dealResponse(&_markets, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.Markets), nil
}

func (m *markets) GetTicker(symbol string) (*types.Ticker, error) {
	if symbol == "" {
		return nil, constants.ErrSymbolIsNull
	}

	response, err := m.client.R().Get(fmt.Sprintf("%s/api/v1/ticker?symbol=%s", m.url, strings.ToUpper(symbol)))
	if err != nil {
		return nil, err
	}
	result, err := dealResponse(new(types.Ticker), response, err)
	if err != nil {
		return nil, err
	}
	fmtRes := result.(types.Ticker)
	return &fmtRes, nil
}

func (m *markets) GetTickers() ([]types.Ticker, error) {
	response, err := m.client.R().Get(fmt.Sprintf("%s/api/v1/tickers", m.url))
	if err != nil {
		return nil, err
	}
	var tickers []types.Ticker
	result, err := dealResponse(&tickers, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.Ticker), nil
}

func (m *markets) GetDepth(symbol string) (*types.DepthData, error) {
	response, err := m.client.R().Get(fmt.Sprintf("%s/api/v1/depth?symbol=%s", m.url, strings.ToUpper(symbol)))
	if err != nil {
		return nil, err
	}
	result, err := dealResponse(new(types.DepthData), response, err)
	if err != nil {
		return nil, err
	}
	fmtRes := result.(types.DepthData)
	return &fmtRes, nil
}

func (m *markets) GetKline(symbol string, interval constants.IntervalTy, startTime, endTime int64) ([]types.Kline, error) {
	if symbol == "" {
		return nil, constants.ErrSymbolIsNull
	}

	if interval != constants.Interval1m && interval != constants.Interval3m && interval != constants.Interval5m &&
		interval != constants.Interval15m && interval != constants.Interval30m && interval != constants.Interval1h &&
		interval != constants.Interval2h && interval != constants.Interval4h && interval != constants.Interval6h &&
		interval != constants.Interval8h && interval != constants.Interval12h && interval != constants.Interval1d &&
		interval != constants.Interval3d && interval != constants.Interval1w && interval != constants.Interval1month {
		return nil, fmt.Errorf("[%s] interval is unSupported", interval)
	}

	if endTime < startTime {
		return nil, fmt.Errorf("endTime should ge startTime")
	}

	url := fmt.Sprintf("%s/api/v1/klines?symbol=%s&interval=%s", m.url, strings.ToUpper(symbol), interval)
	if startTime != 0 {
		url += fmt.Sprintf("&startTime=%d", startTime)
	}
	if endTime != 0 {
		url += fmt.Sprintf("&endTime=%d", endTime)
	}

	response, err := m.client.R().Get(url)
	if err != nil {
		return nil, err
	}
	var depthData []types.Kline
	result, err := dealResponse(&depthData, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.Kline), nil
}

func dealResponse(data interface{}, r *req.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("error StatusCode %d, please check request host", r.StatusCode)
	}
	err = json.Unmarshal(r.Bytes(), data)
	if err != nil {
		return nil, err
	}

	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Ptr:
		val = val.Elem() // Dereference the pointer
	default:
		return nil, fmt.Errorf("expected a pointer, got %T", data)
	}

	switch val.Kind() {
	case reflect.Slice:
		return handleSlice(val)
	case reflect.Array:
		return handleArray(val)
	case reflect.Struct:
		return handleStruct(val)
	case reflect.Map:
		return handleMap(val)
	default:
		return nil, fmt.Errorf("unsupported type: %s", val.Kind().String())
	}
}

func handleSlice(val reflect.Value) (interface{}, error) {
	newSlice := reflect.MakeSlice(val.Type(), val.Len(), val.Len())
	reflect.Copy(newSlice, val)
	return newSlice.Interface(), nil
}

func handleArray(val reflect.Value) (interface{}, error) {
	newSlice := reflect.New(reflect.ArrayOf(val.Len(), val.Type().Elem())).Elem()
	newSlice.Set(val)
	return newSlice.Interface(), nil
}

func handleStruct(val reflect.Value) (interface{}, error) {
	newStruct := reflect.New(val.Type()).Elem()
	newStruct.Set(val)
	return newStruct.Interface(), nil
}

func handleMap(val reflect.Value) (interface{}, error) {
	newMap := reflect.MakeMap(val.Type())
	for _, key := range val.MapKeys() {
		newMap.SetMapIndex(key, val.MapIndex(key))
	}
	return newMap.Interface(), nil
}
