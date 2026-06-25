package main

import (
	"context"
	"log"

	"gomind/internal/config"
	"gomind/internal/controller"
	"gomind/internal/dao"
	"gomind/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

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

	redisClient, err := dao.InitRedis(ctx, cfg.Redis, cfg.RedisAddr())
	if err != nil {
		log.Fatalf("init redis failed: %v", err)
	}
	defer redisClient.Close()

	milvusClient, err := dao.InitMilvus(ctx, cfg.Milvus, cfg.MilvusAddr())
	if err != nil {
		log.Fatalf("init milvus failed: %v", err)
	}
	defer func() {
		if err := milvusClient.Close(ctx); err != nil {
			log.Printf("close milvus client failed: %v", err)
		}
	}()

	gin.SetMode(cfg.Server.Mode)
	router := gin.Default()

	healthService := service.NewHealthService()
	healthController := controller.NewHealthController(healthService)
	healthController.RegisterRoutes(router)

	if err := router.Run(cfg.Addr()); err != nil {
		log.Fatalf("start GoMind server failed: %v", err)
	}
}
