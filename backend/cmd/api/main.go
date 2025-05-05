package main

import (
	"context"
	"log"
	"os"

	"github.com/azevedoMairon/decidr-app/infra/mongo"
	"github.com/azevedoMairon/decidr-app/infra/mongo/migrations"
	"github.com/azevedoMairon/decidr-app/internal/handlers"
	"github.com/azevedoMairon/decidr-app/internal/repositories"
	"github.com/azevedoMairon/decidr-app/internal/services"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		log.Fatal("MONGO_URL not defined")
	}

	mongoClient, err := mongo.Connect(mongoURI)
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
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	router.GET("/api/participants", participantHandler.GetParticipants)

	router.POST("/api/vote", voteHandler.PostVote)

	router.GET("/api/results", voteHandler.GetResults)

	router.Run(":8080")
}
