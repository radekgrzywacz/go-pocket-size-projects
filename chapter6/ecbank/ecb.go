package ecbank

import (
	"fmt"
	"learngo-pockets/moneyconverter/money"
	"net/http"
)

type Client struct {
	url string
}

func (c Client) FetchExchangeRates(source, target money.Currency) (money.ExchangeRate, error) {
	const path = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	if c.url == "" {
		c.url = path
	}

	resp, err := http.Get(path)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", http.ErrServerClosed, err.Error())
	}

	defer resp.Body.Close()
	if err = checkStatusCode(resp.StatusCode); err != nil {
		return money.ExchangeRate{}, err
	}

	rate, err := readRateFromResponse(source.Code(), target.Code(), resp.Body)
	if err != nil {
		return money.ExchangeRate{}, err
	}

	return rate, nil
}

const (
	clientErrorClass = 4
	serverErrorClass = 5
)

const (
	ErrCallingServer      = ecbankError("error calling server")
	ErrUnexpectedFormat   = ecbankError("unexpected response format")
	ErrChangeRateNotFound = ecbankError("couldn't find the exchange rate")
	ErrClientSide         = ecbankError("client side error when contacting ECB")
	ErrServerSide         = ecbankError("server side error when contacting ECB")
	ErrUnknownStatusCode  = ecbankError("unknown status code contacting ECB")
)

func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		return fmt.Errorf("%w: %d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w: %d", ErrServerSide, statusCode)
	default:
		return fmt.Errorf("%w: %d", ErrUnknownStatusCode, statusCode)
	}
}

func httpStatusClass(status int) int {
	const httpStatusSize = 100
	return status / httpStatusSize
}
