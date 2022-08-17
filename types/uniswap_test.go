package types

import (
	"testing"

	"example.com/m/consts"
)

func TestUniswapExchange(t *testing.T) {
	tests := []struct {
		name         string
		m            MainExchange
		amountInETH  float64
		isProfitable bool
	}{
		{
			name: "Overtrade with arbitrage opportunity",
			m:            InitMainExchange(10, 30, 10, 20),
			amountInETH:  10,
			isProfitable: false,
		},
		{
			name: "Undertrade with arbitrage opportunity",
			m: InitMainExchange(1000000, 3000000, 1000000, 2000000),
			amountInETH: 100,
			isProfitable: true,
		},
		{
			name: "No arbitrage opportunity",
			m: InitMainExchange(10, 30, 10, 30),
			amountInETH: 10,
			isProfitable: false,
		},
	}

	for _, tt := range tests {
		result := PerformArbitrage(&tt.m, consts.UNISWAPV1, consts.UNISWAPV2, consts.ETH, consts.DAI, tt.amountInETH)
		if (result > tt.amountInETH) != tt.isProfitable {
			t.Errorf("Expected %v profitability for %v. %vETH in, %vETH out.", tt.isProfitable, tt.name, tt.amountInETH, result)
		}
	}
}
