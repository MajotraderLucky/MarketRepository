package connect

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func GetApi() {
	for range time.Tick(time.Second * 30) {
		tickerName := "BTCUSDT"

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
		res, err := futuresClient.NewDepthService().Symbol(tickerName).Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	}
}
