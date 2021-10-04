package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"fatihy.space/todos-api/db"
	"fatihy.space/todos-api/routes"
)

// MongoDB Connection String
const connectionString = "mongodb+srv://dbUser:dbUserPass@cluster0.3bhbu.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func startServer(isTest bool) {
	dbName, port := getInitVars(isTest)
	dbCon := db.OpenConnection(connectionString, dbName)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: routes.Routes(dbCon),
	}

	fmt.Printf("Serving on :%v\n", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(fmt.Sprintf("Error on listen and serve: %v", err))
	}
}
func main() {
	startServer(false)
}

func getInitVars(isTest bool) (string, string) {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "4000"
	}
	if isTest {
		fmt.Println("Starting server in DEV mode")
		return "test", "4000"
	} else {
		return "todoApp", port
	}
}
