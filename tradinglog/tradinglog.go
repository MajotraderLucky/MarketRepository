package tradinglog

import (
	"log"

	"github.com/MajotraderLucky/MarketRepository/klinesdata"
	"github.com/MajotraderLucky/MarketRepository/orderinfolog"
)

func Hello() {
	log.Println("Hello, tradinglog!")
}

func GetFiboLevelStartTrade() (string, error) {
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

	threshold := 5
	isHigherCorridor, err := klinesdata.IsCorridorHigher(threshold)
	if err != nil {
		log.Printf("Error getting corridor: %v", err)
		return "", err
	}

	newOpenPos382 := isAskPriceHigherThanLongFib382 && isHigherCorridor && !orderinfolog.CheckIfOpenOrdersExist()
	newOpenPos500 := isAskPriceHigherThanLongFib500 && isHigherCorridor && !orderinfolog.CheckIfOpenOrdersExist()
	newOpenPos618 := isAskPriceHigherThanLongFib618 && isHigherCorridor && !orderinfolog.CheckIfOpenOrdersExist()
	newOpenPos786 := isAskPriceHigherThanLongFib786 && isHigherCorridor && !orderinfolog.CheckIfOpenOrdersExist()

	if isAskPriceHigherThanLongFib236 {
		log.Println("stopTrade236")
		return "stopTrade236", nil
	}
	if newOpenPos382 {
		log.Println("startTrade382")
		return "startTrade382", nil
	}
	if newOpenPos500 {
		log.Println("startTrade500")
		return "startTrade500", nil
	}
	if newOpenPos618 {
		log.Println("startTrade618")
		return "startTrade618", nil
	}
	if newOpenPos786 {
		log.Println("startTrade786")
		return "startTrade786", nil
	}
	return "", nil
}
