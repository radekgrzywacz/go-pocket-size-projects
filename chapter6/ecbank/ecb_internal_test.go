package ecbank

import (
	"errors"
	"fmt"
	"learngo-pockets/moneyconverter/money"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestClient_FetchExchangeRate_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope><Cube><Cube>
			<Cube currency='USD' rate='2'/>
			<Cube currency='RON' rate='6'/>
		</Cube></Cube></gesmes:Envelope>`)
	}))

	defer ts.Close()

	proxyUrl, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("failed to parse proxy URL: %v", err)
	}

	ecb := Client{
		client: http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
			Timeout: time.Second,
		},
	}

	got, err := ecb.FetchExchangeRates(mustParseCurrency(t, "USD"), mustParseCurrency(t, "RON"))
	want := mustParseDecimal(t, "3")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if money.Decimal(got) != want {
		t.Errorf("FetchExchangeRate() got = %v, want %v", money.Decimal(got), want)
	}
}

func TestClient_FetchExchangeRates_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 5)
	}))
	defer ts.Close()

	proxyUrl, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("Failed to parse proxy URL: %v", err)
	}

	ecb := Client{
		client: http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
			Timeout: time.Second,
		},
	}

	_, err = ecb.FetchExchangeRates(mustParseCurrency(t, "USD"), mustParseCurrency(t, "RON"))
	if !errors.Is(err, ErrTimeout) {
		t.Errorf("Unexpeccted error: %v, expected %v", err, ErrTimeout)
	}
}

func mustParseCurrency(t *testing.T, code string) money.Currency {
	t.Helper()

	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("cannot parse currency %s code", code)
	}

	return currency
}

func mustParseDecimal(t *testing.T, decimal string) money.Decimal {
	t.Helper()

	dec, err := money.ParseDecimal(decimal)
	if err != nil {
		t.Fatalf("cannot parse decimal %s", decimal)
	}

	return dec
}
