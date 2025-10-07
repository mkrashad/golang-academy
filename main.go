package main

import (
	"fmt"
	"golang-academy/internal"
	"golang-academy/internal/entities"
)

func displayCharactersMovies(db *internal.Database) {
	for _, cm := range db.CharacterMovies {
		fmt.Printf("Movie: %s Year: (%d) Character: %s\n\n", cm.Movie.Title, cm.Movie.Year, cm.Character.Name)
	}
}

func main() {
	db := &internal.Database{}

	movie1 := &entities.Movie{Title: "Lord of the rings", Year: 2001}
	movie2 := &entities.Movie{Title: "Harry Potter", Year: 2001}
	movie3 := &entities.Movie{Title: "Fight Club", Year: 1999}
	movie4 := &entities.Movie{Title: "Batman Begins", Year: 2005}

	char1 := &entities.Character{Name: "Frodo"}
	char2 := &entities.Character{Name: "Harry"}
	char3 := &entities.Character{Name: "Tyler"}
	char4 := &entities.Character{Name: "Bruce"}

	db.Create(movie1, char1)
	db.Create(movie2, char2)
	db.Create(movie3, char3)
	db.Create(movie4, char4)

	fmt.Print("____Characters & Movies____\n\n")
	displayCharactersMovies(db)

	db.Update(0, entities.CharacterMovie{Movie: movie1, Character: char2})
	fmt.Println("____After Update____\n")
	displayCharactersMovies(db)

	db.Delete(2)
	fmt.Println("____After delete____\n")
	displayCharactersMovies(db)
}
