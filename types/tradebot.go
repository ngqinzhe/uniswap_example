package types

import (
	"math"
	"sync"

	"example.com/m/consts"
)

type TradeBot struct {
	WaitGroup *sync.WaitGroup
}

func InitTradeBot() TradeBot {
	return TradeBot {
		WaitGroup: &sync.WaitGroup{},
	}
}

func (t *TradeBot) StartTrade(m *MainExchange, firstExchange, secondExchange, tokenIn, tokenOut string) float64 {
	defer t.WaitGroup.Done()
	amountIn := GetTradeInValue(*m, firstExchange, tokenIn, tokenOut, GetFairPrice(*m, tokenIn, tokenOut))
	amountOut := PerformArbitrage(m, firstExchange, secondExchange, tokenIn, tokenOut, amountIn)
	return amountOut
	// fmt.Printf("Traded %v%v for %v%v, profit: %vETH\n", amountIn, tokenIn, amountOut, tokenOut, amountOut-amountIn)

}

// ScanPrices will send go routines to trade if a profitable trade is spotted
func (t *TradeBot) ScanPrices(m *MainExchange, firstExchange, secondExchange, tokenIn, tokenOut string) {
	for {
		exchangeAPrice := GetConversionPrice(*m, firstExchange, tokenIn, tokenOut, 1)
		exchangeBPrice := GetConversionPrice(*m, secondExchange, tokenIn, tokenOut, 1)

		if exchangeAPrice > consts.PROFITMARGIN*exchangeBPrice {
			t.WaitGroup.Add(1)
			go t.StartTrade(m, consts.UNISWAPV1, consts.UNISWAPV2, consts.ETH, consts.DAI)
			t.WaitGroup.Wait()
		} else if exchangeBPrice > consts.PROFITMARGIN*exchangeAPrice {
			t.WaitGroup.Add(1)
			go t.StartTrade(m, consts.UNISWAPV2, consts.UNISWAPV1, consts.ETH, consts.DAI)
			t.WaitGroup.Wait()
		} else {
			return
		}
	}
}

// In our calculation functions, we want to pass in a copy of the main exchange
// so that we do not alter the original values
func GetConversionPrice(m MainExchange, exchange, tokenIn, tokenOut string, amount float64) float64 {
	return m.CheckSwap(exchange, tokenIn, tokenOut, amount)
}

func GetTradeInValue(m MainExchange, exchange, tokenIn, tokenOut string, price float64) float64 {
	// calculate unit price for conversion
	amountOut := GetConversionPrice(m, exchange, tokenIn, tokenOut, 1)
	// eth kx of tokenOut c = kx^2
	// change to tokenIn k/2x of dai c = (k/2)x^2
	exchangePtr := *m.getExchangePtr(exchange)
	k := exchangePtr.PoolConstant
	originalAmount := math.Sqrt(k / amountOut)
	resultAmount := math.Sqrt(k / price)
	return resultAmount - originalAmount
}

func GetFairPrice(m MainExchange, tokenIn, tokenOut string) float64 {
	// simple get fair price of exchange
	return (GetConversionPrice(m, consts.UNISWAPV1, tokenIn, tokenOut, 1) + GetConversionPrice(m, consts.UNISWAPV2, tokenIn, tokenOut, 1)) / 2
}

func PerformArbitrage(m *MainExchange, firstExchange, secondExchange, tokenIn, tokenOut string, amountIn float64) float64 {
	amountOut := m.Swap(firstExchange, tokenIn, tokenOut, amountIn)
	return m.Swap(secondExchange, tokenOut, tokenIn, amountOut)
}
