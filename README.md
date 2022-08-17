# UNISWAP EXAMPLE

### types/uniswap.go
The Uniswap Exchange stores data about its ETH and DAI reserves and also the pool constant. Every `add` will increase liquidity and hence the PoolConstant in the exchange.

To govern a synchronous swap between two Uniswap Exchanges, a `MainExchange` is created to govern the lock acquisition of both exchanges to ensure data consistency in the event of multiple threads. 

### types/tradebot.go
A `TradeBot` struct is created to store a waitgroup object to synchronize the process of scanning for available arbitrage and sending a trade. 