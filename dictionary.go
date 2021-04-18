package main

import (
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type FoodType string

const ( // possible types of food
	APPETIZER   FoodType = "🥗 Appetizer"
	MAIN_COURSE FoodType = "🍔 Main Course"
	DESSERT     FoodType = "🧁 Dessert"
)

// a hashmap that relates a food to its type
var typesOfFoods = make(map[string]string)

// search `typesOfFoods` for a particular food
func search(qry string) []string {
	var res []string                 // foods that match `qry`
	for food := range typesOfFoods { // loop through all foods
		if fuzzy.MatchNormalizedFold(qry, food) { // checks if food matches (Fuzzy Search)
			res = append(res, food)
		}
	}

	return res // returns all the matching foods
}
