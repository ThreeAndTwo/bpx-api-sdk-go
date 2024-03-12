package types

import "github.com/ThreeAndTwo/bpx-api-sdk-go/constants"

type Asset struct {
	Symbol string `json:"symbol"`
	Tokens []struct {
		Blockchain        string `json:"blockchain"`
		DepositEnabled    bool   `json:"depositEnabled"`
		MinimumDeposit    string `json:"minimumDeposit"`
		WithdrawEnabled   bool   `json:"withdrawEnabled"`
		MinimumWithdrawal string `json:"minimumWithdrawal"`
		MaximumWithdrawal string `json:"maximumWithdrawal"`
		WithdrawalFee     string `json:"withdrawalFee"`
	} `json:"tokens"`
}

type Markets struct {
	Symbol      string `json:"symbol"`
	BaseSymbol  string `json:"baseSymbol"`
	QuoteSymbol string `json:"quoteSymbol"`
	Filters     struct {
		Price struct {
			MinPrice string `json:"minPrice"`
			MaxPrice string `json:"maxPrice"`
			TickSize string `json:"tickSize"`
		} `json:"price"`
		Quantity struct {
			MinQuantity string `json:"minQuantity"`
			MaxQuantity string `json:"maxQuantity"`
			StepSize    string `json:"stepSize"`
		} `json:"quantity"`
		Leverage struct {
			MinLeverage string `json:"minLeverage"`
			MaxLeverage string `json:"maxLeverage"`
			StepSize    string `json:"stepSize"`
		} `json:"leverage"`
	} `json:"filters"`
}

type Ticker struct {
	Symbol             string `json:"symbol"`
	FirstPrice         string `json:"firstPrice"`
	LastPrice          string `json:"lastPrice"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	High               string `json:"high"`
	Low                string `json:"low"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	Trades             int64  `json:"trades"`
}

type DepthData struct {
	Asks         [][]string `json:"asks"`
	Bids         [][]string `json:"bids"`
	LastUpdateId string     `json:"lastUpdateId"`
}

type Kline struct {
	Start  string `json:"start"`
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Close  string `json:"close"`
	End    string `json:"end"`
	Volume string `json:"volume"`
	Trades string `json:"trades"`
}

type SystemStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type RecentTrade struct {
	Id            int64  `json:"id"`
	Price         string `json:"price"`
	Quantity      string `json:"quantity"`
	QuoteQuantity string `json:"quoteQuantity"`
	Timestamp     int64  `json:"timestamp"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
}

type TradeHistory struct {
	Id            int64  `json:"id"`
	Price         string `json:"price"`
	Quantity      string `json:"quantity"`
	QuoteQuantity string `json:"quoteQuantity"`
	Timestamp     int64  `json:"timestamp"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
}

type Signature struct {
	Base64 string
	Bytes  []byte
	Ts     string
}

type TokenCapital struct {
	Available string `json:"available"`
	Locked    string `json:"locked"`
	Staked    string `json:"staked"`
}

type Deposit struct {
	Id                      int64  `json:"id"`
	ToAddress               string `json:"toAddress"`
	FromAddress             string `json:"fromAddress"`
	ConfirmationBlockNumber int64  `json:"confirmationBlockNumber"`
	ProviderId              string `json:"providerId"`
	Source                  string `json:"source"`
	Status                  string `json:"status"`
	TransactionHash         string `json:"transactionHash"`
	SubaccountId            int64  `json:"subaccountId"`
	Symbol                  string `json:"symbol"`
	Quantity                string `json:"quantity"`
	CreatedAt               string `json:"createdAt"`
}

type DepositAddress struct {
	Address string `json:"address"`
}

type Withdrawal struct {
	Id              int64  `json:"id"`
	Blockchain      string `json:"blockchain"`
	ClientId        string `json:"clientId"`
	Identifier      string `json:"identifier"`
	Quantity        string `json:"quantity"`
	Fee             string `json:"fee"`
	Symbol          string `json:"symbol"`
	Status          string `json:"status"`
	SubaccountId    int64  `json:"subaccountId"`
	ToAddress       string `json:"toAddress"`
	TransactionHash string `json:"transactionHash"`
	CreatedAt       string `json:"createdAt"`
}

type WithdrawParams struct {
	Address        string            `json:"address"`
	Blockchain     constants.ChainTy `json:"blockchain"`
	ClientId       string            `json:"clientId"`
	Quantity       string            `json:"quantity"`
	Symbol         string            `json:"symbol"`
	TwoFactorToken string            `json:"twoFactorToken"`
}

type WithdrawResult struct {
	Id              int64  `json:"id"`
	Blockchain      string `json:"blockchain"`
	ClientId        string `json:"clientId"`
	Identifier      string `json:"identifier"`
	Quantity        string `json:"quantity"`
	Fee             string `json:"fee"`
	Symbol          string `json:"symbol"`
	Status          string `json:"status"`
	SubaccountId    int64  `json:"subaccountId"`
	ToAddress       string `json:"toAddress"`
	TransactionHash string `json:"transactionHash"`
	CreatedAt       string `json:"createdAt"`
}

type OrderHistoryParams struct {
	Symbol string
	Offset int64
	Limit  int64
}

type OrderHistoryResult struct {
	Id                  string `json:"id"`
	OrderType           string `json:"orderType"`
	Symbol              string `json:"symbol"`
	Side                string `json:"side"`
	Price               string `json:"price"`
	TriggerPrice        string `json:"triggerPrice"`
	Quantity            string `json:"quantity"`
	QuoteQuantity       string `json:"quoteQuantity"`
	TimeInForce         string `json:"timeInForce"`
	SelfTradePrevention string `json:"selfTradePrevention"`
	PostOnly            bool   `json:"postOnly"`
	Status              string `json:"status"`
}

type FillHistoryParams struct {
	OrderId string
	Symbol  string
	Limit   int64
	Offset  int64
}

type FillHistoryResult struct {
	TradeId   int64  `json:"tradeId"`
	OrderId   string `json:"orderId"`
	Symbol    string `json:"symbol"`
	Side      string `json:"side"`
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	Fee       string `json:"fee"`
	FeeSymbol string `json:"feeSymbol"`
	IsMaker   bool   `json:"isMaker"`
	Timestamp string `json:"timestamp"`
}

type OrderParams struct {
	ClientId int64  `json:"clientId"`
	OrderId  string `json:"orderId"`
	Symbol   string `json:"symbol"`
}

type OrderResult struct {
	OrderType             string `json:"orderType"`
	Id                    string `json:"id"`
	ClientId              int64  `json:"clientId"`
	Symbol                string `json:"symbol"`
	Side                  string `json:"side"`
	Quantity              string `json:"quantity"`
	ExecutedQuantity      string `json:"executedQuantity"`
	QuoteQuantity         string `json:"quoteQuantity"`
	ExecutedQuoteQuantity string `json:"executedQuoteQuantity"`
	TriggerPrice          string `json:"triggerPrice"`
	TimeInForce           string `json:"timeInForce"`
	SelfTradePrevention   string `json:"selfTradePrevention"`
	Status                string `json:"status"`
	CreatedAt             int64  `json:"createdAt"`
}

type ExecuteOrderParams struct {
	ClientId            int64                           `json:"clientId"`
	OrderType           constants.OrderTypeTy           `json:"orderType"`
	PostOnly            bool                            `json:"postOnly"`
	Price               string                          `json:"price"`
	Quantity            string                          `json:"quantity"`
	QuoteQuantity       string                          `json:"quoteQuantity"`
	SelfTradePrevention constants.SelfTradePreventionTy `json:"selfTradePrevention"`
	Side                constants.OrderSideTy           `json:"side"`
	Symbol              string                          `json:"symbol"`
	TimeInForce         constants.TimeInForceTy         `json:"timeInForce"`
	TriggerPrice        string                          `json:"triggerPrice"`
}

type ExecuteOrderResult struct {
	OrderType             constants.OrderTypeTy           `json:"orderType"`
	Id                    string                          `json:"id"`
	ClientId              int64                           `json:"clientId"`
	Symbol                string                          `json:"symbol"`
	Side                  constants.OrderSideTy           `json:"side"`
	Quantity              string                          `json:"quantity"`
	ExecutedQuantity      string                          `json:"executedQuantity"`
	QuoteQuantity         string                          `json:"quoteQuantity"`
	ExecutedQuoteQuantity string                          `json:"executedQuoteQuantity"`
	TriggerPrice          string                          `json:"triggerPrice"`
	TimeInForce           constants.TimeInForceTy         `json:"timeInForce"`
	SelfTradePrevention   constants.SelfTradePreventionTy `json:"selfTradePrevention"`
	Status                string                          `json:"status"`
	CreatedAt             int64                           `json:"createdAt"`
}
