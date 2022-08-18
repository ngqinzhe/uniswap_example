package types

import (
	"sync"
	"testing"

	"example.com/m/consts"
)

func BenchmarkTradeBot(b *testing.B) {
	m := InitMainExchange(10000, 30000, 10000, 20000)

	var _wg sync.WaitGroup

	tradeBot := TradeBot{
		WaitGroup: &_wg,
	}
	tradeBot.WaitGroup.Add(1)
	go tradeBot.ScanPrices(&m, consts.UNISWAPV1, consts.UNISWAPV2, consts.ETH, consts.DAI)
	tradeBot.WaitGroup.Wait()
}
