package types

import (
	"fmt"
	"sync"
	"testing"

	"example.com/m/consts"
)

func BenchmarkTradeBot(b *testing.B) {
	m := MainExchange{
		UniswapV1: NewUniswapExchange(1000, 30000),
		UniswapV2: NewUniswapExchange(1000, 20000),
	}

	var _wg sync.WaitGroup

	tradeBot := TradeBot{
		WaitGroup: &_wg,
	}

	for n := 0; n < b.N; n++ {
		tradeBot.Scan(&m, consts.UNISWAPV1, consts.UNISWAPV2, consts.ETH, consts.DAI)

		fmt.Printf("u1 price: %v, u2 price: %v\n", GetConversionPrice(m, consts.UNISWAPV1, consts.ETH, consts.DAI, 1),
			GetConversionPrice(m, consts.UNISWAPV2, consts.ETH, consts.DAI, 1))
	}
}
