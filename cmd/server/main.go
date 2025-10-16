package main

import (
	"log"
	"os"

	"golang-academy/internal/certificates"
	"golang-academy/internal/client"
	"golang-academy/internal/db"
	"golang-academy/internal/handlers"
	"golang-academy/internal/server"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	certsDir := os.Getenv("CERTS_DIR")
	if _, err := os.Stat(certsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(certsDir, 0755); err != nil {
			log.Fatalf("Failed to create CERTS_DIR: %v", err)
		}
	}

	caCertPath := os.Getenv("CA_CERT_PATH")
	caKeyPath := os.Getenv("CA_KEY_PATH")

	os.Stat(caCertPath)
	os.Stat(caKeyPath)

	certificates.GenerateCA()

	app := fx.New(
		fx.Provide(
			server.NewHTTPServer,
			db.New,
			client.NewClient,
			zap.NewExample,
		),
		fx.Invoke(handlers.RegisterRoutes),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
	app.Run()
}
