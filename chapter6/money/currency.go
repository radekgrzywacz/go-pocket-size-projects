package money

type Currency struct {
	code      string
	precision byte
}

const ErrInvalidCurrenyCode = Error("invalid currency code")

func ParseCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrenyCode
	}

	for i := 0; i < 3; i++ {
		if !(code[i] >= 'A' && code[i] <= 'Z') {
			return Currency{}, ErrInvalidCurrenyCode
		}
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND":
		return Currency{code: code, precision: 3}, nil
	default:
		return Currency{code: code, precision: 2}, nil
	}
}

func (c Currency) String() string {
	return c.code
}

func (c Currency) Code() string {
	return c.code
}
