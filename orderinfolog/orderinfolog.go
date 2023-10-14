package orderinfolog

import (
	"context"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
)

func Hello() {
	log.Println("orderinfolog package has been started")
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

func GetOpenOrdersInfo() {
	apiKey, secretKey := GetAPIKeys()

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

	openOrders, err := futuresClient.NewListOpenOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	for _, o := range openOrders {
		log.Println(o)
		log.Println(len(openOrders), "orders have been opened")
	}
}
