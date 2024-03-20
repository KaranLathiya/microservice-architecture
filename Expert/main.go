package main

import (
	"expert/database"
	"expert/routes"
	"fmt"
	"net/http"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Server started")
	router := routes.InitializeRouter()
	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)

}
