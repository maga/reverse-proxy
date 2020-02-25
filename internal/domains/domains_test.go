package domains

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestFetchRecipeById(t *testing.T) {
	id := 1
	expected := seedRecipeJSON(id)

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

type RequestParam struct {
	Ids      string
	Expected []int
}

func TestGetIdsParam(t *testing.T) {
	urlString := "http://localhost:8080/recipes?ids="
	// url, _ := url.Parse("http://localhost:8080/recipes?ids=1,x,3,8.")

	params := []RequestParam{
		RequestParam{Ids: "1,2,3", Expected: []int{1, 2, 3}},
		RequestParam{Ids: "1,q,3", Expected: []int{1, 3}},
		RequestParam{Ids: "", Expected: []int{}},
	}

	for _, param := range params {
		url, _ := url.Parse(fmt.Sprintf("%s%s", urlString, param.Ids))

		req := &http.Request{URL: url}
		idsFromReq := getIdsParam(req)

		ok := reflect.DeepEqual(idsFromReq, param.Expected)
		if !ok {
			t.Errorf("Ids parsing. failed, got: %v, param.Expected: %v", idsFromReq, param.Expected)
		}
	}
}

func TestSort(t *testing.T) {
	recipes := []*Recipe{
		&Recipe{Id: "1", PrepTime: "PT5M"},
		&Recipe{Id: "2", PrepTime: "PT4M"},
		&Recipe{Id: "3", PrepTime: "PT2-3M"},
		&Recipe{Id: "4", PrepTime: "PT1M"},
		&Recipe{Id: "5", PrepTime: ""},
	}

	Sort(recipes)

	expected := []string{"4", "3", "2", "1", "5"}

	for i, recipe := range recipes {
		if recipe.Id != expected[i] {
			t.Errorf("Recipe is sorted wrongly, got: %s, want: %s.", recipe.Id, expected[i])
		}
	}
}
