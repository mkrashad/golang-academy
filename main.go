package main

import (
	"fmt"
	"golang-academy/models"
)

func displayCharactersMovies(db *models.Database) {
	for _, character := range db.Characters {
		fmt.Printf("Character: %s\n\n", character.Name)
		for _, link := range db.CharacterMovies {
			if link.CharacterId == character.CharacterId {
				for _, movie := range db.Movies {
					if movie.MovieId == link.MovieId {
						fmt.Printf("Movie: %s (%d)\n\n", movie.Title, movie.Year)
					}
				}
			}
		}
	}
}

func main() {

	db := &models.Database{}

	movie1 := &models.Movie{MovieId: 1, Title: "Lord of the rings", Year: 2001}
	movie2 := &models.Movie{MovieId: 2, Title: "Harry Potter", Year: 2001}
	movie3 := &models.Movie{MovieId: 3, Title: "Fight Club", Year: 1999}
	movie4 := &models.Movie{MovieId: 4, Title: "Batman Begins", Year: 2005}

	db.CreateMovie(*movie1)
	db.CreateMovie(*movie2)
	db.CreateMovie(*movie3)
	db.CreateMovie(*movie4)

	char1 := models.Character{CharacterId: 1, Name: "Frodo"}
	char2 := models.Character{CharacterId: 2, Name: "Harry"}
	char3 := models.Character{CharacterId: 3, Name: "Tyler"}
	char4 := models.Character{CharacterId: 4, Name: "Bruce"}

	db.CreateCharacter(char1)
	db.CreateCharacter(char2)
	db.CreateCharacter(char3)
	db.CreateCharacter(char4)

	db.LinkCharacterToMovie(char1.CharacterId, movie1.MovieId)
	db.LinkCharacterToMovie(char2.CharacterId, movie2.MovieId)
	db.LinkCharacterToMovie(char3.CharacterId, movie3.MovieId)
	db.LinkCharacterToMovie(char4.CharacterId, movie4.MovieId)

	fmt.Print("____Characters & Movies____\n\n")
	displayCharactersMovies(db)

	err := db.UpdateCharacter(char1.CharacterId, models.Character{CharacterId: 1, Name: "Gandalf"})

	if err != nil {
		fmt.Println("Error:", err)
	}
	err = db.UpdateMovie(movie1.MovieId, models.Movie{MovieId: 1, Title: "The Lord of the Rings: The Fellowship of the Ring", Year: 2001})
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("____After Update____\n")
	displayCharactersMovies(db)

	err = db.DeleteCharacter(char2.CharacterId)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = db.DeleteMovie(movie3.MovieId)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("____After delete____\n")
	displayCharactersMovies(db)

}
