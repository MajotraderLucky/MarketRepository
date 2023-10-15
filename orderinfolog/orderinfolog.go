package orderinfolog

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
)

func Hello() {
	log.Println("---------------------------------------")
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

func GetOpenOrdersInfoJson() {
	apiKey, secretKey := GetAPIKeys()

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

	openOrders, err := futuresClient.NewListOpenOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	filePath := "logs/orders.json"

	// Create a file for writing
	file, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// Use json.NewEncoder to write each value to the JSON file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t") // for pretty JSON output

	if err := encoder.Encode(openOrders); err != nil {
		log.Println(err)
	}
}

func GetOpenOrdersInfoJsonTest(futuresClient *binance.FuturesClient) error {
	openOrders, err := futuresClient.NewListOpenOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	filePath := "logs/orders.json"

	// Create a file for writing
	file, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	// Use json.NewEncoder to write each value to the JSON file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t") // for pretty JSON output

	if err := encoder.Encode(openOrders); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
