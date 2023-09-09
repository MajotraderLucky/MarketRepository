package initlog

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/joho/godotenv"
)

// func Init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No.env file found.")
// 	}
// 	log.Println(".env file loaded.")

// 	apiKey, exists := os.LookupEnv("BINANCE_API_KEY")
// 	if !exists {
// 		log.Fatal("BINANCE_API_KEY not set")
// 	}

// 	secretKey, exexists := os.LookupEnv("BINANCE_SECRET_KEY")
// 	if !exexists {
// 		log.Fatal("BINANCE_SECRET_KEY not set")
// 	}

// 	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

// 	accountBalance, err := futuresClient.NewGetAccountService().Do(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	accVar, _ := json.Marshal(accountBalance)

// 	type Account struct {
// 		FeeTier                     int    `json:"feeTier"`
// 		CanTrade                    bool   `json:"canTrade"`
// 		CanDeposit                  bool   `json:"canDeposit"`
// 		CanWithdraw                 bool   `json:"canWithdraw"`
// 		UpdateTime                  int64  `json:"updateTime"`
// 		TotalInitialMargin          string `json:"totalInitialMargin"`
// 		TotalMaintMargin            string `json:"totalMaintMargin"`
// 		TotalWalletBalance          string `json:"totalWalletBalance"`
// 		TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
// 		TotalMarginBalance          string `json:"totalMarginBalance"`
// 		TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
// 		TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
// 		TotalCrossWalletBalance     string `json:"totalCrossWalletBalance"`
// 		TotalCrossUnPnl             string `json:"totalCrossUnPnl"`
// 		AvailableBalance            string `json:"availableBalance"`
// 		MaxWithdrawAmount           string `json:"maxWithdrawAmount"`
// 	}

// 	var acc Account
// 	json.Unmarshal(accVar, &acc)
// 	accountStart := 18.149229049682617 + 7.53667852 + 11.86 + 11.97 - 5
// 	accountNowString := acc.AvailableBalance
// 	if accountNowFloat, err := strconv.ParseFloat(accountNowString, 32); err == nil {
// 		log.Println("-------------------Balance usdt----------------------")
// 		log.Println(accountStart, "- start")
// 		log.Println(accountNowFloat, "- now")
// 		log.Println("-----------------------------------------------------")
// 	}
// }

// func NewFuturesClient() *futures.Client {
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No.env file found.")
// 	}
// 	log.Println(".env file loaded.")

// 	apiKey, exists := os.LookupEnv("BINANCE_API_KEY")
// 	if !exists {
// 		log.Fatal("BINANCE_API_KEY not set")
// 	}

// 	secretKey, exexists := os.LookupEnv("BINANCE_SECRET_KEY")
// 	if !exexists {
// 		log.Fatal("BINANCE_SECRET_KEY not set")
// 	}

// 	return binance.NewFuturesClient(apiKey, secretKey)
// }

type FuturesBinanceClient struct {
	futuresClient *futures.Client
}

type Account struct {
	FeeTier                     int    `json:"feeTier"`
	CanTrade                    bool   `json:"canTrade"`
	CanDeposit                  bool   `json:"canDeposit"`
	CanWithdraw                 bool   `json:"canWithdraw"`
	UpdateTime                  int64  `json:"updateTime"`
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
}

func NewFuturesClient() (*FuturesBinanceClient, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No.env file found.")
	}
	log.Println(".env file loaded.")

	apiKey, exists := os.LookupEnv("BINANCE_API_KEY")
	if !exists {
		log.Panicln("BINANCE_API_KEY not set")
	}

	secretKey, exexists := os.LookupEnv("BINANCE_SECRET_KEY")
	if !exexists {
		log.Panicln("BINANCE_SECRET_KEY not set")
	}

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)
	return &FuturesBinanceClient{futuresClient: futuresClient}, nil
}

func (f *FuturesBinanceClient) Init() error {
	accountBalance, err := f.futuresClient.NewGetAccountService().Do(context.Background())
	if err != nil {
		return err
	}
	accVar, _ := json.Marshal(accountBalance)

	var acc Account
	json.Unmarshal(accVar, &acc)
	accountStart := 18.149229049682617 + 7.53667852 + 11.86 + 11.97 - 5
	accountNowString := acc.AvailableBalance
	if accountNowFloat, err := strconv.ParseFloat(accountNowString, 32); err == nil {
		log.Println("-------------------Balance usdt----------------------")
		log.Println(accountStart, "- start")
		log.Println(accountNowFloat, "- now")
		log.Println("-----------------------------------------------------")
	}
	return nil
}
