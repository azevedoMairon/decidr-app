package logger

import (
	"io"
	"log/slog"
	"os"
)

func Init() {
	_ = os.MkdirAll("./logs", os.ModePerm)
	logFile, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Erro ao abrir arquivo de log: " + err.Error())
	}
	handler := slog.NewJSONHandler(io.MultiWriter(os.Stdout, logFile), &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	slog.SetDefault(slog.New(handler))
	slog.Info("Logger inicializado", "destino", "./logs/app.log")
}
