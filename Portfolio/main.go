package main

import (
	"fmt"
	"net/http"
	"portfolio/database"
	"portfolio/routes"
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
	http.ListenAndServe(":8082", nil)
}
