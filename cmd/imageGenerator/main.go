package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ptsgr/ImageGeneratorBot/pkg/http"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error loading configs: %s", err.Error())
	}

	handlers := http.InitHandlers()
	server := new(http.Server)

	go func() {
		if err := server.Run(viper.GetString("imageGeneratorPort"), handlers); err != nil {
			log.Fatalf("Error imageGenerator running http server: %s", err.Error())
		}
	}()

	log.Println("imageGenerator server started on port: ", viper.GetString("imageGeneratorPort"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("imageGenerator server Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error imageGenerator server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}
