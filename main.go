package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/miguelToscano/restaurants-microservice/adapters/inbound/rest"
	"github.com/miguelToscano/restaurants-microservice/adapters/inbound/tui"
)

const (
	TUI_MODE  = "TUI"
	REST_MODE = "REST"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	runMode := os.Getenv("RUN_MODE")

	switch runMode {
	case TUI_MODE:
		tui.Start()
	case REST_MODE:
		rest.Start()
	default:
		fmt.Println("Invalid RUN_MODE env variable")
	}

	// rest.Start()
	tui.Start()
}
