package account

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/adshao/go-binance/v2"
)

func Account() {
	tickerName := "BTCUSDT"
	fmt.Println(tickerName, "- bot")

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

	resAcc, err := futuresClient.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(resAcc)

	accVar, _ := json.Marshal(resAcc)
	// fmt.Println(accVar)

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

	var account Account
	json.Unmarshal(accVar, &account)
	fmt.Println("----------------------")

	accountStart := 18.149229049682617 + 7.53667852 + 11.86
	accountNowString := account.AvailableBalance
	if accountNowFloat, err := strconv.ParseFloat(accountNowString, 32); err == nil {
		fmt.Println(accountStart, "- start")
		fmt.Println(accountNowFloat, "- now")
		fmt.Print("proffit($) = ", accountNowFloat-accountStart, "$", "\n")
		if accountNowFloat < accountStart {
			fmt.Print("proffit(%) = -", (accountNowFloat/accountStart)*100, "%")
		} else {
			fmt.Print("proffit(%) = ", (accountNowFloat/accountStart)*100, "%")
		}
	}
	fmt.Println()
}
