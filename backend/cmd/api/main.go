package main

import (
	"context"
	"os"

	"log/slog"

	"github.com/azevedoMairon/decidr-app/internal/http"
	"github.com/azevedoMairon/decidr-app/internal/infra/mongo"
	"github.com/azevedoMairon/decidr-app/internal/infra/mongo/migrations"
	"github.com/azevedoMairon/decidr-app/pkg/logger"
)

func main() {
	logger.Init()

	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		slog.Error("MONGO_URL not defined")
		return
	}

	client, err := mongo.Connect(mongoURI)
	if err != nil {
		slog.Error("Erro ao conectar no MongoDB", "err", err)
		return
	}
	db := client.Database("decidr_db")

	if err := migrations.SeedParticipants(context.Background(), db); err != nil {
		slog.Error("Erro ao rodar seed", "err", err)
		return
	}

	router := http.NewRouter(db)
	slog.Info("Servidor iniciado", "port", 8080)
	router.Run(":8080")
}
