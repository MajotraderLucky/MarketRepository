package findtheorderlevel

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/adshao/go-binance/v2"
)

func Hello() {
	fmt.Println("-------------------------------------")
	fmt.Println("Hello from package findtheorderlevel!")
	fmt.Println("-------------------------------------")
	fmt.Println()
}

func FindOrderLevel(ch chan int) {
	tickerName := "BTCUSDT"
	fmt.Println(tickerName, "- ticker")

	fmt.Println("----------------------")
	apiKey, exists := os.LookupEnv("BINANCE_API_KEY")
	if exists {
		fmt.Println("apiKey exist")
	}

	secretKey, exexists := os.LookupEnv("BINANCE_SECRET_KEY")
	if exexists {
		fmt.Println("secretKey exist")
		fmt.Println("----------------------")
	}

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

	res, err := futuresClient.NewDepthService().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(res)

	depthVar, _ := json.Marshal(res)
	// fmt.Println(string(depthVar))

	type AutoGenerated struct {
		LastUpdateID int64 `json:"lastUpdateId"`
		E            int64 `json:"E"`
		T            int64 `json:"T"`
		Bids         []struct {
			Price    string `json:"Price"`
			Quantity string `json:"Quantity"`
		} `json:"bids"`
		Asks []struct {
			Price    string `json:"Price"`
			Quantity string `json:"Quantity"`
		} `json:"asks"`
	}

	var autoGenerated AutoGenerated
	json.Unmarshal(depthVar, &autoGenerated)
	fmt.Println("----------------------")
	fmt.Println("----------------------")
	fmt.Println("ASK:", autoGenerated.Asks[0].Price, "-", autoGenerated.Asks[0].Quantity, reflect.TypeOf(autoGenerated.Asks[0].Price))
	fmt.Println("BID:", autoGenerated.Bids[0].Price, "-", autoGenerated.Bids[0].Quantity)
	fmt.Println("----------------------")

	klines, err := futuresClient.NewKlinesService().Symbol("BTCUSDT").Interval("15m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	klinesVar, _ := json.Marshal(klines)

	type AutoGeneratedKlines []struct {
		OpenTime                 int64  `json:"openTime"`
		Open                     string `json:"open"`
		High                     string `json:"high"`
		Low                      string `json:"low"`
		Close                    string `json:"close"`
		Volume                   string `json:"volume"`
		CloseTime                int64  `json:"closeTime"`
		QuoteAssetVolume         string `json:"quoteAssetVolume"`
		TradeNum                 int    `json:"tradeNum"`
		TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
		TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
	}

	var autoGeneratedKlines AutoGeneratedKlines
	json.Unmarshal(klinesVar, &autoGeneratedKlines)

	// Make high slice float64
	var nextHighFloat float64
	var highSliceFloat64 []float64

	for l := 0; l < len(autoGeneratedKlines); l++ {
		nextHighString := autoGeneratedKlines[l].High
		if s2, err := strconv.ParseFloat(nextHighString, 32); err == nil {
			nextHighFloat = s2
			highSliceFloat64 = append(highSliceFloat64, nextHighFloat)
		}
	}

	max := highSliceFloat64[0]
	for _, number := range highSliceFloat64 {
		if number > max {
			max = number
		}
	}

	fmt.Println("Highest price   =", max)
	fmt.Println("----------------------")

	startLowString := autoGeneratedKlines[0].Low
	var startLowFloat float64
	if s, err := strconv.ParseFloat(startLowString, 32); err == nil {
		startLowFloat = s
	}
	var nextLowFloat float64
	var lowSliceFloat64 []float64
	lowSliceFloat64 = append(lowSliceFloat64, startLowFloat)

	for i := 1; i < len(autoGeneratedKlines); i++ {
		nextLowString := autoGeneratedKlines[i].Low
		if s1, err := strconv.ParseFloat(nextLowString, 32); err == nil {
			nextLowFloat = s1
			lowSliceFloat64 = append(lowSliceFloat64, nextLowFloat)
		}
	}

	min := lowSliceFloat64[0]
	for _, number := range lowSliceFloat64 {
		if number < min {
			min = number
		}
	}

	fmt.Println("Lowest price    =", min)
	fmt.Println("----------------------")

	longFib236 := max - ((max - min) * 0.236)
	fmt.Println("long Fibo 236 =", longFib236)
	longFib382 := max - ((max - min) * 0.382)
	fmt.Println("long Fibo 382 =", longFib382)
	longFib500 := max - ((max - min) * 0.500)
	fmt.Println("long Fibo 500 =", longFib500)
	longFib618 := max - ((max - min) * 0.618)
	fmt.Println("long Fibo 618 =", longFib618)
	longFib786 := max - ((max - min) * 0.786)
	fmt.Println("long Fibo 786 =", longFib786)
	fmt.Println("----------------------")

	var askPriceFloat float64

	if askPriceFloat, err = strconv.ParseFloat(autoGenerated.Asks[0].Price, 32); err != nil {
		fmt.Println(err)
	}

	priceAbove382 := askPriceFloat > longFib382 && askPriceFloat < longFib236
	fmt.Println("Price above the level 382 -", priceAbove382)
	priceAbove500 := askPriceFloat > longFib500 && askPriceFloat < longFib382
	fmt.Println("Price above the level 500 -", priceAbove500)
	priceAbove618 := askPriceFloat > longFib618 && askPriceFloat < longFib500
	fmt.Println("Price above the level 618 -", priceAbove618)
	priceAbove786 := askPriceFloat > longFib786 && askPriceFloat < longFib618
	fmt.Println("Price above the level 786 -", priceAbove786)
	fmt.Println("----------------------")

	var openPositionLevel int
	var countPriceLevel int

	if priceAbove382 {
		openPositionLevel = 382
		countPriceLevel++
	}
	if priceAbove500 {
		openPositionLevel = 500
		countPriceLevel++
	}
	if priceAbove618 {
		openPositionLevel = 618
		countPriceLevel++
	}
	if priceAbove786 {
		openPositionLevel = 786
		countPriceLevel++
	}
	if countPriceLevel != 1 {
		openPositionLevel = 0
	}
	fmt.Println("Setup position level -", openPositionLevel)
	ch <- openPositionLevel
}

func SetSetupOrderLevel(ch chan int) {
	setSetupOrderLevel := <-ch
	fmt.Println("Input chan for the var setSetupOrderLevel -", setSetupOrderLevel)
}
