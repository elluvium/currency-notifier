package main

import (
	"fmt"
	"log"
	"ua-currency-notifier/pkg/currency"
)

func main() {
	currencies, err := currency.StateInitializer()
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
	for _, v := range currencies.Organizations {
		fmt.Println(v)
	}
}
