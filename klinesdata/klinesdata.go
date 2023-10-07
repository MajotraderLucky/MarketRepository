package klinesdata

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"
)

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

// Convert milliseconds to time.Time
func MillisecondsToTime(milliseconds int64) time.Time {
	return time.Unix(0, milliseconds*int64(time.Millisecond))
}

func GetAPIKeys() (string, string) {
	apiKey, exists := os.LookupEnv("BINANCE_API_KEY")
	if exists {
		log.Println("apiKey checked")
	}

	secretKey, exexists := os.LookupEnv("BINANCE_SECRET_KEY")
	if exexists {
		log.Println("secretKey checked")
	}
	return apiKey, secretKey
}

func GetDebthData() {
	apiKey, secretKey := GetAPIKeys()

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)
	res, err := futuresClient.NewDepthService().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	depthVar, _ := json.Marshal(res)

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
	log.Println("----------------------")
	log.Println("ASK:", autoGenerated.Asks[0].Price, "-", autoGenerated.Asks[0].Quantity)
	log.Println("BID:", autoGenerated.Bids[0].Price, "-", autoGenerated.Bids[0].Quantity)
	log.Println("----------------------")
}

func KlinesInfo() {
	apiKey, secretKey := GetAPIKeys()

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

	klines, err := futuresClient.NewKlinesService().Symbol("BTCUSDT").Interval("15m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	klinesVar, _ := json.Marshal(klines)

	var autoGeneratedKlines AutoGeneratedKlines
	json.Unmarshal(klinesVar, &autoGeneratedKlines)
	t := MillisecondsToTime(autoGeneratedKlines[498].CloseTime)
	log.Println("Last kline:")
	log.Println(t)
	log.Println("15m open :", autoGeneratedKlines[498].Open)
	log.Println("15m close:", autoGeneratedKlines[498].Close)
	log.Println("15m high :", autoGeneratedKlines[498].High)
	log.Println("15m low  :", autoGeneratedKlines[498].Low)
	log.Println("----------------------")

	tStart := MillisecondsToTime(autoGeneratedKlines[0].CloseTime)
	log.Println("Start history:")
	log.Println(tStart)
	log.Println("15m open :", autoGeneratedKlines[0].Open)
	log.Println("15m close:", autoGeneratedKlines[0].Close)
	log.Println("15m high :", autoGeneratedKlines[0].High)
	log.Println("15m low  :", autoGeneratedKlines[0].Low)
	log.Println("----------------------")
}

func FindMinMaxInfo() {
	apiKey, secretKey := GetAPIKeys()

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

	klines, err := futuresClient.NewKlinesService().Symbol("BTCUSDT").Interval("15m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	klinesVar, _ := json.Marshal(klines)

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
	log.Println("----------------------")
	log.Println("Highest price   =", max)
	log.Println("----------------------")

	startLowString := autoGeneratedKlines[0].Low
	var startLowFloat float64
	if s, err := strconv.ParseFloat(startLowString, 32); err == nil {
		startLowFloat = s
	}
	// Make low slice float64
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

	log.Println("Lowest price    =", min)
	log.Println("----------------------")
}
