package main

import (
	"fmt"
	"net/http"

	"github.com/AlastorTh/go_backend/config"
	"github.com/AlastorTh/go_backend/database"
	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	db := database.ConnectDB(conf.Mongo)

	fmt.Println(db)

	router := mux.NewRouter()
	http.ListenAndServe(":8080", router)
}
