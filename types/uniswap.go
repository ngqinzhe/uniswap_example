package types

import (
	"sync"

	"example.com/m/consts"
)

type MainExchange struct {
	UniswapV1 UniswapExchange
	UniswapV2 UniswapExchange
	lock      sync.Mutex
}

type UniswapExchange struct {
	EthReserve   float64
	DaiReserve   float64
	PoolConstant float64
}

func InitMainExchange(ex1ETH, ex1DAI, ex2ETH, ex2DAI float64) MainExchange {
	return MainExchange{
		UniswapV1: NewUniswapExchange(ex1ETH, ex1DAI),
		UniswapV2: NewUniswapExchange(ex2ETH, ex2DAI),
	}
}

func NewUniswapExchange(_ethReserve, _daiReserve float64) UniswapExchange {
	return UniswapExchange{
		EthReserve:   _ethReserve,
		DaiReserve:   _daiReserve,
		PoolConstant: _daiReserve * _ethReserve,
	}
}

func (m *MainExchange) Add(exchange, token string, amount float64) {
	m.lock.Lock()
	exchangePtr := m.getExchangePtr(exchange)
	tokenPtr := getTokenReservePtr(token, exchangePtr)
	*tokenPtr += amount
	exchangePtr.PoolConstant = exchangePtr.EthReserve * exchangePtr.DaiReserve
	m.lock.Unlock()
}

func (m *MainExchange) Remove(exchange, token string, amount float64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	exchangePtr := m.getExchangePtr(exchange)
	tokenPtr := getTokenReservePtr(token, exchangePtr)

	*tokenPtr -= amount
	exchangePtr.PoolConstant = exchangePtr.EthReserve * exchangePtr.DaiReserve
}

func (m *MainExchange) Swap(exchange, tokenIn, tokenOut string, amount float64) float64 {
	exchangePtr := m.getExchangePtr(exchange)
	tokenInReservePtr := getTokenReservePtr(tokenIn, exchangePtr)
	tokenOutReservePtr := getTokenReservePtr(tokenOut, exchangePtr)

	m.Add(exchange, tokenIn, amount*consts.FEE)

	m.lock.Lock()
	*tokenInReservePtr += amount * (1 - consts.FEE)
	tokenAmountOut := *tokenOutReservePtr - (exchangePtr.PoolConstant / *tokenInReservePtr)
	*tokenOutReservePtr -= tokenAmountOut
	m.lock.Unlock()
	return tokenAmountOut
}


func (m *MainExchange) getExchangePtr(exchange string) *UniswapExchange {
	if exchange == consts.UNISWAPV1 {
		return &m.UniswapV1
	} else {
		return &m.UniswapV2
	}
}

func getTokenReservePtr(token string, u *UniswapExchange) *float64 {
	if token == consts.DAI {
		return &u.DaiReserve
	} else {
		return &u.EthReserve
	}
}
