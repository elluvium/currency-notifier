package main

import (
	"fmt"
	"log"

	"gopkg.in/resty.v1"
)

func getCurrencies(url string) (string, error) {
	resp, err := resty.R().Get(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func jsonParser() error {
	return nil
}

func main() {
	resp, err := getCurrencies("http://resources.finance.ua/ua/public/currency-cash.json")
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
	fmt.Print(resp)
}
