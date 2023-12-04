package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/robertogsf/CRUD-Go/database"
	"github.com/robertogsf/CRUD-Go/handlers"
)

func main() {

	database.ConnectDB()

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", handlers.GetTasks)

	fmt.Println("Running in port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
