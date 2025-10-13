package handlers

import (
	"golang-academy/internal/db"
	"golang-academy/internal/entities"
	"golang-academy/internal/generated"
	"strconv"

	//"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type MovieHandler struct {
	log *zap.Logger
	db  *db.Database
}

func NewMovieHandler(log *zap.Logger, db *db.Database) *MovieHandler {
	return &MovieHandler{log: log, db: db}
}

func (h *MovieHandler) GetMovie(c echo.Context) error {
	movies := h.db.GetAll()
	return c.JSON(200, movies)
}

func (h *MovieHandler) GetMovieId(c echo.Context, id int) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid or missing id parameter"})
	}
	movie, err := h.db.GetById(id)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "No movie with this id exists"})
	}
	return c.JSON(200, movie)
}

func (h *MovieHandler) PostMovie(c echo.Context) error {
	var movieCharacter entities.CharacterMovie
	if err := c.Bind(&movieCharacter); err != nil {
		h.log.Error("Failed to bind request body", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Invalid JSON body"})
	}

	h.db.Create(movieCharacter.Movie, movieCharacter.Character)
	return c.JSON(201, map[string]string{"message": "Movie and character added successfully"})
}

func (h *MovieHandler) DeleteMovieId(c echo.Context, id int) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid or missing id parameter"})
	}
	h.db.Delete(id)
	return c.JSON(200, map[string]string{"message": "Movie with ID " + idStr + " deleted successfully"})
}

func (h *MovieHandler) PutMovieId(c echo.Context, id int) error {
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

func RegisterRoutes(e *echo.Echo, log *zap.Logger, db *db.Database) {
	handler := NewMovieHandler(log, db)
	// swagger, err := generated.GetSwagger()
	// if err != nil {
	// 	log.Fatal("Failed to load OpenAPI spec", zap.Error(err))
	// }

	//  Validation middleware
	//e.Use(middleware.OapiRequestValidator(swagger))

	generated.RegisterHandlers(e, handler)
}
