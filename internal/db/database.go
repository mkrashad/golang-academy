package db

import (
	"fmt"
	"golang-academy/internal/entities"
)

type Database struct {
	CharacterMovies []entities.CharacterMovie
}

func New() *Database {
	return &Database{}
}

func (db *Database) Get() []entities.CharacterMovie {
	return db.CharacterMovies
}

func (db *Database) GetById(id int) (entities.CharacterMovie, error) {
	if id < 0 || id >= len(db.CharacterMovies) {
		return entities.CharacterMovie{}, fmt.Errorf("index out of bound")
	}
	return db.CharacterMovies[id], nil
}

func (db *Database) Create(movie *entities.Movie, character *entities.Character) {
	db.CharacterMovies = append(db.CharacterMovies, entities.CharacterMovie{
		Character: character,
		Movie:     movie,
	})
}

func (db *Database) Update(index int, c entities.CharacterMovie) error {
	if index < 0 || index >= len(db.CharacterMovies) {
		return fmt.Errorf("index out of bound")
	}
	db.CharacterMovies[index] = c
	return nil
}

func (db *Database) Delete(index int) error {
	if index < 0 || index >= len(db.CharacterMovies) {
		return fmt.Errorf("index out of bound")
	}
	db.CharacterMovies = append(db.CharacterMovies[:index], db.CharacterMovies[index+1:]...)
	return nil
}
