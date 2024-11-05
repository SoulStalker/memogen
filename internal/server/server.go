package main

import (
	"log"
	"memogen/internal/service/image"
	"memogen/internal/service/telegram"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	imageService := image.NewImageService()
	telegramService := telegram.NewTelegramService("", imageService)
	httpServer := server.NewServer(imageService)

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Print("Telegram service starting...")
		telegramService.Start()
	}()

	go func() {
		httpServer.Start()
	}()
	<-done

	log.Print("Telegram service stopping...")
	telegramService.Stop()
	log.Print("Telegram service stopped.")

}
