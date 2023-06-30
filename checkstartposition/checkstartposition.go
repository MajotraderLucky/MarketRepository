package checkstartposition

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Convert milliseconds to time.Time
func MillisecondsToTime(milliseconds int64) time.Time {
	return time.Unix(0, milliseconds*int64(time.Millisecond))
}

func Checkstartposition() {
	fmt.Println("-------------------------")
	fmt.Println("Hello checkstartposition!")
	fmt.Println("-------------------------")

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

	priceCorridor := max - min
	priceCorridorPercent := (priceCorridor / max) * 100
	priceCorridorCondition := priceCorridorPercent > 7
	priceCorridorPercentRound := fmt.Sprintf("%.2f", priceCorridorPercent)
	fmt.Print("Price corridor condition - ", priceCorridorCondition, " - ", priceCorridorPercentRound, "%")
	fmt.Println()

	// Check open positions

	accServ, err := futuresClient.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	accServVar, _ := json.Marshal(accServ)
	// fmt.Println(accServVar, reflect.TypeOf(accServVar))

	fileJson, err := json.Marshal(accServ)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("fileJson.json", fileJson, 0644)
	if err != nil {
		panic(err)
	}

	type AutoGeneratedPos struct {
		Assets []struct {
			Asset                  string `json:"asset"`
			InitialMargin          string `json:"initialMargin"`
			MaintMargin            string `json:"maintMargin"`
			MarginBalance          string `json:"marginBalance"`
			MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
			OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
			PositionInitialMargin  string `json:"positionInitialMargin"`
			UnrealizedProfit       string `json:"unrealizedProfit"`
			WalletBalance          string `json:"walletBalance"`
		} `json:"assets"`
		FeeTier                     int    `json:"feeTier"`
		CanTrade                    bool   `json:"canTrade"`
		CanDeposit                  bool   `json:"canDeposit"`
		CanWithdraw                 bool   `json:"canWithdraw"`
		UpdateTime                  int    `json:"updateTime"`
		TotalInitialMargin          string `json:"totalInitialMargin"`
		TotalMaintMargin            string `json:"totalMaintMargin"`
		TotalWalletBalance          string `json:"totalWalletBalance"`
		TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
		TotalMarginBalance          string `json:"totalMarginBalance"`
		TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
		TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
		TotalCrossWalletBalance     string `json:"totalCrossWalletBalance"`
		TotalCrossUnPnl             string `json:"totalCrossUnPnl"`
		AvailableBalance            string `json:"availableBalance"`
		MaxWithdrawAmount           string `json:"maxWithdrawAmount"`
		Positions                   []struct {
			Isolated               bool   `json:"isolated"`
			Leverage               string `json:"leverage"`
			InitialMargin          string `json:"initialMargin"`
			MaintMargin            string `json:"maintMargin"`
			OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
			PositionInitialMargin  string `json:"positionInitialMargin"`
			Symbol                 string `json:"symbol"`
			UnrealizedProfit       string `json:"unrealizedProfit"`
			EntryPrice             string `json:"entryPrice"`
			MaxNotional            string `json:"maxNotional"`
			PositionSide           string `json:"positionSide"`
			PositionAmt            string `json:"positionAmt"`
			Notional               string `json:"notional"`
			IsolatedWallet         string `json:"isolatedWallet"`
			UpdateTime             int64  `json:"updateTime"`
		} `json:"positions"`
	}

	var autoGeneratedpos AutoGeneratedPos
	json.Unmarshal(accServVar, &autoGeneratedpos)

	var positionBTCindex int

	for k := 0; k < len(autoGeneratedpos.Positions); k++ {
		if autoGeneratedpos.Positions[k].Symbol == "BTCUSDT" {
			positionBTCindex = k
		}
	}
	positionStr := autoGeneratedpos.Positions[positionBTCindex].PositionAmt
	fmt.Println("Position size", positionStr)
	fmt.Println("Type -", reflect.TypeOf(positionStr))

	positionFloat, err := strconv.ParseFloat(positionStr, 64)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}
	fmt.Printf("%v", positionFloat, reflect.TypeOf(positionFloat))
	fmt.Println()
	fmt.Println("----------------------")
}
