package apiserver

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go_skill_uper/pkg/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router  *mux.Router
	logger  *logrus.Logger
	storage *storage.Storage
	config  *Config
}

func newServer(storage *storage.Storage, config *Config) *server {
	s := &server{
		router:  mux.NewRouter(),
		logger:  logrus.New(),
		storage: storage,
		config:  config,
	}

	s.configureRouter()

	return s
}

// Start ...
func Start(config *Config, storage *storage.Storage) error {
	srv := newServer(storage, config)
	return http.ListenAndServe(config.BindAddr, srv.router)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/player", s.getPlayerList).Methods("GET")
	s.router.HandleFunc("/player", s.addNewPlayer).Methods("POST")

	//s.router.HandleFunc("/player", s.removePlayer).Methods("DELETE")
	//s.router.HandleFunc("/newgame", s.getNewLineUp).Methods("GET")
	//s.router.HandleFunc("/endgame", s.getManiskaWasher).Methods("GET")

}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) getPlayerList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := s.storage.GetList()
	json.NewEncoder(w).Encode(result)
}

func (s *server) addNewPlayer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	if err := s.storage.AddPlayer(storage.Player{ID: keyVal["ID"], Firstname: keyVal["Firstname"], Lastname: keyVal["Lastname"]}); err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
}

/*
func (s *server) removePlayer(w http.ResponseWriter, r *http.Request) {

}

func (s *server) getNewLineUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test NewLineUp route")
}

func (s *server) getManiskaWasher(w http.ResponseWriter, r *http.Request) {

}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.storage.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

*/
