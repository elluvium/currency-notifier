package main

import (
	"currency-notifier/pkg/currency"
	"currency-notifier/pkg/telegram"
	"fmt"
	"log"
	"os"
)

func main() {
	tg := getEnvOrErr("TELEGRAM_TOKEN")

	currencies, err := currency.StateInitializer()
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
	for _, v := range currencies.Organizations {
		fmt.Println(v)
	}

	bot, updates, err := telegram.BotInit(tg)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	err = telegram.Run(bot, updates)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

func getEnvOrErr(key string) string {
	v, success := os.LookupEnv(key)
	if !success {
		log.Fatalf("Environment variable %s cannot be empty", key)
	}
	return v
}
