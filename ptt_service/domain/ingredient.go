package domain

// IngredientType is an enum of types of ingredients
type ingredientType int

const (
	// SingleUse symbolizes something that disappears on use
	SingleUse ingredientType = iota
	// PantryItem defines an item that may be used multiple times before being used up
	PantryItem
)

type ingredientUnit int

const (
	dash ingredientUnit = iota
	pinch
	teaspoon
	tablespoon
	cup
	pint
	quart
	gallon
	ounce
	pound
	milliliter
	liter
	deciliter
	milligram
	gram
	kilogram
)

// Ingredient is the main data structure that defines an element of a recipe
type Ingredient struct {
	ID         int
	Name       string
	Type       ingredientType
	Perishable bool
}

// IngredientService defines the operations available to store and retrieve ingredients
type IngredientService interface {
	Ingredient(id int) (*Ingredient, error)
	CreateIngredient(ingredient *Ingredient) error
	DeleteIngredient(ingredient *Ingredient) error
}
