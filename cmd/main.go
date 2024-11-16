package main

import (
	doodocsrestapi "doodocs_rest_api"
	"doodocs_rest_api/internal/handler"
	"doodocs_rest_api/internal/service"
	"log"
)

func main() {
	services := service.NewService()
	handlers := handler.NewHandler(services)
	srv := new(doodocsrestapi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running  http server: %s", err.Error())
	}
}
