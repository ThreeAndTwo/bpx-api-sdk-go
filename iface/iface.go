package iface

import (
	"github.com/ThreeAndTwo/bpx-api-sdk-go/constants"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/types"
)

type (
	IClient interface {
		Markets() IMarket
		System() ISystem
		Trades() ITrades
		Capital() ICapital
		History() IHistory
		Order() IOrder
	}

	IMarket interface {
		GetAssets() ([]types.Asset, error)
		GetMarkets() ([]types.Markets, error)
		GetTicker(symbol string) (*types.Ticker, error)
		GetTickers() ([]types.Ticker, error)
		GetDepth(symbol string) (*types.DepthData, error)
		GetKline(symbol string, interval constants.IntervalTy, startTime, endTime int64) ([]types.Kline, error)
	}

	ISystem interface {
		GetStatus() (*types.SystemStatus, error)
		Ping() error
		GetTime() error
	}

	ITrades interface {
		GetRecentTrades(symbol string, limit int64) ([]types.RecentTrade, error)
		GetHistoricalTrades(symbol string, limit, offset int64) ([]types.TradeHistory, error)
	}

	ICapital interface {
		GetBalances() (map[string]types.TokenCapital, error)
		GetDeposits() ([]types.Deposit, error)
		GetDepositAddress(chain constants.ChainTy) (*types.DepositAddress, error)
		GetWithdrawals(limit, offset int64) ([]types.Withdrawal, error)
		Withdrawal(params *types.WithdrawParams) (*types.WithdrawResult, error)
	}

	IHistory interface {
		GetOrderHistory(params *types.OrderHistoryParams) ([]types.OrderHistoryResult, error)
		GetFills(params *types.FillHistoryParams) ([]types.FillHistoryResult, error)
	}

	IOrder interface {
		GetOrder(params *types.OrderParams) (*types.OrderResult, error)
		ExecuteOrder(params *types.ExecuteOrderParams) (*types.ExecuteOrderResult, error)
		CancelOrder(params *types.OrderParams) (*types.OrderResult, error)
		GetOrders(symbol string) ([]types.OrderResult, error)
		CancelOrders(symbol string) ([]types.OrderResult, error)
	}
)
