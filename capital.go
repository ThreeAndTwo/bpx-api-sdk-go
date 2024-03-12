package bpx_api_sdk_go

import (
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/constants"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/types"
	"github.com/imroc/req/v3"
)

type capital struct {
	url        string
	apiKey     string
	privateKey *ed25519.PrivateKey
	client     *req.Client
}

func NewCapital(url string, apiKey string, privateKey *ed25519.PrivateKey, client *req.Client) *capital {
	return &capital{url: url, apiKey: apiKey, privateKey: privateKey, client: client}
}

func (c *capital) GetBalances() (map[string]types.TokenCapital, error) {
	signRes := Sign(c.privateKey, constants.InstructionBalanceQuery, nil, nil)
	headers := SetHeaders(c.apiKey, signRes)

	response, err := c.client.R().SetHeaders(headers).Get(fmt.Sprintf("%s/api/v1/capital", c.url))
	if err != nil {
		return nil, err
	}

	tokenCapitals := make(map[string]types.TokenCapital)
	result, err := dealResponse(&tokenCapitals, response, err)
	if err != nil {
		return nil, err
	}
	return result.(map[string]types.TokenCapital), nil
}

func (c *capital) GetDeposits() ([]types.Deposit, error) {
	signRes := Sign(c.privateKey, constants.InstructionDepositQueryAll, nil, nil)
	headers := SetHeaders(c.apiKey, signRes)
	response, err := c.client.R().SetHeaders(headers).Get(fmt.Sprintf("%s/wapi/v1/capital/deposits", c.url))
	if err != nil {
		return nil, err
	}
	var deposits []types.Deposit
	result, err := dealResponse(&deposits, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.Deposit), nil

}

func (c *capital) GetDepositAddress(chain constants.ChainTy) (*types.DepositAddress, error) {
	if chain != constants.ChainSolana && chain != constants.ChainEthereum &&
		chain != constants.ChainPolygon && chain != constants.ChainBitcoin {
		return nil, fmt.Errorf("[%s] chain unSupported", chain)
	}

	queryParams := make(map[string]interface{})
	queryParams["blockchain"] = chain

	signRes := Sign(c.privateKey, constants.InstructionDepositAddressQuery, queryParams, nil)
	headers := SetHeaders(c.apiKey, signRes)
	response, err := c.client.R().SetHeaders(headers).
		Get(fmt.Sprintf("%s/wapi/v1/capital/deposit/address?blockchain=%s", c.url, chain))
	if err != nil {
		return nil, err
	}
	result, err := dealResponse(new(types.DepositAddress), response, err)
	if err != nil {
		return nil, err
	}

	fmtRes := result.(types.DepositAddress)
	return &fmtRes, nil
}

func (c *capital) GetWithdrawals(limit, offset int64) ([]types.Withdrawal, error) {
	limit = formatLimit(limit)
	offset = formatOffset(offset)

	queryParams := make(map[string]interface{})
	queryParams["limit"] = fmt.Sprintf("%d", limit)
	queryParams["offset"] = fmt.Sprintf("%d", offset)
	signRes := Sign(c.privateKey, constants.InstructionWithdrawalQueryAll, queryParams, nil)
	headers := SetHeaders(c.apiKey, signRes)

	url := fmt.Sprintf("%s/wapi/v1/capital/withdrawals?limit=%d&offset=%d", c.url, limit, offset)
	response, err := c.client.R().SetHeaders(headers).Get(url)
	if err != nil {
		return nil, err
	}

	var withdrawals []types.Withdrawal
	result, err := dealResponse(&withdrawals, response, err)
	if err != nil {
		return nil, err
	}
	return result.([]types.Withdrawal), nil
}

func (c *capital) Withdrawal(params *types.WithdrawParams) (*types.WithdrawResult, error) {
	if params == nil {
		return nil, fmt.Errorf("withdraw params is null")
	}

	bodyParams, err := StructToMap(*params)
	if err != nil {
		return nil, err
	}

	signRes := Sign(c.privateKey, constants.InstructionWithdrawal, nil, bodyParams)

	headers := SetHeaders(c.apiKey, signRes)
	url := fmt.Sprintf("%s/wapi/v1/capital/withdrawals", c.url)

	_body, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, fmt.Errorf("marshal body params error: %s", err)
	}

	response, err := c.client.R().SetHeaders(headers).SetBody(_body).Post(url)
	if err != nil {
		return nil, err
	}
	result, err := dealResponse(new(types.WithdrawResult), response, err)
	if err != nil {
		return nil, err
	}
	fmtRes := result.(types.WithdrawResult)
	return &fmtRes, nil
}
