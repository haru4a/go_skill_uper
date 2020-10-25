package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getConfig() {

}

//Нужно будет потом раскидать функции на разные модули

func getPlayerList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Player route")
}

func addNewPlayer(w http.ResponseWriter, r *http.Request) {

}

func removePlayer(w http.ResponseWriter, r *http.Request) {

}

func getNewLineUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test NewLineUp route")
}

func getManiskaWasher(w http.ResponseWriter, r *http.Request) {

}

func main() {
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

}
