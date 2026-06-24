package main

import (
	"log"

	"gomind/internal/controller"
	"gomind/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	healthService := service.NewHealthService()
	healthController := controller.NewHealthController(healthService)
	healthController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("start GoMind server failed: %v", err)
	}
}
