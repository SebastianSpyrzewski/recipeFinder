package structures

type Ingredient struct {
	Id   int
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
