package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hellofreshdevtests/maga-golang-test/internal/domains"
	"github.com/hellofreshdevtests/maga-golang-test/internal/handlers"
	"github.com/hellofreshdevtests/maga-golang-test/internal/repository"
)

func main() {
	repo := "https://s3-eu-west-1.amazonaws.com/test-golang-recipes/"

	recipesAdapter := domains.NewRecipesAdapter(repo)

	appPort := os.Getenv("APP_PORT")
	log.Printf("Starting Proxy Server on port %s ...", appPort)

	http.HandleFunc("/recipes", handlers.NewRecipesHandler(recipesAdapter))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
