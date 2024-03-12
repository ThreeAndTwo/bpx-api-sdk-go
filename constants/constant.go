package constants

type IntervalTy string

const (
	Interval1m     IntervalTy = "1m"
	Interval3m     IntervalTy = "3m"
	Interval5m     IntervalTy = "5m"
	Interval15m    IntervalTy = "15m"
	Interval30m    IntervalTy = "30m"
	Interval1h     IntervalTy = "1h"
	Interval2h     IntervalTy = "2h"
	Interval4h     IntervalTy = "4h"
	Interval6h     IntervalTy = "6h"
	Interval8h     IntervalTy = "8h"
	Interval12h    IntervalTy = "12h"
	Interval1d     IntervalTy = "1d"
	Interval3d     IntervalTy = "3d"
	Interval1w     IntervalTy = "1w"
	Interval1month IntervalTy = "1month"
)

type InstructionTy string

const (
	InstructionBalanceQuery         InstructionTy = "balanceQuery"
	InstructionDepositQueryAll      InstructionTy = "depositQueryAll"
	InstructionDepositAddressQuery  InstructionTy = "depositAddressQuery"
	InstructionWithdrawalQueryAll   InstructionTy = "withdrawalQueryAll"
	InstructionWithdrawal           InstructionTy = "withdraw"
	InstructionOrderHistoryQueryAll InstructionTy = "orderHistoryQueryAll"
	InstructionFillHistoryQueryAll  InstructionTy = "fillHistoryQueryAll"
	InstructionOrderQuery           InstructionTy = "orderQuery"
	InstructionOrderExecute         InstructionTy = "orderExecute"
	InstructionOrderCancel          InstructionTy = "orderCancel"
	InstructionOrderQueryAll        InstructionTy = "orderQueryAll"
	InstructionOrderCancelAll       InstructionTy = "orderCancelAll"
)

type ChainTy string

const (
	ChainSolana   ChainTy = "Solana"
	ChainEthereum ChainTy = "Ethereum"
	ChainPolygon  ChainTy = "Polygon"
	ChainBitcoin  ChainTy = "Bitcoin"
)

type OrderTypeTy string

const (
	OrderMarket OrderTypeTy = "Market"
	OrderLimit  OrderTypeTy = "Limit"
)

type SelfTradePreventionTy string

const (
	SelfTradePreventionRejectTaker SelfTradePreventionTy = "RejectTaker"
	SelfTradePreventionRejectMaker SelfTradePreventionTy = "RejectMaker"
	SelfTradePreventionRejectBoth  SelfTradePreventionTy = "RejectBoth"
	SelfTradePreventionAllow       SelfTradePreventionTy = "Allow"
)

type OrderSideTy string

const (
	SideBid OrderSideTy = "Bid"
	SideAsk OrderSideTy = "Ask"
)

type TimeInForceTy string

const (
	TimeInForceGTC = "GTC"
	TimeInForceIOC = "IOC"
	TimeInForceFOK = "FOK"
)

const (
	DefaultLimit  = 100
	DefaultOffset = 0
	MaxLimit      = 1000

	DefaultWindow = 5000
)
