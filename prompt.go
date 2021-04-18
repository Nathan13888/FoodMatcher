package main

import (
	"errors"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/manifoldco/promptui"
)

func promptFood() (string, error) {
	// validates if an inputed food exists
	validate := func(input string) error {
		if len(search(input)) > 0 {
			return nil
		}
		return errors.New("invalid input")
	}

	// settings for the prompt
	prompt := promptui.Prompt{
		Label:    "Enter a food",
		Validate: validate,
	}

	// open the prompt
	res, err := prompt.Run()

	return res, err // return information about the prompt
}

type FoodItem struct { // definition of a food item in the prompt
	Name string
	Type string
}

func promptSelectFood(selections []string) (string, error) {
	// a list of displayed foods is created
	items := []FoodItem{}
	for _, nameOfFood := range selections {
		items = append(items, FoodItem{
			Name: strings.Title(nameOfFood),
			Type: typesOfFoods[nameOfFood],
		})
	}

	// how the prompt should be displayed
	templates := &promptui.SelectTemplates{
		Label:    "{{ . | white | bold }}",
		Active:   "\U0001F374 {{ .Name | green | bold }} ({{ .Type | cyan }})",
		Inactive: "  {{ .Name | red }} ({{ .Type | cyan }})",
		Selected: "\U0001F374 {{ .Name | green }}",
		Details: `
{{ "   ---      Food      ---   " | bold }}
{{ "Name:" | faint }}	{{ .Name | yellow }}
{{ "Type:" | faint }}	{{ .Type | yellow }}`,
	}

	// for using the SEARCH feature in the prompt
	searcher := func(input string, index int) bool {
		foodItem := items[index]
		name := strings.Replace(strings.ToLower(foodItem.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return fuzzy.MatchNormalizedFold(name, input)
	}

	// settings for the prompt
	prompt := promptui.Select{
		Label: "Select a food:",
		Items: items,
		// IsVimMode: true,
		Templates: templates,
		Size:      9,
		Searcher:  searcher,
	}

	// open prompt
	i, _, err := prompt.Run()

	return strings.ToLower(items[i].Name), err // return results of the prompt
}
