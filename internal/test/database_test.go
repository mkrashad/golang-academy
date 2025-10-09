package test

import (
	"golang-academy/internal"
	"golang-academy/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	db := &internal.Database{}

	movie1 := &entities.Movie{Title: "Lord of the Rings", Year: 2001}
	movie2 := &entities.Movie{Title: "Harry Potter", Year: 2001}
	character1 := &entities.Character{Name: "Frodo"}
	character2 := &entities.Character{Name: "Harry"}

	db.Create(movie1, character1)
	db.Create(movie2, character2)

	assert.Equal(t, 2, len(db.CharacterMovies))
	assert.Equal(t, "Frodo", db.CharacterMovies[0].Character.Name)
	assert.Equal(t, "Lord of the Rings", db.CharacterMovies[0].Movie.Title)

	newCharacterMovie := entities.CharacterMovie{
		Character: character2,
		Movie:     movie1,
	}
	err := db.Update(0, newCharacterMovie)
	assert.NoError(t, err)
	assert.Equal(t, "Harry", db.CharacterMovies[0].Character.Name)
	assert.Equal(t, "Lord of the Rings", db.CharacterMovies[0].Movie.Title)

	err = db.Update(5, newCharacterMovie)
	assert.Error(t, err)

	err = db.Delete(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(db.CharacterMovies))

	err = db.Delete(5)
	assert.Error(t, err)
}
