package domain

type recipeType int

const (
	pinRecipe recipeType = iota
	userRecipe
)

// RecipeIngredient is the link object between a recipe and an ingredient
type RecipeIngredient struct {
	ID           int
	IngredientID int
	Amount       float32
	Unit         int
}

// Recipe is the top level object that defines a recipe
type Recipe struct {
	ID                  int
	Name                string
	Type                RecipeType
	RecipeIngredientIDs []int
}
