package money

type Amount struct {
	quantity Decimal
	currency Currency
}

const (
	ErrTooPrecise = Error("Quantity is too precise")
)

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}

	return Amount{quantity: quantity, currency: currency}, nil
}

func (a Amount) validate() error {
	switch {
	case a.quantity.subunits > maxDecimal:
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		return ErrTooPrecise
	}

	return nil
}
