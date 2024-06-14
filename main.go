package main

import (
	"GOLANG-DEV-LOGIC-CHALLENGE-NOPECODE96/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRouter()
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
