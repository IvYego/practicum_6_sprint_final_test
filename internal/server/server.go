package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type ServerMain struct {git add .
	Logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) error {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("/upload", handlers.HandleUpload)

	srv := &http.Server{
		Addr:         ":8080",          // Порт
		Handler:      mux,              // Наш роутер
		ErrorLog:     logger,           // Логгер
		ReadTimeout:  5 * time.Second,  // Таймаут чтения
		WriteTimeout: 10 * time.Second, // Таймаут записи
		IdleTimeout:  15 * time.Second, // Таймаут простоя
	}

	logger.Println("Сервер запущен на порту 8080")
	return srv.ListenAndServe()

}
