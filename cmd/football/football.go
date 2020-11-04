package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
)

type Player struct {
	ID string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
  }

var db *sql.DB
var err error

func getConfig() {

}

//Нужно будет потом раскидать функции на разные модули

func getPlayerList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	result, err := db.Query("SELECT id, firstname, lastname from players")
	if err != nil {
	  panic(err.Error())
	}

	defer result.Close()

	var players []Player
	for result.Next() {
	  var player Player
	  err := result.Scan(&player.ID, &player.Firstname, &player.Lastname)
	  if err != nil {
		panic(err.Error())
	  }
	  players = append(players, player)
	}
	json.NewEncoder(w).Encode(players)
}

func addNewPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	stmt, err := db.Prepare("INSERT INTO players(firstname, lastname) VALUES('abc','bcd')")
	if err != nil {
	  panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]

	_, err = stmt.Exec(title)
	if err != nil {
	  panic(err.Error())
	}

	fmt.Fprintf(w, "New player was added")
}

func removePlayer(w http.ResponseWriter, r *http.Request) {

}

func getNewLineUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test NewLineUp route")
}

func getManiskaWasher(w http.ResponseWriter, r *http.Request) {

}

func main() {
	db, err = sql.Open("sqlite3", "../test.db")
	if err != nil {
		panic(err.Error())
	}
    statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS players (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	
	r := mux.NewRouter()
	
	r.HandleFunc("/player", getPlayerList).Methods("GET")
	r.HandleFunc("/player", addNewPlayer).Methods("POST")
	r.HandleFunc("/player", removePlayer).Methods("DELETE")

	r.HandleFunc("/newgame", getNewLineUp).Methods("GET")
	r.HandleFunc("/endgame", getManiskaWasher).Methods("GET")

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
	
	defer db.Close()
}
