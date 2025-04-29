package main

import (
	"context"
	"log"

	"github.com/azevedoMairon/decidr-app/infra/mongo"
	"github.com/azevedoMairon/decidr-app/infra/mongo/migrations"
	"github.com/azevedoMairon/decidr-app/internal/handlers"
	"github.com/azevedoMairon/decidr-app/internal/repositories"
	"github.com/azevedoMairon/decidr-app/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoClient, err := mongo.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	db := mongoClient.Database("decidr_db")

	if err := migrations.SeedParticipants(context.Background(), db); err != nil {
		log.Fatal(err)
	}

	repo := repositories.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK!")
	})

	router.GET("/api/participants", handler.GetParticipants)

	router.Run(":8080")
}
