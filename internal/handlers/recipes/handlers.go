package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hellofreshdevtests/maga-golang-test/internal/domains"
)

// closure for returned function that handles the request
func NewRecipesHandler(recipesAdapter domains.RecipesAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)

			return
		}

		recipesHandler(w, r, recipesAdapter)
	}
}

func recipesHandler(w http.ResponseWriter, r *http.Request, adapter domains.RecipesAdapter) {
	var err error
	var recipes []*domains.Recipe

	recipes, err = adapter.FetchRecipes(r)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	json, err := json.Marshal(recipes)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	// Return recipes in JSON
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
