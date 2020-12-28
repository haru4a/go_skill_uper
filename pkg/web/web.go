package web

import (
	"net/http"
)

//Start function start web gui
func Start() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.ListenAndServe(":80", nil)
}
