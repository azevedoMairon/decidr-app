package main

import (
	"context"
	"log"

	"github.com/azevedoMairon/decidr-app/infra/mongo"
	"github.com/azevedoMairon/decidr-app/infra/mongo/migrations"
	"github.com/azevedoMairon/decidr-app/internal/participants"
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

	repo := participants.NewRepository(db)
	service := participants.NewService(repo)
	handler := participants.NewHandler(service)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK!")
	})

	router.GET("/api/participants", handler.GetParticipants)

	router.Run(":8080")
}
