package ecbank

import (
	"learngo-pockets/moneyconverter/money"
	"testing"
)

func TestExchangeRate(t *testing.T) {
	tt := map[string]struct {
		envelope envelope
		source   string
		target   string
		expected money.ExchangeRate
		error    error
	}{
		"nominal": {
			envelope: envelope{Rates: []currencyRate{{Currency: "USD", Rate: 1.5}}},
			source:   "EUR",
			target:   "USD",
			expected: mustParseExchangeRate(t, "1.5"),
			error:    nil,
		},
		"More values in envelope": {
			envelope: envelope{Rates: []currencyRate{{Currency: "USD", Rate: 1.5}, {Currency: "PLN", Rate: 4.3}}},
			source:   "EUR",
			target:   "PLN",
			expected: mustParseExchangeRate(t, "4.3"),
			error:    nil,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := tc.envelope.exchangeRate(tc.source, tc.target)
			if err != nil {
				t.Errorf("unable to marshal: %s", err.Error())
			}
			if got != tc.expected {
				t.Errorf("Expected: %q, got %q", tc.expected, got)
			}
		})
	}
}

func mustParseExchangeRate(t *testing.T, rate string) money.ExchangeRate {
	t.Helper()

	exchangeRate, err := money.ParseDecimal(rate)
	if err != nil {
		t.Fatalf("Unable to parse exchange rate %s", rate)
	}

	return money.ExchangeRate(exchangeRate)

}
