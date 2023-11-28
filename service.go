package main

import (
	"context"
	"fmt"
	"time"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"GG":  100_00.0,
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	time.Sleep(100 * time.Millisecond)
	price, ok := priceMocks[ticker]

	if !ok {
		return price, fmt.Errorf("the giver ticker (%s) is not supported", ticker)
	}

	return price, nil
}
