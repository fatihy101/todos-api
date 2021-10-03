package main

import (
	"fmt"
	"log"
	"net/http"

	"fatihy.space/todos-api/db"
	"fatihy.space/todos-api/routes"
)

// MongoDB Connection String
const connectionString = "mongodb+srv://dbUser:dbUserPass@cluster0.3bhbu.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func startServer() {
	port := ":4000"
	dbCon := db.OpenConnection(connectionString, "todoApp")
	server := &http.Server{
		Addr:    port,
		Handler: routes.Routes(dbCon),
	}

	fmt.Printf("Serving on %v\n", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(fmt.Sprintf("Error on listen and serve: %v", err))
	}
}
func main() {
	startServer()
}
