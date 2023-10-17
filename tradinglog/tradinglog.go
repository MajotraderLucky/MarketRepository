package tradinglog

import (
	"log"

	"github.com/MajotraderLucky/MarketRepository/klinesdata"
)

func Hello() {
	log.Println("Hello, tradinglog!")
}

func GetFiboLevel() {
	var isAskPriceHigherThanLongFib236 bool
	var isAskPriceHigherThanLongFib382 bool
	var isAskPriceHigherThanLongFib500 bool
	var isAskPriceHigherThanLongFib618 bool
	var isAskPriceHigherThanLongFib786 bool

	fibLevel, isHigher := klinesdata.IsAskPriceHigherThanLongFibRetLog()

	if isHigher {
		log.Printf("The ask price is higher than %s Fibonacci retracement level.\n", fibLevel)
		switch fibLevel {
		case "LongFib236":
			isAskPriceHigherThanLongFib236 = true
		case "LongFib382":
			isAskPriceHigherThanLongFib382 = true
		case "LongFib500":
			isAskPriceHigherThanLongFib500 = true
		case "LongFib618":
			isAskPriceHigherThanLongFib618 = true
		case "LongFib786":
			isAskPriceHigherThanLongFib786 = true
		}
	} else {
		isAskPriceHigherThanLongFib236 = false
		isAskPriceHigherThanLongFib382 = false
		isAskPriceHigherThanLongFib500 = false
		isAskPriceHigherThanLongFib618 = false
		isAskPriceHigherThanLongFib786 = false
		log.Println("The ask price is not higher than any Fibonacci retracement level.")
	}
	if isAskPriceHigherThanLongFib236 {
		log.Println("236")
	}
	if isAskPriceHigherThanLongFib382 {
		log.Println("382")
	}
	if isAskPriceHigherThanLongFib500 {
		log.Println("500")
	}
	if isAskPriceHigherThanLongFib618 {
		log.Println("618")
	}
	if isAskPriceHigherThanLongFib786 {
		log.Println("786")
	}
}
