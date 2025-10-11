package api

import (
	"golang-academy/internal/db"
	"golang-academy/internal/entities"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type MoviesHandler struct {
	log *zap.Logger
	db  *db.Database
}

func NewMovieHandler(log *zap.Logger, db *db.Database) *MoviesHandler {
	return &MoviesHandler{log: log, db: db}
}

func (h *MoviesHandler) GetMovies(c echo.Context) error {
	movies := h.db.Get()
	return c.JSON(200, movies)
}

func (h *MoviesHandler) CreateMovie(c echo.Context) error {
	var movieCharacter entities.CharacterMovie
	if err := c.Bind(&movieCharacter); err != nil {
		h.log.Error("Failed to bind request body", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Invalid JSON body"})
	}

	h.db.Create(movieCharacter.Movie, movieCharacter.Character)
	return c.JSON(201, map[string]string{"message": "Movie and character added successfully"})
}

func (h *MoviesHandler) DeleteMovie(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid or missing id parameter"})
	}
	h.db.Delete(id)
	return c.JSON(200, map[string]string{"message": "Movie with ID " + idStr + " deleted successfully"})
}

func (h *MoviesHandler) UpdateMovie(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid or missing id parameter"})
	}

	var movieCharacter entities.CharacterMovie
	if err := c.Bind(&movieCharacter); err != nil {
		h.log.Error("Failed to bind request body", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Invalid JSON body"})
	}

	h.db.Update(id, movieCharacter)
	return c.JSON(200, map[string]string{"message": "Movie with ID " + idStr + " updated successfully"})
}

func RegisterRoutes(e *echo.Echo, log *zap.Logger, db *db.Database){
	handler := NewMovieHandler(log, db)
	e.GET("/movie", handler.GetMovies)
	e.POST("/movie", handler.CreateMovie)
	e.DELETE("/movie/:id", handler.DeleteMovie)
	e.PUT("/movie/:id", handler.UpdateMovie)
}
