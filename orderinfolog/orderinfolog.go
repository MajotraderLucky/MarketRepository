package orderinfolog

import (
	"context"
	"encoding/json"
	"io"
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
		log.Println("Open orders exist")
		return true
	}

	// If there are any open orders returning true
	log.Println("No open orders exist")
	return false
}

func CheckStopMarketOrders(r io.Reader) bool {
	// Декодируем json в слайс структур
	var orders []struct {
		Type string `json:"type"`
	}

	// Устанавливаем позицию чтения в начало
	if seeker, ok := r.(io.Seeker); ok {
		_, err := seeker.Seek(0, io.SeekStart)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := json.NewDecoder(r).Decode(&orders); err != nil {
		log.Fatal(err)
	}

	// Ищем ордер с типом "STOP_MARKET"
	for _, order := range orders {
		if order.Type == "STOP_MARKET" {
			return true // Если нашли, то возвращаем true
		}
	}

	// Если не нашли ни одного ордера с типом "STOP_MARKET", то возвращаем false
	return false
}

func CheckTakeProfitMarketOrders(r io.Reader) bool {
	// Декодируем json в слайс структур
	var orders []struct {
		Type string `json:"type"`
	}

	// Устанавливаем позицию чтения в начало
	if seeker, ok := r.(io.Seeker); ok {
		_, err := seeker.Seek(0, io.SeekStart)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := json.NewDecoder(r).Decode(&orders); err != nil {
		log.Fatal(err)
	}

	// Ищем ордер с типом "TAKE_PROFIT_MARKET"
	for _, order := range orders {
		if order.Type == "TAKE_PROFIT_MARKET" {
			return true // Если нашли, то возвращаем true
		}
	}

	// Если не нашли ни одного ордера с типом "TAKE_PROFIT_MARKET", то возвращаем false
	return false
}

func CheckLimitOrders(r io.Reader) bool {
	// Декодируем json в слайс структур
	var orders []struct {
		Type string `json:"type"`
	}

	// Устанавливаем позицию чтения в начало
	if seeker, ok := r.(io.Seeker); ok {
		_, err := seeker.Seek(0, io.SeekStart)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := json.NewDecoder(r).Decode(&orders); err != nil {
		log.Fatal(err)
	}

	// Ищем ордер с типом "TAKE_PROFIT_MARKET"
	for _, order := range orders {
		if order.Type == "LIMIT" {
			return true // Если нашли, то возвращаем true
		}
	}

	// Если не нашли ни одного ордера с типом "TAKE_PROFIT_MARKET", то возвращаем false
	return false
}

func CreateOrdersConfigFile() {
	// Получение текущей директории
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Ошибка при получении текущей директории:", err)
	}

	// Полный путь к директории "configurations"
	configurationsDir := currentDir + "/configurations"

	// Проверка существования директории "configurations"
	if _, err := os.Stat(configurationsDir); os.IsNotExist(err) {
		// Создание директории "configurations" с правами записи
		err := os.Mkdir(configurationsDir, 0700)
		if err != nil {
			log.Fatal("Не удалось создать директорию 'configurations':", err)
		}
	}

	// Полный путь к файлу "ordersconfig.json"
	ordersConfigFile := configurationsDir + "/ordersconfig.json"

	// Проверка существования файла "ordersconfig.json"
	if _, err := os.Stat(ordersConfigFile); os.IsNotExist(err) {
		// Создание файла "ordersconfig.json" с правами записи и чтения
		file, err := os.OpenFile(ordersConfigFile, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			log.Fatal("Не удалось создать файл 'ordersconfig.json':", err)
		}
		defer file.Close()

		// Запись пустого JSON объекта в файл
		_, err = file.Write([]byte("{}"))
		if err != nil {
			log.Fatal("Не удалось записать данные в файл 'ordersconfig.json':", err)
		}
	}

	log.Println("Файл ordersconfig.json успешно создан")
}
