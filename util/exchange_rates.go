package util

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
)

type currencyRate struct {
	TWD float64 `json:"TWD"`
	JPY float64 `json:"JPY"`
	USD float64 `json:"USD"`
}

type exchangeRates struct {
	Currencies map[string]currencyRate `json:"currencies"`
}

type ExchangeRates map[string]map[string]float64

func LoadExchangeRates(path string) (rates ExchangeRates, err error) {
	fileContent, err := os.ReadFile(filepath.Join(path, "exchange_rates.json"))
	if err != nil {
		return
	}

	var exchangeRates exchangeRates
	if err = json.Unmarshal([]byte(fileContent), &exchangeRates); err != nil {
		return
	}

	rates = transformExchangeRates(exchangeRates)
	return
}

func transformExchangeRates(exchangeRates exchangeRates) ExchangeRates {
	currencies := exchangeRates.Currencies
	rates := map[string]map[string]float64{}

	for source, targetRates := range currencies {
		v := reflect.ValueOf(targetRates)
		rates[source] = map[string]float64{}

		for i := 0; i < v.NumField(); i++ {
			rates[source][v.Type().Field(i).Name] = v.Field(i).Interface().(float64)
		}
	}

	return rates
}
