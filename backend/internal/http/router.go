package http

import (
	"github.com/azevedoMairon/decidr-app/internal/handlers"
	"github.com/azevedoMairon/decidr-app/internal/repositories"
	"github.com/azevedoMairon/decidr-app/internal/services"
	"github.com/azevedoMairon/decidr-app/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database) *gin.Engine {
	participantRepo := repositories.NewParticipantRepository(db)
	voteRepo := repositories.NewVoteRepository(db)

	participantService := services.NewParticipantService(participantRepo)
	voteService := services.NewVoteService(voteRepo, participantRepo)

	participantHandler := handlers.NewParticipantHandler(participantService)
	voteHandler := handlers.NewVoteHandler(voteService)

	router := gin.New()
	router.Use(middleware.GinLoggerJSON())
	router.Use(gin.Recovery())

	prometheus := ginprometheus.NewPrometheus("gin")
	prometheus.Use(router)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK!")
	})

	router.GET("/api/participants", participantHandler.GetParticipants)
	router.POST("/api/vote", voteHandler.PostVote)
	router.GET("/api/results", voteHandler.GetResults)

	return router
}
