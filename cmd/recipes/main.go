package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/hellofreshdevtests/maga-golang-test/internal/domains"
	"github.com/hellofreshdevtests/maga-golang-test/internal/handlers/recipes"
)

func main() {
	repo, _ := url.Parse(os.Getenv("RECIPES_REPOSITORY"))

	recipesAdapter := domains.NewRecipesAdapter(repo)

	appPort := os.Getenv("APP_PORT")

	log.Printf("Starting Proxy Server on port %s ...", appPort)

	http.HandleFunc("/recipes", handlers.NewRecipesHandler(recipesAdapter))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
