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

type Output struct {
	Orders    []Order `json:"orders"`
	NumOpened int     `json:"numOpened"`
}

func GetOpenOrdersInfoJson() {
	apiKey, secretKey := GetAPIKeys()

	futuresClient := binance.NewFuturesClient(apiKey, secretKey)

	// Предположим, что функция NewListOpenOrdersService() возвращает []Order
	openOrders, err := futuresClient.NewListOpenOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	filePath := "logs/orders.json"

	// Создайте файл для записи
	file, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// Создайте экземпляр Output и добавьте заказы и количество открытых заказов
	output := &Output{
		Orders:    openOrders,
		NumOpened: len(openOrders),
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t") // для красивого вывода JSON

	if err := encoder.Encode(output); err != nil {
		log.Println(err)
	}
}
