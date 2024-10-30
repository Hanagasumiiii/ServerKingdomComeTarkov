package ServerKingdomComeTarkov

import (
	"log"
	"os"
)

var (
	Logger *log.Logger
	file   *os.File
)

func Init() {
	var err error
	file, err = os.OpenFile("app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %v", err)
	}
	Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Close() {
	if err := file.Close(); err != nil {
		log.Fatalf("failed closing log file: %v", err)
	}
}
