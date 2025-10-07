package models

type Database struct {
	Characters      []Character
	Movies          []Movie
	CharacterMovies []CharacterMovie
}

func (db *Database) CreateCharacter(c Character) {
	db.Characters = append(db.Characters, c)
}

func (db *Database) CreateMovie(m Movie) {
	db.Movies = append(db.Movies, m)
}

func (db *Database) LinkCharacterToMovie(characterId, movieId int) {
	db.CharacterMovies = append(db.CharacterMovies, CharacterMovie{
		CharacterId: characterId,
		MovieId:     movieId,
	})
}

func (db *Database) GetCharacters() []Character {
	return db.Characters
}

func (db *Database) GetMovies() []Movie {
	return db.Movies
}

func (db *Database) UpdateCharacter(Id int, c Character) error {
	for i, character := range db.Characters {
		if Id == character.CharacterId {
			db.Characters[i] = c
			break
		}
	}
	return nil
}

func (db *Database) UpdateMovie(Id int, m Movie) error {
	for i, movie := range db.Movies {
		if Id == movie.MovieId {
			db.Movies[i] = m
			break
		}
	}
	return nil
}

func (db *Database) DeleteCharacter(Id int) error {
	for i, character := range db.Characters {
		if Id == character.CharacterId {
			db.Characters = append(db.Characters[:i], db.Characters[i+1:]...)
			break
		}
	}
	newLinks := []CharacterMovie{}
	for _, link := range db.CharacterMovies {
		if Id != link.CharacterId {
			newLinks = append(newLinks, link)
		}
	}
	db.CharacterMovies = newLinks
	return nil
}

func (db *Database) DeleteMovie(Id int) error {
	for i, movie := range db.Movies {
		if Id == movie.MovieId {
			db.Movies = append(db.Movies[:i], db.Movies[i+1:]...)
			break
		}
	}
	newLinks := []CharacterMovie{}
	for _, link := range db.CharacterMovies {
		if Id != link.MovieId {
			newLinks = append(newLinks, link)
		}
		db.CharacterMovies = newLinks
	}
	return nil
}
