package api

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-academy/internal/db"
	"golang-academy/internal/entities"
	"net"
	"net/http"
	"strconv"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Route interface {
	http.Handler
	Pattern() string
}

type MoviesHandler struct {
	log *zap.Logger
	db  *db.Database
}

func NewMovieHandler(log *zap.Logger, db *db.Database) *MoviesHandler {
	return &MoviesHandler{log: log, db: db}
}

func (*MoviesHandler) Pattern() string {
	return "/movie"
}

func (h *MoviesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodGet:
		movies := h.db.Get()
		if err := json.NewEncoder(w).Encode(movies); err != nil {
			h.log.Error("Failed to encode response", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

	case http.MethodPost:
		defer r.Body.Close()
		var movieCharacter entities.CharacterMovie
		if err := json.NewDecoder(r.Body).Decode(&movieCharacter); err != nil {
			h.log.Error("Failed to decode request body", zap.Error(err))
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		h.db.Create(movieCharacter.Movie, movieCharacter.Character)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Movie and character added successfully")

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, "Invalid or missing id parameter", http.StatusBadRequest)
			return
		}
		h.db.Delete(id)
		fmt.Fprintf(w, "Movie with ID %d deleted successfully\n", id)

	case http.MethodPut:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid or missing id parameter", http.StatusBadRequest)
			return
		}

		var movieCharacter entities.CharacterMovie
		if err := json.NewDecoder(r.Body).Decode(&movieCharacter); err != nil {
			h.log.Error("Failed to decode request body", zap.Error(err))
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}
		h.db.Update(id, movieCharacter)
		fmt.Fprintf(w, "Movie with ID %d deleted successfully\n", id)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}
