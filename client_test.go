package bpx_api_sdk_go

import (
	"encoding/json"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/constants"
	"github.com/ThreeAndTwo/bpx-api-sdk-go/types"
	"log"
	"testing"
)

const (
	url       = "https://api.backpack.exchange"
	apiKey    = ""
	apiSecret = ""
	isDebug   = false
)

var bpxClient *BackpackExchange

func init() {
	backpack, err := NewBackpack(url, apiKey, apiSecret, isDebug)
	if err != nil {
		panic("init backpack error:" + err.Error())
	}
	bpxClient = backpack
}

func TestMarkets(t *testing.T) {
	markets := bpxClient.Markets()

	assets, err := markets.GetAssets()
	if err != nil {
		log.Fatalf("get assets error: %s", err)
	}
	marshal, err := json.Marshal(assets)
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}
	log.Printf("assests: %s", string(marshal))

	getMarkets, err := markets.GetMarkets()
	if err != nil {
		log.Fatalf("get markets error: %s", err)
	}
	marshal1, err := json.Marshal(getMarkets)
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}
	log.Printf("getMarkets: %s", string(marshal1))

	symbol := "sol_usdc"
	ticker, err := markets.GetTicker(symbol)
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}
	marshal2, err := json.Marshal(ticker)
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}

	log.Printf("ticker: %s", string(marshal2))

	tickers, err := markets.GetTickers()
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}
	marshal3, err := json.Marshal(tickers)
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}

	log.Printf("tickers: %s", string(marshal3))
	depths, err := markets.GetDepth(symbol)
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}
	marshal4, err := json.Marshal(depths)
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}
	log.Printf("depths: %s", string(marshal4))

	klineData, err := markets.GetKline(symbol, constants.Interval1m, 0, 0)
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}
	marshal5, err := json.Marshal(klineData)
	if err != nil {
		log.Fatalf("marshal error: %s", err)
	}
	log.Printf("depths: %s", string(marshal5))
}

func TestSystem(t *testing.T) {
	_system := bpxClient.System()
	err := _system.Ping()
	if err != nil {
		log.Fatalf("ping err: %s", err)
	}

	status, err := _system.GetStatus()
	if err != nil {
		log.Fatalf("status err: %s", err)
	}
	log.Printf("status: %s", status)
	err = _system.GetTime()
	if err != nil {
		log.Printf("system time error: %s", err)
	}
}

func TestTrades(t *testing.T) {
	_tratdes := bpxClient.Trades()
	symbol := "sol_usdc"
	limit := int64(100)
	offset := int64(0)

	trades, err := _tratdes.GetRecentTrades(symbol, limit)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	log.Printf("trades: %v", trades)

	historicalTrades, err := _tratdes.GetHistoricalTrades(symbol, limit, offset)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	log.Printf("historicalTrades: %v", historicalTrades)
}

func TestCapital(t *testing.T) {
	_capital := bpxClient.Capital()
	//balances, err := _capital.GetBalances()
	//if err != nil {
	//	log.Fatalf("err: %s", err)
	//}
	//log.Printf("balances: %v", balances)

	//deposits, err := _capital.GetDeposits()
	//if err != nil {
	//	log.Fatalf("err: %s", err)
	//}
	//log.Printf("deposits: %v", deposits)

	//address, err := _capital.GetDepositAddress(constants.ChainSolana)
	//if err != nil {
	//	log.Fatalf("err: %s", err)
	//}
	//log.Printf("address: %v", address)

	//withdrawals, err := _capital.GetWithdrawals(constants.DefaultLimit, constants.DefaultOffset)
	//if err != nil {
	//	log.Fatalf("err: %s", err)
	//}
	//log.Printf("withdrawals: %v", withdrawals)

	//GetWithdrawals

	// ========================================================
	withdrawParams := &types.WithdrawParams{
		Address:        "",
		Blockchain:     constants.ChainSolana,
		ClientId:       "",
		Quantity:       "1",
		Symbol:         "PYTH_USDC",
		TwoFactorToken: "",
	}
	withdrawal, err := _capital.Withdrawal(withdrawParams)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	log.Printf("withdrawal: %v", withdrawal)
}

func TestHistory(t *testing.T) {
	_history := bpxClient.History()

	symbol := "SOL_USDC"
	orderHistory := &types.OrderHistoryParams{
		Symbol: symbol,
		Offset: constants.DefaultOffset,
		Limit:  constants.DefaultLimit,
	}
	getOrderHistory, err := _history.GetOrderHistory(orderHistory)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Printf("%v", getOrderHistory)

	fillHistory := &types.FillHistoryParams{
		OrderId: "",
		Symbol:  symbol,
		Limit:   constants.DefaultLimit,
		Offset:  constants.DefaultOffset,
	}
	fills, err := _history.GetFills(fillHistory)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Printf("%v", fills)

}

func TestOrder(t *testing.T) {
	_order := bpxClient.Order()

	symbol := "SOL_USDC"
	//
	//params := &types.OrderParams{
	//	ClientId: 0,
	//	OrderId:  "112082126676754432",
	//	Symbol:   symbol,
	//}
	//
	////_order
	//cancelOrder, err := _order.CancelOrder(params)
	//if err != nil {
	//	log.Fatalf("err: %v", err)
	//}
	//marshal, _ := json.Marshal(cancelOrder)
	//log.Printf(string(marshal))

	//params := &types.ExecuteOrderParams{
	//	OrderType:           constants.OrderLimit,
	//	PostOnly:            true,
	//	Price:               "1",
	//	Quantity:            "0.2",
	//	Side:                constants.SideBid,
	//	Symbol:              symbol,
	//}
	//executeOrder, err := _order.ExecuteOrder(params)
	//if err != nil {
	//	t.Fatalf("err: %v", err)
	//}
	//t.Logf("executeOrder: %v", executeOrder)

	orders, err := _order.CancelOrders(symbol)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	t.Logf("cancelOrder: %v", orders)

	//orderParams := &types.OrderParams{
	//	//ClientId: 0,
	//	//OrderId: "112082126676754432",
	//	Symbol: symbol,
	//}
	//getOrder, err := _order.GetOrder(orderParams)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("%v", getOrder)
	////_order.ExecuteOrder()
	//_order.CancelOrder()

	//orders, err := _order.GetOrders("")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//marshal, _ := json.Marshal(orders)
	////log.Printf("%v", orders)
	//log.Printf(string(marshal))
	//_order.CancelOrders()
}
