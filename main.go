package main

import (
	"context"
	"net/http"

	"github.com/AlastorTh/go_backend/config"
	"github.com/AlastorTh/go_backend/database"
	"github.com/AlastorTh/go_backend/handlers"
	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)
	client := &database.TodoClient{
		Col: collection,
		Ctx: ctx,
	}

	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	router.HandleFunc("/task", handlers.SearchTodo(client)).Methods("GET", "OPTIONS")
	router.HandleFunc("/task/{id}", handlers.GetTodo(client)).Methods("GET", "OPTIONS")
	router.HandleFunc("/task", handlers.InsertTodo(client)).Methods("POST", "OPTIONS")
	router.HandleFunc("/task/{id}", handlers.UpdateTodo(client)).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/task/{id}", handlers.DeleteTodo(client)).Methods("DELETE", "OPTIONS")
	http.ListenAndServe(":8080", router)
}
