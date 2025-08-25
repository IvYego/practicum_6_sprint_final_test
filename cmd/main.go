package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)
	
	if err := server.NewServer(logger); err != nil {
		logger.Fatal(err)
	}

}
