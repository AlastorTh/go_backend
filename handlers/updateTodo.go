package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/AlastorTh/go_backend/database"
	"github.com/gorilla/mux"
)

func UpdateTodo(db database.TodoInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		var todo interface{}
		err = json.Unmarshal(body, &todo)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := db.Update(id, todo)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}

// func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	json.NewEncoder(w).Encode(res)
// }
