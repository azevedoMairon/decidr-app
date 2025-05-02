package main

import (
	"context"
	"log"

	"github.com/azevedoMairon/decidr-app/infra/mongo"
	"github.com/azevedoMairon/decidr-app/infra/mongo/migrations"
	"github.com/azevedoMairon/decidr-app/internal/handlers"
	"github.com/azevedoMairon/decidr-app/internal/repositories"
	"github.com/azevedoMairon/decidr-app/internal/services"
	"github.com/gin-contrib/cors"

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

	participantRepo := repositories.NewParticipantRepository(db)
	participantService := services.NewParticipantService(participantRepo)
	participantHandler := handlers.NewParticipantHandler(participantService)

	voteRepo := repositories.NewVoteRepository(db)
	voteService := services.NewVoteService(voteRepo, participantRepo)
	voteHandler := handlers.NewVoteHandler(voteService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK!")
	})

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	router.GET("/api/participants", participantHandler.GetParticipants)

	router.POST("/api/vote", voteHandler.PostVote)

	router.Run(":8080")
}
