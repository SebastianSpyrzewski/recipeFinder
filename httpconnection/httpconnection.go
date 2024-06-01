package httpconnection

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	st "recipeFinder/structures"
	"strings"
)

func httpGet(request string) []byte {
	resp, err := http.Get(request)
	if err != nil {
		fmt.Println("http communication error")
		println(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http communication error")
		println(err.Error())
		os.Exit(1)
	}
	return body
}

func AskAPI(req st.Request) []st.Recipe {
	request := "https://api.spoonacular.com/recipes/findByIngredients?apiKey=317fd2ab50974f61a73c23cb59ff3c6c&ranking=2&ingredients="
	request += strings.Join(req.Ingredients, ",+")
	request += "&number="
	request += fmt.Sprintf("%v", req.NumberOfRecipes)
	body := httpGet(request)
	var recipes []st.Recipe
	err := json.Unmarshal(body, &recipes)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	for i, r := range recipes {
		request = "https://api.spoonacular.com/recipes/" + fmt.Sprintf("%v", r.Id) + "/nutritionWidget.json?apiKey=317fd2ab50974f61a73c23cb59ff3c6c"
		body = httpGet(request)
		err := json.Unmarshal(body, &recipes[i])
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
	return recipes
}
