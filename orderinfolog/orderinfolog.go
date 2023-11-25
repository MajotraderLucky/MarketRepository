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
	log.Println("orders.json file created")
	log.Println(openOrders)
}

// -------------------This function for the test---------------------------
type OpenOrder struct {
	OrderID string
	Symbol  string
}

type OrderInfoLogger interface {
	GetOpenOrders() ([]OpenOrder, error)
}

func GetOpenOrdersInfoJsonTest(orderService OrderInfoLogger, filePath string) error {
	apiKey, secretKey := GetAPIKeys()

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)
	openOrdersFromService, err := orderService.GetOpenOrders()
	if err != nil {
		log.Println(err)
		return err
	}

	openOrdersFromAPI, err := futuresClient.NewListOpenOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

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

	if err := encoder.Encode(openOrdersFromAPI); err != nil {
		log.Println(err)
		return err
	}

	if err := encoder.Encode(openOrdersFromService); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ---------------------------------------------------------------------

func CheckIfOpenOrdersExist() bool {
	apiKey, secretKey := GetAPIKeys()

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

	openOrders, err := futuresClient.NewListOpenOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return false
	}

	// If length of open orders is 0, returning false (No open orders)
	if len(openOrders) == 0 {
		log.Println("No open orders exist")
		return false
	}

	// If there are any open orders returning true
	log.Println("Open orders exist")
	return true
}

// --------------Test CheckIfOpenOrdersExist-------------------------

type ListOpenOrdersService interface {
	Do(ctx context.Context, opts ...binance.RequestOption) (res []*binance.Order, err error)
}

type BinanceService interface {
	NewListOpenOrdersService() ListOpenOrdersService
}

func CheckIfOpenOrdersExistTest(service BinanceService) bool {
	openOrders, err := service.NewListOpenOrdersService().Do(context.Background())
	if err != nil {
		log.Println(err)
		return false
	}

	if len(openOrders) > 0 {
		return true
	} else {
		return false
	}
}

// --------------------------------------------------------------------

func CheckIfOpenOrderOne() bool {
	apiKey, secretKey := GetAPIKeys()

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

	openOrders, err := futuresClient.NewListOpenOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return false
	}

	// If length of open orders is 0, returning false (No open orders)
	if len(openOrders) == 1 {
		log.Println("No open orders exist")
		return false
	}

	// If there are any open orders returning true
	log.Println("Open orders exist")
	return true
}
