package httpconnection

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	st "recipeFinder/structures"
	"strings"

	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {

	err := godotenv.Load("recipeFinder.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

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
	godotenv.Load()

	apiKey := getEnvVariable("API_KEY")
	request := "https://api.spoonacular.com/recipes/findByIngredients?apiKey=" + apiKey + "&ranking=2&ingredients="
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
		request = "https://api.spoonacular.com/recipes/" + fmt.Sprintf("%v", r.Id) + "/nutritionWidget.json?apiKey=" + apiKey
		body = httpGet(request)
		err := json.Unmarshal(body, &recipes[i])
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
	return recipes
}
