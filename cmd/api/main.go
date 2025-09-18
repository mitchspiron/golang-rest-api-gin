package main

import (
	"database/sql"
	"golang-rest-api-gin/internal/database"
	"golang-rest-api-gin/internal/env"
	"log"

	_ "golang-rest-api-gin/docs"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// @title Go REST API with Gin
// @version 1.0
// @description This is a sample server for a REST API built with Go and Gin.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// Apply the security definition to your endpoints
// @security BearerAuth

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 4000),
		jwtSecret: env.GetEnvString("JWT_SECRET", "supersecretkey"),
		models:    models,
	}

	if err := serve(app); err != nil {
		log.Fatal(err)
	}
}
