package currency

import (
"gopkg.in/resty.v1"
	"encoding/json"
	"ua-currency-notifier/models"
)

func getCurrencies(url string) ([]byte, error) {
	resp, err := resty.R().Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func jsonParser() error {
	var state models.CurrencyState

	response, err := getCurrencies("http://resources.finance.ua/ua/public/currency-cash.json")
	if err != nil {
		return nil
	}

	err = json.Unmarshal([]byte(response), &state)
	if err != nil {
		return nil
	}


	return nil
}

func StateInitializer() error {
	err := jsonParser()
	if err != nil {
		return err
	}
	return nil
}
