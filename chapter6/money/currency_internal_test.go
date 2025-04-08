package money

import (
	"errors"
	"testing"
)

func TestParseCurrency_Success(t *testing.T) {
	tt := map[string]struct {
		in       string
		expected Currency
	}{
		"majority EUR":   {in: "EUR", expected: Currency{"EUR", 2}},
		"thousandth BHD": {in: "BHD", expected: Currency{code: "BHD", precision: 3}},
		"tenth VND":      {in: "VND", expected: Currency{code: "VND", precision: 1}},
		"integer IRR":    {in: "IRR", expected: Currency{code: "IRR", precision: 0}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(tc.in)
			if err != nil {
				t.Errorf("Expected no error, got %s", err.Error())
			}

			if got != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestParseCurrency_UnknownCurrency(t *testing.T) {
	_, err := ParseCurrency("INVALID")
	if !errors.Is(err, ErrInvalidCurrenyCode) {
		t.Errorf("Expected error %s, got %v", ErrInvalidCurrenyCode, err)
	}
}
