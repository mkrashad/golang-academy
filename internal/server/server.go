package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, log *zap.Logger) (*echo.Echo, error) {
	e := echo.New()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("starting HTTP server on :8080")
			go func() {
				if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
					log.Fatal("server start failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("shutting down HTTP server")
			return e.Shutdown(ctx)
		},
	})
	return e, nil
}
