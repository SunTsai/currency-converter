package api

import (
	"suntsai/currency-converter/util"
)

type Cache interface {
	ExchangeRates() (util.ExchangeRates, error)
}

type ExchangeRates struct {
}

func NewExchangeRates() ExchangeRates {
	return ExchangeRates{}
}

func (e ExchangeRates) ExchangeRates() (rates util.ExchangeRates, err error) {
	rates, err = util.LoadExchangeRates("..")
	return
}
