package structures

type Ingredient struct {
	Id   int //we want it to be id in our database, not in spoonacular
	Name string
}

type Recipe struct {
	Id                int
	Title             string
	Carbs             string
	Protein           string
	Calories          string
	UsedIngredients   []Ingredient
	MissedIngredients []Ingredient
}

type Request struct {
	Ingredients     []string
	NumberOfRecipes int
}
