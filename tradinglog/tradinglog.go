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

type TradeLevels struct {
	Response string
	Error    error
}

var tradeLevels *TradeLevels

type GetFiboLevelStart func() *TradeLevels // добавить эти две строки.
var GetFiboLevelStartTradeOnce GetFiboLevelStart = getFiboLevelStartTradeOnce

func getFiboLevelStartTradeOnce() *TradeLevels {
	// This will get the values only once and store them
	if tradeLevels == nil {
		response, err := GetFiboLevelStartTrade()
		tradeLevels = &TradeLevels{
			Response: response,
			Error:    err,
		}
	}
	return tradeLevels
}

func IsStopTradeLevel236Met() bool {
	levels := GetFiboLevelStartTradeOnce()
	if levels.Error != nil {
		log.Printf("Error getting Fibonacci retracement level: %v", levels.Error)
		return false
	}
	if levels.Response == "stopTrade236" {
		return true
	}
	return false
}

func IsStartTradeLevel382Met() bool {
	levels := GetFiboLevelStartTradeOnce()
	if levels.Error != nil {
		log.Printf("Error getting Fibonacci retracement level: %v", levels.Error)
		return false
	}
	if levels.Response == "startTrade382" {
		return true
	}
	return false
}

func IsStartTradeLevel500Met() bool {
	levels := GetFiboLevelStartTradeOnce()
	if levels.Error != nil {
		log.Printf("Error getting Fibonacci retracement level: %v", levels.Error)
		return false
	}
	if levels.Response == "startTrade500" {
		return true
	}
	return false
}

func IsStartTradeLevel618Met() bool {
	levels := GetFiboLevelStartTradeOnce()
	if levels.Error != nil {
		log.Printf("Error getting Fibonacci retracement level: %v", levels.Error)
		return false
	}
	if levels.Response == "startTrade618" {
		return true
	}
	return false
}

func IsStartTradeLevel786Met() bool {
	levels := GetFiboLevelStartTradeOnce()
	if levels.Error != nil {
		log.Printf("Error getting Fibonacci retracement level: %v", levels.Error)
		return false
	}
	if levels.Response == "startTrade786" {
		return true
	}
	return false
}
