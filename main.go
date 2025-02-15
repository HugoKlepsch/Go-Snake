package main

import (
	"fmt"
	"github.com/DrakeEsdon/Go-Snake/api"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", api.HandleIndex)
	http.HandleFunc("/latestlog", api.HandleLatestLog)
	http.HandleFunc("/start", api.HandleStart)
	http.HandleFunc("/move", api.HandleMove)
	http.HandleFunc("/end", api.HandleEnd)

	fmt.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
