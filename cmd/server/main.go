package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data": gin.H{
				"service": "gomind",
				"status":  "running",
			},
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("start GoMind server failed: %v", err)
	}
}
