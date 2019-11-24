package proxyserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Start() {
	appPort := os.Getenv("APP_PORT")

	log.Printf("Starting Proxy Server on port %s ...", appPort)

	http.HandleFunc("/recipes", handleRecipes)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}

func handleRecipes(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("handler")
}
