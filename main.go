package main

import (
	"fmt"

	"example.com/m/consts"
	"example.com/m/types"
)

func main() {
	m := types.InitMainExchange(10, 30, 10, 20)
	tradeBot := types.InitTradeBot()

	tradeBot.ScanPrices(&m, consts.UNISWAPV1, consts.UNISWAPV2, consts.ETH, consts.DAI)

	fmt.Printf("u1 price: %v, u2 price: %v\n", types.GetConversionPrice(m, consts.UNISWAPV1, consts.ETH, consts.DAI, 1),
		types.GetConversionPrice(m, consts.UNISWAPV2, consts.ETH, consts.DAI, 1))

}
