package domains

import (
	"fmt"
	"log"
)

// Recipe is the recipe model structure.
type Recipe struct {
  ID          string       `json:"id"`
  Name        string       `json:"name"`
  Headline    string       `json:"headline"`
  Description string       `json:"description"`
  Difficulty  int          `json:"difficulty"`
  PrepTime    string       `json:"prepTime"`
  ImageLink   string       `json:"imageLink"`
  Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
  Name      string `json:"name"`
  ImageLink string `json:"imageLink"`
}

// NewRecipesAdapter creates a new repository instance.
func NewRecipesAdapter(repo string) RecipesAdapter {
  return recipesAdapter{
    repo: repo
  }
}

type recipesAdapter struct {
  repo string
}

// RecipesAdapter is the recipes repository.
type RecipesAdapter interface {
  FetchByIds(ids []string) (Recipe, error)
  FetchAll() ([]Recipe, error)
}

func (r recipesAdapter) FetchAll() ([]Recipe, error) {

}

func (r recipesAdapter) FetchByIds(ids []string) ([]Recipe, error) {

}
