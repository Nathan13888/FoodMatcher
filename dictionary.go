package main

import (
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type FoodType string

const (
	APPETIZER   FoodType = "🥗 Appetizer"
	MAIN_COURSE FoodType = "🍔 Main Course"
	DESSERT     FoodType = "🧁 Dessert"
)

var typesOfFoods = make(map[string]string)

func search(qry string) []string {
	var res []string
	for food := range typesOfFoods {
		if fuzzy.MatchNormalizedFold(qry, food) {
			res = append(res, food)
		}
	}

	return res
}
