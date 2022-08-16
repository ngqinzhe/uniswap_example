package main

import (
	"fmt"
	"sync"

	"example.com/m/types"
	"example.com/m/consts"
)

func main() {
	m := types.MainExchange{
		UniswapV1: types.NewUniswapExchange(1000, 30000),
		UniswapV2: types.NewUniswapExchange(1000, 20000),
	}
	
	var _wg sync.WaitGroup

	tradeBot := types.TradeBot{
		WaitGroup: &_wg,
	}

	tradeBot.Scan(&m, consts.UNISWAPV1, consts.UNISWAPV2, consts.ETH, consts.DAI)

	_wg.Wait()
	fmt.Printf("u1 price: %v, u2 price: %v\n", types.GetConversionPrice(m, consts.UNISWAPV1, consts.ETH, consts.DAI, 1), 
		types.GetConversionPrice(m, consts.UNISWAPV2, consts.ETH, consts.DAI, 1))
	// fmt.Printf("pool1: %v, pool2: %v\n", GetConversionPrice(u1, types.ETH, 1), GetConversionPrice(u2, types.ETH, 1))
	// tradeEth := GetTradeInAmountForPoolBalanceRatio(u1, types.ETH, GetFairPrice(u1, u2, types.ETH))
	// fmt.Println("expected price:", GetFairPrice(u1, u2, types.ETH))
	// fmt.Println("trading eth:", tradeEth, "ETH...")
	// fmt.Printf("ETh received: %v\n", ArbitrageETHDAI(&u1, &u2, tradeEth))
	// fmt.Println(u1.GetETHReserves(), " ", u1.GetDAIReserves())
	// fmt.Println("pool constant in p1:", u1.GetPoolConstant(), "pool constant in p2:", u2.GetPoolConstant())
	// fmt.Printf("pool1: %v, pool2: %v\n", GetConversionPrice(u1, types.ETH, 1), GetConversionPrice(u2, types.ETH, 1))

}
