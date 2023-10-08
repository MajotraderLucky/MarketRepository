package klinesdata

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func FindMinMaxInfo() (float64, float64, error) {
	apiKey, secretKey := GetAPIKeys()

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

	klines, err := futuresClient.NewKlinesService().Symbol("BTCUSDT").Interval("15m").Do(context.Background())
	if err != nil {
		return 0, 0, err
	}
	klinesVar, _ := json.Marshal(klines)

	var autoGeneratedKlines AutoGeneratedKlines
	json.Unmarshal(klinesVar, &autoGeneratedKlines)

	var nextHighFloat, max, min, nextLowFloat float64
	var highSliceFloat64, lowSliceFloat64 []float64

	for l := 0; l < len(autoGeneratedKlines); l++ {
		nextHighString := autoGeneratedKlines[l].High
		if s2, err := strconv.ParseFloat(nextHighString, 32); err == nil {
			nextHighFloat = s2
			highSliceFloat64 = append(highSliceFloat64, nextHighFloat)
		}
		if number := highSliceFloat64[l]; number > max {
			max = number
		}
	}
	log.Println("----------------------")
	log.Println("Highest price   =", max)
	log.Println("----------------------")

	for i := 0; i < len(autoGeneratedKlines); i++ {
		nextLowString := autoGeneratedKlines[i].Low
		if s1, err := strconv.ParseFloat(nextLowString, 32); err == nil {
			nextLowFloat = s1
			lowSliceFloat64 = append(lowSliceFloat64, nextLowFloat)
		}
		if number := lowSliceFloat64[i]; number < min || i == 0 {
			min = number
		}
	}
	log.Println("Lowest price    =", min)
	log.Println("----------------------")

	return max, min, nil
}

// The data refers to the FindMinMaxInfo function and the subsequent
// test of this function

type KlinesServiceGetter interface {
	NewKlinesService() KlinesService
}

type Kline struct {
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

type RequestOption func(*http.Request)

type KlinesService interface {
	Symbol(symbol string) KlinesService
	Interval(interval string) KlinesService
	Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error)
}

func FindMinMaxInfoTest(client KlinesServiceGetter) (float64, float64, error) {
	klines, err := client.
		NewKlinesService().
		Symbol("BTCUSDT").
		Interval("15m").
		Do(context.Background())
	if err != nil {
		return 0, 0, err
	}
	klinesVar, _ := json.Marshal(klines)

	var autoGeneratedKlines AutoGeneratedKlines
	json.Unmarshal(klinesVar, &autoGeneratedKlines)

	var nextHighFloat, max, min, nextLowFloat float64
	var highSliceFloat64, lowSliceFloat64 []float64

	for l := 0; l < len(autoGeneratedKlines); l++ {
		nextHighString := autoGeneratedKlines[l].High
		if s2, err := strconv.ParseFloat(nextHighString, 32); err == nil {
			nextHighFloat = s2
			highSliceFloat64 = append(highSliceFloat64, nextHighFloat)
		}
		if number := highSliceFloat64[l]; number > max {
			max = number
		}
	}
	log.Println("----------------------")
	log.Println("Highest price   =", max)
	log.Println("----------------------")

	for i := 0; i < len(autoGeneratedKlines); i++ {
		nextLowString := autoGeneratedKlines[i].Low
		if s1, err := strconv.ParseFloat(nextLowString, 32); err == nil {
			nextLowFloat = s1
			lowSliceFloat64 = append(lowSliceFloat64, nextLowFloat)
		}
		if number := lowSliceFloat64[i]; number < min || i == 0 {
			min = number
		}
	}
	log.Println("Lowest price    =", min)
	log.Println("----------------------")

	return max, min, nil
}

func GetFibonacciLevels() {
	max, min, err := FindMinMaxInfo()
	if err != nil {
		log.Fatalf("Error getting min and max info: %v", err)
		return
	}

	longFib236 := max - ((max - min) * 0.236)
	log.Println("long Fibo 236 =", longFib236)
	longFib382 := max - ((max - min) * 0.382)
	log.Println("long Fibo 382 =", longFib382)
	longFib500 := max - ((max - min) * 0.500)
	log.Println("long Fibo 500 =", longFib500)
	longFib618 := max - ((max - min) * 0.618)
	log.Println("long Fibo 618 =", longFib618)
	longFib786 := max - ((max - min) * 0.786)
	log.Println("long Fibo 786 =", longFib786)
	log.Println("----------------------")
	return
}
