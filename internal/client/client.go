package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"crypto-service/internal/currency"
)

type HITBTC interface {
	GetCurrency(symbol string) (currency.Currency, error)
	GetAllCurrency() (currency.Currencies, error)
}

type client struct {
	URL string
}

func NewClient(url string) client {
	return client{URL: url}
}

func (c client) GetCurrency(symbol string) (currency.Currency, error) {

	resp, err := http.Get(c.URL + "/ticker/" + symbol)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return currency.Currency{}, err
	}

	var currency currency.Currency
	err = json.Unmarshal(body, &currency)
	if err != nil {
		return currency, err
	}

	return currency, nil
}

func (c client) GetAllCurrency() (currency.Currencies, error) {

	resp, err := http.Get(c.URL + "/ticker")
	if err != nil {
		return currency.Currencies{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return currency.Currencies{}, err
	}

	var currencies currency.Currencies
	err = json.Unmarshal(body, &currencies)
	if err != nil {
		return currency.Currencies{}, err
	}

	return currencies, nil
}
