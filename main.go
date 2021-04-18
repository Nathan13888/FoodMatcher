package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func displayResult(food string, foodType string) {
	var foodTypeColour *color.Color
	switch foodType { // change the color based on the type of food
	case string(APPETIZER):
		foodTypeColour = color.New(color.FgHiGreen).Add(color.Bold)
	case string(MAIN_COURSE):
		foodTypeColour = color.New(color.FgHiMagenta).Add(color.Bold)
	case string(DESSERT):
		foodTypeColour = color.New(color.FgHiYellow).Add(color.Bold)
	default:
		panic(errors.New("did not expect food type " + foodType))
	}

	// DISPLAY THE RESULT using various colours and formatting
	color.New(color.FgHiCyan).Add(color.Bold).Print(strings.Title(food))
	fmt.Print(" is a ")
	foodTypeColour.Print(foodType)
	fmt.Println()
}

// program starts HERE
func main() {
	setupConfig()            // configures defaults, setups up Viper (library for configs), loads configs
	res, err := promptFood() // prompt user to enter a food
	if err != nil {          // handle the error if there is one
		panic(err)
	}

	foundFoods := search(res) // search for the inputed food
	if len(foundFoods) == 1 { // display the result if there is only one result
		displayResult(foundFoods[0], typesOfFoods[foundFoods[0]])
	} else if len(foundFoods) > 1 { // allow the user to specify the food they want to know about
		color.Blue("\nFound more than one result.\n")
		res, err := promptSelectFood(foundFoods) // opens the prompt to select the available search results
		if err != nil {
			panic(err)
		}
		displayResult(res, typesOfFoods[res]) // displays the selected result
	} else { // shouldn't get here under normal circumstances - promptFood() filters invalid results
		panic(errors.New("unexpected number of search results"))
	}
}
