package money

import "fmt"

type ratesFetcher interface {
	FetchExchangeRates(source, target Currency) (ExchangeRate, error)
}

func Convert(amount Amount, to Currency, rates ratesFetcher) (Amount, error) {
	r, err := rates.FetchExchangeRates(amount.currency, to)
	if err != nil {
		return Amount{}, fmt.Errorf("cannot get exchange rate: %w", err)
	}

	convertedValue := applyExchangeRate(amount, to, r)
	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}

	return convertedValue, nil
}

type ExchangeRate Decimal

// applyExchangeRate returns a new Amount representing the input multiplied by the rate.
// The precision of the returned value is that of the target Currency.
// This function does not guarantee that the output amount is supported.
func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) Amount {
	converted := multiply(a.quantity, rate)

	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:
		converted.subunits = converted.subunits * pow10(target.precision-converted.precision)
	}

	converted.precision = target.precision

	return Amount{
		currency: target,
		quantity: converted,
	}
}

func multiply(d Decimal, r ExchangeRate) Decimal {
	dec := Decimal{
		subunits:  d.subunits * r.subunits,
		precision: d.precision + r.precision,
	}

	return dec
}
