package main

import (
	"choice-tech-project/config"
	"choice-tech-project/internal/api"
	"choice-tech-project/internal/repository"
	"choice-tech-project/internal/service"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	mysqlRepo, err := repository.NewMySQLRepository(cfg.MySQLDSN)
	if err != nil {
		log.Fatalf("MySQL connection error: %v", err)
	}
	if err := mysqlRepo.CreateTable(); err != nil {
		log.Fatalf("Table creation error: %v", err)
	}

	redisRepo, err := repository.NewRedisRepository(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		log.Fatalf("Redis connection error: %v", err)
	}

	svc := service.NewService(mysqlRepo, redisRepo)
	handler := api.NewHandler(svc)
	router := api.SetupRouter(handler)

	log.Println("Server running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
