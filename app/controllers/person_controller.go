package controllers

import (
	"encoding/json"
	"go-sql-crud-api/app/models"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) People(w http.ResponseWriter, r *http.Request) {
	var personModel models.Person
	var result *[]models.Person

	result, err := personModel.GetPeople(server.Context, server.DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	jsonData, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (server *Server) Person(w http.ResponseWriter, r *http.Request) {
	var personModel models.Person
	var result *models.Person

	vars := mux.Vars(r)

	if vars["id"] == "" {
		http.Error(w, "Please provide an id of a person", http.StatusBadRequest)

		return
	}

	result, err := personModel.GetPerson(server.Context, server.DB, vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	jsonData, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (server *Server) AddPerson(w http.ResponseWriter, r *http.Request) {
	var personModel models.Person

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_ = json.Unmarshal(body, &personModel)

	if personModel.Name == "" || personModel.Age == 0 {
		http.Error(w, "Nama dan usia tidak boleh kosong", http.StatusBadRequest)

		return
	}

	err = personModel.CreatePerson(server.Context, server.DB, &personModel)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (server *Server) EditPerson(w http.ResponseWriter, r *http.Request) {
	var newData, mp models.Person

	vars := mux.Vars(r)

	if vars["id"] == "" {
		http.Error(w, "Please provide an id of a person", http.StatusBadRequest)

		return
	}

	_, err := mp.GetPerson(server.Context, server.DB, vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_ = json.Unmarshal(body, &newData)

	if newData.Name == "" || newData.Age == 0 {
		http.Error(w, "Nama dan usia tidak boleh kosong", http.StatusBadRequest)

		return
	}

	err = mp.UpdatePerson(server.Context, server.DB, vars["id"], &newData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (server *Server) DeletePerson(w http.ResponseWriter, r *http.Request) {
	var data, mp models.Person
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_ = json.Unmarshal(body, &data)

	if data.ID == "" {
		http.Error(w, "id tidak boleh kosong", http.StatusBadRequest)

		return
	}

	err = mp.RemovePerson(server.Context, server.DB, data.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
