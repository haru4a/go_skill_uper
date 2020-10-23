package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/player", getPlayerList).Methods("GET")
    r.HandleFunc("/player", addNewPlayer).Methods("POST")
    r.HandleFunc("/player", removePlayer).Methods("DELETE")

    r.HandleFunc("/newgame", getNewLineUp).Methods("GET")
    r.HandleFunc("/endgame", getManiskaWasher).Methods("GET")



}

func getConfig(){
    
}

//Нужно будет потом раскидать функции на разные модули

func getPlayerList(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Test Player route")
    return
}
func addNewPlayer(w http.ResponseWriter, r *http.Request){
    return
}

func removePlayer(w http.ResponseWriter, r *http.Request){
    return
}



func getNewLineUp(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Test NewLineUp route")
    return
}

func getManiskaWasher(w http.ResponseWriter, r *http.Request){
    return
}
