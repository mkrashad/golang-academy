package main

import (
	"golang-academy/internal/api"
	"golang-academy/internal/db"
	"net/http"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(api.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func main() {
	app := fx.New(
		fx.Provide(
			api.NewHTTPServer,
			fx.Annotate(
				api.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			AsRoute(api.NewMovieHandler),
			db.New,
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
	app.Run()
}

// [
//   {
//     "Character": {
//       "Name": "Frodo"
//     },
//     "Movie": {
//       "Title": "Lord of the rings",
//       "Year": 2001
//     }
//   },
//   {
//     "Character": {
//       "Name": "Bruce"
//     },
//     "Movie": {
//       "Title": "Batman Begins",
//       "Year": 2005
//     }
//   }
// ]
