package main

import (
	"log"

	"gomind/internal/config"
	"gomind/internal/controller"
	"gomind/internal/dao"
	"gomind/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load("config/config.toml")
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	mysqlDB, err := dao.InitMySQL(cfg.MySQL, cfg.MySQLDSN())
	if err != nil {
		log.Fatalf("init mysql failed: %v", err)
	}
	sqlDB, err := mysqlDB.DB()
	if err != nil {
		log.Fatalf("get mysql connection failed: %v", err)
	}
	defer sqlDB.Close()

	gin.SetMode(cfg.Server.Mode)
	router := gin.Default()

	healthService := service.NewHealthService()
	healthController := controller.NewHealthController(healthService)
	healthController.RegisterRoutes(router)

	if err := router.Run(cfg.Addr()); err != nil {
		log.Fatalf("start GoMind server failed: %v", err)
	}
}
