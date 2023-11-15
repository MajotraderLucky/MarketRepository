package transactions

import (
	"context"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

func Hello() {
	log.Println("Startng transaction...")
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

func CreateLimitOrder(quantity string, price string) {
	apiKey, secretKey := GetAPIKeys()
	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

	limitOrder, err := futuresClient.NewCreateOrderService().Symbol("BTCUSDT").
		Side(futures.SideTypeBuy).Type(futures.OrderTypeLimit).
		TimeInForce(futures.TimeInForceTypeGTC).Quantity(quantity).
		Price(price).Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(limitOrder)
}
