package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// HandleRoot возвращает HTML из файла index.html

func HandleRoot(res http.ResponseWriter, req *http.Request) {

	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(res, "file not found", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Write(data)

	log.Println("Сервер запущен на :8080")

}

func HandleUpload(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mr, err := req.MultipartReader()
	if err != nil {
		http.Error(res, "failed to read multipart data", http.StatusInternalServerError)
		return
	}

	part, err := mr.NextPart()
	if err != nil {
		http.Error(res, "no file in request", http.StatusBadRequest)
		return
	}
	defer part.Close()

	fileData, err := io.ReadAll(part)
	if err != nil {
		http.Error(res, "error reading file", http.StatusInternalServerError)
		return
	}

	convertedData, err := service.AutoConvert(string(fileData))
	if err != nil {
		http.Error(res, "error converting data", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(part.FileName())
	newFileName := fmt.Sprintf("file_%s%s", time.Now().UTC().Format("20060102_150405"), ext)

	if err := os.WriteFile(newFileName, []byte(convertedData), 0644); err != nil {
		http.Error(res, "failed to write file", http.StatusInternalServerError)
		return
	}

	log.Printf("Файл сохранен: %s", newFileName)
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Write([]byte(convertedData))
}
