package main

import (
	"github.com/joho/godotenv"
	"log"

	"memogen/internal/service/image"
	"memogen/internal/service/telegram"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Читаем переменную TOKEN из среды
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("Переменная TOKEN не установлена")
	}

	imageService := image.NewImageService()
	telegramService := telegram.NewTelegramService(token, imageService)
	//httpServer := server.NewServer(imageService)

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Print("Telegram service starting...")
		telegramService.Start()
	}()

	go func() {
		//httpServer.Start()
	}()
	<-done

	log.Print("Telegram service stopping...")
	telegramService.Stop()
	log.Print("Telegram service stopped.")

}
