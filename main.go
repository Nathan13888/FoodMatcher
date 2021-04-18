package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func displayResult(food string, foodType string) {
	var foodTypeColour *color.Color
	switch foodType {
	case string(APPETIZER):
		foodTypeColour = color.New(color.FgHiGreen).Add(color.Bold)
	case string(MAIN_COURSE):
		foodTypeColour = color.New(color.FgHiMagenta).Add(color.Bold)
	case string(DESSERT):
		foodTypeColour = color.New(color.FgHiYellow).Add(color.Bold)
	default:
		panic(errors.New("did not expect food type " + foodType))
	}
	color.New(color.FgHiCyan).Add(color.Bold).Print(strings.Title(food))
	fmt.Print(" is a ")
	foodTypeColour.Print(foodType)
	fmt.Println()
}

func main() {
	setupConfig()
	res, err := promptFood()
	if err != nil {
		panic(err)
	}

	foundFoods := search(res)
	if len(foundFoods) == 1 {
		displayResult(foundFoods[0], typesOfFoods[foundFoods[0]])
	} else if len(foundFoods) > 1 {
		color.Blue("\nFound more than one result.\n")
		res, err := promptSelectFood(foundFoods)
		if err != nil {
			panic(err)
		}
		displayResult(res, typesOfFoods[res])
	} else { // shouldn't get here under normal circumstances
		panic(errors.New("unexpected number of search results"))
	}
}
