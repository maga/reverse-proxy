package domains

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type recipesAdapter struct {
	Repo url.URL
}

// Recipe is the recipe domain structure.
type Recipe struct {
	Id          string       `json:"id"`
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

// RecipesAdapter is the recipes repository.
type RecipesAdapter interface {
	FetchRecipes(req *http.Request) ([]*Recipe, error)
}

// NewRecipesAdapter creates a new repository instance.
func NewRecipesAdapter(repo *url.URL) RecipesAdapter {
	return recipesAdapter{
		Repo: *repo,
	}
}

func (ra recipesAdapter) FetchRecipes(req *http.Request) ([]*Recipe, error) {
	skip := getSkipParam(req)
	top := getTopParam(req)
	ids := getIdsParam(req)

	if len(ids) == 0 {
		for i := skip + 1; i <= skip+top; i++ {
			ids = append(ids, i)
		}
	}

	var recipes = []*Recipe{}
	recipes, err := fetchRecipesByIds(ids)
	if err != nil {
		fmt.Errorf("Error occured while requesting recipes by ids : [%v]", ids)
	}

	if len(ids) > 0 {
		Sort(recipes)
	}

	return recipes, nil
}

func fetchRecipesByIds(ids []int) ([]*Recipe, error) {
	concurrencyLimit, _ := strconv.Atoi(os.Getenv("CONCURRENCY_LIMIT"))

	// creating bounded channel to limit concurrent calls
	sem := make(chan struct{}, concurrencyLimit)
	// channel to collect the responses
	responseCh := make(chan *Recipe)
	// channel to collect the response. errors
	errorCh := make(chan error)

	// closing channels after they are used
	defer func() {
		close(sem)
		close(responseCh)
		close(errorCh)
	}()

	recipes := []*Recipe{}

	for _, id := range ids {
		go func(id int) {
			select {
			// stop routine if response is longer than 1s
			case <-time.After(1 * time.Second):
				errorCh <- errors.New("Timeout 1s")
				return
			case sem <- struct{}{}:
				recipe, err := fetchRecipeById(id)
				if err != nil {
					errorCh <- err
				} else {
					responseCh <- recipe
				}

				<-sem
			}
		}(id)
	}

	for range ids {
		select {
		case recipe := <-responseCh:
			recipes = append(recipes, recipe)
		case err := <-errorCh:
			log.Println("Error : ", err)
		}
	}

	return recipes, nil
}

func fetchRecipeById(id int) (*Recipe, error) {
	url := fmt.Sprintf("%s/%d", os.Getenv("RECIPES_REPOSITORY"), id)

	client := http.Client{
		Timeout: time.Duration(1 * time.Second),
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Errorf("Error requesting recipe %d : %v", id, err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("No recipe by id %d, status returned : %s", id, resp.Status)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("No recipe found by id %d", id)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Printf("Error reading response body by id %d : %v", id, err)
		return nil, err
	}

	recipe := &Recipe{}
	err = json.Unmarshal(body, recipe)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshal response body by id %d: %v", id, err)
	}

	return recipe, nil
}

func Sort(recipes []*Recipe) {
	sort.Slice(recipes, func(i, j int) bool {
		//If the prep. time is zero, push the recipe to the end of slice
		if recipes[i].PrepTime == "" {
			return false
		}
		if recipes[j].PrepTime == "" {
			return true
		}

		// RegExp to parse preparation time
		prepTimeRegExp := regexp.MustCompile(`^PT(\d+).*M$`)

		si := prepTimeRegExp.FindStringSubmatch(recipes[i].PrepTime)
		sj := prepTimeRegExp.FindStringSubmatch(recipes[j].PrepTime)

		// Extracting mins from the list got above
		ti, err := strconv.Atoi(si[1])
		if err != nil {
			return false
		}

		tj, err := strconv.Atoi(sj[1])
		if err != nil {
			return true
		}

		return ti < tj
	})
}

func getTopParam(req *http.Request) int {
	defaultTop, _ := strconv.Atoi(os.Getenv("DEFAULT_TOP"))
	topVal := defaultTop

	topKeys, ok := req.URL.Query()["top"]
	if ok || len(topKeys) > 0 {
		if top, err := strconv.Atoi(topKeys[0]); err == nil {
			topVal = top
		}
	}

	return topVal
}

func getIdsParam(req *http.Request) []int {
	ids, ok := req.URL.Query()["ids"]
	if !ok || len(ids[0]) < 1 {
		return []int{}
	}

	arr := []int{}
	vals := strings.Split(ids[0], ",")

	for _, v := range vals {
		id, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		arr = append(arr, id)
	}

	return arr
}

func getSkipParam(req *http.Request) int {
	defaultSkip, _ := strconv.Atoi(os.Getenv("DEFAULT_SKIP"))
	skipVal := defaultSkip

	skipKeys, ok := req.URL.Query()["skip"]
	if ok || len(skipKeys) > 0 {
		if skip, err := strconv.Atoi(skipKeys[0]); err == nil {
			skipVal = skip
		}
	}

	return skipVal
}
