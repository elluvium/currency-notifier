package main

import (
	"fmt"
	"log"
	"ua-currency-notifier/pkg/currency"
	"ua-currency-notifier/pkg/telegram"
)

func main() {
	currencies, err := currency.StateInitializer()
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
	for _, v := range currencies.Organizations {
		fmt.Println(v)
	}

	bot, updates, err := telegram.BotInit("")
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	err = telegram.Run(bot, updates)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
}
