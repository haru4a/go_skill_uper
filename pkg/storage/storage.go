package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Storage ...
type Storage struct {
	db *sql.DB
}

// New ...
func New(dbType string, dbPath string) *Storage {
	result, err := sql.Open(dbType, dbPath)
	if err != nil {
		panic(err.Error())
	}
	return &Storage{
		db: result,
	}
}

// AddPlayer ...
func (s *Storage) AddPlayer(data Player) error {
	stmt, err := s.db.Prepare("INSERT INTO players(firstname, lastname) VALUES('abc','bcd')")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err

}

// GetList ...
func (s *Storage) GetList() []Player {
	result, err := s.db.Query(`SELECT id, firstname, lastname from players`)
	if err != nil {
		panic(err.Error())
	}
	var players []Player
	for result.Next() {
		var player Player
		err := result.Scan(&player.ID, &player.Firstname, &player.Lastname)
		if err != nil {
			panic(err.Error())
		}
		players = append(players, player)
	}

	return players
}
