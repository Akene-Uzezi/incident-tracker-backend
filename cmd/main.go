package main

import (
	"log"

	"issueTracking/internal/db"
	"issueTracking/internal/env"
	"issueTracking/internal/logger"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	logger.InitLogger()

	pool, err := db.InitPool()
	if err != nil {
		log.Fatalf("Failed to initialize database connection pool: %v", err)
	}
	defer pool.Close()
	models := db.NewModels(pool)
	app := &application{
		port:      env.GetEnvInt("PORT", 3001),
		jwtsecret: env.GetEnvString("jwtSecret", "someSecret"),
		db:        pool,
		models:    models,
		origins:   env.GetEnvString("allowedOrigins", "http://localhost:3000,http://192.168.9.227:3000"),
	}
	app.serve()
}
