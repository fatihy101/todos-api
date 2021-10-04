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

func startServer(isTest bool) {
	dbName := getDBName(isTest)
	port := ":4000"
	dbCon := db.OpenConnection(connectionString, dbName)
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
	startServer(false)
}

func getDBName(isTest bool) string {
	if isTest {
		fmt.Println("Starting server in DEV mode")
		return "test"
	} else {
		return "todoApp"
	}
}
