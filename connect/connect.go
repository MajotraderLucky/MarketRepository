package connect

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Hello2() {
	fmt.Println("Hello from connect")
}

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func GetApi() {
	fmt.Println("----------------------")
	apiKey, exists := os.LookupEnv("BINANCE_API_KEY")
	if exists {
		fmt.Println("apiKey exist")
	}

	secretKey, exexists := os.LookupEnv("BINANCE_SECRET_KEY")
	if exexists {
		fmt.Println("secretKey exist")
	}
}
