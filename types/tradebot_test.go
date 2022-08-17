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
	tradeBot.ScanPrices(&m, consts.UNISWAPV1, consts.UNISWAPV2, consts.ETH, consts.DAI)

	// fmt.Printf("u1 price: %v, u2 price: %v\n", GetConversionPrice(m, consts.UNISWAPV1, consts.ETH, consts.DAI, 1),
	// 	GetConversionPrice(m, consts.UNISWAPV2, consts.ETH, consts.DAI, 1))
}
