package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/AlastorTh/go_backend/database"
	"github.com/AlastorTh/go_backend/models"
)

func InsertTodo(db database.TodoInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := models.Todo{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
		}

		err = json.Unmarshal(body, &todo)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
		}

		res, err := db.Insert(todo)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
		}

		WriteResponse(w, http.StatusOK, res)
	}
}

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Methods, Access-Control-Allow-Origin")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	//w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
