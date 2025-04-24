package controllers

import "github.com/gorilla/mux"

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()

	server.Router.HandleFunc("/people", server.People).Methods("GET")
	server.Router.HandleFunc("/person/{id}", server.Person).Methods("GET")
	server.Router.HandleFunc("/people", server.AddPerson).Methods("POST")
	server.Router.HandleFunc("/person/{id}", server.EditPerson).Methods("PUT")
	server.Router.HandleFunc("/person", server.DeletePerson).Methods("DELETE")
}
