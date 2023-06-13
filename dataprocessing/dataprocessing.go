package dataprocessing

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
)

func TestHello() {
	fmt.Println("Hello dataprocessing!")
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Convert milliseconds to time.Time
func MillisecondsToTime(milliseconds int64) time.Time {
	return time.Unix(0, milliseconds*int64(time.Millisecond))
}

func DataGenerator(ch chan string) {
	fmt.Println("----------------------")
	apiKey, exists := os.LookupEnv("BINANCE_API_KEY")
	if exists {
		fmt.Println("apiKey exist")
	}

	secretKey, exexists := os.LookupEnv("BINANCE_SECRET_KEY")
	if exexists {
		fmt.Println("secretKey exist")
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
	fmt.Println("ASK:", autoGenerated.Asks[0].Price, "-", autoGenerated.Asks[0].Quantity)
	fmt.Println("BID:", autoGenerated.Bids[0].Price, "-", autoGenerated.Bids[0].Quantity)
	fmt.Println("----------------------")

	klines, err := futuresClient.NewKlinesService().Symbol("BTCUSDT").Interval("1h").Do(context.Background())
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
	fmt.Println("----------------------")
	fmt.Println("----------------------")
	t := MillisecondsToTime(autoGeneratedKlines[498].CloseTime)
	fmt.Println("Last kline from dataprocessing:")
	fmt.Println(t)
	fmt.Println("----------------------")
	fmt.Println("Type var -", reflect.TypeOf(autoGeneratedKlines[498].Low))

	setString := autoGeneratedKlines[498].Low
	ch <- setString
}

func SetTestVarString(ch chan string) {
	setString := <-ch // получаем значение из канала
	fmt.Println("Input chan for the var setString", setString)
}
