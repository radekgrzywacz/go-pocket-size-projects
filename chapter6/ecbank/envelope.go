package ecbank

import (
	"encoding/xml"
	"fmt"
	"io"
	"learngo-pockets/moneyconverter/money"
)

type envelope struct {
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

type currencyRate struct {
	Currency string  `xml:"currency,attr"`
	Rate     float64 `xml:"rate,attr"`
}

const baseCurrencyCode = "EUR"

func (e envelope) exchangeRates() map[string]float64 {
	rates := make(map[string]float64, len(e.Rates)+1)

	for _, c := range e.Rates {
		rates[c.Currency] = c.Rate
	}

	rates[baseCurrencyCode] = 1.

	return rates
}

func (e envelope) exchangeRate(source, target string) (money.ExchangeRate, error) {
	if source == target {
		one, err := money.ParseDecimal("1")
		if err != nil {
			return money.ExchangeRate{}, fmt.Errorf("unable to create a rate of value 1: %w", err)
		}
		return money.ExchangeRate(one), nil
	}

	rates := e.exchangeRates()
	sourceFactor, sourceFound := rates[source]
	if !sourceFound {
		return money.ExchangeRate{}, fmt.Errorf("Failed to found the source currency %s", source)
	}

	targetFactor, targetFound := rates[target]
	if !targetFound {
		return money.ExchangeRate{}, fmt.Errorf("Failed to found the target currency %s", source)
	}

	rate, err := money.ParseDecimal(fmt.Sprintf("%.10f", targetFactor/sourceFactor))
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("unable to parse exchange rate from %s to %s: %w", source, target, err)
	}

	return money.ExchangeRate(rate), nil
}

func readRateFromResponse(source, target string, respBody io.Reader) (money.ExchangeRate, error) {
	decoder := xml.NewDecoder(respBody)

	var ecbMessage envelope
	err := decoder.Decode(&ecbMessage)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w, %s", ErrUnexpectedFormat, err)
	}

	rate, err := ecbMessage.exchangeRate(source, target)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrChangeRateNotFound, err)
	}

	return rate, nil
}
