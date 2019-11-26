package domains

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "net/http"
	// "net/http/httptest"
	// "os"
	"testing"
	// "github.com/hellofreshdevtests/maga-golang-test/internal/domains"
	// "github.com/stretchr/testify/assert"
)

func TestFetchRecipeById(t *testing.T) {
	id := 1
	expected := seedRecipeJSON(id)

	log.Println(expected.Id)

	res, err := fetchRecipeById(id)
	if err != nil {
		t.Errorf("Error requesting recipe %d : %v", id, err)
	}

	if res.Id != expected.Id {
		t.Errorf("Incorrect recipe id. Expected %s, got %s", expected.Id, res.Id)
	}
	if res.Name != expected.Name {
		t.Errorf("Incorrect recipe name. Expected %s, got %s", expected.Name, res.Name)
	}
}

func seedRecipeJSON(id int) Recipe {
	byteVal, err := ioutil.ReadFile(fmt.Sprintf("../seeds/recipe_%d.json", id))
	if err != nil {
		fmt.Println(err)
	}

	var recipe Recipe

	json.Unmarshal(byteVal, &recipe)

	return recipe
}
