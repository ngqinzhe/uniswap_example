# UNISWAP EXAMPLE

## TYPES
### `types/uniswap.go`
The Uniswap Exchange stores data about its ETH and DAI reserves and also the pool constant. Every `Add` will increase liquidity, and every `Remove` will remove liquidity.
The `PoolConstant` will be calculated based on the formula.
```
PoolConstant = ethReserve * daiReserve
```

This is what will happen when trade is executed via the `Swap` call.
```
let FEE = 0.003
let k = PoolConstant
Liquidity added for tokenIn = amountIn * FEE (will increase PoolConstant)

tokenInReserve += amountIn * (1 - FEE) (will not increase PoolConstant)
tokenOutAmount = currentTokenOutReserve - (PoolConstant / tokenInReserve)

```

To ensure no deadlocks in multiple threads, a `MainExchange` will be used as the central exchange to send orders, which ensures that the lock acquisition sequence remains consistent throughout.

### `types/tradebot.go`
A `TradeBot` struct is created to store a `WaitGroup` to synchronise the goroutines when `StartTrade` is called. `Profit` is used to store the total profits gained when running the TradeBot. 

## Calculation of Arbitrage
Although this may not be fairly accurate, we source for arbitrage opportunities by finding price differences between both exchanges. As there is a fee of 0.3% and potential 'slippage' in our trades, I have created a minimum profit margin in `consts`
```
PROFITMARGIN = 1.005
```
This is used to ensure that whenever we scan for prices, there is at least potentially 0.2% of profit for us after arbitrage, if the profit margin is lower than this threshold, we will not take the trade. 

## Calculation of ETH to trade a.k.a amountIn
To calculate how much we should trade to ensure that we do not over trade and still remain profitable, I used a simple calculation to calculate the expected price of tokens using:
```
FairPrice = (ExchangeAPrice + ExchangeBPrice) / 2
```
Although this calculation is only a rough estimation of my own, it allows us to trade through both exchanges and ensure that ETH/DAI prices in both exchanges are relatively similar at the end of our arbitrage, which ensures that markets remain efficient.

## Execution
In `main.go` we will send a `goroutine` to use our `TradeBot` to scan for opportunities.
Our `TradeBot` will continuously scan for arbitrage opportunities between both exchanges, and execute a trade if it deems profitable.

## Benchmarking
Benchmark done on our `TradeBot` scanner.
```
goos: darwin
goarch: amd64
pkg: example.com/m/types
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
1000000000	0.0000443 ns/op	 0 B/op	       0 allocs/op
PASS
ok  	example.com/m/types	0.161s
```