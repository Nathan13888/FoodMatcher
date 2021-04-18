package main

import (
	"github.com/spf13/viper"
)

func setDefaults() {
	viper.SetDefault("foods", map[string]string{
		// 7 Appetizers
		"nigiri sushi":   string(APPETIZER),
		"tofu":           string(APPETIZER),
		"salad":          string(APPETIZER),
		"kimchi":         string(APPETIZER),
		"guacamole":      string(APPETIZER),
		"tortilla chips": string(APPETIZER),
		"sausages":       string(APPETIZER),
		// 6 Main Course Meals
		"steak":           string(MAIN_COURSE),
		"sushimi platter": string(MAIN_COURSE),
		"mac n' cheese":   string(MAIN_COURSE),
		"fried chicken":   string(MAIN_COURSE),
		"roasted duck":    string(MAIN_COURSE),
		"pork belly":      string(MAIN_COURSE),
		// 7 Desserts
		"mango pudding":      string(DESSERT),
		"apple pie":          string(DESSERT),
		"fruit":              string(DESSERT),
		"chocolate cake":     string(DESSERT),
		"tiramisu":           string(DESSERT),
		"red bean ice cream": string(DESSERT),
		"strawberry cupcake": string(DESSERT),
	})
}

func setupConfig() {
	setDefaults()

	viper.SetConfigName("data")      // name of config file (without extension)
	viper.SetConfigType("json")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.fm") // call multiple times to add many search paths
	viper.AddConfigPath(".")         // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			// viper.SafeWriteConfig() // could create a new config with the defaults
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}

	typesOfFoods = viper.GetStringMapString("foods")
}
