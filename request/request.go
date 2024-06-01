/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package request

import (
	"fmt"
	db "recipeFinder/database"
	http "recipeFinder/httpconnection"
	st "recipeFinder/structures"
	"strings"
)

func HandleRequest(ingredients []string, n int) {

	request := st.Request{Ingredients: ingredients, NumberOfRecipes: n}
	recipes := db.Search(request)
	if recipes == nil {
		recipes = http.AskAPI(request)
		db.Update(request, recipes)
	}

	for _, r := range recipes {
		fmt.Println(r.Title)
		fmt.Print("\tIngredients used: ")
		var used []string
		for _, ing := range r.UsedIngredients {
			used = append(used, ing.Name)
		}
		fmt.Println(strings.Join(used, ", "))
		fmt.Print("\tMissing ingredients: ")
		var missing []string
		for _, ing := range r.MissedIngredients {
			missing = append(missing, ing.Name)
		}
		fmt.Println(strings.Join(missing, ", "))
		fmt.Printf("\tCarbs: %s, Proteins: %s, Calories: %s\n", r.Carbs, r.Protein, r.Calories)
	}
}
