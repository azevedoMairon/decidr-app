package main

import (
	"context"
	"log"

	"github.com/azevedoMairon/decidr-app/infra/mongo"
	"github.com/azevedoMairon/decidr-app/infra/mongo/migrations"
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

	router := gin.Default()

	router.Run(":8080")
}
