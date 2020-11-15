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
	database, err := sql.Open(dbType, dbPath)
	if err != nil {
		panic(err.Error())
	}
	//praparate table for players list
	database.Exec("CREATE TABLE IF NOT EXISTS players(id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	return &Storage{
		db: database,
	}
}

// AddPlayer ...
func (s *Storage) AddPlayer(data Player) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO players(firstname, lastname) VALUES(?,?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(data.Firstname, data.Lastname)
	id, _ := result.LastInsertId()

	return id, err

}

// RemovePlayer ...
func (s *Storage) RemovePlayer(data Player) (int64, error) {
	stmt, err := s.db.Prepare("DELETE FROM players WHERE id = ?")

	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(data.ID)

	id, _ := result.RowsAffected()

	return id, err

}

// GetList ...
func (s *Storage) GetList() []Player {
	rows, err := s.db.Query(`SELECT id, firstname, lastname from players`)
	if err != nil {
		panic(err.Error())
	}
	var players []Player
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.ID, &player.Firstname, &player.Lastname)
		if err != nil {
			panic(err.Error())
		}
		players = append(players, player)
	}

	return players
}
