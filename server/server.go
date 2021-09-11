package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux" // надеюсь, это не является "сторонним" решением. без проблем можно переписать и на дефолтный роутер из стд http
)

type Server struct {
	Storage *Storage
	Router  *mux.Router
}

func NewServer() *Server {
	storage := NewStorage(time.Hour*24, time.Second*5)
	server := &Server{
		Storage: storage,
		Router:  mux.NewRouter(),
	}
	return server
}
func (srv *Server) Start() error {
	srv.configureRouter()
	return http.ListenAndServe(":8080", srv.Router)
}
func (srv *Server) configureRouter() {
	http.Handle("/", srv.Router)
	srv.Router.HandleFunc("/note/all", srv.GetAllNotes).Methods(http.MethodGet)
	srv.Router.HandleFunc("/note/first", srv.GetFirstNote).Methods(http.MethodGet)
	srv.Router.HandleFunc("/note/last", srv.GetLastNote).Methods(http.MethodGet)
	srv.Router.HandleFunc("/note/delete/{id}", srv.DeleteNoteByID).Methods(http.MethodDelete)
	srv.Router.HandleFunc("/note/add", srv.AddNote).Methods(http.MethodPost)
}
