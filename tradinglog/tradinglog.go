package tradinglog

import (
	"log"

	"github.com/MajotraderLucky/MarketRepository/klinesdata"
)

func Hello() {
	log.Println("Hello, tradinglog!")
}

func GetFiboLevel() {
	fibLevel, isHigher := klinesdata.IsAskPriceHigherThanLongFibRetLog()
	if isHigher {
		log.Printf("The ask price is higher than %s Fibonacci retracement level.\n", fibLevel)
	} else {
		log.Println("The ask price is not higher than any Fibonacci retracement level.")
	}
}
