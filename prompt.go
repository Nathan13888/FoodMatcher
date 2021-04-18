package main

import (
	"errors"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/manifoldco/promptui"
)

func promptFood() (string, error) {
	validate := func(input string) error {
		if len(search(input)) > 0 {
			return nil
		}
		return errors.New("invalid input")
	}

	prompt := promptui.Prompt{
		Label:    "Enter a food",
		Validate: validate,
	}

	res, err := prompt.Run()

	return res, err
}

func promptSelectFood(selections []string) (string, error) {
	type Food struct {
		Name string
		Type string
	}
	items := []Food{}

	for _, nameOfFood := range selections {
		items = append(items, Food{
			Name: strings.Title(nameOfFood),
			Type: typesOfFoods[nameOfFood],
		})
	}

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

	searcher := func(input string, index int) bool {
		foodItem := items[index]
		name := strings.Replace(strings.ToLower(foodItem.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return fuzzy.MatchNormalizedFold(name, input)
	}

	prompt := promptui.Select{
		Label: "Select a food:",
		Items: items,
		// IsVimMode: true,
		Templates: templates,
		Size:      3,
		Searcher:  searcher,
	}

	_, res, err := prompt.Run()

	return res, err
}
