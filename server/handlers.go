package server

import (
	"encoding/json"
	"io/ioutil"
	"kode-task/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (srv *Server) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes := srv.Storage.GetAll()
	result := reverse(notes)
	data, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Cant marshal JSON"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
func (srv *Server) AddNote(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Bad model"))
		return
	}
	insNote := &models.Note{}
	err = json.Unmarshal(body, insNote)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Cant unmarshal JSON"))
		return
	}
	srv.Storage.Add(insNote.Value, time.Second*time.Duration(insNote.Expiration))
	w.WriteHeader(200)
	w.Write([]byte("Added"))
}
func (srv *Server) DeleteNoteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := vars["id"]; !ok {
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}
	srv.Storage.DeleteNoteByID(vars["id"])
	w.WriteHeader(200)
	w.Write([]byte("Deleted"))
}
func (srv *Server) GetLastNote(w http.ResponseWriter, r *http.Request) {
	note := srv.Storage.GetLastNote()
	data, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Cant marshal JSON"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
func (srv *Server) GetFirstNote(w http.ResponseWriter, r *http.Request) {
	note := srv.Storage.GetFirstNote()
	data, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Cant marshal JSON"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
