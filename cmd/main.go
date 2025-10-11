package main

import (
	"golang-academy/internal/api"
	"golang-academy/internal/db"
	"golang-academy/internal/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)



func main() {
	app := fx.New(
		fx.Provide(
			server.NewHTTPServer,
			db.New,
			zap.NewExample,
		),
		fx.Invoke(api.RegisterRoutes),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
	app.Run()
}
