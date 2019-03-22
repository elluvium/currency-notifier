package main

import (
	"log"
	"ua-currency-notifier/pkg/currency"
)

func main() {
	err := currency.StateInitializer()
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
}
