package handlers

import (
	"http"
)

// closure for returned function that handles the request
func NewRecipesHandler(recipesAdapter domains.RecipesAdapter) http.HandlerFunc {
	if req.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	return func(w http.ResponseWriter, r *http.Request) {
		recipes, err := recipesHandler(req, companiesAdapter)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Return recipes slice as a JSON response
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(recipes)
	}
}

func recipesHandler(request *http.Request, companiesAdapter domains.CompaniesAdapter) (types.Response, error) {
	idParam := r.URL.Query()["ids"]
	//Check if "ids" are passed as query parameters
	if idParam != nil {
		ids := strings.Split(idParam[0], ",")
		h.getRecipesByIds(w, r, ids)
	} else {
		h.getAllRecipes(w, r)
	}
}
